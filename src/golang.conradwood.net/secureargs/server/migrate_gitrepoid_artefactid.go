package main

import (
	"context"
	"fmt"
	af "golang.conradwood.net/apis/artefact"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/secureargs/db"
)

// migrate secureargs with just a repositoryid to artefactid (using gitserver as source)
func Migrate() error {
	ctx := context.Background()
	argstore = db.DefaultDBArg()
	t_rps, err := argstore.All(ctx)
	if err != nil {
		return err
	}
	var xerr error
	to_update := 0
	updated := 0
	for _, arg := range t_rps {
		if arg.ArtefactID == 0 {
			to_update++
			fmt.Printf("Arg #%03d: repo %d, artefactid %d\n", arg.ID, arg.RepositoryID, arg.ArtefactID)
			afid, err := get_artefact_id(arg.RepositoryID)
			if err != nil {
				fmt.Printf("Arg %d failed to get artefactid: %s\n", arg.ID, err)
				xerr = err
				continue
			}
			arg.ArtefactID = afid.ID
			err = argstore.Update(ctx, arg)
			if err != nil {
				fmt.Printf("Arg %d failed to update %s\n", arg.ID, err)
				return err
			}
			fmt.Printf("   Set artefactid to %d\n", arg.ArtefactID)
			updated++
		}
	}
	fmt.Printf("needed to update: %d, actually updated %d\n", to_update, updated)
	return xerr
}

func get_artefact_id(repoid uint64) (*af.ArtefactID, error) {
	ctx := authremote.Context()
	r, err := af.GetArtefactClient().GetArtefactIDForRepo(ctx, &af.ID{ID: repoid})
	if err != nil {
		return nil, err
	}
	return r, nil
}





