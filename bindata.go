// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
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
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
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

// Mode return file modify time
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

	info := bindataFileInfo{name: "html/favicon.ico", size: 6518, mode: os.FileMode(420), modTime: time.Unix(1686133228, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _htmlIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x3a\x7f\x6f\xe3\xb6\x92\xff\xe7\x53\x4c\xb5\xf7\x6a\x19\x90\x69\x49\xfe\xed\xd8\xb9\x7b\x4d\xb7\xb7\x0f\xb7\xfb\xde\xa1\x5b\xb4\x40\x17\x0b\x94\x96\x28\x8b\x1b\x89\x14\x48\xda\xb1\x77\x9b\xef\x7e\x18\x4a\x76\x64\x4b\xce\x26\x2d\x8a\x7b\x01\xda\xa5\x86\x33\xc3\xf9\x3d\x43\x26\x8b\x6f\x62\x19\x99\x7d\xc1\x20\x35\x79\x76\x73\xb5\x28\xff\xb9\x5a\xa4\x8c\xc6\x37\x57\x00\x00\x0b\xc3\x4d\xc6\x6e\xd6\xdc\x6c\xf9\xe7\x45\xbf\xfc\x2a\x77\x32\x2e\xee\x40\xb1\x6c\xe9\xe8\x54\x2a\x13\x6d\x0c\xf0\x48\x0a\x07\x90\xe3\xd2\xe1\x39\x5d\xb3\xfe\xae\x57\xc2\x52\xc5\x92\xa5\x93\xd0\x2d\x7e\x12\x1e\x49\xa7\xe2\xa2\x23\xc5\x0b\x53\xd1\x18\xb6\x33\xfd\x4f\x74\x4b\x4b\xa8\x03\x5a\x45\x4b\x27\x35\xa6\xd0\xf3\x7e\x7f\x23\x8a\xbb\x35\x89\x64\xde\xdf\x72\xdd\x13\xcc\xdc\x4b\x75\xd7\xd7\x86\x8a\x98\x66\x52\xb0\xfe\x26\x8f\xeb\x5b\x24\xe7\x82\x7c\xd2\xce\xcd\xa2\x5f\xf2\x3b\x1c\x69\xf6\x19\xab\x9f\x18\x69\x5d\x89\x83\x3f\x68\x04\xef\xf8\xb5\x92\xf1\x1e\xbe\x1c\x3f\x2d\x02\xe3\xeb\xd4\xcc\x21\xf0\xfd\xbf\x5d\x9f\xec\xdc\xf3\xd8\xa4\x6d\x1b\x39\x55\x6b\x2e\xe6\xe0\x9f\x82\x0b\x1a\xc7\x5c\xac\x1b\xf0\x44\x0a\xd3\x4b\x68\xce\xb3\xfd\x1c\xfe\xae\x38\xcd\x3c\x78\xc3\xb2\x2d\x33\x3c\xa2\x1e\x68\x2a\x74\x4f\x33\xc5\x93\x47\xb2\x87\xab\xe3\xf2\x55\xbe\xaf\x2c\xf0\x52\xc1\x27\x75\x78\x9d\x63\x4c\x0d\x3d\x63\x56\x48\xcd\x0d\x97\x62\x0e\x09\xdf\xb1\xf8\x94\x9f\x2a\x0f\x3a\x53\xcb\xc8\xa2\x01\xfb\xaa\x65\xc8\x88\xe5\xa7\x7b\x2b\xb9\xeb\x69\xfe\xd9\x6e\xaf\xa4\x8a\x99\xea\xad\xe4\xee\xfa\x65\xaa\x0e\xce\xe1\x2b\x1a\xdd\xad\x95\xdc\x88\xb8\x17\xc9\x4c\xaa\x39\xac\x15\xdb\x9f\xe2\xc8\x2d\x53\x49\x26\xef\xe7\xa0\x23\x25\xb3\xec\x74\xb7\x22\xbb\x4f\xb9\x61\x2d\xfe\xd4\xfc\x33\x9b\x43\x40\xc2\xba\x3e\x75\x23\xaf\x36\xc6\x48\xa1\x5f\x62\xe7\x8c\x25\x66\x0e\xe1\xb9\x85\xac\xa1\x2f\x9e\x93\xf0\xcc\x30\x75\x76\x4c\xc6\xb5\xe9\xd9\xcc\xe8\x61\x66\xcc\x41\x48\xc1\xbe\x1e\xad\x35\xbe\xff\x95\xb3\x98\x53\xb4\x0c\x63\x02\xa8\x88\xc1\xcd\xe9\xae\x57\xd9\x7b\xec\xfb\xc5\xae\x7b\x76\xe8\xc5\x48\x85\xa7\x92\xa9\x76\x26\x5c\x08\x4e\xfc\x89\xb9\x2e\x32\xba\x6f\xd3\xe4\x9c\x43\xbb\xe5\x5f\xc6\x84\xac\xb9\x49\x37\xab\x5e\x24\x95\x68\x58\xf7\xab\xac\x4e\x57\x8b\xbe\xf5\xc4\xcd\xd5\xa2\x5f\x96\xe2\xab\x05\xd6\xa1\xaa\x7e\xc5\x7c\x0b\x3c\x5e\x3a\x47\xe3\x61\x85\x8b\xf9\xb6\xda\x2e\x14\xb3\xdb\x68\x16\xe7\xe6\xdb\x57\xd3\xf1\x78\x7a\x0d\x51\xc6\xa3\x3b\x90\x02\x28\x08\x19\x33\x30\x12\x34\x63\x60\x52\x06\x91\x14\x86\x09\xa3\x21\x65\x8a\x2d\xfa\x85\x3a\x54\x78\x5a\x95\xed\x43\x01\x2e\x35\xb4\x15\x58\x71\xf6\x99\xad\xa4\x8e\xd2\x7e\xd9\x1b\x1c\x88\x32\xaa\xf5\xd2\x39\x31\x83\x03\x54\x71\xda\xcb\xe8\x0a\xfb\xc4\xcf\x9c\xdd\x83\x96\x1b\x15\x31\x94\xe4\xbf\xb9\x79\xb3\x59\x39\x37\x0b\xbd\x5d\x37\x53\x74\xe9\x4c\x7d\xa7\xca\xe3\x72\xbd\xe5\xec\xfe\x3b\xb9\x5b\x3a\x3e\xf8\x10\x8e\xec\x7f\xce\x09\xa1\xb5\xda\xd2\x49\x78\x96\xcd\x5f\x25\x49\x72\x5d\xe5\xa4\xcd\xe4\x5a\x26\xd1\x95\x96\xd9\xc6\xb0\xeb\x43\x49\xaa\xea\x88\x5d\x1e\x6b\x57\x25\x7c\xca\xe3\x98\x89\xa5\x63\xd4\x86\xd5\xda\x44\x69\x6b\x6a\x52\x88\x97\xce\x3b\xdf\xf3\xe1\x6d\x10\x8c\xbc\x20\x18\xc1\xdb\x60\xe0\x57\x8b\x61\xe8\x05\xc3\x10\xde\x86\x23\xdf\x43\x91\xed\xc2\x87\x5f\xd1\x65\x48\xdc\xc2\xaf\x19\x38\x4b\xe7\x5d\x10\x4e\xc9\xc0\x0b\xfc\x19\xf1\xe1\x36\x08\x06\x64\xea\xcd\x66\x64\x02\x41\x30\x23\xbe\x37\x9d\x91\x71\x7d\x79\x1b\x84\x21\xae\x43\xc4\x08\x7d\x32\xf2\x26\x53\xc4\x78\x5c\xde\x22\x76\xe8\x4d\x42\xe2\x43\x10\x0e\xc8\xd0\x9b\x8c\xc9\xa0\xbe\xbc\x0d\xc2\x09\x19\x78\x53\x9f\xcc\x20\x08\x47\x64\xe4\x4d\x27\x16\xe3\xb8\xb4\xa7\xcc\xbc\xd9\x04\x59\x0f\x7c\x32\xf6\x02\x3f\x40\xec\xc1\x90\x0c\xbd\xc0\x1f\x90\xd0\x69\x28\x83\xbe\x59\x3a\xd1\x46\x29\x26\xcc\x2d\x7a\xc7\x39\xb8\xcd\x28\x2a\x74\x22\x55\xde\x93\x8a\xdb\xa6\x10\x0c\xfc\x62\x07\x81\x3f\x2e\x76\xd7\xc7\x00\x93\x91\x91\x3d\xaa\xf2\x17\x9b\x30\x18\x11\xeb\x96\xd2\x84\x43\x32\xb3\x1f\x01\x04\xc1\x94\x4c\xbc\x20\x18\x93\x91\x35\xe2\xd4\xc2\x87\xe8\xc6\x01\x6e\xf8\x81\x35\xd8\x60\x8c\xca\xce\x48\x08\xc1\x60\x86\xcb\x29\x19\x42\x30\x0c\x49\x88\x4b\x8b\x81\x6e\x99\x4e\xad\x49\x27\x68\xe9\xa1\xc5\x40\xe8\x68\x6a\x4f\x1d\x4e\xc9\xc8\x1b\x0d\x10\x3c\x1a\x12\xdf\x1b\x05\xc8\x6e\x34\x23\x13\x5c\x22\xc6\xd8\x27\x03\x6f\x38\x43\x8c\xf1\x80\x84\xde\x70\x80\xd6\x9d\x04\x64\xe8\x0d\x7d\x12\xc0\x6d\x6d\x1d\x4c\xc6\x24\xf0\x86\x21\xca\x3d\x99\xe2\x29\x63\x12\xc2\x6d\x30\x1d\x90\x00\x4f\x1c\x43\x30\x9d\x90\xd0\x1b\x07\x64\x0a\xc1\xcc\x27\x33\x6f\x8c\x8a\xdd\x06\xb3\x21\x19\x79\x63\x0c\xa6\x60\x36\x21\x13\x6f\x32\x20\x21\x84\xbe\x4f\x02\x6f\x82\xee\xbc\x0d\x6d\x88\x4d\x7d\x04\x07\x63\x8c\x83\x21\x99\xd5\x97\xb7\x61\x10\x92\x89\x37\x1b\x90\x00\x42\xdf\x9a\x66\x4c\x7c\x08\xfd\x11\x19\xe2\x12\x79\xf8\x23\x12\x78\x81\x1f\x92\x21\x84\xfe\x00\x8d\xef\x4f\xac\x24\x36\x96\x03\x94\xfb\x36\x98\x06\xe8\x88\x70\x8a\x81\x33\xb6\x1b\xa1\x55\x68\x34\xb1\x4e\x19\x5a\x9d\x47\x13\xeb\xad\x31\x22\x8d\xc6\xb8\x11\xda\xb8\x1c\x85\x76\x8d\x02\xbd\x0d\x86\x01\x1e\x31\x18\x5b\xb6\x03\xeb\xc7\xc1\x04\x33\x60\x18\x60\x6c\x0e\x6d\x6c\x0e\x03\x84\xe3\xff\xe1\xd7\xe7\x45\x67\x3d\xec\xb0\x0a\x37\xe3\x6e\xd1\xd7\xdb\xf5\xcd\xa2\x4f\xeb\x03\xe6\xe3\xf6\x69\x67\x98\xa7\x38\x45\x00\x39\x84\xf1\x59\xa7\xa0\x82\xe7\xb4\xac\x52\x88\x11\x51\xd3\xbb\xa7\x5b\x06\xa3\xb1\x9f\x6b\x60\x54\xb3\x1e\x17\x3d\xb9\x31\xad\x2d\xf8\x8e\xed\x13\x45\x73\xa6\x4f\x89\xbf\x9c\xb6\x2a\xff\x6f\xde\xc9\x37\xb6\xd8\x96\x86\x75\xcc\xc7\x39\x28\x69\xa8\x61\xae\xdf\x7d\xaa\x03\x86\xe7\x7c\xc7\xcf\x64\xdb\x0b\x47\x31\x5b\x3f\xc9\x7b\x78\xce\x7b\xfa\x4c\xde\x81\xdf\xc2\xba\xe5\x90\x6a\x7a\xa9\x8d\x2c\xa3\xb6\x89\xe5\x45\xbe\x84\x53\x7f\x62\xdf\x7f\xc1\x04\xf1\x4c\xae\xcf\x8c\x92\x73\xc5\xa1\x3e\x66\x40\x7d\xac\xa8\x86\xa1\x5a\xab\x5b\x1c\x87\x8b\x23\xa4\x44\xb2\xf8\xda\xc8\xc2\x01\x29\xec\x84\x51\x7e\xba\x5d\xe7\xe6\xfd\x4f\xff\xfa\xdf\x45\xbf\xc4\xbb\x4c\x9c\x24\xf7\x71\x8d\x18\x3f\x91\xf8\x87\x1f\x7e\xf9\xbe\x49\x5c\x9b\x72\xec\xe7\x26\x2b\x79\xd8\x79\xf6\xbc\x35\x67\xfc\xa6\x61\xb6\x85\x1d\x44\x9a\x70\xbb\xc7\x45\xb1\x39\x5c\x44\xa3\x94\x45\x77\x2b\xb9\x73\x40\xd0\x9c\x1d\x8f\x80\x2d\xcd\x36\xb8\x2d\xf3\x9c\x1b\x07\x2c\x1a\x8b\xa1\x7f\x81\xa5\x2e\xa8\x38\x74\xb7\xea\x76\xa0\x58\xec\xdc\x94\xf4\x8b\x3e\xee\xb7\x48\xd9\x6f\x11\x73\xd1\x3f\x57\xe8\x2f\xd5\xd0\x28\xc6\xfe\x90\x7e\x6b\x9c\xfe\x9d\x1b\xa4\xff\x77\xd6\x6f\x95\xc9\xd5\x1f\xd2\x6f\x95\xe1\x14\x88\xe4\x7f\x42\xbd\x45\x7f\x93\x3d\x99\x5e\x17\xb4\xbc\xa0\x21\x26\x41\xca\x99\xa2\x2a\x4a\xf7\x47\x15\x6b\x90\x27\xf4\xb4\x3a\xde\x1c\x71\xdb\x94\x6a\x28\x54\xbf\x6d\x94\xcb\x67\xbc\xe5\x3c\x92\x27\xcc\x44\xa9\xdb\xa1\x05\xef\x73\x91\xc8\xce\x69\x75\x26\x26\x65\xc2\x55\x4c\x17\x52\x68\x06\xcb\x1b\x38\xac\xc9\x27\x2d\x85\xdb\x6d\x43\xb7\x57\xc0\xe5\x0d\xc4\x32\xda\xe4\x4c\x18\x62\x1f\xaa\x60\x09\xb8\x41\x62\xae\x58\x64\xa4\xda\x77\x1f\x0b\x6e\x24\x85\x36\x20\x57\x9f\x58\x64\x34\x2c\xe1\x43\x19\x12\x1e\x94\xa1\xef\xc1\x21\xc9\x3f\x5e\x9f\xd1\x28\x96\x94\x04\x78\x35\x43\x44\xc5\x72\x69\x2c\x89\xa1\x6b\xc4\x3f\x23\xa8\xee\xdb\xcb\xb3\x3a\x5e\xf2\x9f\x03\xde\x2b\x4e\x7b\x1b\x4a\xd0\x06\x47\x09\xdb\xe0\x28\x48\x1b\xfc\xcd\xeb\xbf\x7f\xdf\x06\x2f\x05\x6e\x3d\x99\xae\x4b\xf0\x63\xa7\x68\xa8\x63\x64\x8b\x2a\x36\x37\x5c\xf4\x7c\x17\xfd\xd0\xec\x58\x25\x6d\x4e\x8b\x82\x8b\x35\x2c\x41\xb0\x7b\x78\x47\x0b\xf7\x43\x6b\xe6\x7d\xe8\xa0\xae\x1d\x0f\x3a\x98\x6f\x9d\x8f\xde\x05\x2c\xb4\x14\x62\xd9\xaa\x73\x19\xad\x34\x35\x22\x2a\x16\x5f\x46\x43\x3b\x56\xdc\xf6\x4f\x9c\x49\xd7\x5f\x47\x2a\x4d\xfc\x75\x3c\x74\x51\xa9\x26\x8d\xee\x3a\x1f\x1b\x58\x1f\xbb\xd7\x57\x0d\xa0\x62\x66\xa3\xc4\xc1\x98\x64\xcd\x4c\x65\xf9\xdf\x7f\x87\x4e\xc1\xc5\x5d\xe7\xb4\xed\x7b\xa7\x2c\x6c\x36\xcf\xc1\xdd\x59\x57\x55\x29\x40\xb8\x88\xb2\x4d\xcc\xb4\xbb\x23\x25\xb3\xff\x84\x1d\xe1\x31\xd1\x9b\x95\x36\xca\xf5\x3d\x98\x76\x61\x6e\x61\x6d\xf3\x53\xe9\x5f\x21\x63\xa6\x2b\xef\x6e\xb9\x26\xdf\x53\x43\xdf\x33\xe3\x7e\x40\x3d\x5a\x50\xed\x8b\xc0\x29\x3a\x82\x5c\xbb\xe9\xc1\x97\x2a\x77\xe6\x60\x21\x56\xde\x12\xf2\x01\xbf\xad\xa0\x1f\xe1\xa1\x7b\xc6\x9a\xc5\xeb\xcb\x52\x9c\x57\x80\x02\xc7\x28\xdd\x08\xe9\xea\xf5\xc9\xc1\xd9\xd8\x39\x4f\xb7\xf2\x85\xa0\x6d\x8f\x6e\x8c\xfc\x91\x95\x4f\x7d\xcd\xf4\xca\xe8\x5e\x6e\xcc\xbc\x25\x3d\x0e\x05\x98\x47\x34\x6b\xdb\xc7\x1f\x2d\x95\x79\xc7\x4c\x2a\xe3\x39\x74\xca\x9a\xc6\xe2\x4e\x7b\x64\xe9\x94\xde\xb1\x9f\xe4\x3d\x55\xb1\x9e\x43\x47\x49\x69\xf4\x05\xd4\x92\x93\x9d\x25\x3b\x3f\xbe\xbd\x80\xb4\xca\x64\x74\xf7\x3e\xe5\x89\xb1\x2f\x80\x09\xcd\x34\x6b\xc7\x44\xe3\xbf\xe3\x82\xe7\xfc\x73\x35\xa1\x3e\x81\x5c\x50\x7b\xd1\x62\xc2\x28\x9a\x9d\x12\x34\xf0\x1f\xce\x83\xfa\x84\x51\xba\xd7\x3c\xd2\xf3\xf3\xab\x0e\xd8\x27\xe6\xdd\xcf\x2c\x93\x11\x37\x7b\xfb\x9e\xdc\x14\x25\xe7\xa2\x8e\x11\x36\x31\x0c\xcf\x99\x36\xac\x98\x43\xf0\x94\x18\x36\xf2\xda\xfc\x57\x55\x48\xe7\x95\x6f\x7f\x9c\xe6\x09\x55\xc0\xb5\xca\x47\x95\x92\xf7\xad\x7c\xad\x6c\xf2\xd2\x8e\x15\x49\xd0\x55\xc6\x5a\xfb\x43\xfd\x47\x47\x34\x63\x3f\x50\x6c\x93\x56\x88\x56\xc4\x87\x97\x39\xc5\x66\x71\x9b\x6c\x89\x14\xad\x59\x00\x35\x4b\xd9\x67\xf4\xe6\x0d\xfd\xe1\xd2\x8d\xa6\x91\xd9\x91\x14\x86\x72\x61\x3b\xef\x71\x2e\x58\x33\xf3\x3a\x63\xb8\xfc\x6e\xff\x8f\xd8\xed\x1c\xdf\x50\x3b\xe7\x45\xa4\x1c\x29\x2e\x53\xe2\x7e\xa7\x59\x4f\x0e\xef\xd9\x8f\xd5\xe7\x9f\x25\xc4\x3d\xca\x83\x75\xad\x32\xcd\xb1\x08\x7a\x55\xd1\x7a\xf0\x0e\x15\xa9\x56\x30\x0f\xbf\xcf\x92\xc2\xed\xd8\x8b\x52\xc7\x03\xb7\xa0\x8a\xe6\xba\xa5\xdb\xf2\xe4\xb0\x49\x2c\xfb\x0f\xfe\xc7\xf3\x5b\xec\xc1\x39\x28\x25\xa2\xd8\x16\x72\x4e\x73\xdd\x20\x41\xce\x8d\x76\x71\x2c\xc4\xdd\xb6\x53\xe0\x38\xef\xfd\x86\xf3\x5e\x45\xdd\xff\x8f\x2f\x47\xb2\x87\xc3\x07\x8f\x1f\x7e\xeb\x3e\x31\xfd\xe1\x40\xe9\x76\x2b\x8c\xea\xcd\xda\xce\x7c\x38\xe2\x71\x21\x98\xfa\x89\xed\x0c\x2c\x0f\xef\xd9\x2d\x0a\x3c\x00\xcb\x34\xb3\x7a\xe0\x14\xf7\x07\x95\x40\xd2\x13\x0d\xf4\xff\x8f\x0a\x47\x01\x60\xb9\x5c\x82\x83\xa3\x84\xf3\x5c\xf1\x11\xf9\x2f\x17\xf5\x42\xaa\xc2\x71\x84\xa9\xfd\x1e\xa9\x9e\x49\x99\xa4\x31\x2c\x81\xea\xbd\x88\xc0\x6d\x09\xf1\x98\x27\x09\x53\x4c\x44\x18\xc1\x6e\xc2\x95\x36\x1e\x50\x21\x4d\x8a\xd9\x15\xc9\x1c\xfb\x4a\x35\x2d\x28\x6d\x48\x39\x33\xb8\xf6\x8a\xf0\x4d\x85\x47\xb4\xcc\x99\xbb\x42\x50\x45\xe0\x52\x0f\x56\xdd\xee\x99\x26\xf4\x9e\x72\x53\xbf\xb3\xac\x15\x2d\xd2\xce\x53\xb6\xab\xee\x29\x25\x06\x7e\xb4\x8f\xc4\x65\xea\xe1\xa4\xb8\x65\xee\xa3\x4a\xee\x63\x4a\x76\x3d\x40\x72\x52\x8d\x43\xa5\x80\xc8\x8c\x12\x1e\x5b\xaf\xaf\x08\x8f\xcf\x25\x7e\xe4\x4d\xe3\xb8\xce\xb8\xce\xeb\xe4\x90\x4b\x8c\x49\x4e\x0b\x77\x87\x70\xf7\x0b\xf0\xb8\x9c\xfe\xbc\xc3\x00\x69\x24\xb1\x2b\x77\xd7\xf5\xc0\x5e\xb8\x0e\x08\x55\x0d\x37\x92\xd8\xd5\x61\xa4\xf4\xa0\xfc\xd5\x62\xf9\x79\x42\x03\x0f\xd6\xf0\x0d\x3d\x6c\x55\x6c\xb1\x51\x09\xaf\xdb\xc8\x42\x4e\x55\x49\x94\xcc\x2b\x65\xec\xf2\xdb\x6f\x81\x12\xbc\xbe\x58\x90\x91\x6d\x86\x2b\x19\xb7\x19\xae\x3a\xe0\xe4\xe4\x3f\x73\xda\x43\xb7\xf5\xd7\xa8\xda\xc8\x02\xa3\xda\x72\x3d\x54\xfe\x6a\xbc\x21\xb8\xf9\x9e\xe7\x9b\xcc\x4e\x4a\x6e\x8d\x43\x92\xdc\xc7\x47\xb2\xd3\x58\x7b\x3e\x93\x76\x6c\xba\xe2\x19\xff\xcc\xdc\x91\xef\x7f\x1d\x57\x99\x76\xd6\x0f\xa7\xe9\xed\x56\xd9\x71\xaa\x66\xc2\xb1\xe6\x3c\x12\x69\x66\xfe\x21\x0c\x53\x5b\x9a\xb9\x48\xe5\xc1\xc0\xb7\x42\x1c\x31\x5a\xda\xb3\xfe\x6e\xff\x4f\x9a\x33\xf7\xf0\xe4\xd3\x25\x89\x54\xaf\x69\x94\xba\xe5\x8d\xe7\x44\x83\x1d\xba\xfa\xf5\x96\x09\xf3\x96\x6b\xc3\x04\x53\xae\x13\xa5\x54\xac\xf1\x3e\xef\x3e\x7d\x99\xfd\x52\xbe\xb4\x78\xc7\xe7\x95\x07\x58\x02\x23\x86\xaa\x35\x33\xcd\xc8\xaa\x2e\x2e\x96\xe6\x23\x16\xce\x92\xea\x42\xee\xe2\x50\x40\x14\x4b\x14\xd3\xe9\xb9\x8b\x1e\xba\xdd\x27\x2d\x60\x07\x94\xda\xfb\x4f\xf7\xc5\x4a\x1e\x14\xfc\xba\x66\xd5\xb4\x42\xca\xab\x0d\xa9\xdf\x63\x48\x35\x7a\x5e\xd2\xf5\xe0\x75\xcd\xcc\xbf\x4a\x2e\x6e\x73\xf6\x79\x68\x8e\x58\x5f\x80\x63\x50\x50\x7b\x69\x79\x43\x45\x9c\x31\x65\x25\xac\x18\x3e\x12\x63\x97\x6c\xe2\x9e\xf7\xc8\x26\x06\x59\xc9\x78\x4f\x58\xdd\x62\x9a\x48\xf1\x4e\x6e\x34\xfb\x25\x65\x2c\xc3\x4c\xb3\xdb\x17\x02\x04\x0f\xfe\xc6\x22\x90\xc8\xa8\xec\x7f\xd8\xbe\x7b\x68\x79\xd0\xef\x43\x05\x83\x98\x19\xfb\xfc\x44\x21\xa7\x11\x18\xb9\x89\xd2\x82\xc6\x50\x70\x11\xa5\xc0\x05\x48\x71\x6f\x4f\xb3\x9c\x9a\x87\x34\xe5\xae\xcb\x58\x09\x78\x16\x3a\xf5\x8c\x84\xf2\x77\x00\xd5\xdf\x46\x2d\xfa\xe5\xdf\x18\x5c\x2d\xfa\xf6\xef\xc0\xfe\x2f\x00\x00\xff\xff\xf3\x34\x09\xf3\x1e\x26\x00\x00")

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

	info := bindataFileInfo{name: "html/index.html", size: 9758, mode: os.FileMode(420), modTime: time.Unix(1686133872, 0)}
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
