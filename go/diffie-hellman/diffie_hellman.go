package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// Diffie-Hellman-Merkle key exchange
// Private keys should be generated randomly.

func PrivateKey(p *big.Int) *big.Int {
	i := big.NewInt(1)
	key, _ := rand.Int(rand.Reader, p)
	for key.Cmp(i) <= 0 || key.Cmp(p) >= 0 {
		// 1 < key < p
		key, _ = rand.Int(rand.Reader, p)
	}
	return key
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	i := big.NewInt(g)
	i = i.Exp(i, private, p)
	return i
}

func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	a := PrivateKey(p)
	k := PublicKey(a, p, g)
	return a, k
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	sk := big.NewInt(0)
	sk.Exp(public2, private1, p)
	return sk
}
