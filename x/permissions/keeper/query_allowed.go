package keeper

import (
	"context"
	"errors"

	"overgive-chain/x/permissions/types"

	"cosmossdk.io/collections"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) ListAllowed(ctx context.Context, req *types.QueryAllAllowedRequest) (*types.QueryAllAllowedResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	alloweds, pageRes, err := query.CollectionPaginate(
		ctx,
		q.k.Allowed,
		req.Pagination,
		func(_ string, value types.Allowed) (types.Allowed, error) {
			return value, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAllowedResponse{Allowed: alloweds, Pagination: pageRes}, nil
}

func (q queryServer) GetAllowed(ctx context.Context, req *types.QueryGetAllowedRequest) (*types.QueryGetAllowedResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, err := q.k.Allowed.Get(ctx, req.Address)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetAllowedResponse{Allowed: val}, nil
}
