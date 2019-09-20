## do once
checkout `for-testing` once because it is used in a test  
run `go get -u github.com/go-bindata/go-bindata/...` to get `go-bindata`

## do everytime
run `go-bindata html` to include changes in `bindata.go` or run `go-bindata html --debug` to test changes locally