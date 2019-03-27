package addressbook

import (
	eos "github.com/eoscanada/eos-go"
	"github.com/cdyfng/go_eos/config"
)

// upsert is an action to insert an address item
//upsert(name user, std::string first_name, std::string last_name, uint64_t age, std::string street, std::string city, std::string state) {
func Upsert(user eos.AccountName, first_name string,  last_name string, age uint64, street string, city string, state string) *eos.Action {
	a := &eos.Action{
		Account: eos.AccountName(config.ReadEosAccount_1()),
		Name:    eos.ActionName("upsert"),
		Authorization: []eos.PermissionLevel{
			{Actor: user, Permission: eos.PermissionName("active")},
		},
		ActionData: eos.NewActionData(AddressUser{
			User: 			user,
			First_name: 	first_name,
			Last_name: 		last_name,
			Age: 			age,
			Street: 		street,
			City: 			city,
			State: 			state,
		}),
	}
	return a
}


// Upsert represents the `addressbook::upsert` action.
type AddressUser struct {
	User     	eos.AccountName `json:"user"`
	First_name  string        	`json:"first_name"`
	Last_name   string   		`json:"last_name"`
	Age 		uint64          `json:"age"`
	Street    	string    		`json:"street"`
	City    	string			`json:"city"`
	State   	string			`json:"state"`
}
