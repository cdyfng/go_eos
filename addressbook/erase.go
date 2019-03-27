package addressbook

import (
	eos "github.com/eoscanada/eos-go"
	"github.com/cdyfng/go_eos/config"
)

func Erase(user eos.AccountName) *eos.Action {
	a := &eos.Action{
		Account: eos.AccountName(config.ReadEosAccount_1()),
		Name:    eos.ActionName("erase"),
		Authorization: []eos.PermissionLevel{
			{Actor: user, Permission: eos.PermissionName("active")},
		},
		ActionData: eos.NewActionData(User{
			User: 			user,
		}),		
	}
	return a
}


// Upsert represents the `addressbook::upsert` action.
type User struct {
	User     	eos.AccountName `json:"user"`
}
