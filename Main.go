package main

import (
	"encoding/hex"

	"github.com/git-lfs/gitobj"
)

func main() {
}

type Edge struct {
	from string
	to   string
}

func AddEdgesFromTree(id string, tree *gitobj.Tree) (edges []Edge) {
	for _, entry := range tree.Entries {
		edges = append(edges, Edge{from: id, to: hex.EncodeToString(entry.Oid)})
	}

	return
}

func AddEdgesFromCommit(id string, commit *gitobj.Commit) (edges []Edge) {
	edges = append(edges, Edge{from: id, to: hex.EncodeToString(commit.TreeID)})
	return
}
