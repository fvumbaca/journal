package search

import (
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/blevesearch/bleve"
)

type fileInfo struct {
	Filename string
	Content  string
}

const bleveIndexFilename = ".index"

// IndexDir will add/update an index for a folder.
func IndexDir(baseDir string) error {
	indexFilename := filepath.Join(baseDir, bleveIndexFilename)
	index, err := bleve.Open(indexFilename)

	if err == bleve.ErrorIndexPathDoesNotExist {
		mapping := bleve.NewIndexMapping()
		index, err = bleve.New(indexFilename, mapping)
		if err != nil {
			return nil
		}
	}
	if err != nil {
		return err
	}

	defer index.Close()

	err = filepath.WalkDir(baseDir, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		// TODO: whitelist files
		if strings.HasPrefix(path, indexFilename) {
			return nil
		}

		fmt.Println("indexing", path)

		return _indexFile(index, path)
	})

	if err != nil {
		return err
	}

	return err
}

func IndexFile(baseDir string, filename string) error {
	indexFilename := filepath.Join(baseDir, bleveIndexFilename)
	index, err := bleve.Open(indexFilename)

	if err == bleve.ErrorIndexPathDoesNotExist {
		mapping := bleve.NewIndexMapping()
		index, err = bleve.New(indexFilename, mapping)
		if err != nil {
			return nil
		}
	}
	if err != nil {
		return err
	}
	defer index.Close()
	return _indexFile(index, filename)
}

func _indexFile(index bleve.Index, filename string) error {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	// TODO: Concurrency?
	return index.Index(filename, &fileInfo{
		Filename: filename,
		Content:  string(contents),
	})
}

func Search(baseDir string, query string, numResults int, out io.Writer) error {
	indexFilename := filepath.Join(baseDir, bleveIndexFilename)
	index, err := bleve.Open(indexFilename)
	if err != nil {
		return err
	}
	defer index.Close()

	q := bleve.NewQueryStringQuery(query)
	search := bleve.NewSearchRequest(q)

	searchResults, err := index.Search(search)
	if err != nil {
		return err
	}

	maxResults := numResults
	if c := len(searchResults.Hits); c < numResults {
		maxResults = c
	}

	for _, result := range searchResults.Hits[:maxResults] {
		filename := result.ID
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close()

		fmt.Fprintf(os.Stdout, "\n---\n<!-- File: %s -->\n\n", filename)
		_, err = io.Copy(os.Stdout, f)
		if err != nil {
			return err
		}

		f.Close()
	}

	return nil
}
