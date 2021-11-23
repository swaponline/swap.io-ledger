package auth

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
)

func VerifySign(msg string, sign string, pub string) bool {
	publicKeyBytes, err := hex.DecodeString(pub)
	if err != nil {
		return false
	}

	msgBytes, err := hex.DecodeString(msg)
	if err != nil {
		return false
	}
	hash := crypto.Keccak256Hash(msgBytes)

	signature, err := hex.DecodeString(sign)
	if err != nil {
		return false
	}

	return crypto.VerifySignature(publicKeyBytes, hash.Bytes(), signature)
}