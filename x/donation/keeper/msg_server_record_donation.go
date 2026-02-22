package keeper

import (
	"context"
	"encoding/hex"
	"fmt"
	"overgive-chain/x/donation/types"

	"cosmossdk.io/math"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RecordDonation(goCtx context.Context, msg *types.MsgRecordDonation) (*types.MsgRecordDonationResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAllowed, err := k.IsAllowedWriter(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}
	if !isAllowed {
		return nil, errorsmod.Wrap(
			sdkerrors.ErrUnauthorized,
			"not allowed writer",
		)
	}

	if msg.CampaignId == "" {
		return nil, errorsmod.Wrap(
			sdkerrors.ErrInvalidRequest,
			"campaign_id required",
		)
	}

	if _, err := math.LegacyNewDecFromStr(msg.Amount); err != nil {
		return nil, errorsmod.Wrap(
			sdkerrors.ErrInvalidRequest,
			"invalid amount",
		)
	}

	if len(msg.CampaignId) > 64 {
		return nil, errorsmod.Wrap(
			sdkerrors.ErrInvalidRequest,
			"campaign_id too long (max 64 chars)",
		)
	}

	if len(msg.DonationHash) != 64 {
		return nil, errorsmod.Wrap(
			sdkerrors.ErrInvalidRequest,
			"donation_hash must be 64 characters",
		)
	}

	if _, err := hex.DecodeString(msg.DonationHash); err != nil {
		return nil, errorsmod.Wrap(
			sdkerrors.ErrInvalidRequest,
			"donation_hash must be valid hex",
		)
	}

	if msg.Currency == "" {
		return nil, errorsmod.Wrap(
			sdkerrors.ErrInvalidRequest,
			"currency required",
		)
	}

	donation := types.Donation{
		CampaignId:         msg.CampaignId,
		Amount:             msg.Amount,
		Currency:           msg.Currency,
		PaymentReferenceId: msg.PaymentReferenceId,
		DonationHash:       msg.DonationHash,
		MetadataHash:       msg.MetadataHash,
		Creator:            msg.Creator,
		RecordedAt:         ctx.BlockTime().Unix(),
	}

	if _, err := k.DonationsByHash.Get(ctx, donation.DonationHash); err == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "donation_hash already exists")
	}

	id, err := k.AppendDonation(goCtx, donation)
	if err != nil {
		return nil, err
	}

	sdk.NewEvent(
		types.EventTypeDonationRecorded,
		sdk.NewAttribute(types.AttributeKeyDonationID, fmt.Sprintf("%d", id)),
		sdk.NewAttribute(types.AttributeKeyCampaignID, msg.CampaignId),
		sdk.NewAttribute(types.AttributeKeyDonationHash, msg.DonationHash),
		sdk.NewAttribute(types.AttributeKeyCreator, msg.Creator),
	)

	return &types.MsgRecordDonationResponse{
		Id: id,
	}, nil

}
