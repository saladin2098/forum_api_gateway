package token

import (
	"errors"
	// "log"
	// "time"

	// pb "github.com/saladin2098/forum_api_gateway/genproto"
	"github.com/golang-jwt/jwt"
)

const (
	signingKey = "Secret key for forum auth service"
)





func ValidateToken(token string) (bool, error) {
	_, err := ExtractClaim(token)
	if err != nil {
		return false, err
	}

	return true,nil
}

func ExtractClaim(tokenStr string) (jwt.MapClaims, error) {
	var (
		token *jwt.Token
		err   error
	)

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	}
	token, err = jwt.Parse(tokenStr, keyFunc)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
