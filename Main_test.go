package main

import (
	"encoding/hex"
	"testing"

	"github.com/git-lfs/gitobj"
	"github.com/stretchr/testify/assert"
)

func TestObjectToNodeTree(t *testing.T) {
	repo, _ := gitobj.FromFilesystem(".git/objects", "")
	defer repo.Close()

	node, _ := ObjectToNode(repo, "4e84516b47b89c12f2f9bf41f34725ef6ddce099")
	assert.Equal(t, "tree", node.Type)
}

func TestObjectToNodeCommit(t *testing.T) {
	repo, _ := gitobj.FromFilesystem(".git/objects", "")
	defer repo.Close()

	node, _ := ObjectToNode(repo, "9754373abed581d1f4d714a6094f025d8e6cab6f")
	assert.Equal(t, "commit", node.Type)
}

func TestObjectToNodeBlob(t *testing.T) {
	repo, _ := gitobj.FromFilesystem(".git/objects", "")
	defer repo.Close()

	node, _ := ObjectToNode(repo, "eea118847928ac06875446004228e11658bcb789")
	assert.Equal(t, "blob", node.Type)
}

func TestAddEdgesFromTree(t *testing.T) {
	repo, _ := gitobj.FromFilesystem(".git/objects", "")
	defer repo.Close()

	id := "4e84516b47b89c12f2f9bf41f34725ef6ddce099"
	sha, _ := hex.DecodeString(id)

	tree, _ := repo.Tree(sha)
	edges := AddEdgesFromTree(id, tree)
	if edges[0].To != "eea118847928ac06875446004228e11658bcb789" {
		t.Error(edges[0].To)
	}
}

func TestAddEdgesFromCommit(t *testing.T) {
	repo, _ := gitobj.FromFilesystem(".git/objects", "")
	defer repo.Close()

	id := "9754373abed581d1f4d714a6094f025d8e6cab6f"
	sha, _ := hex.DecodeString(id)

	tree, _ := repo.Commit(sha)
	edges := AddEdgesFromCommit(id, tree)
	if edges[0].To != "4e84516b47b89c12f2f9bf41f34725ef6ddce099" {
		t.Error(edges[0].To)
	}
}

func TestFilewalk(t *testing.T) {
	nodes := WalkObjects()
	assert.Contains(t, nodes, "c568e498a51aa3a792956fc3e23d9b239631fbcd")
}
