package api

import (
	"github.com/dgrijalva/jwt-go"
)

func userFromToken(tokenstring string) string {
	token, _ := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		return nil, nil
	})
	if token != nil {
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || claims["sub"] == nil {
			return ""
		}
		userId := claims["sub"].(string)
		return userId
	}
	return ""
}