package keeper_test

import (
	"testing"

	"overgive-chain/testutil/sample"
	donationkeeper "overgive-chain/x/donation/keeper"
	"overgive-chain/x/donation/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stretchr/testify/require"
)

func TestMsgDonate_Success(t *testing.T) {

	k, ctx, mockBank := setupKeeper(t)
	msgServer := donationkeeper.NewMsgServerImpl(*k)

	donor := sample.AccAddress()

	// seed initial balance
	mockBank.balances[donor] =
		sdk.NewCoins(sdk.NewInt64Coin("uovg", 1000))

	coin := sdk.NewInt64Coin("uovg", 100)
	msg := &types.MsgDonate{
		Creator:    donor,
		CampaignId: "campaign-1",
		Amount:     &coin,
		Memmo:      "support",
	}

	res, err := msgServer.Donate(sdk.WrapSDKContext(ctx), msg)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, uint64(1), res.DonationId)

	donation, err := k.Donations.Get(ctx, 1)
	require.NoError(t, err)
	require.Equal(t, "campaign-1", donation.CampaignId)

	// check balance decreased
	finalBalance := mockBank.balances[donor].AmountOf("uovg")
	require.Equal(t, int64(900), finalBalance.Int64())
}

func TestMsgDonate_InvalidDenom(t *testing.T) {

	k, ctx, _ := setupKeeper(t)
	msgServer := donationkeeper.NewMsgServerImpl(*k)

	donor := sample.AccAddress()

	coin := sdk.NewInt64Coin("uovg", 100)

	msg := &types.MsgDonate{
		Creator:    donor,
		CampaignId: "campaign-1",
		Amount:     &coin,
		Memmo:      "",
	}

	_, err := msgServer.Donate(sdk.WrapSDKContext(ctx), msg)
	require.Error(t, err)
}

func TestMsgDonate_InvalidAmount(t *testing.T) {

	k, ctx, _ := setupKeeper(t)
	msgServer := donationkeeper.NewMsgServerImpl(*k)

	donor := sample.AccAddress()

	coin := sdk.NewInt64Coin("uovg", 100)

	msg := &types.MsgDonate{
		Creator:    donor,
		CampaignId: "campaign-1",
		Amount:     &coin,
		Memmo:      "",
	}

	_, err := msgServer.Donate(sdk.WrapSDKContext(ctx), msg)
	require.Error(t, err)
}

func TestMsgDonate_EmptyCampaign(t *testing.T) {

	k, ctx, _ := setupKeeper(t)
	msgServer := donationkeeper.NewMsgServerImpl(*k)

	donor := sample.AccAddress()

	coin := sdk.NewInt64Coin("uovg", 100)

	msg := &types.MsgDonate{
		Creator:    donor,
		CampaignId: "",
		Amount:     &coin,
		Memmo:      "",
	}

	_, err := msgServer.Donate(sdk.WrapSDKContext(ctx), msg)
	require.Error(t, err)
}
