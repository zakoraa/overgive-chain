package keeper

import (
	"overgive-chain/x/permissions/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetParams(ctx sdk.Context) (types.Params, error) {
	return k.Params.Get(ctx)
}

func (k Keeper) IsAllowedWriter(ctx sdk.Context, addr string) (bool, error) {
    params, err := k.Params.Get(ctx)
    if err != nil {
        return false, err
    }

    for _, allowed := range params.AllowedWriters {
        if allowed == addr {
            return true, nil
        }
    }

    return false, nil
}