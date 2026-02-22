package keeper

import (
	"context"
	"errors"
	"overgive-chain/x/donation/types"
)

func (k Keeper) AppendDonation(ctx context.Context, donation types.Donation) (uint64, error) {
	if _, err := k.DonationsByHash.Get(ctx, donation.DonationHash); err == nil {
		return 0, errors.New("duplicate hash")
	}

	id, err := k.DonationSeq.Next(ctx)
	if err != nil {
		return 0, err
	}

	donation.Id = id

	err = k.DonationsByHash.Set(ctx, donation.DonationHash, id)
	if err != nil {
		return 0, err
	}

	err = k.Donations.Set(ctx, id, donation)
	if err != nil {
		return 0, err
	}

	return id, nil
}
