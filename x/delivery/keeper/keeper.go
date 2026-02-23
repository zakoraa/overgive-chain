package keeper

import (
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	corestore "cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"

	deliverytypes "overgive-chain/x/delivery/types"
)

type Keeper struct {
	storeService corestore.KVStoreService
	cdc          codec.Codec
	addressCodec address.Codec
	// Address capable of executing a MsgUpdateParams message.
	// Typically, this should be the x/gov module account.
	authority []byte

	Schema collections.Schema
	Params collections.Item[deliverytypes.Params]


	Deliveries       collections.Map[uint64, deliverytypes.Delivery]
	DeliverySeq      collections.Sequence
	DeliveriesByHash collections.Map[string, uint64]
}

func NewKeeper(
	storeService corestore.KVStoreService,
	cdc codec.Codec,
	addressCodec address.Codec,
	authority []byte,

) Keeper {
	if _, err := addressCodec.BytesToString(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address %s: %s", authority, err))
	}

	sb := collections.NewSchemaBuilder(storeService)

	k := Keeper{
		storeService: storeService,
		cdc:          cdc,
		addressCodec: addressCodec,
		authority:    authority,

		Params: collections.NewItem(sb, deliverytypes.ParamsKey, "params", codec.CollValue[deliverytypes.Params](cdc)),

		Deliveries: collections.NewMap(
			sb,
			deliverytypes.DeliveryKeyPrefix,
			"deliveries",
			collections.Uint64Key,
			codec.CollValue[deliverytypes.Delivery](cdc),
		),

		DeliverySeq: collections.NewSequence(
			sb,
			deliverytypes.DeliverySeqKey,
			"delivery_seq",
		),

		DeliveriesByHash: collections.NewMap(
			sb,
			deliverytypes.DeliveryHashKeyPrefix,
			"deliveries_by_hash",
			collections.StringKey,
			collections.Uint64Value,
		),
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
