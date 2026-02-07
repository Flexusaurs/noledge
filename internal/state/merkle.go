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

	smt.initZeroNodes()

	return smt
}

func (smt *SparseMerkleTree) initZeroNodes() {
	smt.zeroNodes = make([][]byte, smt.depth+1)

	zeroLeaf := smt.hasher.HashLeaf(make([]byte, 32))
	smt.zeroNodes[0] = zeroLeaf

	for i := 1; i <= smt.depth; i++ {
		smt.zeroNodes[i] = smt.hasher.Hash(
			smt.zeroNodes[i-1], smt.zeroNodes[i-1],
		)
	}
}

/* func generateZeroNodes(depth int, hasher MerkleHasher) [][]byte {
	zeros := make([][]byte, depth+1)
	zeros[0] = hasher.HashLeaf([]byte{0})

	for i := 1; i <= depth; i++ {
		zeros[i] = hasher.Hash(zeros[i-1], zeros[i-1])
	}
	return zeros
} */

func nodeKey(level int, path string) string {
	return fmt.Sprintf("%d,%s", level, path)
}

func getBit(key []byte, pos int) byte {
	byteIndex := pos / 8
	bitIndex := 7 - (pos % 8)
	return (key[byteIndex] >> bitIndex) & 1
}

func (smt *SparseMerkleTree) Update(key []byte, leafHash []byte) {

	path := make([]byte, smt.depth)
	currentHash := leafHash

	//build bit path
	for i := 0; i < smt.depth; i++ {
		path[i] = getBit(key, i)
	}

	//store leaf
	smt.nodes[nodeKey(0, string(path))] = currentHash

	for level := 0; level < smt.depth; level++ {
		parentPath := path[:smt.depth-level-1]
		siblingPath := make([]byte, len(path))
		copy(siblingPath, path)
		siblingPath[smt.depth-level-1] ^= 1 //XOR to one

		siblingHash, ok := smt.nodes[nodeKey(level, string(siblingPath))]
		if !ok {
			siblingHash = smt.zeroNodes[level]
		}

		var left, right []byte
		if path[smt.depth-level-1] == 0 {
			left = currentHash
			right = siblingHash
		} else {
			left = siblingHash
			right = currentHash
		}

		parentHash := smt.hasher.Hash(left, right)
		smt.nodes[nodeKey(level+1, string(parentPath))] = parentHash

		currentHash = parentHash

	}
}

func (smt *SparseMerkleTree) Root() []byte {
	root, ok := smt.nodes[nodeKey(smt.depth, "")]
	if !ok {
		return smt.zeroNodes[smt.depth]
	}
	return root
}
