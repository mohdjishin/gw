package server

import (
	"bytes"
	"crypto/ecdsa"
	"io"
	"log"
	"net/http"
	"user_service/ecdsaops"
)

func SignatureVerificationMiddleware(publicKey *ecdsa.PublicKey) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			signature := r.Header.Get("X-Signature")
			if signature == "" {
				http.Error(w, "Signature required", http.StatusUnauthorized)
				return
			}

			body, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Could not read request body", http.StatusBadRequest)
				return
			}
			r.Body = io.NopCloser(bytes.NewReader(body))

			if !ecdsaops.VerifySignature(publicKey, body, signature) {
				http.Error(w, "Invalid signature", http.StatusForbidden)
				return
			}
			log.Println("Signature verified")
			next.ServeHTTP(w, r)
		})
	}
}
