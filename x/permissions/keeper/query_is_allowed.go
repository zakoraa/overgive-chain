package keeper

import (
	"context"

	"overgive-chain/x/permissions/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) IsAllowed(ctx context.Context, req *types.QueryIsAllowedRequest) (*types.QueryIsAllowedResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// TODO: Process the query

	return &types.QueryIsAllowedResponse{}, nil
}
