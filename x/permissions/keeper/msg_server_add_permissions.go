package keeper

import (
	"context"

	"overgive-chain/x/permissions/types"

	errorsmod "cosmossdk.io/errors"
)

func (k msgServer) AddPermissions(ctx context.Context, msg *types.MsgAddPermissions) (*types.MsgAddPermissionsResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	// TODO: Handle the message

	return &types.MsgAddPermissionsResponse{}, nil
}
