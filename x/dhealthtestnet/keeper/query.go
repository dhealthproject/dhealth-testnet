package keeper

import (
	"github.com/dhealthproject/dhealth-testnet/v2/x/dhealthtestnet/types"
)

var _ types.QueryServer = Keeper{}
