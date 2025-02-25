package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/certikfoundation/shentu/v2/app/mva"
)

// BeginBlockForks is intended to be ran in
func BeginBlockForks(ctx sdk.Context, app *ShentuApp) {
	switch ctx.BlockHeight() {
	case mva.UpgradeHeight:
		mva.RunForkLogic(ctx, &app.accountKeeper, app.bankKeeper, &app.stakingKeeper)
	default:
		// do nothing
		return
	}
}
