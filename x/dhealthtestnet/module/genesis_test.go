package dhealthtestnet_test

import (
	"testing"

	keepertest "dhealth-testnet/testutil/keeper"
	"dhealth-testnet/testutil/nullify"
	"dhealth-testnet/x/dhealthtestnet/module"
	"dhealth-testnet/x/dhealthtestnet/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DhealthtestnetKeeper(t)
	dhealthtestnet.InitGenesis(ctx, k, genesisState)
	got := dhealthtestnet.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
