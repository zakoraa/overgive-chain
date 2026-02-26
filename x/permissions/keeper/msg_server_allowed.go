package keeper

import (
	"bytes"
	"context"

	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"overgive-chain/x/permissions/types"
)

func (k msgServer) CreateAllowed(
	ctx context.Context,
	msg *types.MsgCreateAllowed,
) (*types.MsgCreateAllowedResponse, error) {

	// Validate creator address format
	creatorBytes, err := k.addressCodec.StringToBytes(msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, err.Error())
	}

	// Authority check
	if !bytes.Equal(creatorBytes, k.authority) {
		return nil, types.ErrUnauthorized
	}

	// Validate target address
	if _, err := k.addressCodec.StringToBytes(msg.Address); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid target address")
	}

	// Enforce index = address (whitelist design)
	if msg.Index != msg.Address {
		return nil, types.ErrInvalidIndex
	}

	// Check if already exists
	exists, err := k.Allowed.Has(ctx, msg.Index)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, types.ErrAlreadyExists
	}

	allowed := types.Allowed{
		Creator: msg.Creator,
		Index:   msg.Index,
		Address: msg.Address,
	}

	if err := k.Allowed.Set(ctx, msg.Index, allowed); err != nil {
		return nil, err
	}

	return &types.MsgCreateAllowedResponse{}, nil
}

func (k msgServer) UpdateAllowed(
	ctx context.Context,
	msg *types.MsgUpdateAllowed,
) (*types.MsgUpdateAllowedResponse, error) {

	return nil, errorsmod.Wrap(
		sdkerrors.ErrInvalidRequest,
		"update not supported for whitelist module",
	)
}

func (k msgServer) DeleteAllowed(
	ctx context.Context,
	msg *types.MsgDeleteAllowed,
) (*types.MsgDeleteAllowedResponse, error) {

	// Validate creator address format
	creatorBytes, err := k.addressCodec.StringToBytes(msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, err.Error())
	}

	// Authority check
	if !bytes.Equal(creatorBytes, k.authority) {
		return nil, types.ErrUnauthorized
	}

	// Check if exists
	exists, err := k.Allowed.Has(ctx, msg.Index)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, types.ErrNotFound
	}

	if err := k.Allowed.Remove(ctx, msg.Index); err != nil {
		return nil, err
	}

	return &types.MsgDeleteAllowedResponse{}, nil
}
