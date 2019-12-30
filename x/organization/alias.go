package organization

import (
	"github.com/cosmos/e025/organization/x/organization/internal/keeper"
	"github.com/cosmos/e025/organization/x/organization/internal/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewKeeper        = keeper.NewKeeper
	NewQuerier       = keeper.NewQuerier
	NewMsgCreateOrg  = types.NewMsgCreateOrg
	NewOrg           = types.NewOrg
	NewMsgAddUser    = types.NewMsgAddUser
	NewMsgDeleteUser = types.NewMsgDeleteUser
	NewMsgDeleteOrg  = types.NewMsgDeleteOrg
	ModuleCdc        = types.ModuleCdc
	RegisterCodec    = types.RegisterCodec
)

type (
	Keeper          = keeper.Keeper
	MsgCreateOrg    = types.MsgCreateOrg
	MsgAddUser      = types.MsgAddUser
	MsgDeleteUser   = types.MsgDeleteUser
	MsgDeleteOrg    = types.MsgDeleteOrg
	Org             = types.Org
	User            = types.User
	QueryResResolve = types.QueryResResolve
	QueryResNames   = types.QueryResNames
)
