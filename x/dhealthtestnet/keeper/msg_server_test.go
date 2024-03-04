package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/dhealthproject/dhealth-testnet/v2/testutil/keeper"
	"github.com/dhealthproject/dhealth-testnet/v2/x/dhealthtestnet/keeper"
	"github.com/dhealthproject/dhealth-testnet/v2/x/dhealthtestnet/types"
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
