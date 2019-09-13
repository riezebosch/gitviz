package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"

	"github.com/git-lfs/gitobj"
	"github.com/pkg/browser"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {
	repo, _ := gitobj.FromFilesystem(".git/objects", "")
	defer repo.Close()

	r := chi.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"https://riezebosch.github.io"},
		AllowedMethods: []string{"GET"},
	})
	r.Use(cors.Handler)

	r.Get("/v1/graph/", func(w http.ResponseWriter, r *http.Request) {
		output, _ := json.MarshalIndent(Visit(), "", "   ")
		w.Write(output)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		output, _ := Asset("html/index.html")
		w.Write(output)
	})

	r.Route("/v1/objects", func(r chi.Router) {
		r.Get("/blob/{id}", func(w http.ResponseWriter, r *http.Request) {
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
		})

		r.Get("/tree/{id}", func(w http.ResponseWriter, r *http.Request) {
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
		})

		r.Get("/commit/{id}", func(w http.ResponseWriter, r *http.Request) {
			id, err := hex.DecodeString(chi.URLParam(r, "id"))
			if err != nil {
				return
			}

			o, err := repo.Commit(id)
			if err != nil {
				return
			}

			o.Encode(w)
		})
	})

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}

	browser.OpenURL(fmt.Sprintf("http://localhost:%v", listener.Addr().(*net.TCPAddr).Port))
	panic(http.Serve(listener, r))
}
