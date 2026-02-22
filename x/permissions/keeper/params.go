package keeper

import (
	"overgive-chain/x/permissions/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetParams(ctx sdk.Context) (types.Params, error) {
	return k.Params.Get(ctx)
}

func (k Keeper) IsAllowedWriter(ctx sdk.Context, addr string) (bool, error) {
	params, err := k.GetParams(ctx)
	if err != nil {
		return false, err
	}

	for _, writer := range params.AllowedWriters {
		if writer == addr {
			return true, nil
		}
	}
	return false, nil
}
