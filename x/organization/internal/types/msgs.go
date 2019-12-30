package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RouterKey is the module name router key
const RouterKey = ModuleName // this was defined in your key.go file

// MsgSetName defines a SetName message
type MsgCreateOrg struct {
	Name  string         `json:"name"`
	Owner sdk.AccAddress `json:"owner"`
}

// NewMsgSetName is a constructor function for MsgSetName
func NewMsgCreateOrg(name string, owner sdk.AccAddress) MsgCreateOrg {
	return MsgCreateOrg{
		Name:  name,
		Owner: owner,
	}
}

// Route should return the name of the module
func (msg MsgCreateOrg) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateOrg) Type() string { return "set_name" }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateOrg) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.Name) == 0 {
		return sdk.ErrUnknownRequest("Name cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateOrg) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgCreateOrg) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

//-------------------------------------------------------------

// MsgAddUser defines a add user message
type MsgAddUser struct {
	OrgName string         `json:"org_name"`
	User    User           `json:"user"`
	Sender  sdk.AccAddress `json:"sender"`
}

// NewMsgSetName is a constructor function for MsgSetName
func NewMsgAddUser(orgName string, userName string, address sdk.AccAddress, role string, msgSender sdk.AccAddress) MsgAddUser {
	return MsgAddUser{
		OrgName: orgName,
		User: User{
			Name:    userName,
			Address: address,
			Role:    role,
		},
		Sender: msgSender,
	}
}

// Route should return the name of the module
func (msg MsgAddUser) Route() string { return RouterKey }

// Type should return the action
func (msg MsgAddUser) Type() string { return "add_user" }

// ValidateBasic runs stateless checks on the message
func (msg MsgAddUser) ValidateBasic() sdk.Error {
	if msg.Sender.Empty() {
		return sdk.ErrUnknownRequest("Sender cannot be empty")
	}
	if len(msg.OrgName) == 0 {
		return sdk.ErrUnknownRequest("OrgName cannot be empty")
	}
	if msg.User.Address.Empty() {
		return sdk.ErrUnknownRequest("User address cannot be empty")
	}
	if len(msg.User.Name) == 0 {
		return sdk.ErrUnknownRequest("Name cannot be empty")
	}
	if len(msg.User.Role) == 0 {
		return sdk.ErrUnknownRequest("Role cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgAddUser) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgAddUser) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

//-------------------------------------------------------

// MsgAddUser defines a add user message
type MsgDeleteUser struct {
	OrgName  string         `json:"org_name"`
	UserName string         `json:"user_name"`
	Sender   sdk.AccAddress `json:"sender"`
}

// NewMsgSetName is a constructor function for MsgSetName
func NewMsgDeleteUser(orgName string, userName string, msgSender sdk.AccAddress) MsgDeleteUser {
	return MsgDeleteUser{
		OrgName:  orgName,
		UserName: userName,
		Sender:   msgSender,
	}
}

// Route should return the name of the module
func (msg MsgDeleteUser) Route() string { return RouterKey }

// Type should return the action
func (msg MsgDeleteUser) Type() string { return "delete_user" }

// ValidateBasic runs stateless checks on the message
func (msg MsgDeleteUser) ValidateBasic() sdk.Error {
	if msg.Sender.Empty() {
		return sdk.ErrUnknownRequest("Sender cannot be empty")
	}
	if len(msg.OrgName) == 0 {
		return sdk.ErrUnknownRequest("OrgName cannot be empty")
	}
	if len(msg.UserName) == 0 {
		return sdk.ErrUnknownRequest("Address cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgDeleteUser) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgDeleteUser) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

//-------------------------------------------------------------------

type MsgDeleteOrg struct {
	OrgName string         `json:"org_name"`
	Sender  sdk.AccAddress `json:"sender"`
}

func NewMsgDeleteOrg(orgName string, msgSender sdk.AccAddress) MsgDeleteOrg {
	return MsgDeleteOrg{
		OrgName: orgName,
		Sender:  msgSender,
	}
}

func (msg MsgDeleteOrg) Route() string { return RouterKey }

func (msg MsgDeleteOrg) Type() string { return "delete_user" }

func (msg MsgDeleteOrg) ValidateBasic() sdk.Error {
	if msg.Sender.Empty() {
		return sdk.ErrUnknownRequest("Sender cannot be empty")
	}
	if len(msg.OrgName) == 0 {
		return sdk.ErrUnknownRequest("OrgName cannot be empty")
	}
	return nil
}

func (msg MsgDeleteOrg) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgDeleteOrg) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}
