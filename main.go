package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Version information set via ldflags during build
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

// rootCmd is the base command for whatomate CLI
var rootCmd = &cobra.Command{
	Use:   "whatomate",
	Short: "whatomate - WhatsApp automation tool",
	Long: `whatomate is a CLI tool for automating WhatsApp messages.
It allows you to send messages, manage contacts, and automate
common WhatsApp workflows from the command line.

Docs: https://github.com/shridarpatil/whatomate
Fork: https://github.com/me/whatomate`,
	SilenceUsage:  true,
	SilenceErrors: true,
}

// versionCmd prints the current version information
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("whatomate version %s (commit: %s, built: %s)\n", version, commit, date)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		// print the error and exit with a non-zero status code
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
