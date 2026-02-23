package keeper

import (
	"context"
	"crypto/sha256"
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
	creatorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress
	}

	allowed, err := k.permissionsKeeper.IsAllowedWriter(ctx, creatorAddr)
	if err != nil {
		return nil, err
	}

	if !allowed {
		return nil, sdkerrors.ErrUnauthorized.Wrap("not allowed writer")
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

	sum := sha256.Sum256([]byte(msg.PaymentReferenceId))
	donationHash := hex.EncodeToString(sum[:])

	if msg.PaymentReferenceId == "" {
		return nil, errorsmod.Wrap(
			sdkerrors.ErrInvalidRequest,
			"payment_reference_id required",
		)
	}

	amountDec, err := math.LegacyNewDecFromStr(msg.Amount)
	if err != nil {
		return nil, errorsmod.Wrap(
			sdkerrors.ErrInvalidRequest,
			"invalid amount",
		)
	}

	if !amountDec.IsPositive() {
		return nil, errorsmod.Wrap(
			sdkerrors.ErrInvalidRequest,
			"amount must be positive",
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
		DonationHash:       donationHash,
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

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeDonationRecorded,
			sdk.NewAttribute(types.AttributeKeyDonationID, fmt.Sprintf("%d", id)),
			sdk.NewAttribute(types.AttributeKeyCampaignID, msg.CampaignId),
			sdk.NewAttribute(types.AttributeKeyDonationHash, donationHash),
			sdk.NewAttribute(types.AttributeKeyCreator, msg.Creator),
		),
	)

	return &types.MsgRecordDonationResponse{
		Id: id,
	}, nil

}
