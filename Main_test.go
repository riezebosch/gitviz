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

func TestSplitTypeFromContent(t *testing.T) {
	input := "commit 252\000tree 8709fadf421ff79bea117e1195277253074d9bc8\nparent 26faf771e5f7c50d00a03b47e71b8df52ff8a7a7\nauthor Manuel Riezebosch <mriezebosch@gmail.com> 1555269442 +0200\ncommitter Manuel Riezebosch <mriezebosch@gmail.com> 1555269442 +0200\n\nread commit from test\n"
	objectType, content := SplitTypeFromContent(input)

	if objectType != "commit" {
		t.Error(objectType)
	}

	if content != "tree 8709fadf421ff79bea117e1195277253074d9bc8\nparent 26faf771e5f7c50d00a03b47e71b8df52ff8a7a7\nauthor Manuel Riezebosch <mriezebosch@gmail.com> 1555269442 +0200\ncommitter Manuel Riezebosch <mriezebosch@gmail.com> 1555269442 +0200\n\nread commit from test\n" {
		t.Error(content)
	}
}

func TestCommitReferences(t *testing.T) {
	edges := CommitReferences("tree 8709fadf421ff79bea117e1195277253074d9bc8\nparent 26faf771e5f7c50d00a03b47e71b8df52ff8a7a7\nauthor Manuel Riezebosch <mriezebosch@gmail.com> 1555269442 +0200\ncommitter Manuel Riezebosch <mriezebosch@gmail.com> 1555269442 +0200\n\nread commit from test\n")

	if len(edges) != 2 {
		t.Error(edges)
	}

	if edges[0] != "8709fadf421ff79bea117e1195277253074d9bc8" {
		t.Error(edges[0])
	}

	if edges[1] != "26faf771e5f7c50d00a03b47e71b8df52ff8a7a7" {
		t.Error(edges[1])
	}
}

func TestSplitTreeFromContent(t *testing.T) {
	ref, content := SplitRefFromContent("tree 8709fadf421ff79bea117e1195277253074d9bc8\nparent 26faf771e5f7c50d00a03b47e71b8df52ff8a7a7\nauthor Manuel Riezebosch <mriezebosch@gmail.com> 1555269442 +0200\ncommitter Manuel Riezebosch <mriezebosch@gmail.com> 1555269442 +0200\n\nread commit from test\n")
	if ref != "tree 8709fadf421ff79bea117e1195277253074d9bc8" {
		t.Error(ref)
	}

	if content != "parent 26faf771e5f7c50d00a03b47e71b8df52ff8a7a7\nauthor Manuel Riezebosch <mriezebosch@gmail.com> 1555269442 +0200\ncommitter Manuel Riezebosch <mriezebosch@gmail.com> 1555269442 +0200\n\nread commit from test\n" {
		t.Error(content)
	}
}

func TestSplitParentFromContent(t *testing.T) {
	ref, content := SplitRefFromContent("parent 26faf771e5f7c50d00a03b47e71b8df52ff8a7a7\nauthor Manuel Riezebosch <mriezebosch@gmail.com> 1555269442 +0200\ncommitter Manuel Riezebosch <mriezebosch@gmail.com> 1555269442 +0200\n\nread commit from test\n")
	if ref != "parent 26faf771e5f7c50d00a03b47e71b8df52ff8a7a7" {
		t.Error(ref)
	}

	if content != "author Manuel Riezebosch <mriezebosch@gmail.com> 1555269442 +0200\ncommitter Manuel Riezebosch <mriezebosch@gmail.com> 1555269442 +0200\n\nread commit from test\n" {
		t.Error(content)
	}
}

func TestSplitContentFromContent(t *testing.T) {
	ref, content := SplitRefFromContent("author Manuel Riezebosch <mriezebosch@gmail.com> 1555269442 +0200\ncommitter Manuel Riezebosch <mriezebosch@gmail.com> 1555269442 +0200\n\nread commit from test\n")
	if ref != "" {
		t.Error(ref)
	}

	if content != "author Manuel Riezebosch <mriezebosch@gmail.com> 1555269442 +0200\ncommitter Manuel Riezebosch <mriezebosch@gmail.com> 1555269442 +0200\n\nread commit from test\n" {
		t.Error(content)
	}
}
