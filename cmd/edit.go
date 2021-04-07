package cmd

import (
	"time"

	"github.com/fvumbaca/journal/pkg/notes"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var editCMD = &cobra.Command{
	Use:     "edit",
	Short:   "Edit a journal page.",
	Long:    "Edit a journal page. When run with no arguments or flags, this command will open the current day's page in the set editor.",
	Aliases: []string{"e"},
	Args:    cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		monthOffset, _ := cmd.Flags().GetInt("month-offset")
		dayOffset, _ := cmd.Flags().GetInt("day-offset")

		dir := viper.GetString("journalPath")
		editor := viper.GetString("editor")

		day := time.Now().AddDate(0, monthOffset, dayOffset)
		filename := notes.DayFilename(dir, day)
		err := notes.EditNote(editor, filename)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCMD.AddCommand(editCMD)

	editCMD.Flags().IntP("day-offset", "o", 0, "Offset in days for note to load. Can be negative and stacked with other offsets.")
	editCMD.Flags().IntP("month-offset", "m", 0, "Offset in months for note to load. Can be negative and stacked with other offsets.")

	editCMD.Flags().StringP("journal-path", "D", "", "Directory journal entries are stored in.")
	viper.BindPFlag("journalPath", editCMD.Flags().Lookup("journal-path"))
	editCMD.Flags().StringP("editor", "e", "", "Editor to use for notes.")
	viper.BindPFlag("editor", editCMD.Flags().Lookup("editor"))
}
