package keeper

import (
	"context"
	"overgive-chain/x/donation/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// var _ types.QueryServer = queryServer{}

// // NewQueryServerImpl returns an implementation of the QueryServer interface
// // for the provided Keeper.
// func NewQueryServerImpl(k Keeper) types.QueryServer {
// 	return queryServer{k}
// }

// type queryServer struct {
// 	k Keeper
// }

func (k Keeper) Donation(
	goCtx context.Context,
	req *types.QueryGetDonationRequest,
) (*types.QueryGetDonationResponse, error) {

	donation, err := k.Donations.Get(goCtx, req.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetDonationResponse{
		Donation: &donation,
	}, nil
}

func (k Keeper) DonationByHash(
	goCtx context.Context,
	req *types.QueryDonationByHashRequest,
) (*types.QueryDonationByHashResponse, error) {

	id, err := k.DonationsByHash.Get(goCtx, req.DonationHash)
	if err != nil {
		return nil, status.Error(codes.NotFound, "not found")
	}

	donation, err := k.Donations.Get(goCtx, id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryDonationByHashResponse{
		Donation: &donation,
	}, nil
}
