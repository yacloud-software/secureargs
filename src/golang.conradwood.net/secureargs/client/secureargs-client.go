package main

import (
	"context"
	"flag"
	"fmt"
	"strings"
	//	"golang.conradwood.net/apis/common"
	pb "golang.conradwood.net/apis/secureargs"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/utils"
	"os"
)

var (
	debug       = flag.Bool("debug", false, "debug information")
	afid        = flag.Uint64("artefactid", 0, "artefactid to set args for")
	argname     = flag.String("name", "", "name of argument to set")
	argvalue    = flag.String("value", "", "value of argument to set")
	migrate     = flag.Bool("migrate", false, "migrate a service to secargs")
	compromised = flag.Bool("compromised", false, "consider parameters in current git compromised (change them and update backends)")
	view        = flag.Bool("view", false, "view secure args")
)

func main() {
	flag.Parse()
	if *afid == 0 {
		fmt.Printf("-artefactid is required\n")
		os.Exit(10)
	}
	if *compromised {
		Compromised()
		os.Exit(0)
	}
	if *migrate {
		Migrate()
		os.Exit(0)
	}
	if *view {
		View()
		os.Exit(0)
	}
	ctx := authremote.Context()
	ctx = authremote.Context()
	err := Set(ctx, *afid, *argname, parseValue())
	utils.Bail("failed to set arg", err)
	fmt.Printf("Argument set.\n")

	fmt.Printf("Done.\n")
	os.Exit(0)
}
func parseValue() string {
	res := *argvalue
	if strings.HasPrefix(res, "/") {
		f, err := utils.ReadFile(res)
		utils.Bail("failed to read file", err)
		res = string(f)
	}
	return res
}
func View() {
	ctx := authremote.Context()
	svc := pb.GetSecureArgsServiceClient()
	args, err := svc.GetArgs(ctx, &pb.GetArgsRequest{ArtefactID: *afid})
	utils.Bail("failed to get args", err)
	for k, v := range args.Args {
		fmt.Printf("%s: %s\n", k, v)
	}
	fmt.Printf("Done")
}

func Set(ctx context.Context, artefactid uint64, name string, value string) error {
	svc := pb.GetSecureArgsServiceClient()
	req := &pb.SetArgRequest{
		ArtefactID: artefactid,
		Name:       name,
		Value:      value,
	}
	_, err := svc.SetArg(ctx, req)
	return err
}





