package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/fvumbaca/journal/pkg/fs"
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

			filename := fs.DayFilename(baseDir, day)
			err := editNote(editor, filename)
			if err != nil {
				fatal(err)
			}
		},
	}

	cmd.Flags().StringVarP(&editor, "editor", "e", getEnv("EDITOR", "code -w"), "Editor to use for editing notes.")
	cmd.Flags().IntVarP(&offset, "offset", "o", offset, "Offset in days for the note to edit.")
	return &cmd
}

func editNote(editor, filename string) error {
	err := os.MkdirAll(filepath.Dir(filename), 0775)
	if err != nil {
		return err
	}

	fmt.Println("Please make your entry in the opened browser")

	editCmd := strings.Split(editor, " ")
	editCmd = append(editCmd, filename)
	c := exec.Command(editCmd[0], editCmd[1:]...)
	return c.Run()
}

func newMemoCMD() *cobra.Command {
	cmd := cobra.Command{
		Use:     "memo",
		Aliases: []string{"m"},
		Run: func(cmd *cobra.Command, argv []string) {
			now := time.Now()
			filename := fs.DayFilename(baseDir, now)

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
				fmt.Fprintf(f, "\n**%s** ", now.Format("03:04 PM"))
				_, err = io.Copy(f, os.Stdin)
				if err != nil {
					fatal(err)
				}
				_, err = fmt.Fprintln(f)
			} else {
				_, err = fmt.Fprintf(f, "\n**%s** %s\n", now.Format("03:04 PM"), strings.Join(argv, " "))
			}
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
