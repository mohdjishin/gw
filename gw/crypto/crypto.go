package crypto

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"encoding/base64"
	"encoding/pem"
	"math/big"
	"os"

	"github.com/gofiber/fiber/v2/log"
)

type ECDSAHelper struct {
	PrivateKey *ecdsa.PrivateKey
}

type ECDSASignature struct {
	R, S *big.Int
}

func NewECDSAHelper() *ECDSAHelper {
	return &ECDSAHelper{}
}

func (h *ECDSAHelper) VerifySignatureECDSA(publicKey *ecdsa.PublicKey, data []byte, signatureBase64 string) bool {
	signatureBytes, err := base64.StdEncoding.DecodeString(signatureBase64)
	if err != nil {
		return false
	}

	var sig ECDSASignature
	_, err = asn1.Unmarshal(signatureBytes, &sig)
	if err != nil {
		return false
	}

	hasher := sha256.New()
	hasher.Write(data)
	hashed := hasher.Sum(nil)

	return ecdsa.Verify(publicKey, hashed, sig.R, sig.S)
}

func (h *ECDSAHelper) LoadPrivateKey(pemFile string) {
	pemBytes, err := os.ReadFile(pemFile)
	if err != nil {
		log.Fatalf("Failed to read PEM file: %v", err)
		return
	}
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		log.Fatalf("Failed to decode PEM block")
		return
	}
	priv, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
		return
	}
	switch priv := priv.(type) {
	case *ecdsa.PrivateKey:
		h.PrivateKey = priv
	default:
		log.Fatalf("Unsupported key type")
	}
}

func (h *ECDSAHelper) SignDataECDSA(data []byte) (string, error) {
	hasher := sha256.New()
	hasher.Write(data)
	hashed := hasher.Sum(nil)

	r, s, err := ecdsa.Sign(rand.Reader, h.PrivateKey, hashed)
	if err != nil {
		return "", err
	}

	signature, err := asn1.Marshal(ECDSASignature{r, s})
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(signature), nil
}
