package types

import fmt "fmt"

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
		if addr == "" {
			return fmt.Errorf("allowed_writers contains empty address")
		}
	}
	return nil
}
