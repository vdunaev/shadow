// Code generated by go-bindata.
// sources:
// templates/views/index.html
// assets/js/index.min.js
// DO NOT EDIT!

package workers

import (
	"github.com/elazarl/go-bindata-assetfs"
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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesViewsIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\x4f\x6f\xdb\x3a\x12\xbf\xe7\x53\x10\x2a\xb0\x6d\x0f\xaa\xfa\x6f\x2f\xad\xad\x45\xd0\x3d\xb4\x68\xb6\x58\xa4\x7d\x28\xde\x29\x18\x8b\x23\x8b\x31\x45\x0a\x24\x65\x3b\x08\xf2\xdd\x1f\x48\x51\xb6\x64\x53\x8e\xda\xda\x0f\x69\x5f\x4f\x96\xc5\x99\xe1\xf0\xc7\xf9\xf3\xa3\x78\x7b\x4b\x28\xe6\x4c\x20\x89\x32\x29\x0c\x0a\x13\x91\xbb\xbb\xb3\x09\x65\x4b\x92\x71\xd0\x7a\x1a\x55\x30\xc7\xd8\x30\xc3\x31\x4a\xcf\x08\x21\xa4\x3b\xe8\xde\x5f\x71\xcc\x8d\x1f\x74\x02\xc5\xab\xf4\xab\x54\x0b\x54\x7a\x92\x14\xaf\xbc\x56\x42\xd9\x32\x3d\x1b\xb0\xa0\xd8\xbc\xe8\x99\xe8\x48\xe4\x52\x95\xf1\x5c\xc9\xba\x22\x55\xcd\x79\xbc\x2b\xbb\x2b\xcf\x44\x55\x9b\x46\x61\x47\xca\x49\x72\x98\x21\xdf\x7f\xef\xc6\x9c\x2a\x31\x37\x15\x4e\xa3\xac\xc0\x6c\x31\x93\xeb\xa8\xb5\x7b\xad\x63\xbd\x62\x26\x2b\x22\xc2\xe8\x34\x82\xda\x48\x85\xb9\x42\x5d\x44\x24\x49\xc9\xf9\xf6\xff\xfe\xa4\x49\x60\x56\x8f\x48\xe0\xaf\x7f\x6c\x21\xeb\x2e\x2e\xe3\x08\x2a\x67\xeb\x28\x0d\x8d\x2a\xb9\x22\x46\x56\x57\x86\x71\xd4\x81\xed\x02\xc1\x4a\x30\x48\x49\xce\x59\xf5\x41\xfc\x49\x32\xc9\x63\x3e\x8f\x5f\xbb\x87\x92\xfa\x07\x5d\xfa\x87\xb5\x8e\x5f\xbc\x1c\xd8\x16\x3b\x47\xac\x0d\x18\x7d\x68\x2f\x32\x29\xa2\x74\xc2\x36\x7b\x09\x24\x87\x98\x22\xe4\xf6\xf7\xdf\x6e\x19\x2c\xdd\x41\x62\xd7\x48\x26\x6b\x1b\x97\x16\x74\xce\xb4\x41\x81\x4a\xc7\xcd\xcb\xf4\x79\x48\xb7\x78\x95\x5e\xb4\x82\xdb\x08\x0c\x83\xfc\x40\x21\xaa\x35\x2a\xfd\x7d\x18\xad\x9a\xd4\xbb\x0f\xa1\xbd\x0c\xfd\xa9\xf0\x31\xa0\x17\xdf\x89\x8f\x53\x8d\x57\xc0\xcc\xbd\x10\x01\x33\xc4\x89\xdf\x87\x52\x2f\x19\x7b\x81\x1a\xc8\xc2\xf5\x55\x05\x02\xf9\x00\x28\xeb\xab\x6e\xb1\xdd\xba\xf3\x32\x98\x46\xcd\xe2\x49\x37\xdc\x5f\xee\x68\xd6\xbc\x55\x13\xb0\x24\x02\x96\x33\x50\x4d\x11\x25\xce\x8f\x2b\x23\x25\xb7\x85\x2e\x54\x2c\x59\x3a\x81\x2d\x86\x9c\x43\xa5\x31\xe6\x4c\x2c\xf6\xb7\x24\x2b\x70\xa9\xa4\x88\x6d\xd9\x6d\x76\x04\xd2\x49\xc2\x59\xd0\x6c\xab\x4a\x95\xac\xa8\x5c\x89\xc0\xe4\x4e\x12\x48\xa1\x30\x9f\x46\x8f\xa2\x5d\x8d\xd8\xc8\xf9\x9c\x63\x44\x28\x18\xf0\x7f\x3a\xf6\x88\x92\xf6\xff\xac\x36\x46\x8a\x88\x80\x62\x10\xe3\xba\x02\x41\x91\x5a\x8f\xb9\xc6\xfd\x25\xac\x14\x8a\xac\xd8\xba\x1f\xf6\x69\x0b\xe8\xc6\x97\x12\x45\xdd\x4e\xe9\x9e\xc3\xaa\x2d\xa6\x83\x83\xbd\x35\x5f\xc3\x12\x74\xa6\x58\x65\xde\x2c\x25\xa3\x4f\x9e\x3f\xdd\x59\x6c\x29\x29\xf0\xf6\x1d\xa8\x39\x9a\x69\xf4\xa8\xfb\xd2\x3d\x37\xcd\x7b\x1a\xbd\x93\x22\x67\xaa\x24\x0a\x4b\xb9\x44\x02\x9c\x93\x6d\x94\x76\xe5\x67\x92\xde\x4c\xa3\xaf\x8c\x73\x32\xc3\xa0\x38\xc1\x75\x86\x95\xb1\xdc\x01\x6a\x6e\x36\x03\xcf\xc8\xb9\x42\x72\x23\x6b\xa2\x6b\x85\xff\xe9\x59\xcd\x80\xf3\x19\x64\x8b\x4e\x6e\x5c\x3a\xd3\x4f\x9e\xbe\x3d\x00\xd7\x06\x95\xcd\x56\xcd\xf9\x4d\x55\xd8\xaa\x40\x36\x4f\xb1\x51\xa0\xfd\xbe\x91\xcb\x8d\xc3\x87\x61\x1e\xda\xdf\x66\x70\x68\x97\x26\x49\x1d\x60\x0f\xfb\xf2\xfb\x72\x07\xdb\xf8\x56\xad\xff\xb7\x57\x18\x5a\x8a\x36\x6c\xd7\xc0\x8c\x63\xac\x50\x57\x52\x68\xb6\xdc\xad\x22\x4e\xdc\xc9\xf4\x14\x48\xa3\x56\xc8\x25\xaa\xa1\x44\x34\x05\x02\x1d\x1a\x53\x07\x90\x34\x45\xfa\x09\x4a\x9c\x24\xa6\x38\x2c\xf5\x4e\xa1\x6d\x2c\x23\x04\x6d\xd5\x26\xba\xce\x32\xd4\xfa\x7e\xf1\x0b\xd0\xdf\x20\xdd\x18\xcf\x81\xf1\x31\xae\x38\xdb\x63\x85\xcf\x33\xc3\xa4\x38\xe0\xc3\x24\x19\x42\xd2\xea\x1c\xc0\xdf\x66\xec\xa0\x62\x78\x70\x92\xb8\x4d\xff\x11\x4a\xda\x21\x1b\x27\xec\x72\x8e\x09\xf9\xd4\xde\x72\x96\xdf\x4d\xee\x17\x6e\x72\x5d\x16\xab\x0b\xb9\xfa\xe1\xf6\x80\x37\x18\xcb\x0a\x85\x0f\xa3\xcf\x85\x5c\xb9\x86\xe6\xb8\xdd\x29\xba\xc4\x71\x31\x28\x18\x0d\x55\xf2\x6f\xc7\x20\xe3\xd2\x45\x84\x05\xe1\x3d\xa3\xf8\xb7\x80\xb0\x89\x23\xb6\x64\xd4\x36\x98\x53\xc2\x76\x14\x7e\xa4\xd1\x38\x64\xda\xe2\x16\xa6\x31\x7e\xf4\xd2\x8a\x1f\x83\xc3\xb4\x1f\x13\x5a\x16\xe3\xbd\x78\x98\xe1\x79\x04\x9c\x17\x96\x60\x82\xa0\x5d\x86\x39\x0a\xf1\x8f\x8c\xf3\x63\x93\xc6\x8f\xce\x99\x93\xa2\xfd\xc0\xf2\xa0\x2d\x2e\x40\xe9\x0f\x23\x59\xf1\xba\x6d\xd1\xe7\x94\xfa\x5d\xfc\xcd\xbe\xc7\xb1\xef\x53\x30\xee\x0f\xff\x3d\x22\xdf\xfe\x6c\xc0\xd4\x23\xa8\xf3\x17\xd0\x8b\x7f\x26\x0b\x76\x2d\xf4\x84\x1c\xd8\xdb\x6f\x38\x70\xf7\xa3\xd4\xc3\xa6\xc1\x3f\x53\x16\x9e\xfa\x0c\x3c\x26\x23\xc7\x9d\x93\xc7\xa6\xe3\xff\x15\x93\x8a\x99\x9b\x11\x29\x69\x0c\x96\x95\x19\x7b\x96\x46\xa5\xa4\x3a\x62\x7d\xf9\xd5\x2a\x82\x0f\x2b\xd7\x75\x49\x0e\x14\xfb\x6c\xde\x36\x5c\x1b\x6d\x4c\x50\x5c\x4f\xa3\xf8\x45\x7b\xa6\xa2\x0c\xb8\x9c\xfb\x63\x5c\xc1\x28\x45\x31\x8d\x8c\xaa\x43\x17\x6f\x0d\x35\xf2\x1a\xe1\x5c\xf1\xec\xe9\xde\x7c\x69\xe4\x2c\x8c\xc1\xc8\x9f\x34\xa7\x4b\x7f\x2b\xd6\x1e\x35\x37\x39\x6c\x0f\x14\x0d\xb3\xa0\x4c\x97\x6c\x63\x30\xb8\x8c\x7f\x19\x56\xa2\x7e\x3b\x49\x1a\x33\x81\xc9\x8a\xd7\x7d\xb7\x7c\x85\xdc\xb2\x0a\x5b\xf5\x5e\x1f\xdc\x9d\xf0\xfa\xec\x6e\x87\x56\x97\x4b\x55\x0e\x84\x48\xf0\x0e\xf2\xe0\xa1\x17\x66\xc8\x49\x2e\x55\x6f\xb3\xfd\x2d\xc3\xb6\xb2\x0a\xa3\x24\x8f\x9d\x70\x94\x7e\xaa\xcb\x19\x2a\x22\x70\xd5\x2e\xf0\x4d\xf0\xb6\xb0\x37\x51\xf7\x9e\x52\x38\x03\x11\x29\x99\x98\x46\x2f\xa2\x9e\xc7\x7e\xae\xbd\xf8\x6b\x2f\x3e\x06\x52\x63\x0f\xce\xe6\xf5\x3e\x54\xa3\x90\xcf\xa5\x34\xdf\x1e\x59\x33\x23\xc8\xcc\x88\xd8\x7f\x62\x0e\xc7\x58\xfa\xce\xc6\xdf\x81\x70\xea\xcd\xa0\xeb\x59\xc9\xcc\xde\x0c\xfe\xd3\xe0\xc0\x0c\xe7\x94\x86\xed\x8f\xad\x09\xb7\xb7\x04\x05\x25\x77\x77\x67\x67\x9d\xdb\xf6\x6b\xdd\x5c\xb4\x37\xbc\xdd\x3b\x08\x55\xc5\x59\x06\xb6\x1a\x26\x5b\x52\xef\x91\x5b\x82\x6a\xbf\xb7\xb7\xb7\x3d\xb6\x5f\x90\x29\x79\x7c\x7b\x4b\x9e\x85\x86\xee\xee\x1e\xbf\x3d\x9b\x24\x8d\x19\xe7\x89\x36\x60\x58\xf6\xfe\xcb\xff\x2e\xc8\x93\xe6\xf9\x8f\xcb\x0b\x12\x25\x3e\x36\x12\xd0\x1a\x8d\x4e\xae\x75\xe2\xaa\xd3\x33\xeb\xa6\x4d\xdd\xa7\xd6\xd9\xcd\x4a\xfe\x0a\x00\x00\xff\xff\x49\xed\x61\x3b\x36\x20\x00\x00")

func templatesViewsIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesViewsIndexHtml,
		"templates/views/index.html",
	)
}

