package types

// DONTCOVER

import (
	"cosmossdk.io/errors"
)

// x/donation module sentinel errors
var (
	ErrInvalidSigner = errors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")

	// Error when campaign ID not valid
	ErrInvalidCampaignID = errors.Register(ModuleName, 1101, "invalid campaign id")

	// Error when memo too long
	ErrInvalidMemo = errors.Register(ModuleName, 1102, "invalid memo")

	// Error when amount not valid
	ErrInvalidAmount = errors.Register(ModuleName, 1103, "invalid amount")

	// Error when denom not native
	ErrInvalidDenom = errors.Register(ModuleName, 1104, "invalid denom")
)
