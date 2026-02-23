package permissions

import (
	"cosmossdk.io/core/address"
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/core/store"
	"cosmossdk.io/depinject"
	"cosmossdk.io/depinject/appconfig"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"overgive-chain/x/permissions/keeper"
	"overgive-chain/x/permissions/types"

)

var _ depinject.OnePerModuleType = AppModule{}

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (AppModule) IsOnePerModuleType() {}

func init() {
	appconfig.Register(
		&types.Module{},
		appconfig.Provide(ProvideModule),
	)
}

type ModuleInputs struct {
	depinject.In

	Config       *types.Module
	StoreService store.KVStoreService
	Cdc          codec.Codec
	AddressCodec address.Codec

	AuthKeeper types.AuthKeeper
	BankKeeper types.BankKeeper
}

type ModuleOutputs struct {
	depinject.Out

	PermissionsKeeper keeper.Keeper
	Module            appmodule.AppModule
}

func ProvideModule(in ModuleInputs) ModuleOutputs {
	// default to governance authority if not provided
	// authority := authtypes.NewModuleAddress(types.GovModuleName)
	// if in.Config.Authority != "" {
	// 	authority = authtypes.NewModuleAddressOrBech32Address(in.Config.Authority)
	// }
	 authorityStr := "overgive1gncccpp3mgxqag79jsc6yksmzqyq8g4ylktn0g" // address overgive-admin

    authority, err := sdk.AccAddressFromBech32(authorityStr)
    if err != nil {
        panic(err)
    }
	k := keeper.NewKeeper(
		in.StoreService,
		in.Cdc,
		in.AddressCodec,
		authority,
	)
	m := NewAppModule(in.Cdc, k, in.AuthKeeper, in.BankKeeper)

	return ModuleOutputs{PermissionsKeeper: k, Module: m}
}
