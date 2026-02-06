package state

import (
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/crypto/sha3"
)

type MerkleHasher interface {
	Hash(left, right []byte) []byte
	HashLeaf(data []byte) []byte
}

type Keccak256Hasher struct{}

func (h *Keccak256Hasher) Hash(left, right []byte) []byte {
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(left)
	hasher.Write(right)

	return hasher.Sum(nil)
}

func (h *Keccak256Hasher) HashLeaf(data []byte) []byte {
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(data)
	return hasher.Sum(nil)
}

func HashAddress(addr common.Address) []byte {
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(addr.Bytes())
	return hasher.Sum(nil)
}
