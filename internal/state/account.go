package state

type Account struct {
	Address []byte //20 byte, ETH style length
	Balance uint64
	Nonce   uint64
}
