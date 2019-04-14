package Main

import (
	"bytes"
	"compress/zlib"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	dat, err := ioutil.ReadFile(".git/objects/26/faf771e5f7c50d00a03b47e71b8df52ff8a7a7")
	if err != nil {
		t.Error(err)
	}

	b := bytes.NewReader(dat)
	r, err := zlib.NewReader(b)
	if err != nil {
		t.Error(err)
	}
	io.Copy(os.Stdout, r)
	r.Close()
}
