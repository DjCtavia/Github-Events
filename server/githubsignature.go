package server

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

type GithubSignature struct {
	SecretToken string
}

func (ghSignature *GithubSignature) Verify(body []byte, signature string) error {
	expectedSignature := "sha1=" + hex.EncodeToString(hmac.New(sha1.New, []byte(ghSignature.SecretToken)).Sum(body))
	if !hmac.Equal([]byte(expectedSignature), []byte(signature)) {
		return fmt.Errorf("Signature don't match")
	}
	return nil
}
