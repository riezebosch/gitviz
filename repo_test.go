package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVisitObjectTree(t *testing.T) {
	graph := Visit("test-repo.git")
	assert.Contains(t, graph.Nodes, Node{Type: "tree", ID: "9a96926442f477a2b03ae06cfffcab8393748218"})
	assert.Contains(t, graph.Edges, Edge{From: "9a96926442f477a2b03ae06cfffcab8393748218", To: "f07c09068550591ffd7efda7814ec1dfda4a0da8"})
}

func TestVisitObjectCommit(t *testing.T) {
	graph := Visit("test-repo.git")
	assert.Contains(t, graph.Nodes, Node{Type: "commit", ID: "5b4ddc33ef3da7a248025cc228bc9ef7e860740a"})
	assert.Contains(t, graph.Edges, Edge{From: "5b4ddc33ef3da7a248025cc228bc9ef7e860740a", To: "9a96926442f477a2b03ae06cfffcab8393748218"})
	assert.Contains(t, graph.Edges, Edge{From: "5b4ddc33ef3da7a248025cc228bc9ef7e860740a", To: "408a3337cc79ba20939e223697fb6276628992a4"})
}

func TestVisitObjectBlob(t *testing.T) {
	graph := Visit("test-repo.git")
	assert.Contains(t, graph.Nodes, Node{ID: "f07c09068550591ffd7efda7814ec1dfda4a0da8", Type: "blob"})
}

func TestVisitBranches(t *testing.T) {
	graph := Visit("test-repo.git")
	assert.Contains(t, graph.Nodes, Node{ID: "simple-merge", Type: "head"})
	assert.Contains(t, graph.Edges, Edge{From: "simple-merge", To: "6de25b8c5cd0cd49dc40d91e96f8e1cc9c2d07d8"})
}

func TestVisitRemoteTrackingBranches(t *testing.T) {
	graph := Visit("test-repo.git")
	assert.Contains(t, graph.Nodes, Node{ID: "github/master", Type: "remote"})
}

func TestVisitTags(t *testing.T) {
	graph := Visit("test-repo.git")
	assert.Contains(t, graph.Nodes, Node{ID: "R0.1", Type: "tag"})
	assert.Contains(t, graph.Edges, Edge{From: "R0.1", To: "07870fcf1cae67fcee108e7e0bac81a4c69842d0"})
}

func TestVisitHead(t *testing.T) {
	graph := Visit("test-repo.git")
	assert.Contains(t, graph.Nodes, Node{ID: "HEAD", Type: "HEAD"})
	assert.Contains(t, graph.Edges, Edge{From: "HEAD", To: "master"})
}

func TestVisitAll(t *testing.T) {
	graph := Visit("test-repo.git")
	assert.Contains(t, graph.Nodes, Node{ID: "simple-merge", Type: "head"})
	assert.Contains(t, graph.Nodes, Node{ID: "R0.1", Type: "tag"})
}
