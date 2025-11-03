package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func ComputeSecretHash(username, clientID, clientSecret string) string {
	mac := hmac.New(sha256.New, []byte(clientSecret))
	mac.Write([]byte(username + clientID))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}
