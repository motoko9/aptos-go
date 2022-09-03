package crypto

type Signer interface {
	Sign(data []byte) ([]byte, error)
	PublicKey() PublicKey
}
