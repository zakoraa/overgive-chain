package keeper

import (
	"context"

	"overgive-chain/x/permissions/types"

	errorsmod "cosmossdk.io/errors"
)

func (k msgServer) UpdatePermissions(ctx context.Context, msg *types.MsgUpdatePermissions) (*types.MsgUpdatePermissionsResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}
	params, err := k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	if !isAllowedWriter(msg.Creator, params.AllowedWriters) {
		return nil, types.ErrUnauthorizedWriter
	}

	return &types.MsgUpdatePermissionsResponse{}, nil
}

func isAllowedWriter(addr string, allowed []string) bool {
	for _, a := range allowed {
		if a == addr {
			return true
		}
	}
	return false
}
