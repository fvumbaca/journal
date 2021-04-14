package cmd

import (
	"time"

	"github.com/fvumbaca/journal/pkg/notes"
	"github.com/spf13/cobra"
)

var editCMD = &cobra.Command{
	Use:     "edit [archive_filename]",
	Short:   "Edit a journal page.",
	Long:    "Edit a journal page. When run with no arguments or flags, this command will open the current day's page in the set editor.",
	Aliases: []string{"e"},
	Args:    cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		dir, _ := cmd.Flags().GetString("journal-path")
		editor, _ := cmd.Flags().GetString("editor")

		var filename string

		if len(args) == 1 {
			filename = notes.ArchivePath(dir, args[0])
			if err := notes.EnsureArchiveDir(dir); err != nil {
				return err
			}
		} else {
			monthOffset, _ := cmd.Flags().GetInt("month-offset")
			dayOffset, _ := cmd.Flags().GetInt("day-offset")

			day := time.Now().AddDate(0, monthOffset, dayOffset)
			filename = notes.DayFilename(dir, day)
		}

		err := notes.EditNote(editor, filename)
		if err != nil {
			return err
		}
		return notes.IndexFile(dir, filename)
	},
}

func init() {
	rootCMD.AddCommand(editCMD)

	editCMD.Flags().IntP("day-offset", "o", 0, "Offset in days for note to load. Can be negative and stacked with other offsets.")
	editCMD.Flags().IntP("month-offset", "m", 0, "Offset in months for note to load. Can be negative and stacked with other offsets.")

}
