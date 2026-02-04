package state

import "fmt"

type SparseMerkleTree struct {
	depth     int
	nodes     map[string][]byte
	zeroNodes [][]byte
	hasher    MerkleHasher
}

func NewSparseMerkleTree(depth int, hasher MerkleHasher) *SparseMerkleTree {
	smt := &SparseMerkleTree{
		depth:  depth,
		nodes:  make(map[string][]byte),
		hasher: hasher,
	}

	smt.zeroNodes = generateZeroNodes(depth, hasher)

	return smt
}

func generateZeroNodes(depth int, hasher MerkleHasher) [][]byte {
	zeros := make([][]byte, depth+1)
	zeros[0] = hasher.HashLeaf([]byte{0})

	for i := 1; i <= depth; i++ {
		zeros[i] = hasher.Hash(zeros[i-1], zeros[i-1])
	}
	return zeros
}

func nodeKey(level int, index uint32) string {
	return fmt.Sprintf("%d,%d", level, index)
}

func (smt *SparseMerkleTree) Update(index uint32, leafHash []byte) {
	currentHash := leafHash
	idx := index

	//store leaf
	smt.nodes[nodeKey(0, idx)] = leafHash

	for level := 0; level < smt.depth; level++ {
		siblingIndex := idx ^ 1
		siblingKey := nodeKey(level, siblingIndex)
		siblingHash, ok := smt.nodes[siblingKey]
		if !ok {
			siblingHash = smt.zeroNodes[level]
		}

		var left, right []byte
		if idx%2 == 0 {
			left = currentHash
			right = siblingHash
		} else {
			left = siblingHash
			right = currentHash
		}

		parentHash := smt.hasher.Hash(left, right)
		idx = idx / 2

		parentKey := nodeKey(level+1, idx)
		smt.nodes[parentKey] = parentHash
		currentHash = parentHash
	}
}

func (smt *SparseMerkleTree) Root() []byte {
	rootKey := nodeKey(smt.depth, 0)
	if root, ok := smt.nodes[rootKey]; ok {
		return root
	}
	return smt.zeroNodes[smt.depth]
}
