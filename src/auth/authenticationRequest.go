package auth

import (
	"errors"
	"net/http"
)

func (a *Auth) AuthenticationRequest(request *http.Request) (int, error) {
	tokenInfo := request.URL.Query()["token"]
	if len(tokenInfo) == 0 {
		return -1, errors.New("not exist token")
	}
	user, err := a.DecodeAccessToken(tokenInfo[0])
	if !err {
		return -1, errors.New("not valid token")
	}

	return user.Id, nil
}
