package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"

	"github.com/git-lfs/gitobj"
)

// Visit repository and return directed acyclic graph
func Visit(path string) Graph {
	var nodes = []Node{}
	var edges = []Edge{}
	nodes, edges = visitObjects(path, nodes, edges)
	nodes, edges = visitRefs(path, nodes, edges)
	nodes, edges = visitHead(path, nodes, edges)

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

func visitRefs(root string, nodes []Node, edges []Edge) ([]Node, []Edge) {
	filepath.Walk(root+"/refs/heads", func(path string, info os.FileInfo, e error) error {
		if e != nil {
			return e
		}

		if info.Mode().IsRegular() {
			nodes, edges = visitRef(root, path, nodes, edges, "head")
		}

		return nil
	})

	var files, _ = filepath.Glob(root + "/refs/remotes/*/*")
	for _, file := range files {
		nodes, edges = visitRef(root, file, nodes, edges, "remote")
	}

	files, _ = filepath.Glob(root + "/refs/tags/*")
	for _, file := range files {
		nodes, edges = visitRef(root, file, nodes, edges, "tag")
	}

	return nodes, edges
}

func visitRef(path string, file string, nodes []Node, edges []Edge, t string) ([]Node, []Edge) {
	id := refID(file[len(path)+1:])
	nodes = append(nodes, Node{ID: id, Type: t})

	to := refID(readFirstLine(file))
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

func visitObjects(path string, nodes []Node, edges []Edge) ([]Node, []Edge) {
	repo, _ := gitobj.FromFilesystem(filepath.Join(path, "objects"), "")
	defer repo.Close()

	objects, _ := filepath.Glob(path + "/objects/??/*")
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

func visitHead(path string, nodes []Node, edges []Edge) ([]Node, []Edge) {
	to := refID(readFirstLine(filepath.Join(path, "HEAD")))
	return append(nodes, Node{ID: "HEAD", Type: "HEAD"}), append(edges, Edge{From: "HEAD", To: to})
}
