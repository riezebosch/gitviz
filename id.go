package main

import (
	"path/filepath"
	"strings"
)

func refID(id string) string {
	id = filepath.ToSlash(id)
	id = strings.Replace(id, "ref: ", "", 1)

	if strings.HasPrefix(id, "refs/heads/") {
		return id[11:]
	}

	if strings.HasPrefix(id, "refs/remotes/") {
		return id[13:]
	}

	if strings.HasPrefix(id, "refs/tags/") {
		return id[10:]
	}

	return id
}

func objectID(path string) string {
	return path[13:15] + path[16:]
}
