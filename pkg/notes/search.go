package notes

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
		if d == nil || d.IsDir() {
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
	return index.Index(filename, &struct {
		Filename string
		Content  string
		Kind     string
	}{
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

	for i, result := range searchResults.Hits[:maxResults] {
		filename := result.ID
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close()

		if i == 0 {
			fmt.Fprintf(out, "\n<!-- File: %s -->\n\n", filename)
		} else {
			fmt.Fprintf(out, "\n---\n<!-- File: %s -->\n\n", filename)
		}

		_, err = io.Copy(out, f)
		if err != nil {
			return err
		}
		f.Close()
	}

	return nil
}

// func SearchWeb(baseDir string, query string, page, numPage int) ([]string, error) {
// 	var results []string

// 	indexFilename := filepath.Join(baseDir, bleveIndexFilename)
// 	index, err := bleve.Open(indexFilename)
// 	if err != nil {
// 		return results, err
// 	}
// 	defer index.Close()

// 	q := bleve.NewQueryStringQuery(query)
// 	search := bleve.NewSearchRequest(q)

// 	search.Highlight = bleve.NewHighlightWithStyle("html")
// 	search.

// 	return results, nil
// }

func WebSearch(baseDir string, query string, numResults int, out io.Writer) error {
	indexFilename := filepath.Join(baseDir, bleveIndexFilename)
	index, err := bleve.Open(indexFilename)
	if err != nil {
		return err
	}
	defer index.Close()

	q := bleve.NewQueryStringQuery(query)
	search := bleve.NewSearchRequest(q)

	search.Highlight = bleve.NewHighlightWithStyle("html")

	searchResults, err := index.Search(search)
	if err != nil {
		return err
	}

	maxResults := numResults
	if c := len(searchResults.Hits); c < numResults {
		maxResults = c
	}

	for i, result := range searchResults.Hits[:maxResults] {
		filename := result.ID
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close()

		relFilename, err := filepath.Rel(baseDir, filename)
		if err != nil {
			relFilename = filename
		}

		if i == 0 {
			fmt.Fprintf(out, "> From File: [%s](%s)\n\n", relFilename, filename)
		} else {
			fmt.Fprintf(out, "\n---\n> From File: [%s](%s)\n\n", relFilename, filename)
		}

		_, err = io.Copy(out, f)
		if err != nil {
			return err
		}
		f.Close()
	}

	return nil
}
