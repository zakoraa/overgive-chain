package keeper

import (
	"overgive-chain/x/permissions/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetParams(ctx sdk.Context) (types.Params, error) {
	return k.Params.Get(ctx)
}