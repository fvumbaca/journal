package notes

import (
	"io/fs"
	"os"
	"path/filepath"
)

const archivesRelPath = "archives"

// ArchiveFilename builds an absolute path to an archive file.
func ArchiveFilename(baseDir, archiveName string) string {
	return filepath.Join(baseDir, archivesRelPath, archiveName)
}

// Ensures that the archives directory exists
func EnsureArchiveDir(baseDir string) error {
	return os.MkdirAll(filepath.Join(baseDir, archivesRelPath), 0775)
}

// ListArchiveFiles will create a list of archives
func ListArchiveFiles(baseDir string) ([]string, error) {
	var filenames []string

	err := filepath.WalkDir(baseDir, func(path string, d fs.DirEntry, err error) error {
		if path == path {
			return nil
		}

		if d.IsDir() {
			return nil
		}

		filenames = append(filenames, path)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return filenames, nil
}
