package main

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRefId(t *testing.T) {
	id := refID("refs/heads/master")
	assert.Equal(t, "master", id)
}

func TestRefIdRemote(t *testing.T) {
	id := refID("refs/remotes/origin/master")
	assert.Equal(t, "origin/master", id)
}

func TestRefIdTag(t *testing.T) {
	id := refID("refs/tags/0.1")
	assert.Equal(t, "0.1", id)
}

func TestObjectId(t *testing.T) {
	path := ".git/objects/00/507eabbf76528884df48a1c9fe30434825bf57"
	assert.Equal(t, "00507eabbf76528884df48a1c9fe30434825bf57", objectID(path))
}

func TestRefIdWindows(t *testing.T) {
	id := refID(filepath.Join("refs", "tags", "0.1"))
	assert.Equal(t, "0.1", id)
}

func TestRefIdWindowsRemote(t *testing.T) {
	id := refID(filepath.Join("refs", "remotes", "origin", "master"))
	assert.Equal(t, "origin/master", id)
}

func TestRefSpecInRefId(t *testing.T) {
	path := "ref: refs/heads/master"
	assert.Equal(t, "master", refID(path))
}
