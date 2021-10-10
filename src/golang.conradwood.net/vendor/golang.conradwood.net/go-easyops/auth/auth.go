package auth

import (
	"fmt"
	apb "golang.conradwood.net/apis/auth"
	//	"golang.conradwood.net/go-easyops/client"
	"golang.conradwood.net/go-easyops/common"
	"golang.conradwood.net/go-easyops/rpc"
	//	"golang.conradwood.net/go-easyops/tokens"
	"golang.org/x/net/context"
)

func IsRoot(ctx context.Context) bool {
	return IsRootUser(rpc.CallStateFromContext(ctx).User())
}
func IsRootUser(user *apb.User) bool {
	return IsInGroupByUser(user, "1")
}
func IsInGroupByUser(user *apb.User, groupid string) bool {
	if user == nil || groupid == "" || user.Groups == nil {
		return false
	}
	for _, g := range user.Groups {
		if g.ID == groupid {
			return true
		}
	}

	return false
}
func IsInGroup(ctx context.Context, groupid string) bool {
	u := GetUser(ctx)
	return IsInGroupByUser(u, groupid)
}
func GetUser(ctx context.Context) *apb.User {
	cs := rpc.CallStateFromContext(ctx)
	if cs == nil {
		return nil
	}
	return cs.User()
}

func PrintUser(u *apb.User) {
	if u == nil {
		return
	}
	fmt.Printf("User ID: %s\n", u.ID)
	fmt.Printf("  Email: %s\n", u.Email)
	fmt.Printf("  Abbrev:%s\n", u.Abbrev)
}
func PrintSignedUser(uu *apb.SignedUser) {
	u := common.VerifySignedUser(uu)
	if u == nil {
		return
	}

	fmt.Printf("User ID: %s\n", u.ID)
	fmt.Printf("  Email: %s\n", u.Email)
	fmt.Printf("  Abbrev:%s\n", u.Abbrev)
}

func GetService(ctx context.Context) *apb.User {
	cs := rpc.CallStateFromContext(ctx)
	if cs == nil {
		return nil
	}
	return cs.CallerService()
}

// one line description of the user/caller
func Description(user *apb.User) string {
	if user == nil {
		return "ANONYMOUS"
	}
	if user.Abbrev != "" {
		return user.Abbrev
	}
	if user.Email != "" {
		return user.Email
	}
	return "user #" + user.ID
}
func UserIDString(user *apb.User) string {
	if user == nil {
		return "ANONYMOUS"
	}
	if user.Abbrev != "" {
		return "#" + user.ID + " (" + user.Abbrev + ")"
	}
	if user.Email != "" {
		return "#" + user.ID + " (" + user.Email + ")"
	}
	return "user #" + user.ID
}

func CurrentUserString(ctx context.Context) string {
	u := GetUser(ctx)
	if u == nil {
		return "ANONYMOUS"
	}
	return fmt.Sprintf("User #%s (%s)", u.ID, u.Email)
}
