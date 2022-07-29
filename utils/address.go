package utils

import (
	"encoding/hex"
	"golang.org/x/crypto/sha3"
)

func Pubkey2Address(pk PublicKey) string {
	hash := sha3.New256()
	hash.Write(pk[:])
	hash.Write([]byte{0})
	return "0x" + hex.EncodeToString(hash.Sum(nil))
}
