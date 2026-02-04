package state

import (
	"crypto/sha256"
)

type MerkleHasher interface {
	Hash(left, right []byte) []byte
	HashLeaf(data []byte) []byte
}

type SHA256Hasher struct{}

func (h *SHA256Hasher) Hash(left, right []byte) []byte {
	concat := append(left, right...)
	hash := sha256.Sum256(concat)
	return hash[:]
} //array to slice conversion

func (h *SHA256Hasher) HashLeaf(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}
