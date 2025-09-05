package cmd

import (
	"os"

	"github.com/0xBl4nk/FuzzSwarm2/scripts"
	"github.com/0xBl4nk/FuzzSwarm2/src"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var script string

var rootCmd = &cobra.Command{
  Use:   "FuzzSwarm",
	Short: "FuzzSwarm is a multi-threaded fuzzing tool for brute-forcing HTTP endpoints.",
	Long: `FuzzSwarm 2.0 is a powerful multi-threaded fuzzing tool designed for cybersecurity 
professionals to discover vulnerabilities in web applications and APIs. 

Features:
• Multi-threaded fuzzing with customizable concurrency
• Support for GET, POST, PUT, DELETE, PATCH methods
• Built-in security scripts (SSTI, SQLi, XSS)
• Proxy and SSL/TLS support for secure testing
• Multiple output formats (text, JSON, CSV)
• Advanced filtering and response analysis
• User-agent randomization for stealth testing
• Progress tracking for long-running scans

Examples:
  # Basic directory fuzzing
  FuzzSwarm -u http://target.com/FUZZ -W wordlist.txt

  # API endpoint testing with number ranges
  FuzzSwarm -X POST -u http://api.com/users/FUZZ -R 1-1000,3 -H "Content-Type: application/json"

  # SQL injection testing
  FuzzSwarm --script sqli -u "http://target.com/search?q=FUZZ" -v

  # SSTI vulnerability detection
  FuzzSwarm --script ssti -u "http://target.com/template?input=FUZZ" --output-format json

  # XSS payload testing
  FuzzSwarm --script xss -u "http://target.com/comment?text=FUZZ" -f 200 -v
`,
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

    if script != ""{
      scripts.InitScripts(script, &cfg)
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
  rootCmd.Flags().StringP("url", "u", "", "Target URL with FUZZ placeholder (required)")
  rootCmd.Flags().StringP("headers", "H", "", "Custom headers (format: 'Header1: value1,Header2: value2')")
  rootCmd.Flags().StringP("range", "R", "", "Numeric range for FUZZ replacement (format: start-end,digits, e.g., 1-10000,3)")
  rootCmd.Flags().StringP("method", "X", "GET", "HTTP method: GET, POST, PUT, DELETE, PATCH (default: GET)")
  rootCmd.Flags().StringP("wordlist", "W", "", "Path to wordlist file for FUZZ replacement")
  rootCmd.Flags().String("headers-file", "", "Path to file containing custom headers")
  rootCmd.Flags().Bool("use-ssl", false, "Use SSL client certificate from .env file")
  rootCmd.Flags().IntP("threads", "t", 10, "Number of concurrent threads (default: 10, max recommended: 1000)")
  rootCmd.Flags().IntP("filter-size", "f", 0, "Filter out responses with specific size (useful for eliminating error pages)")
  rootCmd.Flags().Int("timeout", 10, "Request timeout in seconds (default: 10)")
  rootCmd.Flags().Int("rate-limit", 0, "Rate limit between requests in milliseconds (0 = no limit)")
  rootCmd.Flags().Bool("use-proxy", false, "Enable HTTP proxy from .env file")
  rootCmd.Flags().BoolP("verbose", "v", false, "Show detailed output including response previews")
  rootCmd.Flags().StringP("data", "d", "", "Request body data for POST/PUT/PATCH methods")
  rootCmd.Flags().String("output-format", "text", "Output format: text (colored), json, csv")
  rootCmd.Flags().Bool("follow-redirects", true, "Follow HTTP redirects (up to 10 redirects)")

  rootCmd.Flags().StringVar(&script, "script", "", "Use built-in security scripts: ssti, sqli, xss")
}
