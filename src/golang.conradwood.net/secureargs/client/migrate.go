package main

import (
	"fmt"
	dm "golang.conradwood.net/apis/deploymonkey"
	sa "golang.conradwood.net/apis/secureargs"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/utils"
	//	"gopkg.in/yaml.v2"
	"bytes"
	"path/filepath"
	"strings"
)

var (
	replaced []*replaceValue
)

/*
we find various "known sensitive" args in the deploy.yaml files, for example dbpw and token.
they are then replaced with ${SECURE-XXX} tags (so the autodeplyer resolves them)
identical parameter name/value pairs in the same repository receive the same SECURE-XXX tag
*/
func Migrate() {
	fmt.Printf("Migrating...\n")
	files, err := findDeployFiles()
	utils.Bail("failed to find deploy files", err)
	for _, f := range files {
		err := parseFile(f)
		utils.Bail("failed to parse file", err)
	}
	for _, r := range replaced {
		err := createSecArg(r)
		utils.Bail("failed to create replacer", err)
	}
	for _, f := range files {
		err := migrateFile(f)
		utils.Bail("failed to migrate file", err)
	}
}

// load file and replace all args
func migrateFile(filename string) error {
	b, err := utils.ReadFile(filename)
	if err != nil {
		return err
	}
	s := string(b)
	var res bytes.Buffer
	for _, line := range strings.Split(s, "\n") {
		l := line
		for _, r := range replaced {
			o := "-" + r.arg + "=" + r.value
			n := "-" + r.arg + "=" + r.varName()
			l = strings.ReplaceAll(l, o, n)
		}
		res.WriteString(l + "\n")
	}
	fmt.Printf("file: %s\n%s\n", filename, res.String())
	err = utils.WriteFile(filename, []byte(res.String()))
	return err
}

// finds all args that need to be migrated (and fills 'replaced' array)
func parseFile(filename string) error {
	ctx := authremote.Context()
	fmt.Printf("File: %s\n", filename)
	r, err := utils.ReadFile(filename)
	if err != nil {
		return err
	}
	dmc := GetDeployMonkeyClient()
	pcr, err := dmc.ParseConfigFile(ctx, &dm.ParseRequest{Config: string(r)})
	if err != nil {
		return err
	}
	for _, gd := range pcr.GroupDef {
		for _, a := range gd.Applications {
			fmt.Printf("    Binary: %s\n", filepath.Base(a.Binary))
			for _, arg := range a.Args {
				if strings.Contains(arg, "${SECURE") {
					continue
				}
				if !strings.Contains(arg, "=") {
					continue
				}

				k, v := splitArg(arg)
				if k == "token" || k == "dbpw" || strings.Contains(k, "private") || strings.Contains(k, "assw") {
					fmt.Printf("      Arg: %s = \"%s\"\n", k, v)
					rpl := findReplacer(k)
					if rpl != nil {
						if rpl.value != v {
							panic(fmt.Sprintf("key %s has two values", k))
						}
					}
					if rpl == nil {
						rpl = &replaceValue{
							artefactid: *afid,
							key:        strings.ToUpper(k),
							value:      v,
							arg:        k}
						replaced = append(replaced, rpl)
					}
				}
			}
		}
	}
	for _, r := range replaced {
		fmt.Printf("artefact %d: %s=%s -> %s=%s\n", r.artefactid, r.arg, r.value, r.arg, r.key)
	}
	return nil
}

func splitArg(n string) (string, string) {
	if len(n) < 2 {
		fmt.Printf("Too short: %s\n", n)
		return n, "TOOSHORT"
	}
	n = n[1:] // remove the -
	eq := strings.Index(n, "=")
	if eq == -1 {
		fmt.Printf("argument has no equal sign (%s)\n", n)
		return n, "true"
	}
	k := n[:eq]
	v := n[eq+1:]
	return k, v
}

/********************************* replace a line *********************/
type replaceValue struct {
	key        string // "TOKEN1"
	arg        string // "token"
	value      string // "foobar"
	artefactid uint64
}

func findReplacer(arg string) *replaceValue {
	for _, r := range replaced {
		if r.arg == arg {
			return r
		}
	}
	return nil
}
func (r *replaceValue) varName() string {
	return fmt.Sprintf("${SECURE-%s}", r.key)
}

func createSecArg(r *replaceValue) error {
	ctx := authremote.Context()
	sar := &sa.SetArgRequest{
		ArtefactID: r.artefactid,
		Name:       r.key,
		Value:      r.value,
	}
	_, err := sa.GetSecureArgsClient().SetArg(ctx, sar)
	if err != nil {
		return err
	}

	return nil
}


