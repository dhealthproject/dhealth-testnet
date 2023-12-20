package types

const (
	// ModuleName defines the module name
	ModuleName = "dhealthtestnet"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_dhealthtestnet"
)

var (
	ParamsKey = []byte("p_dhealthtestnet")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
