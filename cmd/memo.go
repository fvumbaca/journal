package cmd

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fvumbaca/journal/pkg/notes"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var memoCMD = &cobra.Command{
	Use:     "memo [message...]",
	Short:   "Add a memo to a journal page with a timestamp",
	Long:    "Add a memo note to a journal page with a timestamp. Note that when this command is run with no arguments, it will read from stdin.",
	Aliases: []string{"m"},
	RunE: func(cmd *cobra.Command, args []string) error {
		monthOffset, _ := cmd.Flags().GetInt("month-offset")
		dayOffset, _ := cmd.Flags().GetInt("day-offset")
		dir := viper.GetString("journalPath")

		day := time.Now().AddDate(0, monthOffset, dayOffset)
		filename := notes.DayFilename(dir, day)

		err := os.MkdirAll(filepath.Dir(filename), 0775)
		if err != nil {
			return err
		}

		f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0775)
		if err != nil {
			return err
		}
		defer f.Close()

		if len(args) == 0 {
			err = notes.AppendMemo(f, day, os.Stdin)
		} else {
			err = notes.AppendMemo(f, day, strings.NewReader(strings.Join(args, " ")))
		}
		if err != nil {
			return err
		}

		return notes.IndexFile(dir, filename)
	},
}

func init() {
	rootCMD.AddCommand(memoCMD)

	memoCMD.Flags().IntP("day-offset", "o", 0, "Offset in days for note to load. Can be negative and stacked with other offsets.")
	memoCMD.Flags().IntP("month-offset", "m", 0, "Offset in months for note to load. Can be negative and stacked with other offsets.")
}
