package main

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/oauth2-proxy/mockoidc"
)

// ProxyHandler wraps the mockoidc server and rewrites URLs from 127.0.0.1 to a domain name
type ProxyHandler struct {
	mockServer *mockoidc.MockOIDC
	domain     string
	port       string
}

func (p *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Log all requests
	fmt.Printf("[REQUEST] %s %s\n", r.Method, r.URL.Path)

	// Handle /oidc/token requests specially to support both Basic Auth and POST body auth
	if r.URL.Path == "/oidc/token" && r.Method == "POST" {
		fmt.Printf("[DEBUG] Token endpoint request received\n")
		// Check if using Basic Auth
		username, password, ok := r.BasicAuth()
		fmt.Printf("[DEBUG] Basic Auth: ok=%v, username=%s\n", ok, username)
		if ok && username != "" {
			// Read the existing body
			bodyBytes, err := io.ReadAll(r.Body)
			fmt.Printf("[DEBUG] Read body: %d bytes, err=%v\n", len(bodyBytes), err)
			if err == nil {
				// Parse the form data
				formData, err := url.ParseQuery(string(bodyBytes))
				fmt.Printf("[DEBUG] Parsed form: %v, err=%v\n", formData, err)
				if err == nil {
					// Add client_id and client_secret to form if not already present
					if formData.Get("client_id") == "" {
						formData.Set("client_id", username)
						fmt.Printf("[DEBUG] Added client_id=%s\n", username)
					}
					if formData.Get("client_secret") == "" {
						formData.Set("client_secret", password)
						fmt.Printf("[DEBUG] Added client_secret\n")
					}
					// Re-encode the form data
					newBody := formData.Encode()
					fmt.Printf("[DEBUG] New body: %s\n", newBody)
					r.Body = io.NopCloser(bytes.NewReader([]byte(newBody)))
					r.ContentLength = int64(len(newBody))
				}
			}
		}
	}

	// Create a new response recorder to capture the mock server's response
	rec := &responseRecorder{
		ResponseWriter: w,
		domain:         p.domain,
		port:           p.port,
	}

	// Forward to the mock server
	p.mockServer.Server.Handler.ServeHTTP(rec, r)
}

type responseRecorder struct {
	http.ResponseWriter
	domain string
	port   string
}

func (r *responseRecorder) Write(b []byte) (int, error) {
	// Rewrite 127.0.0.1 to domain name in JSON responses
	if strings.Contains(r.Header().Get("Content-Type"), "application/json") {
		modified := strings.ReplaceAll(string(b), "http://127.0.0.1:"+r.port, "http://"+r.domain)
		modified = strings.ReplaceAll(modified, "http://[::]:"+r.port, "http://"+r.domain)
		return r.ResponseWriter.Write([]byte(modified))
	}
	return r.ResponseWriter.Write(b)
}

func main() {
	domain := flag.String("domain", "mockoidc.default.svc.cluster.local", "Domain name for OIDC issuer")
	port := flag.String("port", "8080", "Port to listen on")
	timeout := flag.Int("timeout", 3600, "Timeout in seconds")
	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Generate RSA key for mockoidc
	rsaKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	// Create mock OIDC server
	m, err := mockoidc.NewServer(rsaKey)
	if err != nil {
		panic(err)
	}

	// Start mock server on random port (we'll proxy it)
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}

	if err := m.Start(ln, nil); err != nil {
		panic(err)
	}
	defer func() { _ = m.Shutdown() }()

	// Get the actual port mockoidc is using
	actualPort := strings.Split(ln.Addr().String(), ":")[1]
	cfg := m.Config()
	clientSecretBase64 := base64.StdEncoding.EncodeToString([]byte(cfg.ClientSecret))

	// Create proxy server that rewrites responses
	// If domain already contains a port, don't add one
	fullDomain := *domain
	if !strings.Contains(*domain, ":") {
		fullDomain = *domain + ":" + *port
	}

	proxy := &ProxyHandler{
		mockServer: m,
		domain:     fullDomain,
		port:       actualPort,
	}

	proxyServer := &http.Server{
		Addr:    ":" + *port,
		Handler: proxy,
	}

	// Start proxy server
	go func() {
		fmt.Printf("MockOIDC Server with Domain Name Support\n")
		fmt.Printf("=========================================\n")
		fmt.Printf("Listening on: 0.0.0.0:%s\n", *port)
		fmt.Printf("Domain name: %s\n", fullDomain)
		fmt.Printf("Issuer: http://%s/oidc\n", fullDomain)
		fmt.Printf("\nConfiguration:\n")
		fmt.Printf("CLIENT_ID: %s\n", cfg.ClientID)
		fmt.Printf("CLIENT_SECRET: %s\n", cfg.ClientSecret)
		fmt.Printf("\nExport these variables:\n")
		fmt.Printf("export ISSUER=\"http://%s/oidc\"\n", fullDomain)
		fmt.Printf("export CLIENT_ID=%q\n", cfg.ClientID)
		fmt.Printf("export CLIENT_SECRET_BASE64_ENCODED=%q\n", clientSecretBase64)
		fmt.Printf("\nServer will run for %d seconds (or until Ctrl+C)...\n\n", *timeout)

		if err := proxyServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	// Wait for timeout or interrupt
	select {
	case <-time.After(time.Second * time.Duration(*timeout)):
		fmt.Println("\nTimeout reached - shutting down")
	case <-ctx.Done():
		fmt.Println("\nInterrupt received - shutting down")
	}

	// Graceful shutdown
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = proxyServer.Shutdown(shutdownCtx)
}
