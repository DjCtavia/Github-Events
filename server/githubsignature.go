package server

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

const DEFAULTSIGNATURE = "secret-token"

type GithubSignature struct {
	SecretToken string
}

func (ghSignature *GithubSignature) Verify(body []byte, signature string) error {
	expectedSignature := "sha256=" + hex.EncodeToString(hmac.New(sha256.New, []byte(ghSignature.SecretToken)).Sum(body))
	if !hmac.Equal([]byte(expectedSignature), []byte(signature)) {
		return fmt.Errorf("Signature don't match")
	}
	return nil
}
