package keeper

import (
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	corestore "cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"

	"overgive-chain/x/donation/types"

	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

type Keeper struct {
	storeService corestore.KVStoreService
	cdc          codec.Codec
	addressCodec address.Codec
	// Address capable of executing a MsgUpdateParams message.
	// Typically, this should be the x/gov module account.
	authority []byte

	Schema collections.Schema
	Params collections.Item[types.Params]

	// BankKeeper for token transfer
	bankKeeper bankkeeper.Keeper

	// Primary storage: Map[uint64] Donation
	Donations collections.Map[uint64, types.Donation]

	// Autoincrement sequence for Donation ID
	DonationSeq collections.Sequence

	// Secondary index: (donor, id)
	DonationsByDonor collections.Map[collections.Pair[string, uint64], bool]
}

func NewKeeper(
	storeService corestore.KVStoreService,
	cdc codec.Codec,
	addressCodec address.Codec,
	authority []byte,
	bankKeeper bankkeeper.Keeper,

) Keeper {
	if _, err := addressCodec.BytesToString(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address %s: %s", authority, err))
	}

	sb := collections.NewSchemaBuilder(storeService)

	// prefix key for every storage
	donationsPrefix := collections.NewPrefix(0)
	donationSeqPrefix := collections.NewPrefix(1)
	donationsByDonorPrefix := collections.NewPrefix(2)

	// initiate primary map
	donations := collections.NewMap(
		sb,
		donationsPrefix,
		"donations",
		collections.Uint64Key,
		codec.CollValue[types.Donation](cdc),
	)

	// Initiate sequence auto increment
	donationSeq := collections.NewSequence(
		sb,
		donationSeqPrefix,
		"donation_seq",
	)

	// Initiate secondary index donor
	donationsByDonor := collections.NewMap(
		sb,
		donationsByDonorPrefix,
		"donations_by_donor",
		collections.PairKeyCodec(collections.StringKey, collections.Uint64Key),
		collections.BoolValue,
	)

	k := Keeper{
		storeService: storeService,
		cdc:          cdc,
		addressCodec: addressCodec,
		authority:    authority,

		Params:           collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		bankKeeper:       bankKeeper,
		Donations:        donations,
		DonationSeq:      donationSeq,
		DonationsByDonor: donationsByDonor,
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}
	k.Schema = schema

	return k
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() []byte {
	return k.authority
}
