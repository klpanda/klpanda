package keeper

import (
	"github.com/klpanda/klpanda/x/klpanda/types"
)

var _ types.QueryServer = Keeper{}
