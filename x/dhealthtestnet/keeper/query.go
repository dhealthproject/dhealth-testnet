package keeper

import (
	"github.com/dhealthproject/dhealth-testnet/x/dhealthtestnet/types"
)

var _ types.QueryServer = Keeper{}
