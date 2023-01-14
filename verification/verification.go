package verification

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

func VerifyRequest(publicKey string, request events.APIGatewayProxyRequest) bool {
	key, err := hex.DecodeString(publicKey)
	if err != nil {
		fmt.Println("Failed to decode public key")
		return false
	}

	signatureAsString := request.Headers["X-Signature-Ed25519"]

	if signatureAsString == "" {
		return false
	}

	signature, err := hex.DecodeString(signatureAsString)
	if err != nil {
		return false
	}
	timestamp := request.Headers["X-Signature-Timestamp"]
	body := request.Body

	return ed25519.Verify(key, []byte(timestamp+body), signature)
}
