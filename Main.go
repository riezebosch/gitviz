package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/pkg/browser"
)

func main() {
	port := flag.Int("port", 0, "")
	flag.Usage = func() {
		fmt.Println(fmt.Sprintf("%s [flags] [repository]", os.Args[0]))
		fmt.Println()
		fmt.Println("flags:")
		flag.PrintDefaults()
	}
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", *port))
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf("http://localhost:%v", listener.Addr().(*net.TCPAddr).Port)
	fmt.Print(url)

	browser.OpenURL(url)
	path := flag.Arg(0)

	panic(http.Serve(listener, Routes(path)))
}
