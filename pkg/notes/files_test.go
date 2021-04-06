package notes

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestDayFilename(t *testing.T) {
	baseDirectory := "/tmp/journal"
	now, _ := time.Parse("Mon Jan 2 15:04:05 -0700 MST 2006", "Mon Jan 2 15:04:05 -0700 MST 2006")
	r := DayFilename(baseDirectory, now)
	expected := "/tmp/journal/2006/January/02.md"

	if diff := cmp.Diff(expected, r); diff != "" {
		t.Error(diff)
	}

}
