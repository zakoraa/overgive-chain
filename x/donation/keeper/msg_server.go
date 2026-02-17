package keeper

import (
	"context"
	"fmt"
	"overgive-chain/x/donation/types"

	"cosmossdk.io/collections"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// Donate handle donation transaction
func (m msgServer) Donate(
	ctx context.Context,
	msg *types.MsgDonate,
) (*types.MsgDonationResponse, error) {
	// Convert to SDK context
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Validate bech32 address
	donorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	// Validate campaign_id not empty
	if len(msg.CampaignId) == 0 {
		return nil, types.ErrInvalidCampaignID
	}

	// Validate len campaign_id max 64
	if len(msg.CampaignId) > 64 {
		return nil, types.ErrInvalidCampaignID
	}

	// Validate memo max 256 char
	if len(msg.Memmo) > 256 {
		return nil, types.ErrInvalidMemo
	}

	// validate amount not nil
	if msg.Amount == nil {
		return nil, types.ErrInvalidAmount
	}

	// Validate denom must be uovg native
	if msg.Amount.Denom != "uovg" {
		return nil, types.ErrInvalidDenom
	}

	// Validate amount > 0
	coin := sdk.NewCoin(msg.Amount.Denom, msg.Amount.Amount)
	if !coin.IsPositive() {
		return nil, types.ErrInvalidAmount
	}

	// Trnsfer token from user to module account donation
	err = m.bankKeeper.SendCoinsFromAccountToModule(
		sdkCtx,
		donorAddr,
		types.ModuleName,
		sdk.NewCoins(coin),
	)

	if err != nil {
		return nil, err
	}

	// Generate ID auto increment
	id, err := m.DonationSeq.Next(sdkCtx)

	if err != nil {
		return nil, err
	}

	// create Donation object
	donation := types.Donation{
		Id:         id,
		Donor:      msg.Creator,
		CampaignId: msg.CampaignId,
		Amount:     msg.Amount,
		CraetedAt:  sdkCtx.BlockTime().Unix(),
		Memo:       msg.Memmo,
	}

	// Store to primary storage
	err = m.Donations.Set(sdkCtx, id, donation)
	if err != nil {
		return nil, err
	}

	// Store secondary index (donor, id)
	pairKey := collections.Join(msg.Creator, id)
	err = m.DonationsByDonor.Set(sdkCtx, pairKey, true)
	if err != nil {
		return nil, err
	}

	// Emit event for indexer / explorer
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			"donation_created",
			sdk.NewAttribute("id", fmt.Sprintf("%d", id)),
			sdk.NewAttribute("donor", msg.Creator),
			sdk.NewAttribute("campaign_id", msg.CampaignId),
			sdk.NewAttribute("amount", msg.Amount.Amount.String()),
			sdk.NewAttribute("denom", msg.Amount.Denom),
		),
	)

	// Return response with donation ID
	return &types.MsgDonationResponse{
		DonationId: id,
	}, nil
}
