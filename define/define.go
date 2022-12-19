package define

import "github.com/golang-jwt/jwt/v4"

type M map[string]interface{}

type UserClaim struct {
	Id       uint   `json:"id"`
	Identity string `json:"identity"`
	Name     string `json:"name"`
	jwt.RegisteredClaims
}

var JwtKet = "myiot"
