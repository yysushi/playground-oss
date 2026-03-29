package main

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/oauth2-proxy/mockoidc"
)

func main() {
	key, _ := rsa.GenerateKey(rand.Reader, 2048)
	m, _ := mockoidc.NewServer(key)

	ln, _ := net.Listen("tcp", "127.0.0.1:0")
	m.Start(ln, nil)

	port := strings.Split(ln.Addr().String(), ":")[1]

	// Proxy to rewrite URLs
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		m.Server.Handler.ServeHTTP(w, r)
	})

	cfg := m.Config()
	fmt.Printf("CLIENT_ID=%s\n", cfg.ClientID)
	fmt.Printf("CLIENT_SECRET=%s\n", cfg.ClientSecret)
	fmt.Printf("CLIENT_SECRET_BASE64=%s\n", base64.StdEncoding.EncodeToString([]byte(cfg.ClientSecret)))
	fmt.Printf("ISSUER=http://oauth.default.svc.cluster.local/oidc\n")

	log.Println("OAuth Server on :8080, internal on :" + port)
	http.ListenAndServe(":8080", nil)
}
