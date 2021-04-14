package cmd

import (
	"fmt"
	"net/http"

	"github.com/fvumbaca/journal/pkg/web"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serverCMD = &cobra.Command{
	Use:     "server",
	Short:   "Start a web server to host notes.",
	Long:    "Start a web server to view rendered notes in your web browser.",
	Aliases: []string{"serv"},
	Args:    cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := viper.GetString("journalPath")

		h := web.Handler{
			JournalDir: dir,
		}

		port, _ := cmd.Flags().GetInt("port")
		addr := fmt.Sprintf(":%d", port)
		return http.ListenAndServe(addr, h)
	},
}

func init() {
	rootCMD.AddCommand(serverCMD)
	serverCMD.Flags().IntP("port", "p", 3000, "Web server port.")
}
