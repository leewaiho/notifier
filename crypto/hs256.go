package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
)

// HmacSHA256
func HS256(message []byte, key []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(message)
	return h.Sum(nil)
}
