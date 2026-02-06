package state

import (
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

func (s *RollupState) GetAccount(addr common.Address) (*Account, error) {
	acc, ok := s.accounts[addr]

	if !ok {
		return &Account{
			Address: addr,
			Balance: 0,
			Nonce:   0,
		}, nil
	}

	return acc, nil

}

func (s *RollupState) SetAccount(addr common.Address, acc *Account) error {
	s.accounts[addr] = acc

	serialized := SerializeAccount(acc)
	leafHash := s.tree.hasher.HashLeaf(serialized)
	key := HashAddress(addr)
}
