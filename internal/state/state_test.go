package state

import (
	"bytes"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func newTestState() *RollupState {
	hasher := &Keccak256Hasher{}
	tree := NewSparseMerkleTree(256, hasher)
	return NewRollupState(tree)
}

func TestNewAccountInit(t *testing.T) {
	state := newTestState()
	addr := common.HexToAddress("0x1111111111111111111111111111111111111111")
	acc := state.GetAccount(addr)

	if acc.Nonce != 0 {
		t.Fatalf("expected nonce 0, got %d", acc.Nonce)
	}

	if acc.Balance.Cmp(big.NewInt(0)) != 0 {
		t.Fatalf("expected balance 0")
	}
}

func TestSetAccountUpdatesRoot(t *testing.T) {
	state := newTestState()

	addr := common.HexToAddress("0x2222222222222222222222222222222222222222")

	acc := &Account{
		Address: addr,
		Balance: big.NewInt(1000),
		Nonce:   1,
	}

	initialRoot := state.tree.Root()
	state.SetAccount(addr, acc)
	newRoot := state.tree.Root()
	fmt.Println(string(newRoot))

	if bytes.Equal(initialRoot, newRoot) {
		t.Fatal("state root did not change after account update")
	}
}

func TestDeterministicRoot(t *testing.T) {
	addr := common.HexToAddress("0x3333333333333333333333333333333333333333")

	buildState := func() []byte {
		state := newTestState()

		acc := &Account{
			Address: addr,
			Balance: big.NewInt(500),
			Nonce:   2,
		}

		state.SetAccount(addr, acc)
		return state.tree.Root()
	}

	_aroot := buildState()
	fmt.Println("_aroot: ", string(_aroot))
	_broot := buildState()
	fmt.Println("_broot: ", string(_broot))

	if !bytes.Equal(_aroot, _broot) {
		t.Fatal("state roots are not deterministic")
	}
}

func TestMultipleAccounts(t *testing.T) {
	state := newTestState()

	_addra := common.HexToAddress("0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	_addrb := common.HexToAddress("0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")

	fmt.Println("_addra: ", _addra)
	fmt.Println("_addrb: ", _addrb)

	_acca := &Account{
		Address: _addra,
		Balance: big.NewInt(1000),
		Nonce:   0,
	}

	_accb := &Account{
		Address: _addrb,
		Balance: big.NewInt(2000),
		Nonce:   0,
	}

	state.SetAccount(_addra, _acca)
	rootAfterA := state.tree.Root()
	fmt.Println("rootAfterA: ", string(rootAfterA))
	state.SetAccount(_addrb, _accb)
	rootAfterB := state.tree.Root()
	fmt.Println("rootAfterB: ", string(rootAfterB))

	if bytes.Equal(rootAfterA, rootAfterB) {
		t.Fatal("root MUST change after adding second account")
	}
}
