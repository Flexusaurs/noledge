package state

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Transaction struct {
	Source      common.Address
	Destination common.Address
	Amount      *big.Int
	nonce       uint64
	sig         []byte //placeholder
}
