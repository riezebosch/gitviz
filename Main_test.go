package main

import (
	"encoding/hex"
	"io/ioutil"
	"testing"

	"github.com/git-lfs/gitobj"
	"github.com/stretchr/testify/assert"
)

func TestContent(t *testing.T) {
	repo, _ := gitobj.FromFilesystem(".git/objects", "")

	id := "eea118847928ac06875446004228e11658bcb789"
	sha, _ := hex.DecodeString(id)
	blob, _ := repo.Blob(sha)

	content, _ := ioutil.ReadAll(blob.Contents)
	assert.Contains(t, string(content), "CommitReferences")
}
