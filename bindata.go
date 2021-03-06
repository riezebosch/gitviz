// Package main Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// html/favicon.ico
// html/index.html
package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _htmlFaviconIco = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x96\x5b\x68\x1c\x55\x18\xc7\x4f\xb5\x6b\x53\xb1\x4d\xa2\xed\x83\x22\xc9\x26\x73\x3b\x93\x07\x41\x03\x25\xa2\xd0\x1a\x93\xec\xdc\x27\x6d\xc4\xa2\x60\x71\x9f\x7d\xf1\x5d\x5a\x4a\xa9\x22\xbe\x14\xd2\x54\x08\x6a\x10\xda\x87\x0a\xa2\x0f\x45\xf1\x2d\xb6\x5e\xaa\x90\x97\x54\x1f\xa4\x1a\x63\x6b\xbb\x89\xbb\xdd\xec\x65\x76\x76\x67\x26\x97\x23\x67\x76\xa7\x33\x3b\x3b\x3b\xd9\xdd\x6e\x1b\xd0\x7c\xf0\x11\x32\x33\xff\xf3\xfb\x7f\xdf\x39\xe7\x4b\x00\xd8\x01\x1e\x02\x5d\x5d\xf8\x67\x14\xc0\x08\x00\x34\x00\x20\x1a\xad\xfc\xbe\x0f\x00\x25\x02\x00\x04\x00\x74\xe1\xe7\xa0\xfc\xdc\x8e\x08\x68\x22\xa6\x23\x64\x2c\x3d\xd9\xc7\x67\xcf\x00\x30\xbb\xb3\x19\xa5\xad\xe5\x72\x53\xac\x62\x21\x46\x5d\x45\x24\x97\x3f\xdb\xf8\x1a\xd3\x11\x22\x96\x3b\xc7\xaa\x6b\x88\x55\x57\x6d\x3d\x3b\xbe\x81\x48\xae\x70\x16\xbf\x0b\xd7\xce\x45\x08\x3e\x77\x6e\x40\x5d\x43\x8c\xac\xa3\x28\x7f\xe7\x4a\x1f\x97\xf9\x9e\x96\x4b\x88\x1d\x47\x28\x2a\xe4\xa7\xea\xfb\x98\x76\xb5\x8a\x89\x48\x3e\xfb\x1b\x20\xbf\xde\xf5\xe4\xe0\xdc\xa3\xa4\xa0\xfd\x69\xfb\x38\x8c\x10\x29\x14\x02\xd6\x98\xdd\x49\x72\xf9\xa9\x32\xb7\x80\x68\xa5\x84\x28\x51\x5b\xe9\x19\x4b\x0c\xf6\x1e\xfc\x7b\x88\x10\xf5\x1c\xa3\x5a\x88\x55\xd7\xd1\x00\x5e\x43\x32\x3d\xfd\x98\x8b\xe0\xfe\xb0\x8e\xb6\x92\x98\xd5\x17\x4b\x9d\xec\x8b\xa5\x4f\xb2\x13\x98\xab\x15\x49\x51\xbb\x71\xd7\x87\x68\x4e\xda\xf5\xc6\x92\x17\x58\xc5\xac\xd2\xd2\xb2\x5e\xae\x77\xf4\xd6\xf1\xe8\xd8\xf2\x89\x81\xc3\x1b\x88\x18\x49\xbf\xda\x3d\xb8\xd0\x49\xf2\xda\x65\xec\xc1\x5e\x5f\xd0\x3e\xed\x8b\x25\xbf\x64\x55\xcb\xd6\xb8\x59\xb4\xdf\x47\x63\xb7\x8e\x47\x47\x53\xa7\x49\x21\x9f\xe8\x1e\x59\xe9\xec\x7c\xf1\x46\x37\x21\xe4\xbe\x1d\x38\x52\xd6\x13\x5c\xe1\x73\x30\x74\x73\x37\x21\xe4\x67\x98\xf1\x0d\x5b\xe7\x24\x7b\x04\xa1\xfe\xb1\xf4\x7b\x04\x5f\x38\xd6\x33\x9a\x7e\xa1\x9f\xb3\x4e\x90\x92\xb5\xc0\xa8\xeb\xf6\x3b\x42\xb4\x3e\xc6\x5a\xbb\x05\x43\x3f\xee\x26\x04\x7d\x86\x19\x47\x08\xef\x15\xa3\x9a\x88\x14\x8b\x85\xde\xd8\xd2\x21\xfc\x9a\x10\x4a\xa7\xd8\x8a\x67\x5b\x2b\x19\x9f\x80\xde\xc5\x8e\xaa\x2d\x38\xb8\xd8\x41\x08\xc6\x0c\xfe\x06\xfb\x23\x79\xfd\x07\xfc\xd8\xae\x59\xd4\x97\x71\x3f\xf0\x73\x4a\xb0\x6a\xb5\x4e\x3c\x8d\x7d\x94\x2e\x50\x62\x71\xb1\x67\x38\xf9\x1c\xc9\x65\xde\xa6\x44\xf3\x3b\x46\x5d\x2b\x6b\xa5\x10\x6d\x25\x1e\x3f\x70\x7d\xef\xfe\x97\xb2\x04\x29\x94\xde\x75\xfc\xba\xdc\xd9\x50\xad\x27\x76\x10\xbc\x71\x9e\xad\xf4\xd9\xd6\x92\xd7\x77\x35\xa8\x2d\x47\xef\x62\x47\xbf\xa8\x7f\xd6\x2f\xe8\x17\x9b\xd6\xfe\x8f\x02\x56\x66\xe4\x21\xef\x9c\xdc\xd7\xbe\xf5\x69\x29\xff\x2c\x23\xac\x5c\x62\x84\xcc\x17\x50\xd2\x9e\x69\xdf\xca\x0d\xb1\x63\x50\xd2\x90\x37\x19\xb9\x30\xf2\x80\xd8\x63\x01\x6c\x27\x87\xef\x33\x3b\xa8\x6e\x5f\xea\x2f\xdf\x27\x76\x58\xdd\xf6\xdf\x25\x4f\xb6\xb5\x0f\x9b\xb3\xfd\xfc\x22\xce\xb6\xf4\x21\xac\xe7\x64\xec\x36\x22\xb9\x44\x10\xbb\x9c\xca\xbd\x79\xd8\xa4\xee\xd3\x8c\x5c\xd8\xc3\xc8\x7a\x27\x23\xeb\x67\x6a\xd8\x36\xbf\x84\xb3\xa5\xbd\x08\xab\x9b\x12\xee\x5c\xf1\x7f\x4f\xf1\xa9\x5f\x02\xd8\x76\x42\xc5\x68\xaa\x0f\x0d\xec\xf7\x07\x7e\x0d\x23\x17\x3f\xf2\xb3\x61\x99\xed\x64\x43\x7d\x68\xe4\xac\x91\x5c\x62\x85\x91\xf5\x47\x3c\xec\x4e\x92\x4b\x18\x01\x75\x7b\xd2\xc4\x19\xda\x87\x66\xee\x18\xc5\x27\xaf\xb9\xbd\x4f\xfe\xc1\xc8\xa5\x7a\x75\x3b\xec\x50\x0f\xad\xdc\x31\x4f\xfd\x9b\xd5\xed\x49\x0b\x41\xd5\x1a\xf6\xb1\x1b\x99\x6b\x55\x6c\x7c\xf7\x1c\xbd\x7d\x07\x3d\x6c\x8a\xfb\x07\x31\x92\x76\x89\x16\xd2\xd7\x02\xd8\x95\x5c\x1d\xa9\xb0\x9f\x6f\x61\xb6\xd4\xf0\x9d\x9e\x53\x7c\x32\x01\x15\x63\xbf\xf3\x0e\x2a\xe6\x64\x00\xdb\xc9\x03\x34\xb7\xf4\x7b\xb3\xec\x60\xfe\xdd\x9e\x1f\xf3\xef\x2d\xee\x47\x00\x1b\xd1\x62\xe6\x57\x2a\x76\x7b\xb9\x81\x79\x5e\x33\x5b\x30\xd3\x5d\x7f\x69\xdd\xb3\xdf\xaf\xd7\xe7\xaf\x56\x25\x25\xac\xdc\x84\x92\x26\x37\x53\xb7\x73\xd6\xbc\x7c\xa7\x7e\x5a\xcc\xfe\x04\x15\xb3\xea\xbf\x1c\xa8\x5a\xef\x07\xb1\xa1\xba\x86\x53\x28\x9f\xdf\x42\xbc\xb1\xba\xdd\x33\x4e\x71\x89\x0d\x77\x8f\x0d\x3c\x13\xff\xf2\xec\xf9\x2b\x8c\x6c\x9e\xa7\xa5\xdc\xd5\x10\xf6\x1b\x5e\x9f\xae\x87\xf0\xba\x7d\x77\x6c\x07\x54\x8c\x27\x18\x49\xfb\x0a\x2a\x66\x87\xc3\xae\x73\xd6\xea\xb2\x5d\x0f\x7a\xbc\x51\x36\x2d\xa4\x16\xa0\x62\x3c\xec\x70\x2b\xec\x89\x6a\xb6\x9f\x5f\x9f\xed\xf1\xf0\x66\x3d\xb6\x7b\xc7\x96\x35\xa8\x18\x7b\xab\xf6\x59\x31\x8f\xb6\x5a\x77\xad\x87\x62\x3c\xa4\xe7\xb8\x3f\x17\xfd\x1a\x5a\xcc\x5f\x6e\x07\xbb\xca\x43\xdd\x79\x6e\x7c\x58\xcb\xcf\xcc\xb7\xda\xf3\xba\x1e\x94\x52\xdc\x3f\xcf\x69\x31\xf3\x33\x54\xcc\xc7\x9c\x6f\xa0\x6a\xed\x81\xaa\x75\xaa\x5d\x75\xfb\x03\x2a\x46\xdc\x65\x67\xaf\xba\xcf\xcd\xa3\x14\x9f\x4a\xd5\x9b\x2d\xed\x60\x7b\x3c\xbc\xc3\x48\xda\x37\x1e\x76\xcb\x77\xec\xde\xbd\x98\xaf\x6d\x21\xfb\xa9\xad\x62\x57\xf8\x11\xa8\x5a\x6f\x6d\x05\xbb\xca\x87\xba\x1a\xdf\x2a\xb6\xc7\xc3\x04\x2d\x66\xe7\x69\x31\x37\x0f\xd5\xb5\xf1\x07\xc9\xde\x8e\xed\xd8\x8e\xff\x46\xfc\x1b\x00\x00\xff\xff\x3e\xc2\x78\xb2\x76\x19\x00\x00")

