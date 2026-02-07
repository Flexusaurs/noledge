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

func keccak(data ...[]byte) []byte {
	hasher := sha3.NewLegacyKeccak256()
	for _, b := range data {
		hasher.Write(b)
	}
	return hasher.Sum(nil)
}

func (h *Keccak256Hasher) Hash(left, right []byte) []byte {
	return keccak(left, right)
}

func (h *Keccak256Hasher) HashLeaf(data []byte) []byte {
	return keccak(data)
}

func HashAddress(addr common.Address) []byte {
	return keccak(addr.Bytes())
}
