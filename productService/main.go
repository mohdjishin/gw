package main

import (
	"log"
	"net/http"
	"user_service/ecdsaops"
	"user_service/server"
)

func main() {
	publicKey, err := ecdsaops.LoadECDSAPublicKey("/home/muhammedjishinjamaltcp/Key_pairs/ecdsa/public.pem")
	if err != nil {
		log.Fatalf("Failed to load public key: %v", err)
	}

	http.Handle("/products", server.SignatureVerificationMiddleware(publicKey)(http.HandlerFunc(server.GetProducts)))
	log.Println("Product Service is running on port 8082...")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
