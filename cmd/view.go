package cmd

import (
	"io"
	"os"
	"time"

	"github.com/fvumbaca/journal/pkg/notes"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var viewCMD = &cobra.Command{
	Use:     "view",
	Short:   "Read journal pages.",
	Long:    "Read journal pages. With no arguments, this command will only print the current day's page.",
	Aliases: []string{"v"},
	RunE: func(cmd *cobra.Command, args []string) error {
		monthOffset, _ := cmd.Flags().GetInt("month-offset")
		dayOffset, _ := cmd.Flags().GetInt("day-offset")
		dir := viper.GetString("journal-path")

		day := time.Now().AddDate(0, monthOffset, dayOffset)
		filename := notes.DayFilename(dir, day)

		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(os.Stdout, f)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCMD.AddCommand(viewCMD)

	viewCMD.Flags().IntP("day-offset", "o", 0, "Offset in days for note to load. Can be negative and stacked with other offsets.")
	viewCMD.Flags().IntP("month-offset", "m", 0, "Offset in months for note to load. Can be negative and stacked with other offsets.")

	viewCMD.Flags().StringP("journal-path", "D", "", "Directory journal entries are stored in.")
	viper.BindPFlag("journalPath", viewCMD.Flags().Lookup("journal-path"))
}
