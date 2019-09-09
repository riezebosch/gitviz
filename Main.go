package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/git-lfs/gitobj"
)

func main() {
	objects := walkObjects()
	nodes, edges := visitObjects(objects[:])
	output, _ := json.MarshalIndent(Data{Nodes: nodes, Edges: edges}, "", "   ")
	fmt.Print(string(output))
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

type Node struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

func visitObjects(objects []string) (nodes []Node, edges []Edge) {
	repo, _ := gitobj.FromFilesystem(".git/objects", "")
	defer repo.Close()

	for _, id := range objects {
		node, e, _ := visitObject(repo, id)
		nodes = append(nodes, node)
		edges = append(edges, e...)
	}

	return
}

func visitObject(repo *gitobj.ObjectDatabase, id string) (node Node, edges []Edge, err error) {
	sha, _ := hex.DecodeString(id)

	object, err := repo.Object(sha)
	if err != nil {
		return Node{}, edges, err
	}

	switch v := object.(type) {
	case *gitobj.Tree:
		edges = append(addEdgesFromTree(id, v))
		return Node{Type: "tree", Id: id}, edges, nil
	case *gitobj.Commit:
		edges = append(addEdgesFromCommit(id, v))
		return Node{Type: "commit", Id: id}, edges, nil
	case *gitobj.Blob:
		return Node{Type: "blob", Id: id}, edges, nil
	}

	return Node{}, edges, fmt.Errorf("Unkown object type for sha-ish %s", id)
}
