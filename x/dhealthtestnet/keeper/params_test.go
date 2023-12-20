package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "dhealth-testnet/testutil/keeper"
	"dhealth-testnet/x/dhealthtestnet/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.DhealthtestnetKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
