package keeper

import (
	"context"

	"overgive-chain/x/donation/types"

	errorsmod "cosmossdk.io/errors"
)

func (k msgServer) CreateDonation(ctx context.Context, msg *types.MsgCreateDonation) (*types.MsgCreateDonationResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	// TODO: Handle the message

	return &types.MsgCreateDonationResponse{}, nil
}
