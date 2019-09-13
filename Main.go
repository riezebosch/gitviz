package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/pkg/browser"
)

func main() {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}

	browser.OpenURL(fmt.Sprintf("http://localhost:%v", listener.Addr().(*net.TCPAddr).Port))
	panic(http.Serve(listener, Routes()))
}
