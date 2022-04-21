package keeper_test

import (
	"testing"

	testkeeper "github.com/klpanda/klpanda/testutil/keeper"
	"github.com/klpanda/klpanda/x/klpanda/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.KlpandaKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
