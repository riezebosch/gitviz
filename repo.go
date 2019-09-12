package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/git-lfs/gitobj"
)

// Visit repository and return directed acyclic graph
func Visit() Graph {
	var nodes = []Node{}
	var edges = []Edge{}
	nodes, edges = visitObjects(nodes, edges)
	nodes, edges = visitRefs(nodes, edges)
	nodes, edges = visitHead(nodes, edges)

	return Graph{Nodes: nodes, Edges: edges}
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

func visitRefs(nodes []Node, edges []Edge) ([]Node, []Edge) {
	files, _ := filepath.Glob(".git/refs/**/*")
	for _, file := range files {
		nodes, edges = visitRef(file, nodes, edges)
	}

	return nodes, edges
}

func visitRef(path string, nodes []Node, edges []Edge) ([]Node, []Edge) {
	id := refID(path[5:])
	nodes = append(nodes, Node{ID: id, Type: "branch"})

	to := readFirstLine(path)
	edges = append(edges, Edge{From: id, To: to})

	return nodes, edges
}

func readFirstLine(path string) string {
	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	return scanner.Text()
}

func visitObjects(nodes []Node, edges []Edge) ([]Node, []Edge) {
	repo, _ := gitobj.FromFilesystem(".git/objects", "")
	defer repo.Close()

	objects, _ := filepath.Glob(".git/objects/??/*")
	for _, object := range objects {
		node, e, _ := visitObject(repo, objectID(object))
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
		return Node{Type: "tree", ID: id}, edges, nil
	case *gitobj.Commit:
		edges = append(edges, addEdgesFromCommit(id, v)...)
		return Node{Type: "commit", ID: id}, edges, nil
	case *gitobj.Blob:
		return Node{Type: "blob", ID: id}, edges, nil
	}

	return Node{}, edges, fmt.Errorf("Unkown object type for sha-ish %s", id)
}

func visitHead(nodes []Node, edges []Edge) ([]Node, []Edge) {
	to := readFirstLine(".git/HEAD")
	if strings.HasPrefix(to, "ref: ") {
		to = refID(to[5:])
	}

	return append(nodes, Node{ID: "HEAD", Type: "branch"}), append(edges, Edge{From: "HEAD", To: to})
}
