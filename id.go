package main

import "strings"

func refID(path string) string {
	if strings.HasPrefix(path, "refs/heads/") {
		return path[11:]
	}

	if strings.HasPrefix(path, "refs/remotes/") {
		return path[13:]
	}

	if strings.HasPrefix(path, "refs/tags/") {
		return path[10:]
	}

	return path
}

func objectID(path string) string {
	return path[13:15] + path[16:]
}
