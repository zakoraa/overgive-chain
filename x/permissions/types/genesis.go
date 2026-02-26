package types

import "fmt"

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:     DefaultParams(),
		AllowedMap: []Allowed{}}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	allowedAddressMap := make(map[string]struct{})

	for _, elem := range gs.AllowedMap {
		address := fmt.Sprint(elem.Address)
		if _, ok := allowedAddressMap[address]; ok {
			return fmt.Errorf("duplicated Address for allowed")
		}
		allowedAddressMap[address] = struct{}{}
	}

	return gs.Params.Validate()
}
