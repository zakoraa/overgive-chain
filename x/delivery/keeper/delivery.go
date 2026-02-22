package keeper

import (
	"context"
	"errors"
	"overgive-chain/x/delivery/types"
)

func (k Keeper) AppendDelivery(ctx context.Context, delivery types.Delivery) (uint64, error) {
	if _, err := k.DeliveriesByHash.Get(ctx, delivery.DeliveryHash); err == nil {
		return 0, errors.New("duplicate hash")
	}

	id, err := k.DeliverySeq.Next(ctx)
	if err != nil {
		return 0, err
	}

	delivery.Id = id

	err = k.Deliveries.Set(ctx, id, delivery)
	if err != nil {
		return 0, err
	}

	err = k.DeliveriesByHash.Set(ctx, delivery.DeliveryHash, id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
