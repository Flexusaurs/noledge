package state

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type RollupState struct {
	tree     *SparseMerkleTree
	accounts map[common.Address]*Account
}

func NewRollupState(tree *SparseMerkleTree) *RollupState {
	return &RollupState{
		tree:     tree,
		accounts: make(map[common.Address]*Account),
	}
}

func (s *RollupState) GetAccount(addr common.Address) *Account {
	acc, ok := s.accounts[addr]

	if !ok {
		return &Account{
			Address: addr,
			Balance: big.NewInt(0),
			Nonce:   0,
		}
	}

	return acc

}

func (s *RollupState) SetAccount(addr common.Address, acc *Account) error {
	s.accounts[addr] = acc

	serialized := SerializeAccount(acc)
	leafHash := s.tree.hasher.HashLeaf(serialized)
	key := HashAddress(addr)

	s.tree.Update(key, leafHash)
	return nil
}
