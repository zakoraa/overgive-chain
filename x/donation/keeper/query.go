package keeper

import (
	"context"
	"overgive-chain/x/donation/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (q queryServer) Params(
	goCtx context.Context,
	req *types.QueryParamsRequest,
) (*types.QueryParamsResponse, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	params, err := q.k.Params.Get(ctx)
	if err != nil {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryParamsResponse{
		Params: params,
	}, nil
}

func (q queryServer) Donations(
	goCtx context.Context,
	req *types.QueryDonationsRequest,
) (*types.QueryDonationsResponse, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	const maxLimit = 100
	if req.Pagination != nil && req.Pagination.Limit > maxLimit {
		req.Pagination.Limit = maxLimit
	}

	donations, pageRes, err := query.CollectionPaginate(
		ctx,
		q.k.Donations,
		req.Pagination,
		func(_ uint64, value types.Donation) (*types.Donation, error) {
			v := value
			return &v, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryDonationsResponse{
		Donations:  donations,
		Pagination: pageRes,
	}, nil
}

func (q queryServer) Donation(
	goCtx context.Context,
	req *types.QueryGetDonationRequest,
) (*types.QueryGetDonationResponse, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	donation, err := q.k.Donations.Get(goCtx, req.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetDonationResponse{
		Donation: &donation,
	}, nil
}

func (q queryServer) DonationByHash(
	goCtx context.Context,
	req *types.QueryDonationByHashRequest,
) (*types.QueryDonationByHashResponse, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	id, err := q.k.DonationsByHash.Get(goCtx, req.DonationHash)
	if err != nil {
		return nil, status.Error(codes.NotFound, "not found")
	}

	donation, err := q.k.Donations.Get(goCtx, id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryDonationByHashResponse{
		Donation: &donation,
	}, nil
}
