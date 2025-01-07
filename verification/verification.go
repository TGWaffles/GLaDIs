package verification

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
)

const SignatureHeader = "x-signature-ed25519"
const TimestampHeader = "x-signature-timestamp"

func VerifySignature(publicKey string, signature string, timestamp string, body string) (bool, error) {
	key, err := hex.DecodeString(publicKey)
	if err != nil {
		return false, err
	}

	if len(key) == 0 {
		return false, fmt.Errorf("No public key provided")
	}

	sig, err := hex.DecodeString(signature)
	if err != nil {
		return false, err
	}

	if len(sig) == 0 {
		return false, fmt.Errorf("No signature provided")
	}

	if len(timestamp) == 0 {
		return false, fmt.Errorf("No timestamp provided")
	}

	return ed25519.Verify(key, []byte(timestamp+body), sig), nil
}

func VerifyHttpRequest(publicKey string, request *http.Request) (bool, error) {
	var signature = request.Header.Get(SignatureHeader)
	var timestamp = request.Header.Get(TimestampHeader)

	body, err := io.ReadAll(request.Body)

	if err != nil {
		return false, err
	}

	return VerifySignature(publicKey, signature, timestamp, string(body))
}
