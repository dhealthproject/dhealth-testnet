package keeper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/dhealthproject/dhealth-testnet/testutil/keeper"
	"github.com/dhealthproject/dhealth-testnet/x/dhealthtestnet/keeper"
	"github.com/dhealthproject/dhealth-testnet/x/dhealthtestnet/types"
)

func setupMsgServer(t testing.TB) (keeper.Keeper, types.MsgServer, context.Context) {
	k, ctx := keepertest.DhealthtestnetKeeper(t)
	return k, keeper.NewMsgServerImpl(k), ctx
}

func TestMsgServer(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)
}
