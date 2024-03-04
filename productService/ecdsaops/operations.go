package ecdsaops

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"encoding/base64"
	"encoding/pem"
	"math/big"
	"os"
)

type ECDSASignature struct {
	R, S *big.Int
}

func VerifySignature(publicKey *ecdsa.PublicKey, data []byte, signatureBase64 string) bool {
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

func LoadECDSAPublicKey(filePath string) (*ecdsa.PublicKey, error) {
	pemBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, err
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	switch pub := pub.(type) {
	case *ecdsa.PublicKey:
		return pub, nil
	default:
		return nil, err // Unsupported key type
	}
}
