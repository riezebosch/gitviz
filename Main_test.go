package main

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/git-lfs/gitobj"
)

func Test(t *testing.T) {
	repo, err := gitobj.FromFilesystem(".git/objects", "")
	if err != nil {
		panic(err)
	}
	defer repo.Close()

	sha, _ := hex.DecodeString("4e84516b47b89c12f2f9bf41f34725ef6ddce099")

	object, err := repo.Object(sha)
	if err != nil {
		t.Error(err)
	}

	switch v := object.(type) {
	case *gitobj.Tree:
		var entry = v.Entries[0]
		if entry.Name != "Main.go" {
			t.Error(entry)
		}

		fmt.Printf("%x", entry.Oid)
	}
}

func TestAddEdgesFromTree(t *testing.T) {
	repo, err := gitobj.FromFilesystem(".git/objects", "")
	if err != nil {
		t.Error(err)
	}
	defer repo.Close()

	id := "4e84516b47b89c12f2f9bf41f34725ef6ddce099"
	sha, _ := hex.DecodeString(id)

	tree, _ := repo.Tree(sha)
	edges := AddEdgesFromTree(id, tree)
	if edges[0].to != "eea118847928ac06875446004228e11658bcb789" {
		t.Error(edges[0].to)
	}
}

func TestAddEdgesFromCommit(t *testing.T) {
	repo, err := gitobj.FromFilesystem(".git/objects", "")
	if err != nil {
		t.Error(err)
	}
	defer repo.Close()

	id := "9754373abed581d1f4d714a6094f025d8e6cab6f"
	sha, _ := hex.DecodeString(id)

	tree, _ := repo.Commit(sha)
	edges := AddEdgesFromCommit(id, tree)
	if edges[0].to != "4e84516b47b89c12f2f9bf41f34725ef6ddce099" {
		t.Error(edges[0].to)
	}
}
