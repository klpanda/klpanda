package klpanda_test

import (
	"testing"

	keepertest "github.com/klpanda/klpanda/testutil/keeper"
	"github.com/klpanda/klpanda/testutil/nullify"
	"github.com/klpanda/klpanda/x/klpanda"
	"github.com/klpanda/klpanda/x/klpanda/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.KlpandaKeeper(t)
	klpanda.InitGenesis(ctx, *k, genesisState)
	got := klpanda.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
