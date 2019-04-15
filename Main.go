package main

import (
	"bytes"
	"compress/zlib"
	"io/ioutil"
	"strings"
)

func main() {
}

func ReadObject(filename string) (string, error) {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	b := bytes.NewReader(dat)
	r, err := zlib.NewReader(b)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(r)

	r.Close()
	content := buf.String()

	return content, nil
}

func SplitTypeFromContent(content string) (string, string) {
	split := strings.SplitN(content, "\000", 2)
	return strings.SplitN(split[0], " ", 2)[0], split[1]
}

func SplitRefFromContent(content string) (string, string) {
	var line = strings.SplitN(content, "\n", 2)
	if strings.HasPrefix(line[0], "tree") ||
		strings.HasPrefix(line[0], "parent") {
		return line[0], line[1]
	}

	return "", content
}

func CommitReferences(content string) []string {
	var edges []string
	var ref, remainder = SplitRefFromContent(content)
	for ref != "" {
		edges = append(edges, strings.SplitN(ref, " ", 2)[1])
		ref, remainder = SplitRefFromContent(remainder)
	}

	return edges
}
