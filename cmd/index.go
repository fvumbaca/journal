package cmd

import (
	"github.com/fvumbaca/journal/pkg/search"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var indexCMD = &cobra.Command{
	Use:   "index",
	Short: "Re-index all journal pages for searching.",
	Long: `Re-index all journal pages for searching. This operation is
usually done automatically after any edit operation and is here in case
the user wants to trigger it manually.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := viper.GetString("journalPath")
		return search.IndexDir(dir)
	},
}

func init() {
	rootCMD.AddCommand(indexCMD)

	indexCMD.Flags().StringP("journal-path", "D", "", "Directory journal entries are stored in.")
	viper.BindPFlag("journalPath", indexCMD.Flags().Lookup("journal-path"))
}