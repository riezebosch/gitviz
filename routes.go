package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/render"

	"github.com/git-lfs/gitobj"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

// Routes for graph and content
func Routes() *chi.Mux {
	r := chi.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"https://riezebosch.github.io"},
		AllowedMethods: []string{"GET"},
	})
	r.Use(cors.Handler)

	r.Route("/", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			output, _ := Asset("html/index.html")
			w.Write(output)
		})
		r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
			output, _ := Asset("html" + r.URL.EscapedPath())
			w.Write(output)
		})
	})

	r.Route("/api", func(r chi.Router) {
		r.Get("/info", getInfo)
		r.Get("/graph", getGraph)
		r.Route("/objects", func(r chi.Router) {
			r.Get("/blob/{id}", getBlob)
			r.Get("/tree/{id}", getTree)
			r.Get("/commit/{id}", getCommit)
		})
	})

	return r
}

// Info from current repo
type Info struct {
	Directory string `json:"directory"`
}

func getInfo(w http.ResponseWriter, r *http.Request) {
	directory, _ := os.Getwd()
	render.JSON(w, r, Info{Directory: filepath.Base(directory)})
}

func getGraph(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, Visit())
}

func getBlob(w http.ResponseWriter, r *http.Request) {
	repo, _ := gitobj.FromFilesystem(".git/objects", "")
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

func getTree(w http.ResponseWriter, r *http.Request) {
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

func getCommit(w http.ResponseWriter, r *http.Request) {
	repo, _ := gitobj.FromFilesystem(".git/objects", "")
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
