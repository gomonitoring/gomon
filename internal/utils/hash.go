package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func GetSha256(str string) (string, error) {
	hashedByte := sha256.Sum256([]byte(str))
	return hex.EncodeToString(hashedByte[:]), nil
}
