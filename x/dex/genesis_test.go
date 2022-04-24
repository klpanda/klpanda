package dex_test

import (
	"testing"

	keepertest "github.com/klpanda/klpanda/testutil/keeper"
	"github.com/klpanda/klpanda/testutil/nullify"
	"github.com/klpanda/klpanda/x/dex"
	"github.com/klpanda/klpanda/x/dex/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		BuyOrderBookList: []types.BuyOrderBook{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DexKeeper(t)
	dex.InitGenesis(ctx, *k, genesisState)
	got := dex.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	require.ElementsMatch(t, genesisState.BuyOrderBookList, got.BuyOrderBookList)
	// this line is used by starport scaffolding # genesis/test/assert
}
