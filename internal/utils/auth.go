package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/nihal-ramaswamy/GoVid/internal/constants"
	"github.com/nihal-ramaswamy/GoVid/internal/dto"
)

func GenerateToken(user *dto.User) (string, error) {
	secret := GetDotEnvVariable("SECRET_KEY")

	var signingKey = []byte(secret)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(constants.TOKEN_EXPIRY_TIME).Unix()

	return token.SignedString(signingKey)
}
