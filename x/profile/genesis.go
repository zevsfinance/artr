package profile

import (
	"github.com/pkg/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initialize default parameters
func InitGenesis(ctx sdk.Context, k Keeper, data GenesisState) {
	k.Logger(ctx).Info("Starting from genesis...")
	k.SetParams(ctx, data.Params)
	for _, record := range data.ProfileRecords {
		acc := k.AccountKeeper.GetAccount(ctx, record.Address)
		record.Profile.CardNumber = k.CardNumberByAccountNumber(ctx, acc.GetAccountNumber())
		if err := k.SetProfile(ctx, record.Address, record.Profile); err != nil {
			panic(errors.Wrapf(err, "invalid profile %s", record.Address))
		}
	}
}

// ExportGenesis writes the current store values
// to a genesis file, which can be imported again
// with InitGenesis
func ExportGenesis(ctx sdk.Context, k Keeper) (data GenesisState) {
	return NewGenesisState(k.GetParams(ctx), k.ExportProfileRecords(ctx))
}
