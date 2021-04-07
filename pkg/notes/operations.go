package notes

import (
	"fmt"
	"io"
	"time"
)

// AppendMemo appends a memo entry to a writer.
func AppendMemo(w io.Writer, now time.Time, memo io.Reader) error {
	_, err := fmt.Fprintf(w, "\n**%s** ", now.Format("03:04 PM"))
	if err != nil {
		return err
	}
	_, err = io.Copy(w, memo)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(w)
	return err
}
