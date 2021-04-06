package notes

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestAppendMemoToNote(t *testing.T) {
	now, _ := time.Parse("Mon Jan 2 15:04:05 -0700 MST 2006", "Mon Jan 2 15:04:05 -0700 MST 2006")
	var buff bytes.Buffer
	err := AppendMemo(&buff, now, strings.NewReader("this is a test"))
	if err != nil {
		t.Error(err)
	}

	expected := "\n**03:04 PM** this is a test\n"

	if diff := cmp.Diff(expected, buff.String()); diff != "" {
		t.Error(diff)
	}
}
