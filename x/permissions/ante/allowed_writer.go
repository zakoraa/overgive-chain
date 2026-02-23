package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"

	"overgive-chain/x/permissions/keeper"
)

type AllowedWriterDecorator struct {
	PermissionsKeeper keeper.Keeper
}

func NewAllowedWriterDecorator(pk keeper.Keeper) AllowedWriterDecorator {
	return AllowedWriterDecorator{
		PermissionsKeeper: pk,
	}
}

func (awd AllowedWriterDecorator) AnteHandle(
	ctx sdk.Context,
	tx sdk.Tx,
	simulate bool,
	next sdk.AnteHandler,
) (sdk.Context, error) {

	if simulate {
		return next(ctx, tx, simulate)
	}

	txWithSigners, ok := tx.(signing.Tx)
	if !ok {
		return ctx, sdkerrors.ErrTxDecode.Wrap("invalid transaction type")
	}

	signers, err := txWithSigners.GetSigners()
	if err != nil {
		return ctx, err
	}
	for _, signer := range signers {
		allowed, err := awd.PermissionsKeeper.IsAllowedWriter(ctx, signer)
		if err != nil {
			return ctx, err
		}

		if !allowed {
			return ctx, sdkerrors.ErrUnauthorized.Wrap(
				"signer not allowed to write to chain",
			)
		}
	}

	return next(ctx, tx, simulate)
}
