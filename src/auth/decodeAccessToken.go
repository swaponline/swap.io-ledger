package auth

import (
	"encoding/hex"
	"strconv"
	"strings"
	"swap.io-ledger/src/database"
	"time"
)

func (a *Auth) DecodeAccessToken(tokenString string) (*database.User, bool) {
	tokenData := strings.Split(tokenString, ".")
	if len(tokenData) != 3 {
		return nil, false
	}
	tokenLifeTimeHex, signHex, pubKeyHex := tokenData[0], tokenData[1], tokenData[2]

	validTokenLifeTimeBytes, err := hex.DecodeString(tokenLifeTimeHex)
	if err != nil {
		return nil, false
	}
	validTokenLifeTime, err := strconv.ParseInt(
		string(validTokenLifeTimeBytes),
		0,
		64,
	)
	if err != nil || validTokenLifeTime < time.Now().Unix() {
		return nil, false
	}

	ok := VerifySign(tokenLifeTimeHex, signHex, pubKeyHex)
	if !ok {
		return nil, false
	}

	user, err := a.usersManager.GetUserByPubKey(
		pubKeyHex,
	)

	if err != nil {
		return nil, false
	}

	return user, true
}
