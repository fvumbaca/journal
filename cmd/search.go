package cmd

import (
	"os"
	"strings"

	"github.com/fvumbaca/journal/pkg/search"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var searchCMD = &cobra.Command{
	Use:     "search",
	Short:   "Run a search across all journal pages.",
	Long:    "Run a search across all journal pages and return a markdown document of the results.",
	Aliases: []string{"s"},
	RunE: func(cmd *cobra.Command, args []string) error {
		resultsCount, _ := cmd.Flags().GetInt("results")
		dir := viper.GetString("journalPath")

		return search.Search(dir, strings.Join(args, " "), resultsCount, os.Stdout)
	},
}

func init() {
	rootCMD.AddCommand(searchCMD)

	searchCMD.Flags().IntP("results", "r", 5, "Number of results to display.")
}
