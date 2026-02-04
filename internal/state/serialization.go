package state

import "encoding/binary"

func SerializeAccount(acc *Account) []byte {
	buf := make([]byte, 36)
	copy(buf[:20], acc.Address) //first 20 bytes reserved for address
	binary.BigEndian.PutUint64(buf[20:28], acc.Balance)
	binary.BigEndian.PutUint64(buf[28:36], acc.Nonce) //all
	return buf
}
