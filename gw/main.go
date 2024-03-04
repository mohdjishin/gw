package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"gw/config"
	"gw/crypto"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.RequestID, middleware.Logger, middleware.Recoverer, middleware.Throttle(100), customMiddleware)

	cfg := config.LoadConfig()

	ec := crypto.NewECDSAHelper()
	ec.LoadPrivateKey(cfg.PrivateKeyPath)

	for _, service := range cfg.Services {
		parsedURL, err := url.Parse(service.URL)
		if err != nil {
			log.Fatalf("Failed to parse URL '%s': %v", service.URL, err)
		}
		router.Mount("/"+service.Name, http.StripPrefix("/"+service.Name, serviceProxy(parsedURL, ec)))
	}

	log.Println("API Gateway running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func customMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func serviceProxy(target *url.URL, ec *crypto.ECDSAHelper) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		proxy := httputil.NewSingleHostReverseProxy(target)

		var bodyBytes []byte
		if r.Body != nil {
			bodyBytes, _ = io.ReadAll(r.Body)
		}
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		signature, err := ec.SignDataECDSA(bodyBytes)
		if err != nil {
			log.Printf("Error signing data: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		r.Header.Add("X-Signature", signature)

		originalPath := r.URL.Path
		r.URL.Path = target.Path + strings.TrimPrefix(originalPath, target.Path)

		log.Printf("Forwarding request to: %s", r.URL.String())
		proxy.ServeHTTP(w, r)
	})
}
