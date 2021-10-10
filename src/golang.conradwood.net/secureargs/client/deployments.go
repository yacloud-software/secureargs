package main

import (
	"fmt"
	dm "golang.conradwood.net/apis/deploymonkey"
	"golang.conradwood.net/go-easyops/client"
	"golang.conradwood.net/go-easyops/utils"
	"io/ioutil"
	"os"
	"strings"
)

// return absolute filenames of all deployXXX.yaml files in current dir
func findDeployFiles() ([]string, error) {
	depldir, err := findDeploymentDir()
	if err != nil {
		return nil, err
	}
	fmt.Printf("Deploymentdir: %s\n", depldir)
	files, err := deplfiles(depldir)
	if err != nil {
		return nil, err
	}
	return files, nil
}

// return absolute path to deployment dir (relative to current dir)
func findDeploymentDir() (string, error) {
	cur, err := os.Getwd()
	if err != nil {
		return "", err
	}
	if utils.FileExists(cur + "/deployment") {
		return cur + "/deployment", nil
	}
	fmt.Printf("Current dir: %s\n", cur)
	return "", fmt.Errorf("deploymentdir not found in %s", cur)
}

// return absolute filenames of deploy.yamls (in dir)
func deplfiles(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var res []string
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".yaml") {
			res = append(res, dir+"/"+f.Name())
		}
	}
	return res, nil
}

func GetDeployMonkeyClient() dm.DeployMonkeyClient {
	dmc := dm.NewDeployMonkeyClient(client.Connect("deploymonkey.DeployMonkey"))
	return dmc
}
