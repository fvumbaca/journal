package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fvumbaca/journal/pkg/notes"
	bf "github.com/russross/blackfriday/v2"
	"github.com/spf13/cobra"
)

func xmain() {
	f, err := os.Open("example.md")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	contents, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	root := bf.New().Parse(contents)
	fmt.Println("title:", string(root.Title))
}

func main() {
	newRootCMD().Execute()
}

var baseDir = "/Users/frank/.journal"

func newRootCMD() *cobra.Command {
	cmd := cobra.Command{
		Use: "journal",
	}

	home := os.Getenv("HOME")
	if home == "" {
		fatal(fmt.Errorf("Could not determine home-path. Please set your home path with the $HOME environment variable."))
	}

	cmd.Flags().StringVarP(&baseDir, "dir", "D", home+"/.journal", "Base directory for all notes.")

	cmd.AddCommand(newOpenCMD())
	cmd.AddCommand(newMemoCMD())
	cmd.AddCommand(newViewCMD())
	return &cmd
}

func newOpenCMD() *cobra.Command {
	editor := "code -w"
	offset := 0
	cmd := cobra.Command{
		Use:     "open",
		Aliases: []string{"o"},
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, argv []string) {

			day := time.Now().AddDate(0, 0, offset)

			filename := notes.DayFilename(baseDir, day)
			err := notes.EditNote(editor, filename)
			if err != nil {
				fatal(err)
			}
		},
	}

	cmd.Flags().StringVarP(&editor, "editor", "e", getEnv("EDITOR", "code -w"), "Editor to use for editing notes.")
	cmd.Flags().IntVarP(&offset, "offset", "o", offset, "Offset in days for the note to edit.")
	return &cmd
}

func newMemoCMD() *cobra.Command {
	cmd := cobra.Command{
		Use:     "memo",
		Aliases: []string{"m"},
		Run: func(cmd *cobra.Command, argv []string) {
			now := time.Now()
			filename := notes.DayFilename(baseDir, now)

			err := os.MkdirAll(filepath.Dir(filename), 0775)
			if err != nil {
				fatal(err)
			}

			f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0775)
			if err != nil {
				fatal(err)
			}
			defer f.Close()

			if len(argv) == 0 {
				err = notes.AppendMemo(f, now, os.Stdin)
			} else {
				err = notes.AppendMemo(f, now, strings.NewReader(strings.Join(argv, " ")))
			}
			if err != nil {
				fatal(err)
			}
		},
	}

	return &cmd
}

func newViewCMD() *cobra.Command {
	cmd := cobra.Command{
		Use:     "view",
		Aliases: []string{"v"},
		Run: func(cmd *cobra.Command, args []string) {
			now := time.Now()
			filename := notes.DayFilename(baseDir, now)

			f, err := os.Open(filename)
			if err != nil {
				fatal(err)
			}
			defer f.Close()

			_, err = io.Copy(os.Stdout, f)
			if err != nil {
				fatal(err)
			}
		},
	}
	return &cmd
}

func getEnv(name, d string) string {
	v, ok := os.LookupEnv(name)
	if !ok {
		return d
	}
	return v
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
