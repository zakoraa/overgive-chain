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

func (m msgServer) UpdateAllowedWriter(
	goCtx context.Context,
	msg *types.MsgUpdateAllowedWriter,
) (*types.MsgUpdateAllowedWriterResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	authorityAddr, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrap("invalid authority address")
	}

	if !authorityAddr.Equals(m.authority) {
		return nil, sdkerrors.ErrUnauthorized.Wrap("unauthorized")
	}

	addr, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress
	}

	if msg.Add {
		m.SetAllowedWriter(ctx, addr)
	} else {
		m.RemoveAllowedWriter(ctx, addr)
	}

	return &types.MsgUpdateAllowedWriterResponse{}, nil
}
