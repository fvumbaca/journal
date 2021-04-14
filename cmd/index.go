package cmd

import (
	"fmt"

	"github.com/fvumbaca/journal/pkg/notes"
	"github.com/spf13/cobra"
)

var indexCMD = &cobra.Command{
	Use:   "index",
	Short: "Re-index all journal pages for searching.",
	Long: `Re-index all journal pages for searching. This operation is
usually done automatically after any edit operation and is here in case
the user wants to trigger it manually.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		dir, _ := cmd.Flags().GetString("journal-path")
		fmt.Println("Indexing files in ", dir)
		return notes.IndexDir(dir)
	},
}

func init() {
	rootCMD.AddCommand(indexCMD)
}
