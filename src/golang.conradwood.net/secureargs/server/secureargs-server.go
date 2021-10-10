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
	"golang.conradwood.net/go-easyops/sql"
	"golang.conradwood.net/go-easyops/utils"
	"golang.conradwood.net/secureargs/db"
	"google.golang.org/grpc"
	"os"
)

var (
	port     = flag.Int("port", 4100, "The grpc server port")
	argstore *db.DBArg
)

type echoServer struct {
}

func main() {
	flag.Parse()
	fmt.Printf("Starting SecureArgsServiceServer...\n")
	psql, err := sql.Open()
	utils.Bail("cannot open sql", err)
	argstore = db.NewDBArg(psql)
	sd := server.NewServerDef()
	sd.Port = *port
	sd.Register = server.Register(
		func(server *grpc.Server) error {
			e := new(echoServer)
			pb.RegisterSecureArgsServiceServer(server, e)
			return nil
		},
	)
	err = server.ServerStartup(sd)
	utils.Bail("Unable to start server", err)
	os.Exit(0)
}

/************************************
* grpc functions
************************************/

func (e *echoServer) SetArg(ctx context.Context, req *pb.SetArgRequest) (*common.Void, error) {
	su := auth.GetService(ctx)
	if su != nil && su.ID == "3539" {
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
	if req.RepositoryID == 0 {
		return nil, errors.InvalidArgs(ctx, "missing argument repository", "missing argument repository")
	}

	dbs, err := argstore.ByRepositoryID(ctx, req.RepositoryID)
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
	if req.RepositoryID == 0 {
		return nil, errors.InvalidArgs(ctx, "missing repository", "missing repository")
	}
	fmt.Printf("Getting args for repository \"%d\"\n", req.RepositoryID)
	rps, err := argstore.ByRepositoryID(ctx, req.RepositoryID)
	if err != nil {
		return nil, err
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
	if s.ID != "18" {
		return errors.AccessDenied(ctx, "access denied for service %s [%s]", auth.Description(s), s.ID)
	}
	return nil

}
