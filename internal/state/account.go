package state

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Account struct {
	Address common.Address //20 byte, ETH style length
	Balance *big.Int
	Nonce   uint64
}
