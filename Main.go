package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"

	"github.com/git-lfs/gitobj"
)

func main() {
}

type Edge struct {
	From string
	To   string
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
