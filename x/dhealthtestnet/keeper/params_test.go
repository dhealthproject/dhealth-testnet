package keeper_test

import (
	"testing"

	testkeeper "github.com/dhealthproject/dhealth-testnet/testutil/keeper"
	"github.com/dhealthproject/dhealth-testnet/x/dhealthtestnet/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.DhealthtestnetKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
