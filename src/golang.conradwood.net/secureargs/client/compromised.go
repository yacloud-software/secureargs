package main

import (
	"context"
	"fmt"
	"golang.conradwood.net/apis/auth"
	dm "golang.conradwood.net/apis/deploymonkey"
	pm "golang.conradwood.net/apis/postgresmgr"
	sa "golang.conradwood.net/apis/secureargs"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/utils"
	"strings"
)

func Compromised() {
	files, err := findDeployFiles()
	utils.Bail("failed to find deploy files", err)
	for _, f := range files {
		err := checkFile(f)
		utils.Bail("Failed to process file", err)
	}
	fmt.Printf("Done.\n")
}
func checkFile(filename string) error {
	fmt.Printf("Checking args in %s\n", filename)
	ctx := authremote.Context()
	fmt.Printf("File: %s\n", filename)
	r, err := utils.ReadFile(filename)
	if err != nil {
		return err
	}
	pcr, err := GetDeployMonkeyClient().ParseConfigFile(ctx, &dm.ParseRequest{Config: string(r)})
	if err != nil {
		return err
	}
	or := uint64(0)
	fmt.Printf("%d Groupdefs in deployment file\n", len(pcr.GroupDef))
	for _, gd := range pcr.GroupDef {
		for _, a := range gd.Applications {
			fmt.Printf("Dealing with application %s\n", a.Binary)
			repo := a.RepositoryID
			if or != 0 && or != repo {
				panic("cannot deal with deployment yaml files with two repos")
			}
			am := make(map[string]string)
			for _, arg := range a.Args {
				k, v := splitArg(arg)
				am[k] = v
			}
			m := &Changer{repoid: repo, args: am}
			for k, v := range am {
				if !strings.Contains(v, "${SECURE-") {
					continue
				}
				if *debug {
					fmt.Printf("%s = %s\n", k, v)
				}
				err := m.migrateValue(k, v)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type Changer struct {
	repoid uint64
	args   map[string]string
}

func (m *Changer) migrateValue(key, value string) error {
	var err error
	fmt.Printf("Changing value for \"%s\" in repo \"%d\" \n", key, m.repoid)
	// add new handlers for new types here
	if key == "token" {
		err = m.change(value, m.backend_changeToken)
	} else if key == "dbpw" {
		err = m.change(value, m.backend_changePSQL)
	} else {
		fmt.Printf("WARNING - cannot change type \"%s\"\n", key)
	}
	return err
}

func (m *Changer) change(secArgName string, backend func(ctx context.Context, old_value string) (string, error)) error {
	ctx := authremote.Context()
	secname, current_value, err := m.getSecureArg(secArgName)
	if err != nil {
		return err
	}
	if *debug {
		fmt.Printf("Current value of \"%s\": \"%s\"\n", secArgName, current_value)
	}
	new_value, err := backend(ctx, current_value)
	if err != nil {
		fmt.Printf("Backend error: %s\n", err)
		return err
	}
	// store it
	sar := &sa.SetArgRequest{
		RepositoryID: m.repoid,
		Name:         secname,
		Value:        new_value,
	}
	_, err = sa.GetSecureArgsClient().SetArg(ctx, sar)
	if err != nil {
		return err
	}
	if *debug {
		fmt.Printf("New value of \"%s\": \"%s\"\n", secArgName, new_value)
	}
	return nil
}

// getArg is privileged, so this can only be run by root
func (m *Changer) getSecureArg(saname string) (string, string, error) {
	ctx := authremote.Context()
	gar := &sa.GetArgsRequest{RepositoryID: m.repoid}
	resp, err := sa.GetSecureArgsClient().GetArgs(ctx, gar)
	if err != nil {
		return "", "", err
	}

	for k, v := range resp.Args {
		fn := fmt.Sprintf("${SECURE-%s}", k)
		if saname == fn {
			return k, v, nil
		}
		if *debug {
			fmt.Printf("%s = %s\n", fn, v)
		}
	}
	return "", "", fmt.Errorf("\"%s\" not found", saname)
}

/****************************************************************
* if a secure arg becomes compromised we need to not only change
* the variable, but because it is a "shared secret" we need to
* also tell the backend. Each backend has its own handler
* each handler gets invoked with current value and is expected
* to return a new one (which it negotiated with the backend)
****************************************************************/

// change the service token in authmanager
func (m *Changer) backend_changeToken(ctx context.Context, current_value string) (string, error) {
	// change it
	tc := &auth.TokenCompromisedRequest{Token: current_value}
	nt, err := authremote.GetAuthManagerClient().TokenCompromised(ctx, tc)
	if err != nil {
		return "", err
	}
	new_value := nt.Token
	return new_value, nil
}

// change a postgres password
func (m *Changer) backend_changePSQL(ctx context.Context, current_value string) (string, error) {

	dbhost := m.args["dbhost"]
	dbuser := m.args["dbuser"]
	cpr := &pm.ChangePasswordRequest{
		User:     dbuser,
		Host:     dbhost,
		Password: current_value,
	}
	cp, err := pm.GetPostgresMgrClient().ChangePassword(ctx, cpr)
	if err != nil {
		return "", err
	}
	return cp.NewPassword, nil
}




