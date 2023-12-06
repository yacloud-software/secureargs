package main

import (
	"context"
	"flag"
	"fmt"
	"golang.conradwood.net/apis/common"
	pb "golang.conradwood.net/apis/secureargs"
	"golang.conradwood.net/go-easyops/auth"
	"golang.conradwood.net/go-easyops/errors"
	"golang.conradwood.net/go-easyops/server"
	"golang.conradwood.net/go-easyops/utils"
	"golang.conradwood.net/secureargs/db"
	"google.golang.org/grpc"
	"os"
)

var (
	allow_repository_id = flag.Bool("allow_repository_id_only", false, "if true allow access by repositoryid only (use only if using a single system assigning repo id, e.g. gitserver OR gerrit, not both")
	port                = flag.Int("port", 4100, "The grpc server port")
	migrate             = flag.Bool("migrate_from_gitserver_repository_id_to_artefactid", false, "if true do a one-off migration")
	argstore            *db.DBArg
)

type echoServer struct {
}

func main() {
	flag.Parse()
	if *migrate {
		utils.Bail("migration failed", Migrate())
		os.Exit(0)
	}
	fmt.Printf("Starting SecureArgsServiceServer...\n")
	var err error
	argstore = db.DefaultDBArg()
	sd := server.NewServerDef()
	sd.SetPort(*port)
	sd.SetRegister(server.Register(
		func(server *grpc.Server) error {
			e := new(echoServer)
			pb.RegisterSecureArgsServiceServer(server, e)
			return nil
		},
	))
	err = server.ServerStartup(sd)
	utils.Bail("Unable to start server", err)
	os.Exit(0)
}

/************************************
* grpc functions
************************************/

func (e *echoServer) SetArg(ctx context.Context, req *pb.SetArgRequest) (*common.Void, error) {
	su := auth.GetService(ctx)
	if su != nil && su.ID == auth.GetServiceIDByName("repobuilder.RepoBuilder") {
		fmt.Printf("Repobuilder is setting arg %s\n", req.Name)
	} else {
		err := needAuthorisation(ctx)
		if err != nil {
			return nil, err
		}
	}
	if req.Name == "" {
		return nil, errors.InvalidArgs(ctx, "missing argument name", "missing argument name")
	}
	if req.Value == "" {
		return nil, errors.InvalidArgs(ctx, "missing argument value", "missing argument value")
	}
	if !*allow_repository_id {
		if req.RepositoryID != 0 && req.ArtefactID == 0 {
			return nil, errors.InvalidArgs(ctx, "repositoryid obsolete, artefactid required", "repositoryid obsolete, artefactid required")
		}
	}
	if req.RepositoryID == 0 && req.ArtefactID == 0 {
		return nil, errors.InvalidArgs(ctx, "missing argument repository", "missing argument repository")
	}

	dbs, err := argstore.ByArtefactID(ctx, req.ArtefactID)
	if err != nil {
		return nil, err
	}
	var dbarg *pb.Arg
	for _, db := range dbs {
		if db.Name == req.Name {
			dbarg = db
			break
		}
	}
	if dbarg != nil {
		dbarg.Value = req.Value
		err = argstore.Update(ctx, dbarg)
	} else {
		dbarg = &pb.Arg{
			RepositoryID: req.RepositoryID,
			Name:         req.Name,
			Value:        req.Value,
			ArtefactID:   req.ArtefactID,
		}
		_, err = argstore.Save(ctx, dbarg)
	}
	if err != nil {
		return nil, err
	}
	return &common.Void{}, nil
}
func (e *echoServer) GetArgs(ctx context.Context, req *pb.GetArgsRequest) (*pb.ArgsResponse, error) {
	err := needAuthorisation(ctx)
	if err != nil {
		return nil, err
	}
	if !*allow_repository_id {
		if req.RepositoryID != 0 && req.ArtefactID == 0 {
			return nil, errors.InvalidArgs(ctx, "repositoryid obsolete, artefactid required", "repositoryid obsolete, artefactid required")
		}
	}
	if req.RepositoryID == 0 && req.ArtefactID == 0 {
		return nil, errors.InvalidArgs(ctx, "missing repository", "missing repository")
	}
	var rps []*pb.Arg

	// we can get args for both, artefactid and repositoryid in one request

	if req.ArtefactID != 0 {
		// retrieve for artefactid
		fmt.Printf("Getting args for artefact \"%d\"\n", req.ArtefactID)
		t_rps, err := argstore.ByArtefactID(ctx, req.ArtefactID)
		if err != nil {
			return nil, err
		}
		rps = append(rps, t_rps...)
	}

	if *allow_repository_id {
		if req.RepositoryID != 0 {
			// retrieve for repositoryid
			fmt.Printf("Getting args for repository \"%d\"\n", req.RepositoryID)
			t_rps, err := argstore.ByRepositoryID(ctx, req.RepositoryID)
			if err != nil {
				return nil, err
			}
			rps = append(rps, t_rps...)
		}
	}
	res := &pb.ArgsResponse{
		Args: make(map[string]string),
	}
	for _, rp := range rps {
		res.Args[rp.Name] = rp.Value
	}
	fmt.Printf("Returning %d args for repository #%d\n", len(res.Args), req.RepositoryID)
	return res, nil
}
func needAuthorisation(ctx context.Context) error {
	if auth.IsRoot(ctx) {
		return nil
	}
	s := auth.GetService(ctx)
	if s == nil {
		return errors.Unauthenticated(ctx, "access denied")
	}
	if s.ID != auth.GetServiceIDByName("autodeployer.AutoDeployer") {
		return errors.AccessDenied(ctx, "access denied for service %s [%s]", auth.Description(s), s.ID)
	}
	return nil

}

