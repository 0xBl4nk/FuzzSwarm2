package src

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"os"
	"time"
)

func CreateClient(useProxy bool, timeout int, useSSL bool) *http.Client {
  transport := &http.Transport{}

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


    if useSSL {
      sslCertPath := os.Getenv("SSL_CERT_PATH")
      if sslCertPath == "" {
        LogFatal("SSL-Cert-File is missing in the .env file.")
      }
      cert, err := tls.LoadX509KeyPair(sslCertPath, sslCertPath)
      if err != nil {
        LogFatal("Failed to load SSL certificate from %s: %v", sslCertPath, err)
      }
      transport.TLSClientConfig = &tls.Config{Certificates: []tls.Certificate{cert}}
      LogInfo("Using SSL certificate from %s", sslCertPath)
    }
  }

  client := &http.Client{
    Transport: transport,
    Timeout: time.Duration(timeout) * time.Second,
  }

  return client
}

