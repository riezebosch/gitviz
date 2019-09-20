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
