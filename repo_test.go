package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVisitObjectTree(t *testing.T) {
	graph := Visit()
	assert.Contains(t, graph.Nodes, Node{Type: "tree", ID: "4e84516b47b89c12f2f9bf41f34725ef6ddce099"})
	assert.Contains(t, graph.Edges, Edge{From: "4e84516b47b89c12f2f9bf41f34725ef6ddce099", To: "eea118847928ac06875446004228e11658bcb789"})
}

func TestVisitObjectCommit(t *testing.T) {
	graph := Visit()
	assert.Contains(t, graph.Nodes, Node{Type: "commit", ID: "bb4840c0b5dc29bcb7d4e0e2e5d1b9e9dec721e5"})
	assert.Contains(t, graph.Edges, Edge{From: "bb4840c0b5dc29bcb7d4e0e2e5d1b9e9dec721e5", To: "c932e6ae3a32d3cdea9dd8043c165e52495ea4c9"})
	assert.Contains(t, graph.Edges, Edge{From: "bb4840c0b5dc29bcb7d4e0e2e5d1b9e9dec721e5", To: "3dda9e9c4e40b7e1b743793075850eadf5817ab5"})
}

func TestVisitObjectBlob(t *testing.T) {
	graph := Visit()
	assert.Contains(t, graph.Nodes, Node{ID: "eea118847928ac06875446004228e11658bcb789", Type: "blob"})
}

func TestVisitBranches(t *testing.T) {
	graph := Visit()
	assert.Contains(t, graph.Nodes, Node{ID: "for-testing", Type: "branch"})
	assert.Contains(t, graph.Edges, Edge{From: "for-testing", To: "627c86822eaa47167417c2c7fc99ef42c599711a"})
}

func TestVisitRemoteTrackingBranches(t *testing.T) {
	graph := Visit()
	assert.Contains(t, graph.Nodes, Node{ID: "origin/master", Type: "branch"})
}

func TestVisitTags(t *testing.T) {
	graph := Visit()
	assert.Contains(t, graph.Nodes, Node{ID: "tag-for-testing", Type: "branch"})
	assert.Contains(t, graph.Edges, Edge{From: "tag-for-testing", To: "5e0ddee0751a036f9f51585aa7fb7bde5afe5000"})
}

func TestVisitHead(t *testing.T) {
	graph := Visit()
	assert.Contains(t, graph.Nodes, Node{ID: "HEAD", Type: "branch"})
	assert.Contains(t, graph.Edges, Edge{From: "HEAD", To: "master"})
}

func TestVisitAll(t *testing.T) {
	graph := Visit()
	assert.Contains(t, graph.Nodes, Node{ID: "for-testing", Type: "branch"})
	assert.Contains(t, graph.Nodes, Node{ID: "tag-for-testing", Type: "branch"})
}
