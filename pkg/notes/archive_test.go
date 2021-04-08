package notes

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestArchiveFilename(t *testing.T) {
	baseDir := "/Users/tester/journal"
	expected := "/Users/tester/journal/archives/special-note.md"

	if diff := cmp.Diff(expected, ArchiveFilename(baseDir, "special-note.md")); diff != "" {
		t.Error(diff)
	}
}
