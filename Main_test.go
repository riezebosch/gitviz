package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	content, err := ReadObject(".git/objects/ae/710a6ef6cd3145a5366d1ed2b2d918d529e88a")
	if err != nil {
		t.Error(err)
	}

	if content != "commit 252\000tree 8709fadf421ff79bea117e1195277253074d9bc8\nparent 26faf771e5f7c50d00a03b47e71b8df52ff8a7a7\nauthor Manuel Riezebosch <mriezebosch@gmail.com> 1555269442 +0200\ncommitter Manuel Riezebosch <mriezebosch@gmail.com> 1555269442 +0200\n\nread commit from test\n" {
		t.Fail()
	}
}
