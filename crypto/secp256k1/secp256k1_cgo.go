// +build libsecp256k1

package secp256k1

import (
	"crypto/elliptic"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/secp256k1/internal/secp256k1"
)

// Sign creates an ECDSA signature on curve Secp256k1, using SHA256 on the msg.
func (privKey PrivKeySecp256k1) Sign(msg []byte) ([]byte, error) {
	rsv, err := secp256k1.Sign(crypto.Sha256(msg), privKey[:])
	if err != nil {
		return nil, err
	}
	// we do not need v  in r||s||v:
	rs := rsv[:len(rsv)-1]
	return rs, nil
}

func (pubKey PubKeySecp256k1) VerifyBytes(msg []byte, sig []byte) bool {
	return secp256k1.VerifySignature(pubKey[:], crypto.Sha256(msg), sig)
}


func (privKey PrivKeySecp256k1) SignRSV(msg []byte) ([]byte, error) {
	return secp256k1.Sign(crypto.Sha256(msg), privKey[:])
}

func (privKey PrivKeySecp256k1) RecoverRSV(msg []byte, sig []byte) (crypto.PubKey, error) {
	pubkey, err := secp256k1.RecoverPubkey(crypto.Sha256(msg), sig)
	if err != nil {
		return nil, err
	}
	x, y := elliptic.Unmarshal(secp256k1.S256(), pubkey)
	var pubkeyBytes PubKeySecp256k1
	copy(pubkeyBytes[:], secp256k1.CompressPubkey(x, y))
	return pubkeyBytes, nil
}



