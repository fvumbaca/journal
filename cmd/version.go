package cmd

import (
	"fmt"

	"github.com/fvumbaca/journal/version"
	"github.com/spf13/cobra"
)

var versionCMD = &cobra.Command{
	Use:   "version",
	Short: "Print version info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version:", version.Version)
		fmt.Println("SHA:", version.SHA)
	},
}

func init() {
	rootCMD.AddCommand(versionCMD)
}
