package cmd

import (
	"os"

	"github.com/0xBl4nk/FuzzSwarm2/src"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "FuzzSwarm2",
	Short: "A brief description of your application",
	Long: `.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) { 
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

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
  rootCmd.Flags().StringP("url", "u", "", "The target URL.")
  rootCmd.Flags().StringP("range", "R", "", "Range of numbers to use, format start-end,digits (e.g., 1-10000,3).")
  rootCmd.Flags().StringP("method", "X", "GET", "HTTP method to use: GET or POST. (Default: GET)")
  rootCmd.Flags().StringP("wordlist", "W", "", "Path to the wordlist file.")
  rootCmd.Flags().StringP("headers-file", "H", "", "Path to the headers file.")
  rootCmd.Flags().Bool("use-proxy", false, "Enable proxy configuration from .env file.")
}
