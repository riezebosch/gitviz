package main

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/browser"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "https://riezebosch.github.io"},
		AllowedMethods: []string{"GET"},
	})
	r.Use(cors.Handler)

	r.Get("/v1/graph/", func(w http.ResponseWriter, r *http.Request) {
		output, _ := json.MarshalIndent(Visit(), "", "   ")
		w.Write(output)
	})

	browser.OpenURL("https://riezebosch.github.io/gitgraph")
	panic(http.ListenAndServe(":3333", r))
}
