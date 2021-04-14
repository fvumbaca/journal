package web

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"net/http"
	"strings"

	"github.com/fvumbaca/journal/pkg/notes"
	blackfriday "github.com/russross/blackfriday/v2"
)

type Handler struct {
	JournalDir string
}

//go:embed static/*
var staticFS embed.FS

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()
	// mux.Handle("/static/", serveFSFile(staticFS))
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("pkg/web/static"))))
	mux.HandleFunc("/", handleIndex(h.JournalDir))
	mux.ServeHTTP(w, r)
}

//go:embed templates/*
var templateFS embed.FS

type searchResult struct {
	Query         string
	ResultContent template.HTML
}

func handleIndex(baseDir string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// t, err := template.ParseFS(templateFS, "templates/*")
		var result searchResult
		t, err := template.ParseGlob("pkg/web/templates/*")
		if err != nil {
			fmt.Fprintln(w, "Template Error:", err)
			return
		}

		result.Query = r.URL.Query().Get("q")
		var buff bytes.Buffer

		if result.Query != "" {
			err = notes.WebSearch(baseDir, result.Query, 50, &buff)
			if err != nil {
				fmt.Fprintln(w, "Template Error:", err)
				return
			}
			result.ResultContent = template.HTML(blackfriday.Run(buff.Bytes()))
		}
		if err = t.ExecuteTemplate(w, "index.html", result); err != nil {
			fmt.Fprintln(w, "Template Error:", err)
		}
		// t.ExecuteTemplate(w, "index.html", struct {
		// 	ResultDoc string
		// 	Query     string
		// 	Results   []string
		// }{
		// 	Query: query,
		// })

	})
}

func serveFSFile(fs fs.FS) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		path = strings.TrimPrefix(path, "/")
		fmt.Println("serving", path)
		f, err := fs.Open(path)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, "404 - Not Found")
			return
		}
		defer f.Close()
		io.Copy(w, f)
	})
}
