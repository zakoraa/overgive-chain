package keeper

import (
	"context"
	"overgive-chain/x/delivery/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Delivery(
	goCtx context.Context,
	req *types.QueryGetDeliveryRequest,
) (*types.QueryGetDeliveryResponse, error) {

	delivery, err := k.Deliveries.Get(goCtx, req.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetDeliveryResponse{
		Delivery: &delivery,
	}, nil
}

func (k Keeper) DeliveryByHash(
	goCtx context.Context,
	req *types.QueryDeliveryByHashRequest,
) (*types.QueryDeliveryByHashResponse, error) {

	id, err := k.DeliveriesByHash.Get(goCtx, req.DeliveryHash)
	if err != nil {
		return nil, status.Error(codes.NotFound, "not found")
	}

	delivery, err := k.Deliveries.Get(goCtx, id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryDeliveryByHashResponse{
		Delivery: &delivery,
	}, nil

}
