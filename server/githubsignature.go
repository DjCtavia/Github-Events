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
	signatureBytes, err := hex.DecodeString(signature)
	if err != nil {
		return fmt.Errorf("failed to decode signature: %s", err)
	}

	h := hmac.New(sha256.New, []byte(ghSignature.SecretToken))
	_, _ = h.Write(body)

	if !hmac.Equal(signatureBytes, h.Sum(nil)) {
		return fmt.Errorf("signature don't match")
	}
	return nil
}
