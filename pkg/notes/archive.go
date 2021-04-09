package notes

import (
	"io/fs"
	"os"
	"path/filepath"
)

const archivesRelPath = "archives"

// ArchivePath builds an absolute path to a file in the archive directory.
func ArchivePath(baseDir string, archiveName ...string) string {
	parts := []string{baseDir, archivesRelPath}
	parts = append(parts, archiveName...)
	return filepath.Join(parts...)
}

// Ensures that the archives directory exists.
func EnsureArchiveDir(baseDir string) error {
	return os.MkdirAll(filepath.Join(baseDir, archivesRelPath), 0775)
}

// ListArchiveFiles will create a list of archives
func ListArchiveFiles(baseDir string) ([]Note, error) {
	var notes []Note

	archivesPath := ArchivePath(baseDir)
	err := filepath.WalkDir(archivesPath, func(path string, d fs.DirEntry, err error) error {
		if d == nil || d.IsDir() {
			return nil
		}

		relPath, _ := filepath.Rel(archivesPath, path)
		notes = append(notes, Note{
			DIR:  archivesPath,
			Name: relPath,
		})
		return nil
	})

	if err != nil {
		return nil, err
	}

	return notes, nil
}
