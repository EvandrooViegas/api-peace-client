package services

import (
	"fmt"
	"time"

	"github.com/EvandrooViegas/types"
	"github.com/EvandrooViegas/utils"
	"github.com/golang-jwt/jwt/v4"
)

func createAuthJWT(id string) (string, error) {
	secret, err := utils.LoadEnvVariable("AUTH_SIGN_KEY")
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &types.UserTokenClaims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TOKEN_AND_COOKIE_EXPIRATION)),
		},
	})
	if err != nil {
		return "", err
	}
	return t.SignedString([]byte(secret))
}

func ReadPlayerToken(tokenString string) (string, error) {
	secret, err := utils.LoadEnvVariable("AUTH_SIGN_KEY")
	if err != nil {
		return "", nil
	}

	token, err := jwt.ParseWithClaims(tokenString, &types.UserTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*types.UserTokenClaims)
		if !token.Valid {
			return "", fmt.Errorf("Token is not valid")
			} 
			if !ok {
				return "", fmt.Errorf("Couldn't parse the claims")
			}
			return claims.ID, nil
}
 