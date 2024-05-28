package utils

import (
	"math/rand"

	"github.com/nihal-ramaswamy/GoVid/internal/constants"
)

func NewUUID(length int) string {

	allowedChars := constants.GetRuneUuidCharacters()
	lengthAllowedCharacters := len(allowedChars)

	uuidString := make([]byte, length)

	for i := range uuidString {
		uuidString[i] = byte(allowedChars[rand.Intn(lengthAllowedCharacters)])
	}

	return string(uuidString)
}
