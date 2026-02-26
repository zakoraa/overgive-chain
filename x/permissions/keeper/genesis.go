package keeper

import (
	"context"

	"overgive-chain/x/permissions/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx context.Context, genState types.GenesisState) error {
	for _, elem := range genState.AllowedMap {
		if err := k.Allowed.Set(ctx, elem.Index, elem); err != nil {
			return err
		}
	}

	return k.Params.Set(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis.
func (k Keeper) ExportGenesis(ctx context.Context) (*types.GenesisState, error) {
	var err error

	genesis := types.DefaultGenesis()
	genesis.Params, err = k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}
	if err := k.Allowed.Walk(ctx, nil, func(_ string, val types.Allowed) (stop bool, err error) {
		genesis.AllowedMap = append(genesis.AllowedMap, val)
		return false, nil
	}); err != nil {
		return nil, err
	}

	return genesis, nil
}
