package keeper

import (
	"github.com/gogo/protobuf/grpc"

	sdk "github.com/cosmos/cosmos-sdk/types"
	v043 "github.com/cosmos/cosmos-sdk/x/auth/legacy/v043"
	sdktypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/certikfoundation/shentu/v2/x/auth/types"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper      types.AccountKeeper
	queryServer grpc.Server
}

// NewMigrator returns a new Migrator.
func NewMigrator(keeper Keeper, queryServer grpc.Server) Migrator {
	return Migrator{keeper: keeper.ak, queryServer: queryServer}
}

// Migrate1to2 migrates from version 1 to 2.
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	var iterErr error

	m.keeper.IterateAccounts(ctx, func(account sdktypes.AccountI) (stop bool) {
		wb, err := v043.MigrateAccount(ctx, account, m.queryServer)
		if err != nil {
			iterErr = err
			return true
		}

		if wb == nil {
			return false
		}

		m.keeper.SetAccount(ctx, wb)
		return false
	})

	return iterErr
}
