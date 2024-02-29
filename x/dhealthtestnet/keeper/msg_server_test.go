package keeper_test

import (
	"context"
	"testing"

	keepertest "dhealth-testnet/testutil/keeper"
	"dhealth-testnet/x/dhealthtestnet/keeper"
	"dhealth-testnet/x/dhealthtestnet/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.DhealthtestnetKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestMsgServer(t *testing.T) {
	ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
}
