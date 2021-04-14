package cmd

import (
	"fmt"

	"github.com/fvumbaca/journal/pkg/notes"
	"github.com/spf13/cobra"
)

var infoCMD = &cobra.Command{
	Use:   "info",
	Short: "Show info/stats about your journal.",
	Long: `Show info/stats about your journal. Review your writing and note
taking consistency or get a birds eye-view on current configuration settings.`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		dir, _ := cmd.PersistentFlags().GetString("journal-path")

		s, err := notes.NoteStats(dir)
		if err != nil {
			return err
		}
		fmt.Printf("Journal Entries: %d\n", s.JournalEntriesCount)
		fmt.Printf("Archives: %d\n", s.ArchivesCount)

		showVerboseInfo, _ := cmd.Flags().GetBool("verbose")
		if showVerboseInfo {
			fmt.Println("Client Version:", version)
			fmt.Println("Journal Path:", dir)
			editor, _ := cmd.Flags().GetString("editor")
			fmt.Println("Editor Command:", editor)
		}

		return nil
	},
}

func init() {
	rootCMD.AddCommand(infoCMD)

	infoCMD.Flags().BoolP("verbose", "v", false, "Set to show verbose info about your notes setup.")
}
