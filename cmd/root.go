package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Execute() {
	if err := rootCMD.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

var rootCMD = &cobra.Command{
	Use:   "journal",
	Short: "A CLI Journal",
	Long:  `Journal is a fully featured journal for keeping daily and archival notes, all from your terminal.`,
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AddConfigPath("$HOME/.config/journal")
	viper.AddConfigPath("$HOME/.journalrc")

	defaults()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("config err:", err)
	}
}

func defaults() {
	viper.SetDefault("journalDir", os.ExpandEnv("$HOME/.journal"))
	viper.SetDefault("editor", os.ExpandEnv("$EDITOR"))
}
