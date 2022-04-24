package keeper

import (
	"github.com/klpanda/klpanda/x/dex/types"
)

var _ types.QueryServer = Keeper{}
