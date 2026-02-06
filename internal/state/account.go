package state

import "github.com/ethereum/go-ethereum/common"

type Account struct {
	Address common.Address //20 byte, ETH style length
	Balance uint64
	Nonce   uint64
}
