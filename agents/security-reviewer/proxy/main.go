package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	defaultListenAddr       = ":4000"
	defaultOpenAIBaseURL    = "https://api.openai.com"
	defaultAnthropicBaseURL = "https://api.anthropic.com"
	openAIInboundPrefix     = "/openai/"
	anthropicInboundPrefix  = "/anthropic/"
	healthPath              = "/health/liveliness"
	headerAuthorization     = "Authorization"
	headerAnthropicAPIKey   = "X-Api-Key"
)

// providerProxy defines how to forward requests to a specific upstream API.
type providerProxy struct {
	Prefix      string
	Target      *url.URL
	HeaderName  string
	HeaderValue string
	DisplayName string
}

// main configures the proxy service and starts the HTTP server.
func main() {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatalf("proxy configuration error: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc(healthPath, handleHealth)

	mountProxy(mux, cfg.openAIProxy, cfg.clientToken)
	mountProxy(mux, cfg.anthropicProxy, cfg.clientToken)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	server := &http.Server{
		Addr:         cfg.listenAddr,
		Handler:      withLogging(mux),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("proxy listening on %s (OpenAI -> %s, Anthropic -> %s)",
		cfg.listenAddr, cfg.openAIProxy.Target.String(), cfg.anthropicProxy.Target.String())

	if err = server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("proxy server error: %v", err)
	}
}

// proxyConfig captures runtime settings for the reverse proxy.
type proxyConfig struct {
	listenAddr     string
	openAIProxy    providerProxy
	anthropicProxy providerProxy
	clientToken    string
}

// loadConfig reads environment variables and constructs the proxy configuration.
func loadConfig() (proxyConfig, error) {
	listen := firstNonEmpty(os.Getenv("PROXY_LISTEN_ADDR"), defaultListenAddr)

	clientToken := strings.TrimSpace(os.Getenv("PROXY_API_KEY"))
	if clientToken == "" {
		return proxyConfig{}, errors.New("PROXY_API_KEY must be set")
	}

	openAIBase, err := parseBaseURL(firstNonEmpty(os.Getenv("PROXY_OPENAI_BASE_URL"), defaultOpenAIBaseURL))
	if err != nil {
		return proxyConfig{}, fmt.Errorf("parse OpenAI base URL: %w", err)
	}
	anthropicBase, err := parseBaseURL(firstNonEmpty(os.Getenv("PROXY_ANTHROPIC_BASE_URL"), defaultAnthropicBaseURL))
	if err != nil {
		return proxyConfig{}, fmt.Errorf("parse Anthropic base URL: %w", err)
	}

	openAIKey := strings.TrimSpace(os.Getenv("OPENAI_API_KEY"))
	anthropicKey := strings.TrimSpace(os.Getenv("ANTHROPIC_API_KEY"))

	openAIProxy := providerProxy{
		Prefix:      openAIInboundPrefix,
		Target:      openAIBase,
		HeaderName:  headerAuthorization,
		HeaderValue: bearerValue(openAIKey),
		DisplayName: "OpenAI",
	}
	anthropicProxy := providerProxy{
		Prefix:      anthropicInboundPrefix,
		Target:      anthropicBase,
		HeaderName:  headerAnthropicAPIKey,
		HeaderValue: anthropicKey,
		DisplayName: "Anthropic",
	}

	return proxyConfig{
		listenAddr:     listen,
		openAIProxy:    openAIProxy,
		anthropicProxy: anthropicProxy,
		clientToken:    clientToken,
	}, nil
}

// mountProxy attaches a provider proxy to the HTTP mux.
func mountProxy(mux *http.ServeMux, provider providerProxy, clientToken string) {
	handler := buildProviderHandler(provider, clientToken)
	mux.Handle(provider.Prefix, handler)
}

// buildProviderHandler creates an HTTP handler that forwards requests to the provider.
func buildProviderHandler(provider providerProxy, clientToken string) http.Handler {
	proxy := httputil.NewSingleHostReverseProxy(provider.Target)
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		rewriteRequest(req, provider)
	}
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("proxy error [%s]: %v", provider.DisplayName, err)
		http.Error(w, "upstream request failed", http.StatusBadGateway)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, provider.Prefix) {
			http.NotFound(w, r)
			return
		}
		if provider.HeaderValue == "" {
			log.Printf("proxy warning [%s]: request rejected due to missing API key", provider.DisplayName)
			http.Error(w, "upstream API key is not configured", http.StatusServiceUnavailable)
			return
		}
		if !validateClientToken(r.Header.Get(headerAuthorization), clientToken) {
			log.Printf("proxy warning [%s]: request rejected due to missing or invalid client bearer token", provider.DisplayName)
			http.Error(w, "invalid bearer token", http.StatusUnauthorized)
			return
		}

		proxy.ServeHTTP(w, r)
	})
}

