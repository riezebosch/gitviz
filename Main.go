package main

import (
	"bufio"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/browser"

	"github.com/git-lfs/gitobj"
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
		nodes, edges := visitAll()
		output, _ := json.MarshalIndent(Data{Nodes: nodes, Edges: edges}, "", "   ")
		w.Write(output)
	})

	browser.OpenURL("https://riezebosch.github.io/gitgraph")
	panic(http.ListenAndServe(":3333", r))
}

func visitAll() ([]Node, []Edge) {
	var nodes = []Node{}
	var edges = []Edge{}
	nodes, edges = visitObjects(nodes, edges)
	nodes, edges = visitRefs(nodes, edges)
	nodes, edges = visitHead(nodes, edges)

	return nodes, edges
}

type Data struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}

type Edge struct {
	From string `json:"from"`
	To   string `json:"to"`
}

func addEdgesFromTree(id string, tree *gitobj.Tree) (edges []Edge) {
	for _, entry := range tree.Entries {
		edges = append(edges, Edge{From: id, To: hex.EncodeToString(entry.Oid)})
	}

	return
}

func addEdgesFromCommit(id string, commit *gitobj.Commit) (edges []Edge) {
	edges = append(edges, Edge{From: id, To: hex.EncodeToString(commit.TreeID)})
	for _, parent := range commit.ParentIDs {
		edges = append(edges, Edge{From: id, To: hex.EncodeToString(parent)})
	}

	return
}

func walkObjects() (nodes []string) {
	var dir = ""
	filepath.Walk(".git/objects", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {

			nodes = append(nodes, fmt.Sprintf("%s%s", dir, info.Name()))
		} else {
			switch info.Name() {
			case "info", "pack":
				return filepath.SkipDir
			}

			dir = info.Name()
		}
		return nil
	})

	return
}

func visitRefs(nodes []Node, edges []Edge) ([]Node, []Edge) {
	filepath.Walk(".git/refs", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			nodes, edges = visitRef(path, nodes, edges)
		}
		return nil
	})

	return nodes, edges
}

func visitRef(path string, nodes []Node, edges []Edge) ([]Node, []Edge) {
	id := refId(path[5:])
	nodes = append(nodes, Node{Id: id, Type: "branch"})

	to := readFirstLine(path)
	edges = append(edges, Edge{From: id, To: to})

	return nodes, edges
}

func refId(path string) string {
	if strings.HasPrefix(path, "refs/heads/") {
		return path[11:]
	}

	if strings.HasPrefix(path, "refs/remotes/") {
		return path[13:]
	}

	if strings.HasPrefix(path, "refs/tags/") {
		return path[10:]
	}

	return path
}

func readFirstLine(path string) string {
	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	return scanner.Text()
}

type Node struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

func visitObjects(nodes []Node, edges []Edge) ([]Node, []Edge) {
	repo, _ := gitobj.FromFilesystem(".git/objects", "")
	defer repo.Close()

	objects := walkObjects()
	for _, id := range objects {
		node, e, _ := visitObject(repo, id)
		nodes = append(nodes, node)
		edges = append(edges, e...)
	}

	return nodes, edges
}

func visitObject(repo *gitobj.ObjectDatabase, id string) (node Node, edges []Edge, err error) {
	sha, _ := hex.DecodeString(id)

	object, err := repo.Object(sha)
	if err != nil {
		return Node{}, edges, err
	}

	switch v := object.(type) {
	case *gitobj.Tree:
		edges = append(edges, addEdgesFromTree(id, v)...)
		return Node{Type: "tree", Id: id}, edges, nil
	case *gitobj.Commit:
		edges = append(edges, addEdgesFromCommit(id, v)...)
		return Node{Type: "commit", Id: id}, edges, nil
	case *gitobj.Blob:
		return Node{Type: "blob", Id: id}, edges, nil
	}

	return Node{}, edges, fmt.Errorf("Unkown object type for sha-ish %s", id)
}

func visitHead(nodes []Node, edges []Edge) ([]Node, []Edge) {
	to := readFirstLine(".git/HEAD")
	if strings.HasPrefix(to, "ref: ") {
		to = refId(to[5:])
	}

	return append(nodes, Node{Id: "HEAD", Type: "branch"}), append(edges, Edge{From: "HEAD", To: to})
}
