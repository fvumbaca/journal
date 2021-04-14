package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func Execute() {
	if err := rootCMD.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

var rootCMD = &cobra.Command{
	Use:   "journal",
	Short: "A CLI Journal",
	Long: `Journal is an oppinionated journaling utility for keeping daily
and archival notes - all from your terminal.`,
}

func init() {
	rootCMD.PersistentFlags().StringP("journal-path", "D", os.ExpandEnv("$HOME/.journal"), "Directory journal entries are stored in.")

	defaultEditor := os.Getenv("EDITOR")
	if defaultEditor == "" {
		defaultEditor = "open"
	}

	rootCMD.PersistentFlags().StringP("editor", "e", defaultEditor, "Editor command to use for editing pages.")

}
