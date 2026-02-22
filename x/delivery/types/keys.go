package types

import "cosmossdk.io/collections"

const (
	// ModuleName defines the module name
	ModuleName = "delivery"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// GovModuleName duplicates the gov module's name to avoid a dependency with x/gov.
	// It should be synced with the gov module's name if it is ever changed.
	// See: https://github.com/cosmos/cosmos-sdk/blob/v0.52.0-beta.2/x/gov/types/keys.go#L9
	GovModuleName = "gov"
)

var (
	ParamsKey             = collections.NewPrefix(0)
	DeliveryKeyPrefix     = collections.NewPrefix(1)
	DeliverySeqKey        = collections.NewPrefix(2)
	DeliveryHashKeyPrefix = collections.NewPrefix(3)
)

// ParamsKey is the prefix to retrieve all Params
