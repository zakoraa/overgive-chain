package keeper

import (
	"context"
	"overgive-chain/x/donation/types"

	"cosmossdk.io/collections"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/types/query"
)

var _ types.QueryServer = queryServer{}

// NewQueryServerImpl returns an implementation of the QueryServer interface
// for the provided Keeper.
func NewQueryServerImpl(k Keeper) types.QueryServer {
	return queryServer{k}
}

type queryServer struct {
	k Keeper
}

// Donation: get 1 donation by ID
func (q queryServer) Donation(
	ctx context.Context,
	req *types.QueryGetDonationRequest,
) (*types.QueryGetDonationResponse, error) {
	// Validate req is nil
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Get donation from colletions
	donation, err := q.k.Donations.Get(sdkCtx, req.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "donation not found")
	}

	return &types.QueryGetDonationResponse{
		Donation: &donation,
	}, nil
}

// DonationAll return all donations with pagination
func (q queryServer) DonationAll(
	ctx context.Context,
	req *types.QueryGetAllDonationRequest,
) (*types.QueryGetAllDonationResponse, error) {
	// Validate if req is nil
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	var donations []*types.Donation

	// get key-value store
	store := runtime.KVStoreAdapter(q.k.storeService.OpenKVStore(sdkCtx))

	pageRes, err := query.Paginate(
		store,
		req.Pagination,
		func(key []byte, value []byte) error {

			var donation *types.Donation
			if err := q.k.cdc.Unmarshal(value, donation); err != nil {
				return err
			}

			donations = append(donations, donation)
			return nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryGetAllDonationResponse{
		Donations:  donations,
		Pagination: pageRes,
	}, nil
}

// DonationByDonor get donation by donor address
func (q queryServer) DonationByDonor(
	ctx context.Context,
	req *types.QueryDonationByDonorRequest,
) (*types.QueryDonationByDonorResponse, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	var donations []*types.Donation

	// Create prefix range for certain donor
	rng := collections.NewPrefixedPairRange[string, uint64](req.Donor)

	iter, err := q.k.DonationsByDonor.Iterate(sdkCtx, rng)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {

		pair, err := iter.Key()
		if err != nil {
			return nil, err
		}

		// Get donation ID from pair key
		id := pair.K2()

		donation, err := q.k.Donations.Get(sdkCtx, id)
		if err != nil {
			continue
		}
		d := donation

		donations = append(donations, &d)
	}

	return &types.QueryDonationByDonorResponse{
		Donation: donations,
	}, nil
}
