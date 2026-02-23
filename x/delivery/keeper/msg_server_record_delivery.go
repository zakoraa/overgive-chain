package keeper

import (
	"context"
	"encoding/hex"
	"fmt"

	"overgive-chain/x/delivery/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RecordDelivery(
	goCtx context.Context,
	msg *types.MsgRecordDelivery,
) (*types.MsgRecordDeliveryResponse, error) {

	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	isAllowed, err := k.permissionsKeeper.IsAllowedWriter(ctx, msg.Creator)
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

	if msg.Title == "" {
		return nil, errorsmod.Wrap(
			sdkerrors.ErrInvalidRequest,
			"title required",
		)
	}

	if len(msg.CampaignId) > 64 {
		return nil, errorsmod.Wrap(
			sdkerrors.ErrInvalidRequest,
			"campaign_id too long (max 64 chars)",
		)
	}

	if len(msg.Title) > 128 {
		return nil, errorsmod.Wrap(
			sdkerrors.ErrInvalidRequest,
			"title too long (max 128 chars)",
		)
	}

	decodedHash, err := hex.DecodeString(msg.DeliveryHash)
	if err != nil || len(decodedHash) != 32 {
		return nil, errorsmod.Wrap(
			sdkerrors.ErrInvalidRequest,
			"delivery_hash must be valid 32-byte hex",
		)
	}

	if _, err := hex.DecodeString(msg.DeliveryHash); err != nil {
		return nil, errorsmod.Wrap(
			sdkerrors.ErrInvalidRequest,
			"delivery_hash must be valid hex",
		)
	}

	if msg.NoteHash != "" {
		if len(msg.NoteHash) != 64 {
			return nil, errorsmod.Wrap(
				sdkerrors.ErrInvalidRequest,
				"note_hash must be 64 characters",
			)
		}
		if _, err := hex.DecodeString(msg.NoteHash); err != nil {
			return nil, errorsmod.Wrap(
				sdkerrors.ErrInvalidRequest,
				"note_hash must be valid hex",
			)
		}
	}

	if _, err := k.DeliveriesByHash.Get(ctx, msg.DeliveryHash); err == nil {
		return nil, errorsmod.Wrap(
			sdkerrors.ErrInvalidRequest,
			"delivery_hash already exists",
		)
	}

	delivery := types.Delivery{
		CampaignId:   msg.CampaignId,
		Title:        msg.Title,
		NoteHash:     msg.NoteHash,
		DeliveryHash: msg.DeliveryHash,
		Creator:      msg.Creator,
		RecordedAt:   ctx.BlockTime().Unix(),
	}

	id, err := k.AppendDelivery(goCtx, delivery)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeDeliveryRecorded,
			sdk.NewAttribute(types.AttributeKeyDeliveryID, fmt.Sprintf("%d", id)),
			sdk.NewAttribute(types.AttributeKeyCampaignID, msg.CampaignId),
			sdk.NewAttribute(types.AttributeKeyDeliveryHash, msg.DeliveryHash),
			sdk.NewAttribute(types.AttributeKeyCreator, msg.Creator),
		),
	)

	return &types.MsgRecordDeliveryResponse{
		Id: id,
	}, nil
}
