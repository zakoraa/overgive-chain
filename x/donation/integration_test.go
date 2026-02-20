package donation

import (
	"overgive-chain/x/donation/types"
	"testing"

	overgiveapp "overgive-chain/app"

	donationkeeper "overgive-chain/x/donation/keeper"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// TestIntegration_MsgDonate test the donation transaction  through full app
func TestIntegration_MsgDonate(t *testing.T) {

	// initiate the app with in-memory DB
	app := overgiveapp.Setup(false)

	ctx := app.BaseApp.NewContext(false)

	// Create donor address
	donor := sdk.AccAddress([]byte("donor_______________"))

	// Mint token to module bank
	err := app.BankKeeper.MintCoins(
		ctx,
		types.ModuleName,
		sdk.NewCoins(sdk.NewInt64Coin("uovg", 1000)),
	)
	require.NoError(t, err)

	// Send to donor
	err = app.BankKeeper.SendCoinsFromModuleToAccount(
		ctx,
		types.ModuleName,
		donor,
		sdk.NewCoins(sdk.NewInt64Coin("uovg", 1000)),
	)
	require.NoError(t, err)

	coin := sdk.NewInt64Coin("uovg", 100)

	// Create MsgDonate
	msg := &types.MsgDonate{
		Creator:    donor.String(),
		CampaignId: "campaign-1",
		Amount:     &coin,
		Memmo:      "integration-test",
	}

	// Get msg service
	msgServer := donationkeeper.NewMsgServerImpl(app.DonationKeeper)

	// transaksi execution via keeper (service level)
	res, err := msgServer.Donate(sdk.WrapSDKContext(ctx), msg)
	require.NoError(t, err)
	require.NotNil(t, res)

	// donation stored verification
	donation, err := app.DonationKeeper.Donations.Get(ctx, 1)
	require.NoError(t, err)
	require.Equal(t, "campaign-1", donation.CampaignId)

	//  donor saldo decrement verification
	balance := app.BankKeeper.GetBalance(ctx, donor, "uovg")
	require.Equal(t, int64(900), balance.Amount.Int64())

	//  module saldo account increment verification
	moduleAddr := authtypes.NewModuleAddress(types.ModuleName)
	moduleBalance := app.BankKeeper.GetBalance(ctx, moduleAddr, "uovg")
	require.Equal(t, int64(100), moduleBalance.Amount.Int64())
}
