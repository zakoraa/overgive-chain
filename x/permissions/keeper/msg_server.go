package keeper

import (
	"context"
	"overgive-chain/x/permissions/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (m msgServer) AddWriter(
	ctx context.Context,
	msg *types.MsgAddWriter,
) (*types.MsgAddWriterResponse, error) {

	authorityAddr, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrap("invalid authority address")
	}

	if !authorityAddr.Equals(m.authority) {
		return nil, sdkerrors.ErrUnauthorized.Wrap("unauthorized")
	}
	params, err := m.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	for _, w := range params.AllowedWriters {
		if w == msg.Writer {
			return &types.MsgAddWriterResponse{}, nil
		}
	}

	params.AllowedWriters = append(params.AllowedWriters, msg.Writer)

	if err := m.Params.Set(ctx, params); err != nil {
		return nil, err
	}

	return &types.MsgAddWriterResponse{}, nil
}
