package keeper

import (
	"context"
	"errors"
	"fmt"

	"overgive-chain/x/permissions/types"

	"cosmossdk.io/collections"
	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateAllowed(ctx context.Context, msg *types.MsgCreateAllowed) (*types.MsgCreateAllowedResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	// Check if the value already exists
	ok, err := k.Allowed.Has(ctx, msg.Index)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	} else if ok {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var allowed = types.Allowed{
		Creator: msg.Creator,
		Index:   msg.Index,
		Address: msg.Address,
	}

	if err := k.Allowed.Set(ctx, allowed.Index, allowed); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	return &types.MsgCreateAllowedResponse{}, nil
}

func (k msgServer) UpdateAllowed(ctx context.Context, msg *types.MsgUpdateAllowed) (*types.MsgUpdateAllowedResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.Allowed.Get(ctx, msg.Index)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
		}

		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var allowed = types.Allowed{
		Creator: msg.Creator,
		Index:   msg.Index,
		Address: msg.Address,
	}

	if err := k.Allowed.Set(ctx, allowed.Index, allowed); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to update allowed")
	}

	return &types.MsgUpdateAllowedResponse{}, nil
}

func (k msgServer) DeleteAllowed(ctx context.Context, msg *types.MsgDeleteAllowed) (*types.MsgDeleteAllowedResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.Allowed.Get(ctx, msg.Index)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
		}

		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	if err := k.Allowed.Remove(ctx, msg.Index); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to remove allowed")
	}

	return &types.MsgDeleteAllowedResponse{}, nil
}
