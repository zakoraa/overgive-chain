package keeper

import (
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	corestore "cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"overgive-chain/x/permissions/types"
)

type Keeper struct {
	storeService corestore.KVStoreService
	cdc          codec.Codec
	addressCodec address.Codec
	// Address capable of executing a MsgUpdateParams message.
	// Typically, this should be the x/gov module account.
	authority sdk.AccAddress

	Schema         collections.Schema
	Params         collections.Item[types.Params]
	AllowedWriters collections.Map[string, bool]
}

func NewKeeper(
	storeService corestore.KVStoreService,
	cdc codec.Codec,
	addressCodec address.Codec,
	authority sdk.AccAddress,

) Keeper {
	if _, err := addressCodec.BytesToString(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address %s: %s", authority, err))
	}

	sb := collections.NewSchemaBuilder(storeService)

	k := Keeper{
		storeService: storeService,
		cdc:          cdc,
		addressCodec: addressCodec,
		authority:    authority,

		Params: collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),

		AllowedWriters: collections.NewMap(
			sb,
			types.AllowedWritersKeyPrefix,
			"allowed_writers",
			collections.StringKey,
			collections.BoolValue,
		),
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}
	k.Schema = schema

	return k
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() []byte {
	return k.authority
}

func (k Keeper) SetAllowedWriter(ctx sdk.Context, addr sdk.AccAddress) error {
	return k.AllowedWriters.Set(ctx, addr.String(), true)
}

func (k Keeper) RemoveAllowedWriter(ctx sdk.Context, addr sdk.AccAddress) error {
	return k.AllowedWriters.Remove(ctx, addr.String())
}

func (k Keeper) IsAllowedWriter(ctx sdk.Context, addr sdk.AccAddress) (bool, error) {
	return k.AllowedWriters.Has(ctx, addr.String())
}
