package cmd

import (
	"os"

	"github.com/0xBl4nk/FuzzSwarm2/src"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "FuzzSwarm",
	Short: "FuzzSwarm is a fuzzing tool designed for brute-forcing HTTP endpoints.",
	Long: `FuzzSwarm is a fuzzing tool designed for brute-forcing HTTP endpoints. It supports optional proxy usage, SSL configuration, and response size filtering to focus on significant results. `,
	Run: func(cmd *cobra.Command, args []string) { 
    
    err := godotenv.Load()
    if err != nil {
      src.LogFatal("No .env file found or failed to load. Continuing without environment variables.")
    }

    cfg, err := src.LoadConfig(cmd)
    if err != nil {
      src.LogFatal("Failed to load configuration: %v", err)
    }

    cfg, err = src.ParseConfig(cfg)
    if err != nil {
      src.LogFatal("Failed to validate configuration: %v", err)
    }
    src.StartFuzzing(cfg)
  },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
  rootCmd.Flags().StringP("url", "u", "", "The target URL.")
  rootCmd.Flags().StringP("headers", "H", "", "Custom header.")
  rootCmd.Flags().StringP("range", "R", "", "Range of numbers to use, format start-end,digits (e.g., 1-10000,3).")
  rootCmd.Flags().StringP("method", "X", "GET", "HTTP method to use: GET or POST. (Default: GET)")
  rootCmd.Flags().StringP("wordlist", "W", "", "Path to the wordlist file.")
  rootCmd.Flags().String("headers-file", "", "Path to the headers file.")
  rootCmd.Flags().String("use-ssl", "", "Enable SSL certificate from .env file.")
  rootCmd.Flags().IntP("threads", "t", 10, "Number of threads to use for fuzzing.")
  rootCmd.Flags().IntP("filter-size", "f", 0, "Filter responses by size (skip responses with this size).")
  rootCmd.Flags().Int("timeout", 10, "Set timeout.")
  rootCmd.Flags().Int("rate-limit", 0, "Rate limit in milliseconds between requests.")
  rootCmd.Flags().Bool("use-proxy", false, "Display verbose output including response preview.")
  rootCmd.Flags().BoolP("verbose", "v", false, "Enable proxy configuration from .env file.")
  rootCmd.Flags().StringP("data", "d", "", "POST data.")
}
