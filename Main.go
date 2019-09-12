package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/pkg/browser"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {
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

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}

	browser.OpenURL(fmt.Sprintf("http://localhost:%v", listener.Addr().(*net.TCPAddr).Port))
	panic(http.Serve(listener, r))
}
