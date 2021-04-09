package notes

import (
	"io"
	"os"
	"path/filepath"
)

// type fileInfo struct {
// 	Filename string
// 	Content  string
// }

type Note struct {
	DIR  string
	Name string
}

func (n Note) Open() (io.ReadCloser, error) {
	return os.Open(filepath.Join(n.DIR, n.Name))
}
