package main

import (
	"bytes"
	"compress/zlib"
	"io/ioutil"
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
