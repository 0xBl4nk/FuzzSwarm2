package src

import (
	//"strings"
	"errors"
	"strings"

	"net/url"

	"github.com/spf13/cobra"
)


type Config struct {
  URL       string
  Headers   string
  Method    string
  Data      string
}

func LoadConfig(cmd *cobra.Command) (Config, error) {
  var cfg Config

  cfg.URL, _ = cmd.Flags().GetString("url")
  cfg.Data, _ = cmd.Flags().GetString("data")
  cfg.Method, _ = cmd.Flags().GetString("method")
  cfg.Headers, _ = cmd.Flags().GetString("headers")

  // URL Validate / Parser
  if cfg.URL == "" {
    return cfg, errors.New("the --url flag is required")
  }
  
  parsedURL, err := url.Parse(cfg.URL)
    if err != nil || !(parsedURL.Scheme == "http" || parsedURL.Scheme == "https") {
      return cfg, errors.New("invalid URL format. Ensure it starts with http:// or https://")
    }

  // Validate HTTP Method
  cfg.Method = strings.ToUpper(cfg.Method)
  if cfg.Method != "GET" && cfg.Method != "POST" {
    return cfg, errors.New("Invalid HTTP method. Only GET and POST are supported")
  }

  return cfg, nil
}