func templatesViewsIndexHtml() (*asset, error) {
	bytes, err := templatesViewsIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/views/index.html", size: 8246, mode: os.FileMode(420), modTime: time.Unix(1505252732, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsJsIndexMinJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x58\x51\x6f\xdb\x38\x0c\xfe\x2b\x9e\x56\x34\x16\xe2\x38\xdb\xdd\x76\x0f\x89\xed\x62\xc8\x0e\x87\xe1\x8a\xbb\x61\x2b\xb0\x87\x6d\x28\x14\x4b\x49\x74\x91\xed\x40\xa2\x9b\x05\x5d\xfe\xfb\x41\xb2\xec\xc8\x89\xd3\x76\x59\xbb\xa7\x28\x24\x45\x52\x24\x3f\x8a\xf2\xac\xcc\x53\xe0\x45\xee\xcd\x19\x7c\x2a\xe4\x92\xc9\x8f\x40\xa0\x54\xff\x90\x8c\xf9\xca\x2c\xf1\xad\x5a\x73\x48\x17\xcd\xdf\x94\x28\xe6\xbd\x18\x49\x06\xa5\xcc\xd1\x9a\x70\x40\x63\x43\x7b\x59\xd3\x56\xb2\x48\x99\x52\x96\xfc\x5b\x4d\x9e\x96\x6a\x83\xc6\x94\xcd\x48\x29\xa0\x26\x96\xf9\x32\x2f\xd6\x39\xda\x6e\x5d\x5f\xae\x88\x5a\x3e\xa1\x27\xaa\x4c\x1d\xf2\xef\x35\x79\x46\xb8\xb0\xb4\x57\x2e\xcd\x9b\x6e\x3c\xe0\x19\x2b\xca\xda\xc0\xeb\x9a\xbd\xe4\xa2\xde\xf2\x87\xeb\x88\x27\xd9\x8a\x11\x78\xc8\x71\xcb\x15\x25\xc0\x7c\x7c\x7b\x16\x92\xff\xc8\x37\xff\x16\x36\x2b\x36\x42\x7f\xfd\x79\x85\x82\x52\x8a\x11\x1a\xae\x4d\x66\xd4\x50\xb3\x87\x17\xc4\x6c\x8b\x75\x10\x14\x0a\xec\x51\x46\xb5\x3a\x5f\xe2\xdb\x1b\x22\x3d\xb8\x8c\xcf\x7c\xf4\x5c\x70\x05\x2c\x67\x52\x79\x30\x2d\xe8\x06\xe1\x90\x65\x2b\xd8\xf8\x38\x80\x4f\x46\xc0\xea\x3e\x64\x5f\x19\x36\x10\xb5\x3c\x64\x36\x5a\x63\x19\x36\xeb\xeb\xb4\x28\x73\xf8\xfe\xfd\x45\x60\x75\xc6\x32\xb4\xab\x1d\xcb\xe8\xbb\xd6\x11\x8a\x5f\x8c\xf9\xcc\x7f\x23\x25\xd9\x84\x5c\x99\x5f\xdf\xd1\x86\xf1\xed\xac\x90\xbe\x3e\x0a\xf7\x78\xee\xb9\x2c\x73\xc0\xfa\xaf\xeb\xc2\x67\xfe\x35\x10\x93\x18\x45\x20\x93\x08\x68\x82\xfa\x35\x2b\xcc\x49\xc6\xfa\x28\x1a\x02\xb5\x1c\x1d\xf4\xab\xe2\x23\x48\x9e\xcf\xfd\x46\x2c\x95\x8c\x00\xa3\xd7\x04\x70\x4b\x7a\x27\xa0\x4f\x72\xad\x8f\x71\x6d\x43\xdf\x92\xdb\x69\x12\x44\xb5\xe5\xae\x09\x5c\x74\x1b\xed\x12\xc5\x23\x84\xee\x77\x41\x57\x27\xa3\xf7\x7a\x50\x89\x3d\xc4\x81\x46\xd2\xb5\x8f\xc6\xad\x30\x3e\x8b\x6d\x51\x5f\x5a\xaa\xc6\xe8\x85\x98\xf4\xe3\x9e\x76\x21\xa2\xfc\xc6\x4b\x05\x51\x2a\x46\x53\xc8\x07\x73\x59\x94\x2b\xaf\x59\x0d\xbe\x29\x94\x44\xd3\x12\xa0\xc8\x3d\x5d\xea\x31\xaa\xfe\x20\x67\x93\x11\xa7\x24\x9f\x33\x69\x96\x3c\x2d\xf2\x26\xe3\x03\xc9\xb2\xe2\x86\x21\x8f\x12\x20\x03\x28\xe6\x73\xc1\x62\x94\x15\x94\x88\x9a\x46\xe4\x9c\x41\x8c\x9e\xbb\x44\xb3\x1e\x00\x07\x2d\x3d\x29\xf2\x19\x97\x99\x57\xa9\x6a\x54\x7b\xbd\xbd\x8a\xe9\xb5\x36\xa7\x44\x88\x29\x49\x97\x31\x6a\x4a\xee\x83\x51\xe0\x7f\xe9\x1d\xec\xfc\xd2\xc3\x63\x94\x44\xbc\x3e\xd6\x5c\x6c\x56\x0b\x73\x90\x66\x35\x00\x49\xd4\x02\x79\xd6\xa9\x4a\x17\x4a\xa2\x21\x4f\xa2\x61\x15\x95\x24\x1a\x52\x7e\x93\x98\x3c\xf4\x46\x3a\xc8\xc8\x04\xd9\xe4\x25\xa8\xfe\x0f\x41\x26\x28\x80\xcb\x90\xac\x56\x2c\xa7\xbe\x98\xe0\x6d\x27\x44\x43\xc1\xf2\x39\x2c\xb6\x1d\xc0\xb3\x48\x3d\x84\x5d\xcd\x30\xa0\xab\xfe\xec\x80\xad\x01\xa7\xeb\x26\xae\xfe\x87\x7a\x1d\xac\x5b\x08\xb4\x1c\x4e\x8f\xa3\xcf\x8a\x58\xec\xb5\xab\xbe\xeb\x62\xb2\xf2\xf6\x1a\x68\x8a\x74\x3d\xe9\xc7\xda\x81\x8b\xaa\x0c\x1f\x52\x62\x16\x6f\x66\x9d\x72\x99\x0a\xe6\x69\x0d\x03\xb5\x28\xd6\x4d\x2d\xa9\x65\x8c\x7a\x7d\xde\xef\xdd\x97\x4d\xb6\x61\x83\x62\xc5\xf2\x83\x0c\x9a\xdc\xb9\x79\x5b\x3f\x3a\x58\x78\x3e\x2b\x76\x50\xa9\x22\x34\x90\x4c\x31\x78\x24\x9c\x28\x06\x56\xad\xf7\xbc\xe7\x64\xf5\x18\x42\x6c\x85\x7c\xd0\x1b\x0d\x3c\x9c\x2d\x0f\x81\x86\x64\x33\xc9\x5a\xe0\xd0\x67\xd9\x8b\xec\x29\x6d\xc4\xc6\xc6\xdc\xdc\x8f\x11\x1a\xad\xe8\x94\xc8\xfc\xcd\x85\x38\x25\x30\xed\x9e\xa1\xb5\x1c\xef\x18\xa6\x35\xf4\x02\xf8\x54\xb7\x86\xf5\x04\x1b\xc0\x9e\x9f\xef\x68\xbd\x08\xa4\xa7\x60\xa3\xd5\x51\xae\x56\x82\x6c\x46\x79\x91\x33\xe4\x71\x1a\x23\x83\x87\xba\xfc\x81\x7a\x69\x21\xd4\x8a\xe4\x31\x7a\x8d\x92\xa8\x14\xb5\xa7\xba\xc7\x54\x35\x8b\x92\x48\xf0\x43\xf2\x80\x03\xcb\x50\x12\xe9\xcd\x35\x77\x55\x0a\x31\x90\x7c\xbe\x00\x0f\xd8\x37\x18\x64\x25\x30\xea\xa9\x8c\x98\x33\xb1\x2c\xe9\xf5\xb5\x79\x13\x9a\x68\xc8\xb2\x24\x1a\xea\xed\x49\xa4\x40\x16\xf9\x3c\x79\xf7\x36\x1a\xda\x65\x34\x95\xde\x30\x89\x86\x82\x3f\xbe\xf9\xaa\x9f\x77\x39\xa0\xdb\xd1\x53\xbb\x70\x38\x0e\x1b\xa7\xea\xee\xd7\xe9\x57\x25\xfc\x4b\x82\xb3\x92\xbc\x90\x1c\x36\xdd\x8e\xbc\xb7\xdc\x5f\xe2\x0a\x01\xd0\x13\xaa\xea\x76\xe5\x8d\xe5\x3e\xb5\x2b\x55\x7a\xcc\x48\xc5\xa4\x2c\xe4\xc5\xde\xff\x11\x3a\xcf\x28\x51\x8b\x31\x3a\x92\xbc\x4b\xa2\xc0\x33\xa2\x4f\xed\x6a\xce\xd6\xde\x5b\xfd\xf8\x30\x2e\xd6\xf7\x6f\x08\xc5\x65\x91\x12\xc1\xec\xf5\x6c\x6e\xd8\x43\x37\x27\x95\x78\xa7\x8f\xc3\x52\x38\x0d\x08\xe1\xed\xc1\x93\xe0\x8e\x51\x64\xf7\x46\x38\x9c\x46\x1c\xde\xad\xb9\x99\x5d\xd2\x67\xfe\x75\x0c\x57\x75\x5b\x73\x66\x90\xba\x89\xb8\x83\xc5\x0e\xd9\x7b\xe3\xc6\x5d\x60\x3b\x50\xd0\x54\xff\x01\xa7\x29\xc6\xd6\x70\xfe\x03\xb5\x71\x74\x5c\x6a\x25\x4b\x97\x90\x15\xbb\x6f\x96\x50\xd9\x49\xb3\x84\xb9\x00\x1e\x7f\xe4\xd6\x6a\xf5\x7d\xd9\x34\xf8\x23\xb7\xa5\xe6\x3b\x43\x76\x23\xfe\xf8\xe3\x75\x75\x59\xe2\xad\xf3\x46\x75\xab\xab\xae\xd8\xd6\xb3\x7a\x60\x1e\x63\x08\x87\x1a\x62\xfe\xee\x89\x1a\x38\x6f\xeb\xb6\x4c\x3d\x4f\x07\xcd\xf3\xba\x0a\x50\xcd\xdf\x19\xec\x23\xcf\x7c\x4a\x30\x14\xe4\x6e\xd0\xe4\xb6\x56\x07\x17\xae\x69\x2f\xdc\x8d\xb3\x38\x4c\x05\x4f\x97\x7e\xf3\xa5\xa0\x1a\xe9\x59\x7c\xe6\xc3\x82\x2b\x1c\x4c\x63\x16\xce\xb8\x46\x0e\x47\x38\xe0\x31\x0b\x75\x42\x7c\x93\x01\x84\xc7\xd3\x70\x41\xd4\x44\x07\xdb\x47\x1d\x33\x2f\xbe\xf0\xa7\x61\x95\xdb\x3b\x84\x82\x69\x48\x28\xed\x14\x48\x45\xa1\x98\x73\xcc\x01\xea\x73\x1c\x6a\xd7\x7d\x34\x23\x0a\x10\xc6\xa3\xbb\x4d\xd4\x1a\x8e\xdb\xb0\x4e\xb4\x4d\x2c\x38\x65\x8d\x89\x2d\xde\x6e\xf1\xee\xe3\xcc\xfe\x3b\x8f\xd3\xfd\xef\x34\xef\xff\xfd\x78\xf7\x87\x9a\x5d\xad\x58\x10\x05\x3a\xac\xa3\x5b\x4e\x47\x9c\x6e\x9b\x0f\x38\xd5\x67\x20\xd7\xb6\x53\xf9\x27\x98\x75\x61\xfb\x50\x8b\xee\x64\x7a\x82\xc9\xba\xde\xcd\x68\xfd\x83\x36\xab\x77\xc2\x4f\x18\xad\xde\x3a\xf7\x58\x3d\xf3\x69\x91\x96\x19\xcb\x01\x87\x92\x11\xba\x71\xd1\xe0\x42\xf6\x18\x62\xba\xb1\x35\x5a\x10\xe5\xf3\xf0\xb0\xd6\x70\xa3\x03\x6f\xdb\x3d\x41\x17\xdd\x4f\x1a\x30\xe5\x7e\xdc\x02\xa1\xd4\xab\x9a\xdc\x67\xd3\xed\x55\x39\xcd\x38\x7c\xed\x34\x7a\x5a\xc8\x09\xa5\x75\xc0\x4d\x33\x1a\xed\x99\x6f\x5a\xd4\x0d\x11\x3e\xee\x48\x07\x0e\xea\x8f\x9f\x63\xdd\x8a\x48\x09\x85\x7d\xf6\xc5\x79\x29\xc4\x58\xeb\x73\x88\x5d\xae\xeb\xd6\x15\xa6\x0b\x96\x2e\x19\xbd\xd0\x9b\xe2\x38\x76\xb6\x9c\x9f\xfb\xb5\x89\xc0\x55\xbf\xe6\x39\x2d\xd6\xa1\x62\xf0\x2e\x07\x26\xb5\x83\x95\x5c\xf0\x92\xbd\xc2\x78\xa4\x35\x3d\xdb\xd7\x64\x37\xa5\x82\x11\xd9\x6c\x73\x44\xda\x26\xb4\x0a\xdd\x4e\xf0\xf8\xff\x00\x00\x00\xff\xff\xd3\xd5\xb7\x48\x66\x17\x00\x00")

func assetsJsIndexMinJsBytes() ([]byte, error) {
	return bindataRead(
		_assetsJsIndexMinJs,
		"assets/js/index.min.js",
	)
}

func assetsJsIndexMinJs() (*asset, error) {
	bytes, err := assetsJsIndexMinJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/js/index.min.js", size: 5990, mode: os.FileMode(420), modTime: time.Unix(1505253211, 0)}
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
	"templates/views/index.html": templatesViewsIndexHtml,
	"assets/js/index.min.js": assetsJsIndexMinJs,
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
	"assets": &bintree{nil, map[string]*bintree{
		"js": &bintree{nil, map[string]*bintree{
			"index.min.js": &bintree{assetsJsIndexMinJs, map[string]*bintree{}},
		}},
	}},
	"templates": &bintree{nil, map[string]*bintree{
		"views": &bintree{nil, map[string]*bintree{
			"index.html": &bintree{templatesViewsIndexHtml, map[string]*bintree{}},
		}},
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


func assetFS() *assetfs.AssetFS {
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: k}
	}
	panic("unreachable")
}
