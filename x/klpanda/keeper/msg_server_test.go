package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/klpanda/klpanda/testutil/keeper"
	"github.com/klpanda/klpanda/x/klpanda/keeper"
	"github.com/klpanda/klpanda/x/klpanda/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.KlpandaKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
