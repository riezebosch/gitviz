package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInfo(t *testing.T) {
	assert.Contains(t, request(t, "test-repo.git", "/api/info"), "\"test-repo.git\"")
}

func TestInfoCurrentDir(t *testing.T) {
	assert.Contains(t, request(t, "", "/api/info"), "\"gitviz\"")
}

func TestGraph(t *testing.T) {
	assert.Contains(t, request(t, "test-repo.git", "/api/graph"), "f07c09068550591ffd7efda7814ec1dfda4a0da8")
}

func TestGraphCurrentDir(t *testing.T) {
	assert.Contains(t, request(t, "", "/api/graph"), "\"from\":\"HEAD\"")
}

func TestObjectCommit(t *testing.T) {
	assert.Contains(t, request(t, "test-repo.git", "/api/objects/commit/5b4ddc33ef3da7a248025cc228bc9ef7e860740a"), "author Manuel Riezebosch <mriezebosch@gmail.com>")
}

func TestObjectBlob(t *testing.T) {
	assert.Contains(t, request(t, "test-repo.git", "/api/objects/blob/f07c09068550591ffd7efda7814ec1dfda4a0da8"), "<html>")
}

func TestRef(t *testing.T) {
	assert.Contains(t, request(t, "test-repo.git", "/api/refs/heads/simple-merge"), "6de25b8c5cd0cd49dc40d91e96f8e1cc9c2d07d8")
}

func TestRefTag(t *testing.T) {
	assert.Contains(t, request(t, "test-repo.git", "/api/refs/tags/R0.1"), "07870fcf1cae67fcee108e7e0bac81a4c69842d0")
}

func TestRefRemoteTrackingBranch(t *testing.T) {
	assert.NotEmpty(t, request(t, "test-repo.git", "/api/refs/remotes/github/master"))
}

func TestRefHead(t *testing.T) {
	assert.Contains(t, request(t, "test-repo.git", "/api/refs/HEAD"), "ref: refs/heads/master")
}

func TestIndex(t *testing.T) {
	assert.Contains(t, request(t, "test-repo.git", ""), "<html>")
}

func TestFavicon(t *testing.T) {
	assert.NotContains(t, request(t, "test-repo.git", "/favicon.ico"), "404")
}
func TestObjectTree(t *testing.T) {
	assert.Contains(t, request(t, "test-repo.git", "/api/objects/tree/4e84516b47b89c12f2f9bf41f34725ef6ddce099"), "Main.go")
}

func request(t *testing.T, repo string, path string) string {
	ts := httptest.NewServer(Routes(repo))
	defer ts.Close()

	req, err := http.NewRequest("GET", ts.URL+path, nil)
	assert.NoError(t, err)

	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)

	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)

	return string(body)
}
