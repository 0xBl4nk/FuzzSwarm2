/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

  "github.com/ArthurHydr/FuzzSwarm2"
)

// reqCmd represents the req command
var reqCmd = &cobra.Command{
	Use:   "req",
	Short: "A brief description of your command",
	Long: `req`,
	Run: func(cmd *cobra.Command, args []string) {
		NewRequest()
	},
}

func init() {
	rootCmd.AddCommand(reqCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// reqCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// reqCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
