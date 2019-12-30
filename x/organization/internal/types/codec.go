package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// ModuleCdc is the codec for the module
var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateOrg{}, "organization/CreateOrg", nil)
	cdc.RegisterConcrete(MsgAddUser{}, "organization/AddUser", nil)
	cdc.RegisterConcrete(MsgDeleteUser{}, "organization/DeleteUser", nil)
	cdc.RegisterConcrete(MsgDeleteOrg{}, "organization/DeleteOrg", nil)
}
