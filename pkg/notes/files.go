package notes

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// DayFilename will build the absolute path to the file containing the day's
// notes.
func DayFilename(base string, t time.Time) string {
	return filepath.Join(base, fmt.Sprintf("%d/%s/%02d", t.Year(), t.Month().String(), t.Day())) + ".md"
}

// EditNote will open the specified filename with the passed editor command.
// Note that this function blocks until the editor command returns. This is
// usually intentional since it allows post-processing of the file after they
// have been edited.
func EditNote(editor, filename string) error {
	err := os.MkdirAll(filepath.Dir(filename), 0775)
	if err != nil {
		return err
	}

	fmt.Println("Please make your entry in the opened editor")

	editCmd := strings.Split(editor, " ")
	editCmd = append(editCmd, filename)
	c := exec.Command(editCmd[0], editCmd[1:]...)
	return c.Run()
}
