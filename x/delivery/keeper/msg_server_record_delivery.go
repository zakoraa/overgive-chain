package keeper

import (
	"context"

	"overgive-chain/x/delivery/types"

	errorsmod "cosmossdk.io/errors"
)

func (k msgServer) RecordDelivery(ctx context.Context, msg *types.MsgRecordDelivery) (*types.MsgRecordDeliveryResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	// TODO: Handle the message

	return &types.MsgRecordDeliveryResponse{}, nil
}
