package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/git-lfs/gitobj"
)

func main() {
}

type Edge struct {
	From string `json:"from"`
	To   string `json:"to"`
}

func AddEdgesFromTree(id string, tree *gitobj.Tree) (edges []Edge) {
	for _, entry := range tree.Entries {
		edges = append(edges, Edge{From: id, To: hex.EncodeToString(entry.Oid)})
	}

	return
}

func AddEdgesFromCommit(id string, commit *gitobj.Commit) (edges []Edge) {
	edges = append(edges, Edge{From: id, To: hex.EncodeToString(commit.TreeID)})
	return
}

func WalkObjects() (nodes []string) {
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

func ObjectToNode(repo *gitobj.ObjectDatabase, id string) (Node, error) {
	sha, _ := hex.DecodeString(id)

	object, err := repo.Object(sha)
	if err != nil {
		return Node{}, err
	}

	switch object.(type) {
	case *gitobj.Tree:
		return Node{Type: "tree", Id: id}, nil
	case *gitobj.Commit:
		return Node{Type: "commit", Id: id}, nil
	case *gitobj.Blob:
		return Node{Type: "blob", Id: id}, nil
	}

	return Node{}, errors.New(fmt.Sprintf("Unkown object type for sha-ish %s", id))
}
