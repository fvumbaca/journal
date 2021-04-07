package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "unreleased"

var versionCMD = &cobra.Command{
	Use:   "version",
	Short: "Print version info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version:", version)
	},
}

func init() {
	rootCMD.AddCommand(versionCMD)
}
