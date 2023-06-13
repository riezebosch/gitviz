package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/git-lfs/gitobj"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/gobuffalo/packr/v2"
)

// Routes for graph and content
func Routes(wd string) *chi.Mux {
	r := chi.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"https://riezebosch.github.io"},
		AllowedMethods: []string{"GET"},
	})
	r.Use(cors.Handler)

	r.Route("/", func(r chi.Router) {
		box := packr.New("html", "./html")

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			html, _ := box.Find("index.html")
			w.Write(html)
		})
		r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
			output, _ := box.Find(r.URL.EscapedPath())
			w.Write(output)
		})
	})

	if wd == "" {
		wd, _ = os.Getwd()
	}

	repo := repo(wd)

	r.Route("/api", func(r chi.Router) {
		r.Get("/info", get(info, wd))
		r.Get("/graph", get(graph, repo))
		r.Route("/refs", func(r chi.Router) {
			r.Get("/{type}/*", get(ref, repo))
			r.Get("/remotes/{remote}/{name}", get(remote, repo))
			r.Get("/HEAD", get(head, repo))
		})
		r.Route("/objects", func(r chi.Router) {
			r.Get("/blob/{id}", get(blob, repo))
			r.Get("/tree/{id}", get(tree, repo))
			r.Get("/commit/{id}", get(commit, repo))
		})
	})

	return r
}

// Info from current repo
type Info struct {
	Directory string `json:"directory"`
}

func info(w http.ResponseWriter, r *http.Request, wd string) {
	render.JSON(w, r, Info{Directory: filepath.Base(wd)})
}

func graph(w http.ResponseWriter, r *http.Request, path string) {
	render.JSON(w, r, Visit(path))
}

func ref(w http.ResponseWriter, r *http.Request, path string) {
	t := chi.URLParam(r, "type")
	n := chi.URLParam(r, "*")
	w.Write([]byte(readFirstLine(filepath.Join(path, "refs", t, n))))
}

func remote(w http.ResponseWriter, r *http.Request, path string) {
	t := chi.URLParam(r, "remote")
	n := chi.URLParam(r, "name")
	w.Write([]byte(readFirstLine(filepath.Join(path, "refs", "remotes", t, n))))
}

func head(w http.ResponseWriter, r *http.Request, path string) {
	w.Write([]byte(readFirstLine(filepath.Join(path, "HEAD"))))
}

func blob(w http.ResponseWriter, r *http.Request, path string) {
	repo, _ := gitobj.FromFilesystem(filepath.Join(path, "objects"), "")
	defer repo.Close()

	id, err := hex.DecodeString(chi.URLParam(r, "id"))
	if err != nil {
		return
	}

	o, err := repo.Blob(id)
	if err != nil {
		return
	}

	content, err := ioutil.ReadAll(o.Contents)
	if err != nil {
		return
	}

	w.Write(content)
}

func tree(w http.ResponseWriter, r *http.Request, path string) {
	repo, _ := gitobj.FromFilesystem(".git/objects", "")
	defer repo.Close()

	id, err := hex.DecodeString(chi.URLParam(r, "id"))
	if err != nil {
		return
	}

	o, err := repo.Tree(id)
	if err != nil {
		return
	}

	for _, entry := range o.Entries {
		w.Write([]byte(fmt.Sprintf("%o %s %x\n", entry.Filemode, entry.Name, entry.Oid)))
	}
}

func commit(w http.ResponseWriter, r *http.Request, path string) {
	repo, _ := gitobj.FromFilesystem(filepath.Join(path, "objects"), "")
	defer repo.Close()

	id, err := hex.DecodeString(chi.URLParam(r, "id"))
	if err != nil {
		return
	}

	o, err := repo.Commit(id)
	if err != nil {
		return
	}

	o.Encode(w)
}

func get(f func(http.ResponseWriter, *http.Request, string), path string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		f(w, r, path)
	}
}

func repo(wd string) string {
	if strings.HasSuffix(wd, ".git") != true {
		return filepath.Join(wd, ".git")
	}

	return wd
}
