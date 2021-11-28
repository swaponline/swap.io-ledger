package auth

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
)

func VerifySign(msg string, sign string) ([]byte, bool) {
	msgBytes, err := hex.DecodeString(msg)
	if err != nil {
		return nil, false
	}
	hash := crypto.Keccak256Hash(msgBytes)

	signature, err := hex.DecodeString(sign)
	if err != nil || len(signature) != 64 {
		return nil, false
	}

	// recoveryId
	signature = append(signature, 0)
	pub, err := crypto.SigToPub(hash.Bytes(), signature)
	if err != nil {
		return nil, false
	}
	publicKeyBytes := crypto.FromECDSAPub(pub)

	return publicKeyBytes,
		crypto.VerifySignature(
			publicKeyBytes,
			hash.Bytes(),
			signature[:64],
		)
}