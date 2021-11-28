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
	if len(tokenData) != 2 {
		return nil, false
	}

	validTokenTimeBytes, err  := hex.DecodeString(tokenData[0])
	if err != nil {
		return nil, false
	}
	validTokenTime, err := strconv.ParseInt(
		string(validTokenTimeBytes),
		0,
		64,
	)
	if err != nil || validTokenTime < time.Now().Unix() {
		return nil, false
	}

	pubKey, ok := VerifySign(tokenData[0], tokenData[1])
	if !ok {
		return nil, false
	}

	user, err := a.usersManager.GetUserByPubKey(string(pubKey))
	if err != nil {
		return nil, false
	}

	return user, true
}
