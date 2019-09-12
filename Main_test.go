package main

import (
	"encoding/hex"
	"io/ioutil"
	"testing"

	"github.com/git-lfs/gitobj"
	"github.com/stretchr/testify/assert"
)

func TestVisitObjectTree(t *testing.T) {
	nodes, edges := visitAll()
	assert.Contains(t, nodes, Node{Type: "tree", Id: "4e84516b47b89c12f2f9bf41f34725ef6ddce099"})
	assert.Contains(t, edges, Edge{From: "4e84516b47b89c12f2f9bf41f34725ef6ddce099", To: "eea118847928ac06875446004228e11658bcb789"})
}

func TestVisitObjectCommit(t *testing.T) {
	nodes, edges := visitAll()
	assert.Contains(t, nodes, Node{Type: "commit", Id: "bb4840c0b5dc29bcb7d4e0e2e5d1b9e9dec721e5"})
	assert.Contains(t, edges, Edge{From: "bb4840c0b5dc29bcb7d4e0e2e5d1b9e9dec721e5", To: "c932e6ae3a32d3cdea9dd8043c165e52495ea4c9"})
	assert.Contains(t, edges, Edge{From: "bb4840c0b5dc29bcb7d4e0e2e5d1b9e9dec721e5", To: "3dda9e9c4e40b7e1b743793075850eadf5817ab5"})
}

func TestVisitObjectBlob(t *testing.T) {
	nodes, _ := visitAll()
	assert.Contains(t, nodes, Node{Id: "eea118847928ac06875446004228e11658bcb789", Type: "blob"})
}

func TestVisitBranches(t *testing.T) {
	nodes, edges := visitAll()
	assert.Contains(t, nodes, Node{Id: "for-testing", Type: "branch"})
	assert.Contains(t, edges, Edge{From: "for-testing", To: "627c86822eaa47167417c2c7fc99ef42c599711a"})
}

func TestVisitTags(t *testing.T) {
	nodes, edges := visitAll()
	assert.Contains(t, nodes, Node{Id: "tag-for-testing", Type: "branch"})
	assert.Contains(t, edges, Edge{From: "tag-for-testing", To: "5e0ddee0751a036f9f51585aa7fb7bde5afe5000"})
}

func TestVisitHead(t *testing.T) {
	nodes, edges := visitAll()
	assert.Contains(t, nodes, Node{Id: "HEAD", Type: "branch"})
	assert.Contains(t, edges, Edge{From: "HEAD", To: "master"})
}

func TestVisitAll(t *testing.T) {
	nodes, _ := visitAll()
	assert.Contains(t, nodes, Node{Id: "for-testing", Type: "branch"})
	assert.Contains(t, nodes, Node{Id: "tag-for-testing", Type: "branch"})
}

func TestContent(t *testing.T) {
	repo, _ := gitobj.FromFilesystem(".git/objects", "")

	id := "eea118847928ac06875446004228e11658bcb789"
	sha, _ := hex.DecodeString(id)
	blob, _ := repo.Blob(sha)

	content, _ := ioutil.ReadAll(blob.Contents)
	assert.Contains(t, string(content), "CommitReferences")
}

func TestRefId(t *testing.T) {
	id := refId("refs/heads/master")
	assert.Equal(t, id, "master")
}

func TestRefIdRemote(t *testing.T) {
	id := refId("refs/remotes/origin/master")
	assert.Equal(t, id, "origin/master")
}

func TestRefIdTag(t *testing.T) {
	id := refId("refs/tags/0.1")
	assert.Equal(t, id, "0.1")
}

func TestObjectId(t *testing.T) {
	path := ".git/objects/00/507eabbf76528884df48a1c9fe30434825bf57"
	assert.Equal(t, "00507eabbf76528884df48a1c9fe30434825bf57", objectId(path))
}
