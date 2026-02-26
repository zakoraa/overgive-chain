package keeper

import (
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	corestore "cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"

	donationtypes "overgive-chain/x/donation/types"
	permissionstypes "overgive-chain/x/permissions/types"
)

type Keeper struct {
	storeService corestore.KVStoreService
	cdc          codec.Codec
	addressCodec address.Codec
	// Address capable of executing a MsgUpdateParams message.
	// Typically, this should be the x/gov module account.
	authority []byte

	permissionsKeeper permissionstypes.PermissionsKeeper

	Schema collections.Schema
	Params collections.Item[donationtypes.Params]

	Donations       collections.Map[uint64, donationtypes.Donation]
	DonationSeq     collections.Sequence
	DonationsByHash collections.Map[string, uint64]
}

func NewKeeper(
	storeService corestore.KVStoreService,
	cdc codec.Codec,
	addressCodec address.Codec,
	authority []byte,
	permissionsKeeper permissionstypes.PermissionsKeeper,
) Keeper {
	if _, err := addressCodec.BytesToString(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address %s: %s", authority, err))
	}

	sb := collections.NewSchemaBuilder(storeService)

	k := Keeper{
		storeService:      storeService,
		cdc:               cdc,
		addressCodec:      addressCodec,
		authority:         authority,
		permissionsKeeper: permissionsKeeper,

		Params: collections.NewItem(sb, donationtypes.ParamsKey, "params", codec.CollValue[donationtypes.Params](cdc)),

		Donations: collections.NewMap(
			sb,
			donationtypes.DonationKeyPrefix,
			"donations",
			collections.Uint64Key,
			codec.CollValue[donationtypes.Donation](cdc),
		),

		DonationSeq: collections.NewSequence(
			sb,
			donationtypes.DonationSeqKey,
			"donation_seq",
		),

		DonationsByHash: collections.NewMap(
			sb,
			donationtypes.DonationHashKeyPrefix,
			"donations_by_hash",
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
