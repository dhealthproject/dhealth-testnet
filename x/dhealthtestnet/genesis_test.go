package dhealthtestnet_test

import (
	"testing"

	keepertest "github.com/dhealthproject/dhealth-testnet/v2/testutil/keeper"
	"github.com/dhealthproject/dhealth-testnet/v2/testutil/nullify"
	"github.com/dhealthproject/dhealth-testnet/v2/x/dhealthtestnet"
	"github.com/dhealthproject/dhealth-testnet/v2/x/dhealthtestnet/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DhealthtestnetKeeper(t)
	dhealthtestnet.InitGenesis(ctx, *k, genesisState)
	got := dhealthtestnet.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
