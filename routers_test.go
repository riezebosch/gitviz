package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInfo(t *testing.T) {
	assert.Contains(t, request(t, "/api/info"), "\"gitviz\"")
}

func TestGraph(t *testing.T) {
	assert.Contains(t, request(t, "/api/graph"), "nodes")
}

func TestObjectCommit(t *testing.T) {
	assert.Contains(t, request(t, "/api/objects/commit/627c86822eaa47167417c2c7fc99ef42c599711a"), "author Manuel Riezebosch <mriezebosch@gmail.com>")
}

func TestObjectBlob(t *testing.T) {
	assert.Contains(t, request(t, "/api/objects/blob/eea118847928ac06875446004228e11658bcb789"), "package main")
}

func TestIndex(t *testing.T) {
	assert.Contains(t, request(t, ""), "<html>")
}

func TestFavicon(t *testing.T) {
	assert.NotContains(t, request(t, "/favicon.ico"), "404")
}
func TestObjectTree(t *testing.T) {
	assert.Contains(t, request(t, "/api/objects/tree/4e84516b47b89c12f2f9bf41f34725ef6ddce099"), "Main.go")
}

func request(t *testing.T, path string) string {
	ts := httptest.NewServer(Routes())
	defer ts.Close()

	req, err := http.NewRequest("GET", ts.URL+path, nil)
	assert.NoError(t, err)

	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)

	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)

	return string(body)
}
