package dollar_test

import (
	"testing"

	keepertest "dollar/testutil/keeper"
	"dollar/testutil/nullify"
	dollar "dollar/x/dollar/module"
	"dollar/x/dollar/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DollarKeeper(t)
	dollar.InitGenesis(ctx, k, genesisState)
	got := dollar.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
