package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/e025/organization/x/organization/internal/types"
)

type Keeper struct {
	CoinKeeper bank.Keeper

	storeKey sdk.StoreKey // Unexposed key to access store from sdk.Context

	cdc *codec.Codec // The wire codec for binary encoding/decoding.
}

func NewKeeper(coinKeeper bank.Keeper, storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		CoinKeeper: coinKeeper,
		storeKey:   storeKey,
		cdc:        cdc,
	}
}

func (k Keeper) GetOrg(ctx sdk.Context, name string) types.Org {
	store := ctx.KVStore(k.storeKey)
	if !k.IsNamePresent(ctx, name) {
		return types.Org{}
	}
	bz := store.Get([]byte(name))
	var org types.Org
	k.cdc.MustUnmarshalBinaryBare(bz, &org)
	return org
}

func (k Keeper) GetOrgOwnerAddress(ctx sdk.Context, name string) sdk.Address {
	return k.GetOrg(ctx, name).Owner
}

func (k Keeper) SetOrg(ctx sdk.Context, name string, org types.Org) {
	if org.Owner.Empty() {
		return
	}
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(name), k.cdc.MustMarshalBinaryBare(org))
}

// Deletes the entire Whois metadata struct for a name
func (k Keeper) DeleteOrg(ctx sdk.Context, name string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(name))
}

func (k Keeper) HasOwner(ctx sdk.Context, name string) bool {
	return !k.GetOrg(ctx, name).Owner.Empty()
}

// ResolveName - returns the string that the name resolves to
func (k Keeper) ResolveName(ctx sdk.Context, name string) string {
	return k.GetOrg(ctx, name).Name
}

func (k Keeper) GetNamesIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, nil)
}

// Check if the name is present in the store or not
func (k Keeper) IsNamePresent(ctx sdk.Context, name string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(name))
}
