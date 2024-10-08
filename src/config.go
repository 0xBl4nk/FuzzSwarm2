package src

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

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
  cfg.Retries = 3
  cfg.Script, _ = cmd.Flags().GetString("script")

 return cfg, nil
}

func ParseConfig(cfg Config) (Config, error) {
  
  // URL Validate.
  if cfg.URL == "" {
    return cfg, errors.New("the --url flag is required")
  }

  // URL Parsing.
  parsedURL, err := url.Parse(cfg.URL)
  if err != nil || !(parsedURL.Scheme == "http" || parsedURL.Scheme == "https") {
    return cfg, errors.New("invalid URL format. Ensure it starts with http:// or https://")
  }

  // Validate HTTP Method.
  cfg.Method = strings.ToUpper(cfg.Method)
  if cfg.Method != "GET" && cfg.Method != "POST" {
    return cfg, errors.New("Invalid HTTP method. Only GET and POST are supported")
  }

  // If method is POST, data must be provided.
  if cfg.Method == "POST" && cfg.Data == "" {
    return cfg, errors.New("Post method requires --data flag")
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
