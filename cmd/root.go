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
    config, err := src.LoadConfig(cmd)
    if err != nil {
      src.LogFatal("Failed to load configuration: %v", err)
    }
    src.StartFuzzing(config)
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
  rootCmd.Flags().StringP("method", "X", "", "The target URL.")
}


