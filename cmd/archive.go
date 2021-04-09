package cmd

import (
	"fmt"

	"github.com/fvumbaca/journal/pkg/notes"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var archivesCMD = &cobra.Command{
	Use:   "archives",
	Short: "List archive files.",
	Long:  "List archive files.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := viper.GetString("journalPath")

		err := notes.EnsureArchiveDir(dir)
		if err != nil {
			return err
		}

		filenames, err := notes.ListArchiveFiles(dir)
		if err != nil {
			return err
		}

		for _, filename := range filenames {
			fmt.Println(filename)
		}

		return nil
	},
}

func init() {
	rootCMD.AddCommand(archivesCMD)
}
