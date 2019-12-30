package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var MinNamePrice = sdk.Coins{sdk.NewInt64Coin("orgtoken", 1)}

type User struct {
	Name    string         `json:"name"`
	Address sdk.AccAddress `json:"address"`
	Role    string         `json:"role"`
}

type Org struct {
	Name  string         `json:"name"`
	Owner sdk.AccAddress `json:"owner"`
	Users []User         `json:"users"`
}

func NewOrg() Org {
	return Org{}
}

func (o Org) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Owner: %s
Value: %s
Price: %s`, o.Name, o.Owner))
}
