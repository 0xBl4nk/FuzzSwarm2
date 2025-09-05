package src

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"os"
	"time"
	"math/rand"
)

// Common user agents for stealth
var userAgents = []string{
  "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36",
  "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36",
  "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:122.0) Gecko/20100101 Firefox/122.0",
  "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2.1 Safari/605.1.15",
  "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36",
  "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Edge/121.0.0.0",
}

// GetRandomUserAgent returns a random user agent string for stealth
func GetRandomUserAgent() string {
  return userAgents[rand.Intn(len(userAgents))]
}

func CreateClient(useProxy bool, timeout int, useSSL bool, followRedirects bool) *http.Client {
  transport := &http.Transport{
    MaxIdleConns:       100,
    IdleConnTimeout:    90 * time.Second,
    DisableCompression: false,
    ForceAttemptHTTP2:  true,
  }

  if useProxy {
    proxy := os.Getenv("HTTP_PROXY")
    if proxy == "" {
      LogFatal("Proxy configuration is missing in the .env file.")
    }
    
    proxyURL, err := url.Parse(proxy)
    if err != nil {
      LogFatal("Invalid proxy URL: %v", err)
    }

    transport.Proxy = http.ProxyURL(proxyURL)
    LogInfo("Using proxy: %s", proxyURL)
  }

  if useSSL {
    sslCertPath := os.Getenv("SSL_CERT_PATH")
    if sslCertPath == "" {
      LogFatal("SSL-Cert-File is missing in the .env file.")
    }
    cert, err := tls.LoadX509KeyPair(sslCertPath, sslCertPath)
    if err != nil {
      LogFatal("Failed to load SSL certificate from %s: %v", sslCertPath, err)
    }
    
    if transport.TLSClientConfig == nil {
      transport.TLSClientConfig = &tls.Config{}
    }
    transport.TLSClientConfig.Certificates = []tls.Certificate{cert}
    transport.TLSClientConfig.InsecureSkipVerify = false // Ensure secure connections
    LogInfo("Using SSL certificate from %s", sslCertPath)
  } else {
    // Set up secure TLS config even without custom cert
    transport.TLSClientConfig = &tls.Config{
      MinVersion: tls.VersionTLS12, // Enforce minimum TLS 1.2
    }
  }

  client := &http.Client{
    Transport: transport,
    Timeout:   time.Duration(timeout) * time.Second,
  }

  // Configure redirect policy
  if followRedirects {
    client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
      // Follow up to 10 redirects
      if len(via) >= 10 {
        return http.ErrUseLastResponse
      }
      return nil
    }
  } else {
    client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
      return http.ErrUseLastResponse // Don't follow redirects
    }
  }

  return client
}

