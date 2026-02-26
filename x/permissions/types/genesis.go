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
	allowedIndexMap := make(map[string]struct{})

	for _, elem := range gs.AllowedMap {
		index := fmt.Sprint(elem.Index)
		if _, ok := allowedIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for allowed")
		}
		allowedIndexMap[index] = struct{}{}
	}

	return gs.Params.Validate()
}
