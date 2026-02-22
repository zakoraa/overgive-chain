package keeper

import (
	"context"
	"overgive-chain/x/delivery/types"
)

func (k Keeper) AppendDelivery(ctx context.Context, delivery types.Delivery) (uint64, error) {
	id, err := k.DeliverySeq.Next(ctx)
	if err != nil {
		return 0, err
	}

	delivery.Id = id

	err = k.Deliveries.Set(ctx, id, delivery)
	if err != nil {
		return 0, err
	}

	return id, nil
}
