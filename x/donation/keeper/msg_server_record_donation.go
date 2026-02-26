package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"

	"overgive-chain/x/donation/types"
	permissionstypes "overgive-chain/x/permissions/types"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RecordDonation(
	goCtx context.Context,
	msg *types.MsgRecordDonation,
) (*types.MsgRecordDonationResponse, error) {

	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(
			sdkerrors.ErrInvalidAddress,
			"invalid creator address",
		)
	}

	isAllowed, err := k.permissionsKeeper.HasAllowed(goCtx, msg.Creator)
	if err != nil {
		return nil, err
	}
	if !isAllowed {
		return nil, errorsmod.Wrap(
			permissionstypes.ErrUnauthorized,
			"not allowed writer",
		)
	}

	if msg.CampaignId == "" {
		return nil, errorsmod.Wrap(
			sdkerrors.ErrInvalidRequest,
			"campaign_id required",
		)
	}

	if len(msg.CampaignId) > 64 {
		return nil, errorsmod.Wrap(
			sdkerrors.ErrInvalidRequest,
			"campaign_id too long (max 64 chars)",
		)
	}

	if msg.PaymentReferenceId == "" {
		return nil, errorsmod.Wrap(
			sdkerrors.ErrInvalidRequest,
			"payment_reference_id required",
		)
	}

	if msg.Currency == "" {
		return nil, errorsmod.Wrap(
			sdkerrors.ErrInvalidRequest,
			"currency required",
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

	sum := sha256.Sum256([]byte(msg.PaymentReferenceId))
	donationHash := hex.EncodeToString(sum[:])

	donationHash = strings.ToLower(donationHash)

	exists, err := k.DonationsByHash.Has(goCtx, donationHash)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errorsmod.Wrap(
			sdkerrors.ErrInvalidRequest,
			"donation_hash already exists",
		)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

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