// rewriteRequest adjusts the outbound request before it is sent upstream.
func rewriteRequest(req *http.Request, provider providerProxy) {
	req.URL.Scheme = provider.Target.Scheme
	req.URL.Host = provider.Target.Host
	req.Host = provider.Target.Host

	// Strip the provider prefix from the incoming path.
	trimmedPath := strings.TrimPrefix(req.URL.Path, provider.Prefix)
	if !strings.HasPrefix(trimmedPath, "/") {
		trimmedPath = "/" + trimmedPath
	}

	// Combine the target's base path with the trimmed path.
	basePath := provider.Target.Path
	req.URL.Path = joinURLPath(basePath, trimmedPath)
	req.URL.RawPath = req.URL.EscapedPath()

	stripSensitiveHeaders(req.Header)

	if provider.HeaderName == headerAuthorization {
		req.Header.Set(headerAuthorization, provider.HeaderValue)
	} else if provider.HeaderName != "" {
		req.Header.Set(provider.HeaderName, provider.HeaderValue)
	}
}

// stripSensitiveHeaders removes inbound authentication headers that should not propagate upstream.
func stripSensitiveHeaders(header http.Header) {
	header.Del(headerAuthorization)
	header.Del(headerAnthropicAPIKey)
}

// joinURLPath concatenates base and additional path segments.
func joinURLPath(basePath, extraPath string) string {
	switch {
	case basePath == "" || basePath == "/":
		return singleLeadingSlash(extraPath)
	case extraPath == "" || extraPath == "/":
		return singleLeadingSlash(basePath)
	default:
		return singleLeadingSlash(strings.TrimSuffix(basePath, "/") + "/" + strings.TrimPrefix(extraPath, "/"))
	}
}

// singleLeadingSlash ensures the provided path has a leading slash.
func singleLeadingSlash(path string) string {
	if path == "" {
		return "/"
	}
	if !strings.HasPrefix(path, "/") {
		return "/" + path
	}
	return path
}

// withLogging wraps the handler with structured request logging.
func withLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		remote := remoteAddr(r.Context(), r.RemoteAddr)
		if r.URL.Path != healthPath {
			log.Printf("proxy request method=%s path=%s remote=%s duration=%s",
				r.Method, r.URL.Path, remote, duration)
		}
	})
}

// remoteAddr normalizes the remote address for logging.
func remoteAddr(ctx context.Context, fallback string) string {
	if peer, ok := ctx.Value(http.LocalAddrContextKey).(net.Addr); ok {
		return peer.String()
	}
	return fallback
}

// handleHealth responds to health check requests.
func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write([]byte("ok"))
}

// parseBaseURL validates and normalizes the upstream base URL.
func parseBaseURL(raw string) (*url.URL, error) {
	parsed, err := url.Parse(raw)
	if err != nil {
		return nil, err
	}
	if parsed.Scheme == "" || parsed.Host == "" {
		return nil, fmt.Errorf("invalid URL %q (must include scheme and host)", raw)
	}
	if !strings.HasSuffix(parsed.Path, "/") {
		parsed.Path += "/"
	}
	return parsed, nil
}

// bearerValue formats the bearer token header.
func bearerValue(key string) string {
	key = strings.TrimSpace(key)
	if key == "" {
		return ""
	}
	return "Bearer " + key
}

// firstNonEmpty returns the first non-empty string in candidates.
func firstNonEmpty(values ...string) string {
	for _, v := range values {
		if strings.TrimSpace(v) != "" {
			return v
		}
	}
	return ""
}

func validateClientToken(headerValue, expectedToken string) bool {
	if expectedToken == "" {
		return false
	}
	parts := strings.SplitN(headerValue, " ", 2)
	if len(parts) != 2 {
		return false
	}
	if !strings.EqualFold(parts[0], "bearer") {
		return false
	}
	return strings.TrimSpace(parts[1]) == expectedToken
}