func htmlFaviconIcoBytes() ([]byte, error) {
	return bindataRead(
		_htmlFaviconIco,
		"html/favicon.ico",
	)
}

func htmlFaviconIco() (*asset, error) {
	bytes, err := htmlFaviconIcoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "html/favicon.ico", size: 6518, mode: os.FileMode(420), modTime: time.Unix(1568978470, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _htmlIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x59\xff\x8f\xe2\x38\xb2\xff\xbd\xff\x8a\xda\xcc\x9b\x25\x48\xc1\xd8\x21\xe1\x5b\x03\x6f\xdf\xf6\x7e\x95\x66\xdf\x9e\x6e\x57\x7b\xd2\x8e\x46\x5a\x93\x38\xe0\xe9\x24\x8e\x6c\x43\xd3\xd3\xcb\xff\x7e\x2a\x27\x40\xa0\xe9\x51\x6b\xa4\x3b\xa4\x69\xec\xaa\x4f\x95\xab\xca\x65\x57\x99\x99\x7d\x95\xaa\xc4\x3e\x56\x02\xd6\xb6\xc8\x17\x37\xb3\xc3\x97\xe0\xe9\xe2\x06\x60\x66\xa5\xcd\xc5\x62\x25\xed\x56\x7e\x9a\xf5\xeb\x19\xd2\x73\x59\xde\x83\x16\xf9\xdc\x33\x6b\xa5\x6d\xb2\xb1\x20\x13\x55\x7a\x80\xca\xe6\x9e\x2c\xf8\x4a\xf4\x77\xbd\x9a\xb6\xd6\x22\x9b\x7b\x19\xdf\xe2\x94\xc8\x44\x79\x4e\x87\x49\xb4\xac\x6c\x23\x61\xc5\xce\xf6\x3f\xf2\x2d\xaf\xa9\x1e\x18\x9d\xcc\xbd\xb5\xb5\x95\x99\xf6\xfb\x9b\xb2\xba\x5f\x91\x44\x15\xfd\xad\x34\xbd\x52\xd8\x07\xa5\xef\xbf\xc9\xb9\x15\xc6\xf6\x53\x69\x6c\x9b\x4e\x0a\x59\x92\x8f\xc6\x5b\xcc\xfa\xb5\xb2\x93\xc5\xb5\x25\x5f\xae\x35\x31\xc6\x6b\xdc\xb6\x8f\xb9\x30\x6b\x21\xac\xd7\xf6\xc0\x01\xfa\xb5\x7b\x88\xb8\xe4\x21\x07\x5c\xb0\x83\xa5\x4a\x1f\xe1\xc9\xcd\x1d\x4d\xc8\xd5\xda\x4e\x81\x51\xfa\xf6\xf6\x48\x7d\x90\xa9\x5d\x5f\x12\x0b\xae\x57\xb2\x9c\x02\x3d\x91\x2a\x9e\xa6\xb2\x5c\x1d\x69\xfb\x1b\xf7\xf5\xa6\x78\x6c\xcc\x7f\xed\x4a\xa3\x03\xed\xa0\x21\xe5\x96\xb7\x84\x2b\x65\xa4\x95\xaa\x9c\x42\x26\x77\x22\x3d\xc9\xeb\x5a\x69\xcb\x26\xab\xaa\xb3\xf9\x67\xcd\x26\xb1\x28\x4e\xf4\xa5\xda\xf5\x8c\xfc\xe4\x58\x4b\xa5\x53\xa1\x7b\x4b\xb5\xbb\x7d\x9d\x0b\x83\x36\x6d\xc9\x93\xfb\x95\x56\x9b\x32\xed\x25\x2a\x57\x7a\x0a\x2b\x2d\x1e\x4f\x7c\xb5\x15\x3a\xcb\xd5\xc3\x14\x4c\xa2\x55\x9e\x9f\x38\x0d\xfc\x61\x2d\xad\x38\x51\x33\x55\x5a\x34\x4d\x4c\x81\x91\xf0\x60\xf3\x21\x58\xcb\x8d\xb5\xaa\x34\xaf\x89\x57\x2e\x32\x3b\x85\xb0\xed\xb5\x0b\x58\x4b\x27\xc0\xac\xef\xb2\x68\x71\x33\xeb\xd7\x47\x72\x86\x59\x53\x27\xd1\x2c\x95\x5b\x90\xe9\xdc\x3b\xee\x31\x66\x7c\x2a\xb7\x0d\xbb\xd2\xc2\xb1\x71\x03\xbd\xc5\xd7\x6f\xc6\xc3\xe1\xf8\x16\x92\x5c\x26\xf7\xa0\x4a\xe0\x50\xaa\x54\x80\x55\x60\x84\x00\xbb\x16\x90\xa8\xd2\x8a\xd2\x1a\x58\x0b\x2d\x66\xfd\x4a\x8b\x46\x13\xbf\x38\x37\x2b\x69\xd7\x9b\xa5\x3b\x38\x5a\x8a\x4f\x62\xa9\x4c\xb2\xee\xd7\x97\x84\x07\x49\xce\x8d\x99\x7b\x35\xa8\x97\x28\x5d\x0a\xed\x01\xd7\x92\xf7\x72\xbe\xc4\xb3\xf3\x87\x14\x0f\x60\xd4\x46\x27\x02\x2d\xf9\x51\xda\x9f\x36\x4b\x6f\x31\x33\xdb\x55\xbd\x87\x73\x6f\x4c\xbd\x66\x93\xeb\xf1\x56\x8a\x87\x6f\xd5\x6e\xee\x51\xa0\x10\xc6\xee\x9f\x07\x2e\x38\x73\x2f\x93\x79\x3e\x7d\x93\x65\xd9\x6d\xb3\x69\x6e\x8b\x5b\xa1\xe7\x4b\xa3\xf2\x8d\x15\xb7\x87\x9c\x6c\x92\xca\x0d\x8f\x89\xdb\x18\xb9\x96\x69\x2a\xca\xb9\x67\xf5\x46\x78\x8b\x59\xc5\xed\x1a\xd2\xb9\xf7\x0b\x0d\x28\xbc\x63\x2c\x0e\x18\x8b\xe1\x1d\x1b\xd0\x66\x10\x85\x01\x8b\x42\x78\x17\xc6\x34\x40\xcb\xdc\x80\xc2\x9f\xb8\x1b\x28\xdc\x52\xc1\xc2\x31\x19\x04\x8c\x4e\x08\x85\x3b\xc6\x06\x64\x1c\x4c\x26\x64\x04\x8c\x4d\x08\x0d\xc6\x13\x32\x6c\x0f\xef\x58\x18\xe2\x38\x44\x44\x48\x49\x1c\x8c\xc6\x88\x38\x0d\xef\x10\x1d\x06\xa3\x90\x50\x60\xe1\x80\x44\xc1\x68\x48\x06\xed\xe1\x1d\x0b\x47\x64\x10\x8c\x29\x99\x00\x0b\x63\x12\x07\xe3\x91\x43\x1c\x87\x6e\x95\x49\x30\x19\xa1\xea\x01\x25\xc3\x80\x51\x86\xe8\x41\x44\xa2\x80\xd1\x01\x09\x3d\xc0\x08\xcf\xbd\x64\xa3\xb5\x28\xed\x1d\xc6\xf8\x18\x7c\xab\x79\x69\x32\xa5\x8b\x9e\xd2\xd2\x9d\x6f\x36\xa0\xd5\x0e\x18\x1d\x56\xbb\xdb\x63\x3a\xa8\xc4\xaa\x1e\xd7\xc5\xb5\xa8\xb0\x98\xb8\x60\xd6\x51\x89\xc8\xc4\x4d\x18\x30\x36\x26\xa3\x80\xb1\x21\x89\x5d\x5c\xc6\x8e\x1e\x61\xf0\x07\xc8\xa0\xcc\xc5\x60\x30\x44\xfb\x27\x24\x04\x36\x98\xe0\x70\x4c\x22\x60\x51\x48\x42\x1c\x3a\x04\x46\x7a\x3c\x76\x51\x1a\x61\xf0\x22\x87\x40\x6a\x3c\x76\xab\x46\x63\x12\x07\xf1\x00\xc9\x71\x44\x68\x10\x33\x54\x17\x4f\xc8\x08\x87\x88\x18\x52\x32\x08\xa2\x09\x22\x86\x03\x12\x06\xd1\x00\x03\x36\x62\x24\x0a\x22\x4a\x18\xdc\xb5\xc6\x6c\x34\x24\x2c\x88\x42\xb4\x7b\x34\xc6\x55\x86\x24\x84\x3b\x36\x1e\x10\x86\x2b\x0e\x81\x8d\x47\x24\x0c\x86\x8c\x8c\x81\x4d\x28\x99\x04\x43\x74\xec\x8e\x4d\x22\x12\x07\x43\xcc\x0f\x36\x19\x91\x51\x30\x1a\x90\x10\x42\x4a\x09\x0b\x46\xb8\x43\x77\xa1\xcb\x9a\x31\x45\x32\x1b\xe2\xd6\x46\x64\xd2\x1e\xde\x85\x2c\x24\xa3\x60\x32\x20\x0c\x42\xea\x42\x33\x24\x14\x42\x1a\x93\x08\x87\xa8\x83\xc6\x84\x05\x8c\x86\x24\x82\x90\x0e\x30\xf8\x74\xe4\x2c\x71\xe9\xc9\xd0\xee\x3b\x36\x66\xb8\x11\xe1\x18\x73\x61\xe8\x18\xa1\x73\x28\x1e\xb9\x4d\x89\x9c\xcf\xf1\xc8\xed\xd6\x10\x41\xf1\x10\x19\xa1\x4b\xb5\x38\x74\x63\x34\xe8\x1d\x8b\x18\x2e\x31\x18\x3a\xb5\x03\xb7\x8f\x83\x11\x26\x75\xc4\x30\xdd\x22\x97\x6e\x11\x43\x3a\xfe\x85\x3f\xaf\x27\x5c\x3b\x93\xf0\x1a\x3c\xa5\x52\xdf\x6c\x57\x8b\x59\x9f\x2f\xea\xa2\xbb\x20\x67\x97\xcf\x74\x8d\x97\x3c\x90\x43\x0a\x3e\xf1\x52\x16\xdc\xdd\x0c\x48\x4a\xb8\xed\x3d\xf0\xad\x80\x78\x48\x0b\x03\x82\x1b\xd1\x93\x65\x4f\x6d\xec\xfe\x9b\x7b\xf1\x98\x69\x5e\x08\x03\x6d\xe4\x13\x7d\x1b\x60\xe5\x79\x3a\xe6\xfe\x54\x2b\xcb\xad\xf0\x69\x77\x1f\xd2\xb7\xc1\xf0\x1a\xaf\x17\xc6\xa9\x58\x75\xf7\x11\x7d\x1b\x8c\xaf\x01\x18\x75\xfc\xfd\x37\x85\x48\x25\x07\xbf\xe0\xbb\x5e\x5d\xd2\x62\x4a\xab\x5d\xf7\xe9\xd5\x6e\x95\xaa\x14\xfb\x73\xf4\x97\xb8\xbf\x3f\x96\x9f\xb3\x7a\xd3\x14\xb9\xa6\x97\x71\xac\x9a\xe4\xb8\xd5\xfa\xd1\xc8\xc4\x78\xa0\x4a\x57\x6b\xb0\x55\x52\x95\xdf\xf5\x16\xbf\xfd\xfe\xeb\x3f\x66\xfd\x1a\x7a\x5d\x36\xcb\x1e\xd2\x96\x20\x4e\x51\xf0\x87\x1f\xfe\xf5\xdd\xb9\x60\x53\xe7\xea\xf1\x67\xfb\xc8\x1a\x9f\x09\x9b\xac\xfd\x0e\xaf\x64\x5f\x96\x99\xea\x74\x89\x5d\x8b\xd2\xd7\xc2\x54\xaa\x34\x02\xe6\x0b\x38\x8c\xc9\x47\xa3\x4a\xbf\xdb\x20\x5c\xf7\x33\x5f\x40\xaa\x92\x4d\x21\x4a\x4b\x5c\x0b\x0c\x73\x40\x06\x49\xa5\x16\x89\x55\xfa\xb1\x7b\x74\xe7\xa6\x6e\x1d\x4a\x63\x41\x2d\x3f\x8a\xc4\x1a\x98\xc3\x7b\x6f\x99\xab\xa5\x17\x80\x67\xb5\x10\xf8\x9d\xa8\xa2\x90\xd6\xfb\x70\xdb\xc2\x6b\x91\xd5\x60\x2c\xf5\x08\xd2\xa2\x50\xd6\xc1\x2d\x5f\x21\xf6\xa6\xee\x12\xdc\x79\x80\x39\xf8\xe8\x72\x17\xcd\x7b\x6a\x75\x2d\xa8\xaa\xe0\x55\x25\xcb\x15\xcc\xa1\x14\x0f\xf0\x0b\xaf\xfc\xf7\x47\x04\x7e\xde\x77\xd0\xa0\x4e\x00\x9d\x65\xbe\x11\x9d\x0f\xc1\x05\x17\xcd\x44\xee\x4a\x0b\x51\x3e\x67\xd7\xd6\x23\x40\x8b\xf4\x39\x1b\xed\x6f\xa4\x1f\xaf\xe8\xe6\xab\x97\x99\xb5\xcb\x2f\xf3\x7f\xfa\xfe\xff\xbe\xab\xcd\xe6\xc9\x7d\xe7\xc3\x91\xfb\xa1\x7b\xdb\x42\x9e\x9a\x52\x61\x37\xba\x3c\x84\x83\xac\x84\x6d\x62\xf6\xf7\xdf\xd0\xa9\x64\x79\xdf\x69\xb7\x6f\x56\xbd\xc3\x06\x05\x23\xbb\x73\x61\x6d\x76\x90\xc8\x32\xc9\x37\xa9\x30\xfe\x8e\xd4\xe2\xff\x0b\x3b\x22\x53\x62\x36\x4b\x63\xb5\x4f\x03\x18\x77\x61\xea\x68\xb5\x19\xee\xcf\x96\x6b\xd7\x63\x99\x66\x1b\xb6\xd2\x90\xef\xb8\xe5\xbf\x09\xeb\xbf\x3f\x18\x8c\x20\x91\xae\x5e\x06\xb5\x13\xaa\xc2\x53\x8b\xc8\xa7\xcb\x9e\xd7\xc3\xfb\xc8\x0b\x9e\x35\xc8\x97\x74\xbe\xb1\xea\x9f\xa2\xee\x5f\xb1\xc1\x39\x71\x9a\x73\x3b\x6d\xe9\xc6\x8f\x28\xf9\x32\x17\xe9\x25\x1a\x5c\x6b\xad\x4b\x61\x7e\xda\xd8\x4b\x19\xfc\xac\x34\xdf\x4a\xeb\xae\x19\x9e\xdf\xa1\xf9\xbc\xb4\x53\xe8\x31\x4a\x29\x0d\x9e\xc1\x13\x51\x5a\xcd\xf3\x1f\x9d\xd4\xe3\x14\xae\x40\x52\x5e\x54\xee\x41\xc0\x9e\xf3\xf8\x56\xc9\xf4\xd7\xad\xd0\x39\xaf\xae\x02\x4c\xa5\x65\xb9\x7a\x27\xca\x95\x7b\x4e\xc5\x17\xfa\xf7\xe7\xd3\x82\xef\xfe\x10\xb9\x4a\x6a\x53\x48\x7c\xc1\x95\x65\x8b\x7b\xce\xb3\xb2\x10\xc6\x0a\x34\xe2\xe6\x8a\x72\xb7\xd3\x97\xe1\x32\x85\x52\x68\xd5\xf3\x20\x1e\x83\x9f\xf1\xdc\x88\xe7\x5e\x61\x32\x4e\xa1\x83\x8d\xbb\x2c\x37\x6a\x63\x3a\x9f\xf3\xaa\x79\xd0\x78\x6f\xa8\xfb\x78\xe7\xdc\x26\x91\x9e\x79\xcb\xb5\x56\x0f\xcf\x6c\x76\xab\xab\x6b\x54\xf8\x7c\xd2\x1c\x9d\x4e\x78\x2e\x7e\xe0\x78\x7b\xba\x45\x9f\x81\xf6\x37\xd7\x67\x2d\xa7\xdc\xe1\xba\xb4\x01\x1f\x68\xd7\xec\x3a\x78\xef\xde\x73\xde\x4b\xca\xeb\xeb\xe0\xec\xd4\x61\x74\xb9\xc4\x12\x3a\x3f\x95\x81\x95\xb0\xdf\xe7\x02\x87\xdf\x3e\xfe\x9c\xfa\x9d\xe3\x1b\xac\xd3\x6d\x89\xd6\xd5\xe3\x65\x29\xe4\x77\xba\xed\x1a\x70\x78\xad\x9f\x2e\x83\xff\xaf\x29\xfe\xd1\x8c\x00\x9e\x6a\xcf\x83\xe6\xe2\xd8\x07\x87\x9b\xa1\x51\x75\xf8\xc5\x42\x95\x7e\xc7\x95\xd2\x4e\x00\x7e\xc5\x35\x2f\xcc\x45\xb5\x90\xd9\x81\x41\x9c\xca\xf7\xf4\x43\xf7\x3c\x74\xee\x91\x38\xaf\x17\x74\xf7\xe7\x25\xfc\xf6\x0c\x8e\x0a\x9f\xdd\x9a\x88\xad\x2f\xce\xee\x95\x8d\xa9\x6b\xf3\x5f\x58\x9b\x1b\xc9\xfe\xff\x3c\x1d\x45\xf6\x87\x89\x4c\xf7\x7f\x7d\xae\x6c\x63\xe1\x3f\x96\xed\xe6\x29\xeb\x2a\x37\x16\x6a\x59\x96\x42\xff\x2e\x76\x16\xe6\x87\x67\xee\x85\xe1\x7b\x10\xb9\x11\xce\x7e\xac\xc3\x5f\x60\x3c\x8a\x9d\x59\x6e\xfe\xfb\xa6\x1f\x17\x87\xf9\x7c\x0e\x1e\x56\x4b\xef\x35\x66\x23\xf0\x3f\x6a\xe2\xe9\x84\x5d\x16\xe7\xe6\xc7\x8e\x43\xa5\x4b\x65\x96\x09\x2d\xca\x04\xb3\xce\xcf\xa4\x36\x36\x00\x5e\x2a\xbb\xc6\xcc\x4f\x54\x51\x71\x5d\xf7\x3c\x8e\x47\x32\x99\x5b\xa1\x7d\xd7\xa4\x7d\xd5\xe0\x88\x51\x85\xf0\x97\x48\x6a\x04\x7c\x1e\xc0\xb2\xdb\x3d\x2c\x92\x6d\xca\x04\x0f\x0c\xe4\x8a\xa7\x7e\x3b\x40\xad\x46\x71\xa5\x79\xb5\x7e\x7d\xa7\x88\x93\xf3\xc3\x75\x38\x3e\x86\x60\x57\xb3\x15\xfe\xc9\x37\xff\x74\x9e\xba\x01\xa0\x28\x69\x8e\x74\x6d\x29\x2a\xe2\x44\xa6\x6e\x1b\x97\x44\xa6\xdd\x8b\x80\xd6\xf2\x3c\x4d\xdb\x4a\xdb\x7a\xce\x16\x78\x49\x29\x29\x78\xe5\xef\x90\xee\x3f\x81\x4c\xeb\x0e\x26\x00\xf7\x33\xcd\x14\xec\x5a\x1a\xd2\xf4\x44\xfe\xae\x1b\x80\xeb\x7d\x0f\xa0\xe6\x3e\x6d\x40\xae\x25\x3d\xb4\x47\x41\x53\x98\xea\xe9\x99\x1c\xec\x4f\xbb\x70\xf8\xb8\x5b\xec\x4a\x8c\x6a\x7a\x3b\x46\x8e\x72\xee\x4e\xa6\x55\xd1\x38\xe4\x86\x5f\x7f\x0d\x9c\x58\xd5\x90\xac\xba\x0c\x5c\xad\xf4\x5a\xe0\x1a\xe5\x67\xab\x7e\xe9\x4a\xfb\xee\xd9\xef\x82\xc7\x7c\xab\x9f\x42\xad\x1c\x39\x5c\xd5\x4d\xf7\x45\x10\xf0\x9b\x2c\x36\xb9\xeb\x9c\xfc\x0b\x35\xd9\x43\x8a\xa7\xe2\xe2\x12\x7f\x9d\x8e\xeb\x48\xbe\x94\xb9\xfc\x24\xfc\x98\xd2\xcf\xe3\xb4\x7d\xd1\xac\xfa\x10\xd5\x24\x23\xec\xcf\xa5\x15\x7a\xcb\x73\x1f\xe9\x01\x0c\xe8\x51\xf5\xe9\xf7\xf8\x59\xbf\xfe\x1d\x73\xd6\x77\xff\xe1\xf0\xef\x00\x00\x00\xff\xff\x4d\xaa\x7d\xa9\x87\x18\x00\x00")

func htmlIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_htmlIndexHtml,
		"html/index.html",
	)
}

func htmlIndexHtml() (*asset, error) {
	bytes, err := htmlIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "html/index.html", size: 6279, mode: os.FileMode(420), modTime: time.Unix(1582397588, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"html/favicon.ico": htmlFaviconIco,
	"html/index.html":  htmlIndexHtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"html": &bintree{nil, map[string]*bintree{
		"favicon.ico": &bintree{htmlFaviconIco, map[string]*bintree{}},
		"index.html":  &bintree{htmlIndexHtml, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
