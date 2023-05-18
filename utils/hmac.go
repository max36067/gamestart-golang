package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func CreateHmac(secretKey string, payload string) string {
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(payload))
	return hex.EncodeToString(h.Sum(nil))
}
