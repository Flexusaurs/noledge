package state

import (
	"bytes"
	"testing"
)

func TestEmptyTreeRootDete(t *testing.T) {
	hasher := &SHA256Hasher{}
	smt1 := NewSparseMerkleTree(16, hasher)
	smt2 := NewSparseMerkleTree(16, hasher)

	if !bytes.Equal(smt1.Root(), smt2.Root()) {
		t.Fatal("empty roots MUST be identical")
	}
}

func TestSingleUpdateChangesRoot(t *testing.T) {
	hasher := &SHA256Hasher{}
	smt := NewSparseMerkleTree(16, hasher)

	initialRoot := smt.Root()
	leaf := hasher.HashLeaf([]byte("account1"))
	smt.Update(42, leaf)

	newRoot := smt.Root()

	if bytes.Equal(initialRoot, newRoot) {
		t.Fatal("root MUST change after update")
	}
}

func TestDeterministicUpdate(t *testing.T) {
	hasher := &SHA256Hasher{}
	smt1 := NewSparseMerkleTree(16, hasher)
	smt2 := NewSparseMerkleTree(16, hasher)

	leaf := hasher.HashLeaf([]byte("account1"))

	smt1.Update(42, leaf)
	smt2.Update(42, leaf)

	if !bytes.Equal(smt1.Root(), smt2.Root()) {
		t.Fatal("same updates MUST produce identical results")
	}
}

func TestMultipleUpdatesOrderMatters(t *testing.T) {
	hasher := &SHA256Hasher{}
	smt := NewSparseMerkleTree(16, hasher)

	leaf1 := hasher.HashLeaf([]byte("account1"))
	leaf2 := hasher.HashLeaf([]byte("account2"))

	smt.Update(10, leaf1)
	root1 := smt.Root()

	smt.Update(11, leaf2)
	root2 := smt.Root()

	if bytes.Equal(root1, root2) {
		t.Fatal("root MUST change after second update")
	}
}
