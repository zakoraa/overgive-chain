package app

import (
	"io"

	"cosmossdk.io/log"

	dbm "github.com/cosmos/cosmos-db"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
)

func Setup(loadLatest bool) *App {
	db := dbm.NewMemDB()

	logger := log.NewNopLogger()

	var traceStore io.Writer

	appOpts := servertypes.AppOptions(nil)

	app := New(
		logger,
		db,
		traceStore,
		loadLatest,
		appOpts,
	)

	return app
}
