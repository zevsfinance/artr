package schedule

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// BeginBlocker check for infraction evidence or downtime of validators
// on every begin block
func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k Keeper) {
	k.PerfomSchedule(ctx, uint64(ctx.BlockHeight()))
}

// EndBlocker called every block, process inflation, update validator set.
func EndBlocker(ctx sdk.Context, k Keeper) {
}
