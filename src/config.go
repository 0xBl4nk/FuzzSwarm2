package src

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"net"

	"github.com/spf13/cobra"
)

// Configuration settings for the fuzzing process.
type Config struct {
  URL           string
  Method        string
  Data          string
  Range         string
  Wordlist      string
  Script        string
  Values        []string
  UseProxy      bool
  Threads       int
  UseSSL        bool
  Timeout       int
  Retries       int
  Verbose       bool
  RateLimit     int
  FilterSize    int
  Headers       string
  ParsedHeaders map[string]string
  OutputFormat  string
  FollowRedirects bool

  HeadersFile   map[string]string
  HeadersPath   string
}

func LoadConfig(cmd *cobra.Command) (Config, error) {
  var cfg Config

  cfg.URL, _ = cmd.Flags().GetString("url")
  cfg.Data, _ = cmd.Flags().GetString("data")
  cfg.Threads, _ = cmd.Flags().GetInt("threads")
  cfg.Range, _ = cmd.Flags().GetString("range")
  cfg.Method, _ = cmd.Flags().GetString("method")
  cfg.Wordlist, _ = cmd.Flags().GetString("wordlist")
  cfg.HeadersPath, _ = cmd.Flags().GetString("headers-file")
  cfg.Timeout, _ = cmd.Flags().GetInt("timeout")
  cfg.FilterSize, _ = cmd.Flags().GetInt("filter-size")
  cfg.RateLimit, _ = cmd.Flags().GetInt("rate-limit")
  cfg.UseProxy, _ = cmd.Flags().GetBool("use-proxy")
  cfg.UseSSL, _ = cmd.Flags().GetBool("use-ssl")
  cfg.Verbose, _ = cmd.Flags().GetBool("verbose")
  cfg.Headers, _ = cmd.Flags().GetString("headers")
  cfg.OutputFormat, _ = cmd.Flags().GetString("output-format")
  cfg.FollowRedirects, _ = cmd.Flags().GetBool("follow-redirects")
  cfg.Retries = 3
  cfg.Script, _ = cmd.Flags().GetString("script")

 return cfg, nil
}

func ParseConfig(cfg Config) (Config, error) {
  
  // URL Validate.
  if cfg.URL == "" {
    return cfg, errors.New("the --url flag is required")
  }

  // Enhanced URL Parsing and Security Validation
  parsedURL, err := url.Parse(cfg.URL)
  if err != nil {
    return cfg, fmt.Errorf("invalid URL format: %v", err)
  }
  
  // Security checks for URL
  if err := validateURLSecurity(parsedURL); err != nil {
    return cfg, fmt.Errorf("URL security validation failed: %v", err)
  }

  // Validate HTTP Method.
  cfg.Method = strings.ToUpper(cfg.Method)
  if cfg.Method != "GET" && cfg.Method != "POST" && cfg.Method != "PUT" && cfg.Method != "DELETE" && cfg.Method != "PATCH" {
    return cfg, errors.New("invalid HTTP method. Supported methods: GET, POST, PUT, DELETE, PATCH")
  }

  // If method is POST, PUT, or PATCH, data should be provided.
  if (cfg.Method == "POST" || cfg.Method == "PUT" || cfg.Method == "PATCH") && cfg.Data == "" {
    LogInfo("Warning: %s method without --data flag may not be effective", cfg.Method)
  }

  // Validate thread count
  if cfg.Threads <= 0 {
    return cfg, errors.New("thread count must be greater than 0")
  }
  if cfg.Threads > 1000 {
    LogInfo("Warning: Using more than 1000 threads may cause performance issues")
  }

  // Validate timeout
  if cfg.Timeout <= 0 {
    return cfg, errors.New("timeout must be greater than 0")
  }

  // Validate rate limit
  if cfg.RateLimit < 0 {
    return cfg, errors.New("rate limit cannot be negative")
  }

  // Validate output format
  validFormats := []string{"text", "json", "csv"}
  validFormat := false
  for _, format := range validFormats {
    if cfg.OutputFormat == format {
      validFormat = true
      break
    }
  }
  if !validFormat {
    return cfg, fmt.Errorf("invalid output format: %s. Valid formats: %v", cfg.OutputFormat, validFormats)
  }

  // Validate Headers File/Path.
  if cfg.HeadersPath != "" {
    headersFile, err := ReadHeaders(cfg.HeadersPath)
    if err != nil {
      return cfg, fmt.Errorf("error reading headers: %v", err)
    }
    cfg.HeadersFile = headersFile
  }

  // Validate if script != ""
  if cfg.Script != ""{

  } else {
  // Validate Wordlist and Range.
    if cfg.Wordlist != "" {
      wordlistValues, err := ReadWordlistValues(cfg.Wordlist)
      if err != nil {
        return cfg, fmt.Errorf("error reading wordlist: %v", err)
      }
      cfg.Values = wordlistValues

    } else if cfg.Range != "" {
      rangeValues, err := parseRange(cfg.Range)
      
      if err != nil {
        return cfg, fmt.Errorf("error parsing range: %v", err)
      } 
      cfg.Values = rangeValues

    } else {
        return cfg, errors.New("either a range or wordlist must be provided")
    }
}
  return cfg, nil
}

// validateURLSecurity performs security checks on the target URL
func validateURLSecurity(parsedURL *url.URL) error {
  // Check for valid scheme
  if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
    return errors.New("URL must use http or https scheme")
  }
  
  // Check for dangerous hosts to prevent SSRF
  host := parsedURL.Hostname()
  if host == "" {
    return errors.New("URL must contain a valid hostname")
  }
  
  // Parse IP address if it's an IP
  if ip := net.ParseIP(host); ip != nil {
    // Check for localhost and private IP ranges
    if ip.IsLoopback() {
      LogInfo("Warning: Target is localhost - ensure this is intentional")
    }
    if ip.IsPrivate() {
      LogInfo("Warning: Target is in private IP range - ensure this is intentional")
    }
    // Block some dangerous ranges
    if ip.IsMulticast() || ip.IsUnspecified() {
      return errors.New("invalid target IP address")
    }
  }
  
  // Check for dangerous URLs
  dangerousPatterns := []string{
    "file://", "ftp://", "gopher://", "dict://", "ldap://",
  }
  
  urlString := parsedURL.String()
  for _, pattern := range dangerousPatterns {
    if strings.Contains(strings.ToLower(urlString), pattern) {
      return fmt.Errorf("potentially dangerous URL scheme detected: %s", pattern)
    }
  }
  
  return nil
}
