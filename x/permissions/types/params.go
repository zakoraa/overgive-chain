package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewParams creates a new Params instance.
func NewParams() Params {
	return Params{}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return Params{
		AllowedWriters: []string{},
	}
}

// Validate validates the set of params.
func (p Params) Validate() error {
	for _, addr := range p.AllowedWriters {
		if _, err := sdk.AccAddressFromBech32(addr); err != nil {
			return fmt.Errorf("invalid allowed_writers address: %s", addr)
		}
	}
	return nil
}
