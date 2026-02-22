package keeper

import (
	"context"
	"overgive-chain/x/delivery/types"

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

func (q queryServer) Deliveries(
	goCtx context.Context,
	req *types.QueryDeliveriesRequest,
) (*types.QueryDeliveriesResponse, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	const maxLimit = 100
	if req.Pagination != nil && req.Pagination.Limit > maxLimit {
		req.Pagination.Limit = maxLimit
	}

	deliveries, pageRes, err := query.CollectionPaginate(
		ctx,
		q.k.Deliveries,
		req.Pagination,
		func(_ uint64, value types.Delivery) (*types.Delivery, error) {
			v := value
			return &v, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryDeliveriesResponse{
		Deliveries: deliveries,
		Pagination: pageRes,
	}, nil
}

func (q queryServer) Delivery(
	ctx context.Context,
	req *types.QueryGetDeliveryRequest,
) (*types.QueryGetDeliveryResponse, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	delivery, err := q.k.Deliveries.Get(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetDeliveryResponse{
		Delivery: &delivery,
	}, nil
}

func (q queryServer) DeliveryByHash(
	goCtx context.Context,
	req *types.QueryDeliveryByHashRequest,
) (*types.QueryDeliveryByHashResponse, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	id, err := q.k.DeliveriesByHash.Get(goCtx, req.DeliveryHash)
	if err != nil {
		return nil, status.Error(codes.NotFound, "not found")
	}

	delivery, err := q.k.Deliveries.Get(goCtx, id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryDeliveryByHashResponse{
		Delivery: &delivery,
	}, nil

}
