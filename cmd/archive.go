package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/fvumbaca/journal/pkg/notes"
	"github.com/spf13/cobra"
)

var archivesCMD = &cobra.Command{
	Use:   "archives",
	Short: "List archive files.",
	Long:  "List archive files.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		dir, _ := cmd.PersistentFlags().GetString("journal-path")

		err := notes.EnsureArchiveDir(dir)
		if err != nil {
			return err
		}

		notes, err := notes.ListArchiveFiles(dir)
		if err != nil {
			return err
		}

		if showAbsolute, _ := cmd.Flags().GetBool("absolute-path"); showAbsolute {
			for _, note := range notes {
				fmt.Println(filepath.Join(note.DIR, note.Name))
			}
		} else {
			for _, note := range notes {
				fmt.Println(note.Name)
			}
		}

		return nil
	},
}

func init() {
	rootCMD.AddCommand(archivesCMD)

	archivesCMD.Flags().BoolP("absolute-path", "a", false, "Show full absolute paths.")
}
