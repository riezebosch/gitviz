package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRefId(t *testing.T) {
	id := refID("refs/heads/master")
	assert.Equal(t, id, "master")
}

func TestRefIdRemote(t *testing.T) {
	id := refID("refs/remotes/origin/master")
	assert.Equal(t, id, "origin/master")
}

func TestRefIdTag(t *testing.T) {
	id := refID("refs/tags/0.1")
	assert.Equal(t, id, "0.1")
}

func TestObjectId(t *testing.T) {
	path := ".git/objects/00/507eabbf76528884df48a1c9fe30434825bf57"
	assert.Equal(t, "00507eabbf76528884df48a1c9fe30434825bf57", objectID(path))
}
