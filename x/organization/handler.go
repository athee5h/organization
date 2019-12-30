package organization

import (
	"fmt"

	"github.com/cosmos/e025/organization/x/organization/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler returns a handler for "orgservice" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgCreateOrg:
			return handleMsgCreateOrg(ctx, keeper, msg)
		case MsgAddUser:
			return handleMsgAddUser(ctx, keeper, msg)
		case MsgDeleteUser:
			return handleMsgDeleteUser(ctx, keeper, msg)
		case MsgDeleteOrg:
			return handleMsgDeleteOrg(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized orgservice Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// Handle a message to set name
func handleMsgCreateOrg(ctx sdk.Context, keeper Keeper, msg MsgCreateOrg) sdk.Result {
	if keeper.HasOwner(ctx, msg.Name) {
		return sdk.ErrUnknownRequest("Address already exist").Result()
	}

	org := types.Org{
		Name:  msg.Name,
		Owner: msg.Owner,
	}
	keeper.SetOrg(ctx, msg.Name, org) // If so, set the name to the value specified in the msg.
	return sdk.Result{}               // return
}

func handleMsgAddUser(ctx sdk.Context, keeper Keeper, msg MsgAddUser) sdk.Result {
	org := keeper.GetOrg(ctx, msg.OrgName)
	if org.Name == "" {
		return sdk.ErrUnknownRequest("Organization does not exist").Result()
	}

	if org.Owner.String() != msg.Sender.String() {
		return sdk.ErrUnknownRequest("Incorrect Owner").Result()
	}

	users := org.Users
	users = append(users, msg.User)
	org.Users = users
	keeper.SetOrg(ctx, msg.OrgName, org)
	return sdk.Result{}
}

func handleMsgDeleteUser(ctx sdk.Context, keeper Keeper, msg MsgDeleteUser) sdk.Result {
	org := keeper.GetOrg(ctx, msg.OrgName)
	if org.Name == "" {
		return sdk.ErrUnknownRequest("Organization does not exist").Result()
	}

	if org.Owner.String() != msg.Sender.String() {
		return sdk.ErrUnknownRequest("Incorrect Owner").Result()
	}

	users := org.Users
	index := -1

	for i, user := range users {
		if user.Name == msg.UserName {
			index = i
			break
		}
	}

	if index != -1 {
		copy(users[index:], users[index+1:])
		users[len(users)-1] = User{}
		users = users[:len(users)-1]
	} else {
		return sdk.ErrUnknownRequest("user not found").Result()
	}

	org.Users = users
	keeper.SetOrg(ctx, msg.OrgName, org)
	return sdk.Result{}
}

func handleMsgDeleteOrg(ctx sdk.Context, keeper Keeper, msg MsgDeleteOrg) sdk.Result{
	org := keeper.GetOrg(ctx, msg.OrgName)
	if org.Name == "" {
		return sdk.ErrUnknownRequest("Organization does not exist").Result()
	}

	if org.Owner.String() != msg.Sender.String() {
		return sdk.ErrUnknownRequest("Incorrect Owner").Result()
	}

	keeper.DeleteOrg(ctx, msg.OrgName)
	return sdk.Result{}
}