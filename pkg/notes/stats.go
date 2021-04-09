package notes

import (
	"io/fs"
	"path/filepath"
	"regexp"
	"strings"
)

type Stats struct {
	JournalEntriesCount int
	ArchivesCount       int
}

func NoteStats(baseDir string) (Stats, error) {
	var s Stats
	indexFilepath := filepath.Join(baseDir, ".index")
	archivesFilepath := filepath.Join(baseDir, "/archives")
	journalsFilepathRegexp := regexp.MustCompile(`^` + regexp.QuoteMeta(baseDir) + `/[0-9]+/[a-zA-Z]+/[0-9]+\..+`)

	err := filepath.WalkDir(baseDir, func(path string, d fs.DirEntry, err error) error {
		if d == nil || d.IsDir() {
			return nil
		}

		if strings.HasPrefix(path, indexFilepath) {
			return nil
		} else if strings.HasPrefix(path, archivesFilepath) {
			s.ArchivesCount++
		} else if journalsFilepathRegexp.MatchString(path) {
			s.JournalEntriesCount++
		}

		return nil
	})

	return s, err
}
