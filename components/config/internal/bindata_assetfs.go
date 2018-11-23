// Code generated by go-bindata.
// sources:
// templates/views/manager.html
// assets/js/manager.min.js
// locales/ru/LC_MESSAGES/config.mo
// locales/ru/LC_MESSAGES/manage.mo
// DO NOT EDIT!

package internal

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

var _templatesViewsManagerHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x5a\x5f\x93\xdb\xb6\x11\x7f\xf7\xa7\xc0\x20\xd7\x8c\x14\x9b\xa2\xed\xba\x99\x36\x91\xce\x93\x71\x33\x71\xa7\x6e\x3a\x93\x38\xee\x83\xeb\xde\x80\xc4\x52\xc2\x95\x02\x64\x00\x94\xee\xaa\xb9\xef\xde\x01\xc0\xff\x02\x25\x8a\x3a\xdd\xd5\xf5\xc3\x99\x22\x80\xfd\xf3\xdb\x5d\x2c\x16\xdc\xed\x16\x51\x48\x18\x07\x84\x63\xc1\x35\x70\x8d\xd1\xdd\xdd\x93\x29\x65\x6b\x14\xa7\x44\xa9\x19\x96\x62\x83\x2f\x9f\x20\x84\x50\xfd\xed\xcd\xd5\x8a\x70\x48\xf3\x91\xdd\x51\xcd\x74\x0a\xb5\x51\x3b\x63\xf1\xf2\x72\xbb\x45\xec\xc5\x1f\x39\xc2\x6f\x04\x4f\xd8\x3c\x93\x44\x33\xc1\x31\x9a\xa0\xbb\xbb\x69\xb8\x78\xd9\x5a\x51\xa3\x19\xa7\x40\x64\xc2\x6e\xf0\xe5\x34\xa4\x6c\x5d\x63\xdc\xfa\xd9\x90\xa3\xd0\xaa\x49\xd7\x48\x91\xa0\x09\x48\x29\xa4\xd1\xb7\x8b\x27\x49\x41\x6a\x64\xff\x06\x94\xf0\x39\xc8\xe2\x07\x53\x4b\xa6\x14\x89\x76\x94\xb4\x24\xa2\x4c\x6b\xc1\x91\xbe\x5d\xc1\x0c\xbb\x1f\xb8\xd2\x43\x28\xc0\x88\x12\x4d\x0a\x32\x39\x23\x8c\x88\x64\x24\x58\x30\x4a\x81\xcf\xb0\x96\x19\xe0\xcb\xaf\x35\x5b\x82\xfa\x7e\x1a\x3a\x32\xbb\xcc\xb6\xdb\x5c\x91\xc9\x8f\x5e\x75\x9a\xe8\xe4\x2b\x80\x53\x33\xb1\x39\x33\x11\x72\x89\xa4\x48\x61\x86\xcd\x23\x46\x24\x36\xc6\x99\xe1\xaf\x30\x5a\x82\x5e\x08\x3a\xc3\x2b\xa1\x34\x46\x8c\xce\x8c\xbf\x24\x6c\xae\x72\x4d\xb4\x98\xcf\xcd\xc2\x35\x49\x19\x25\x5a\x48\x1f\x2c\x06\x59\x47\x5f\x93\xa8\xed\x3f\x8d\x99\x59\x5a\xc0\xc5\xc9\x1a\x71\xb2\x0e\x34\x89\x14\x8a\x88\xbc\x32\x0f\xb8\x22\x93\x32\xd5\xb6\x6e\x4b\x57\x69\x0c\x87\x2e\x38\x59\xc2\x33\x74\xb1\x26\x52\xa1\xef\x66\x68\xb2\x36\x60\x47\x29\xa8\x36\x62\x0d\x41\x52\x96\xb3\x5a\x49\x50\xc0\x75\xee\xae\xce\x81\xe0\xb3\x23\x8b\x72\x30\x4c\xec\x94\xae\x13\x6b\xb6\x06\x5c\x62\xdd\x2d\xa2\xe5\x43\xd0\x42\x42\x32\xc3\x5f\x6d\xb7\x39\xcd\xbb\xbb\xdc\x1f\x8c\x0b\x4b\x91\xaa\x19\x6e\x8c\x95\x08\xb4\x2c\x60\xdf\xd8\x85\x70\xb3\x22\x9c\x02\xb5\x0b\x3b\xc4\x35\x4e\x66\x64\x4c\x95\xa1\x9a\x90\x54\x41\x29\xf2\x1e\x58\x6b\xf0\x16\x12\xed\xd7\x2f\x24\xdd\xc4\xa6\x61\xca\xf6\x5a\x30\xf7\x56\xff\xda\x2c\xbd\x7c\xe2\x1f\xaa\x05\xb2\x26\x51\xe0\xdf\x0a\x5a\xac\x4e\x72\x16\x8f\x7f\xd7\x05\x30\x6f\xf6\x38\x8e\xf3\x18\xc4\x78\x65\x00\x1b\x68\x75\xab\x1f\xf0\xa2\xa6\xc6\x29\x04\x12\xd4\x4a\x70\x65\x3c\xf1\xb0\x31\xa7\x76\x4d\x83\x00\x72\x64\x16\x62\x0d\xbe\x88\xf6\x53\x59\x00\xa1\x7d\xe7\xca\x7e\x13\x73\xc2\x68\xc3\xa8\x5e\xcc\xf0\xcb\xe7\xbf\xc3\x55\x2a\xf9\x99\x2c\x01\xa3\x0b\x9b\x41\xf4\x62\x18\xc1\x3f\xd4\x09\x7e\x20\x69\x76\x2a\xc5\x17\x0d\x8a\x7f\x86\x84\x64\xa9\x3e\x95\x66\x43\xed\x7f\x10\x1d\x2f\x40\xaa\x81\x44\xeb\xc2\xa9\x58\xb2\x95\xdb\xd9\x8e\xa3\x35\x0d\xfb\x5a\xd0\xd0\x3c\xc2\x2f\x22\x41\x6f\xfb\xcd\xad\x82\x96\x3d\x43\x17\x2e\x9e\x4c\xcc\xba\xe0\x3d\xb0\x2f\xd5\x88\x98\xf9\x7f\xa1\x66\xa5\x84\x55\x4a\x62\x28\x88\x4d\x3e\xe4\xc1\x3f\xf9\x2b\xdc\x22\x3c\xc1\x08\x5f\xe1\xbe\x84\x8f\xf4\xf0\x9e\xf8\xd4\xc4\x66\xc9\xae\x9c\x3f\x52\xe6\xa2\xb7\xa7\x90\x25\xff\x94\x44\x50\x66\xde\x3c\xef\x04\xf6\x25\x46\x89\x90\x6e\x2f\x72\x38\x99\xdd\xc9\xe6\x9d\xb9\x14\xd9\xca\xa5\x18\xe3\x4d\x3b\xc2\xfc\x64\xc6\xd1\x45\xb5\xd5\x55\x9b\x99\xa1\xe6\xc3\xd8\x78\xa0\xe5\x7a\x34\x1a\x79\x1e\xbb\x3f\xad\xff\x07\x55\xec\x4e\x86\x5e\xed\xc2\x63\x9c\xea\xd1\x3d\xb0\x96\xc1\xcc\x19\xd4\x41\xdf\x33\xf3\x74\xcb\xf5\x96\xa8\x0f\x0c\x36\x08\x03\xcf\x96\xbd\x63\xb7\x21\x98\x82\x14\x62\xdd\x90\x2d\x77\x15\xe4\x86\x5e\x62\x64\xac\xee\x62\xc4\x6f\xf2\x2a\x9f\x97\x31\x74\xbc\x66\xa8\xb1\xe7\x89\x67\xe8\x42\xd8\xad\xdb\xec\x5c\xa3\x82\xf1\x4f\xa0\x8d\xc2\x7f\x77\x23\xd8\xcd\x50\x78\x3c\x44\x75\xab\x7e\xce\x63\x6d\x32\xa3\x8b\x04\x4e\xe1\xa6\xe4\xfd\xdc\xe8\x52\x9e\x6d\x46\xad\xc1\xf1\x2e\x20\x36\xc5\x9a\x53\x8f\x03\xcf\x1c\x52\x8b\xa7\xda\x99\xb9\x88\xb8\x16\xc1\x17\x63\x5f\xb0\x4d\x43\x37\x3c\x18\xd2\x23\x23\xab\xc4\x26\x74\x92\x0f\x72\x52\xbb\x61\xe5\x27\xc2\x36\x46\xef\x6f\x57\x80\x70\x24\x44\x3a\xcc\x65\x19\x5f\x65\x3a\x2f\x42\xe3\x05\xc4\xff\x8e\xc4\x4d\x79\x20\xbd\x56\x81\xda\x30\x1d\x2f\x06\xb9\x6d\x55\x00\x8d\x56\x92\x71\xdd\x61\xe0\x31\x72\x55\xac\xad\x8b\x8c\x04\x40\x4b\x9c\x51\x78\x12\x60\x42\xa2\x51\x37\x68\x8c\x6b\x3c\x3e\x30\xe1\xdb\x57\xc3\xe2\xa1\x81\x2b\xcf\x96\x11\x48\xec\xdb\x17\xfa\x00\x5b\xc5\x53\x67\x84\xf8\xd0\x5f\x11\xad\x41\xf2\x19\xfe\xd7\xc7\xe0\xe9\xa7\xd7\x1f\x9f\x07\x7f\xfa\xf4\xcd\x05\x3e\x27\xa2\xd9\x41\x48\xb3\xff\x1b\x4c\xef\x0d\xd2\x6e\xb0\x92\x54\x10\x83\xd6\x97\x0f\x56\xe5\x80\xa3\x7f\x4e\xdc\xc3\xf8\xf5\xf9\x80\xa3\xe5\x95\xe1\xa9\xc8\x69\xb8\xd1\x8f\x8c\xdb\x68\x07\xb8\x8f\x24\xf8\xcf\xa7\xa7\xe3\xa7\xa7\x02\xb8\x73\xf2\x59\x11\xa5\x36\x42\xd2\xd3\x71\xab\x28\xf9\x0e\x43\xc5\x68\xa0\x16\x62\x73\x3e\x24\xef\x19\x1d\x4d\xe6\x6a\x10\x32\x46\x28\xea\x8a\xfb\xf7\x70\xa3\x6d\xe1\xe9\x3f\x83\xe5\xb3\x02\xe7\x76\x67\x73\x5e\xe4\x54\x39\x17\xee\x43\xcf\x56\x06\xf4\x3a\x4e\x77\x77\xee\x02\xbc\x06\x4a\xad\xc0\xaa\xcf\xf4\x95\x55\xe5\x21\xe2\x04\x37\xf8\x02\xb7\x8f\x81\xda\x1e\x7f\xac\xf5\x7c\x34\xe8\xc3\x68\x48\xf1\xad\x56\x84\x7b\xab\xe4\x02\x87\x69\x68\xa7\x7c\xd9\x65\xb2\x55\x73\xc7\xa0\x57\x6b\x77\xd1\x68\x43\x61\x8f\x4b\xe4\xb7\x87\xc3\xea\xc5\x8e\x4a\xf8\xa4\x7c\xf0\x8d\xfd\xf7\x90\xa1\x87\x4a\x4d\x08\xa7\x7b\xee\x1b\x46\x66\x78\xe4\xaf\xfb\xc7\xd5\x40\x1b\xda\xf1\xe0\xca\x18\x3d\x46\x39\xde\x44\xa4\x5f\xc5\x5d\x39\xd1\x49\x2c\x0b\xb6\xbd\xcb\xf2\x53\x15\x1c\x58\x95\xdf\x0f\x85\x13\x3d\x36\x27\x71\xdf\xb6\x18\xae\xd3\xd0\x74\x70\xdc\x16\xfc\xa0\x77\x8e\xc5\x67\x90\x61\x99\x67\x08\x82\x79\xac\x6f\x9e\xa1\x8b\x8d\x63\x5e\x3f\xf6\x0d\x95\xa7\x94\xa9\x38\x53\xb8\x6b\x69\xfb\xb7\x38\x27\xb9\xfb\xe4\x9c\xe9\xe4\x57\x91\xc9\x78\x68\x8a\x44\x0f\xe6\x0d\xc3\x38\x3d\x44\x42\xbe\xec\xbc\xd1\xff\x4d\x11\x63\xe2\x13\x77\xb3\x2a\x61\x8d\xb8\xe7\x73\x52\x99\xb4\xb4\xcc\x60\xdc\x9d\xb5\x86\x38\x52\x24\x51\x58\x7d\xe1\xfb\x21\x4d\xc5\x06\xa8\x3b\x76\xaa\xef\xdc\x57\xbe\xe3\x89\x66\x47\x7e\xad\x40\x8f\x91\x1b\xf7\xe4\xc4\xdd\x7c\x35\xf4\x3a\x3c\x65\x97\xde\x3b\xf0\xfd\xad\x0c\x07\xc4\x3e\x25\xd3\x74\x4a\x84\xbe\x5e\x52\xa2\x16\xdf\x1f\x97\xb4\x4f\xd3\xe3\x21\x73\x13\x2a\x9b\x40\xce\xcf\xeb\x7c\x79\xb0\xff\x77\xf4\xe3\x24\x9f\x86\x3d\xbf\xa4\x4f\x43\xbb\x15\x1d\xe8\x2f\xd9\x5f\x16\x1e\x18\x3e\xd4\xc8\xe3\x5d\x9c\xbf\xf6\xf7\x92\x15\x89\x92\x5f\x29\x91\x32\xba\xd3\x19\xe8\x9b\x7c\xf0\xab\x62\xa3\xe7\x50\xa4\xc1\x92\x06\xdf\xa2\xfc\x41\x24\x89\x02\x1d\xfc\x7e\x4f\x19\xd6\xec\xfd\x93\xa0\xa0\xba\x25\x88\x34\x47\x91\xe6\x01\xe3\x89\x28\xef\x00\xdc\x94\x6a\xaf\xfe\xc5\x2d\x71\x0d\x91\x5d\x4d\x7f\x7e\x6e\x2a\x8b\x96\x6c\x97\x9d\xca\xe2\x18\x94\x2a\x39\x2a\xb2\x36\x45\x27\xb3\xed\x8b\x74\x86\x8b\xa7\x9a\x10\xbf\xda\x29\x87\x65\xd8\x6f\xb5\xe6\x2b\x83\xbc\xb7\x6b\x33\x7f\x2c\xfe\xab\xe1\xbf\x14\x94\xa4\x28\x21\x14\xdc\x6d\x88\xfd\xfd\x26\xdf\xab\x34\x89\xec\x56\x36\xc3\xc1\x8b\xa2\x1d\x8e\x32\x92\x8a\x79\xde\xff\x66\x8f\x4e\x29\xd0\xe8\xb6\xb1\xf2\x9d\xfd\xb2\xef\x04\xf1\x34\x5c\x3e\x69\x3b\x81\x5d\x1a\xe4\x94\xfd\x7d\xa6\x6e\x8a\xbf\xc1\x6c\x77\xde\x02\x08\xf5\xb6\x53\x1d\xdf\x37\x6a\x09\x0e\xeb\x1b\x9d\x2e\x5e\x35\xc5\x72\x2d\xbb\x6d\x9c\x1d\x5a\xad\xbe\x5d\xb9\x44\xc6\x89\x50\xec\x6d\xe2\x7d\xd5\xb6\xfc\xae\x33\xec\x60\x62\x76\xa9\x1a\x97\x1f\x24\xa0\x5b\x91\x21\x95\x49\x78\x5d\x10\xee\x41\x26\x11\x42\x1f\x0f\x6d\x11\x28\xc5\x29\xdb\x0b\x72\x0d\x02\x67\x89\xfd\xd1\xd1\x2b\x32\x5d\x4f\xf1\x21\x7e\x87\x82\xb1\xbb\x1f\xba\x19\x59\xb5\xbe\xdf\x5a\xdf\xb7\xf1\xc6\xf2\x12\x68\xbb\x45\x4a\x13\xcd\xe2\xb7\xef\xff\xf6\x0e\x8d\xdc\xf3\x6f\xbf\xbc\x43\x38\x34\xa7\x88\x48\x10\x49\x43\xa2\x14\x68\x15\xae\x81\x53\x21\x55\x68\x84\xb7\xc9\x43\x4d\x38\xe8\x20\x52\x61\xac\xdc\xdb\xf7\xee\x6d\x24\x84\x56\x5a\x92\xd5\x64\xc9\xf8\x24\x36\x3b\x91\x6d\x34\x1d\xdf\x23\xd7\x84\xdd\x00\x75\x81\x55\x48\x60\x5f\xbd\xb5\xaf\xf6\x8b\xe0\xc7\xe5\x5a\xdd\x23\x2a\xe1\xb5\x0a\xaf\x3f\x67\x20\x6f\x27\x35\x60\x8c\x2c\xd7\xe7\x40\x23\x52\x86\x61\xa7\x09\xce\xc2\xb3\x66\x81\x16\xf3\xba\x21\x8e\x64\x9f\x77\x17\xe0\x10\xa3\xc9\x1b\xb1\x5c\x09\x0e\x5c\xff\x6c\x3b\x67\x0b\x71\xae\x55\xb8\x24\x9c\xcc\x41\x1a\xc2\xe3\xbc\x98\xaa\x9b\xf5\xbf\x01\x00\x00\xff\xff\x7c\xba\xd1\xbb\xe5\x30\x00\x00")

func templatesViewsManagerHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesViewsManagerHtml,
		"templates/views/manager.html",
	)
}

func templatesViewsManagerHtml() (*asset, error) {
	bytes, err := templatesViewsManagerHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/views/manager.html", size: 12517, mode: os.FileMode(420), modTime: time.Unix(1543013994, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsJsManagerMinJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x55\x4d\x6f\xe3\x36\x10\xfd\x2b\xf2\x24\xc8\x92\x30\x43\x27\x87\x02\x85\x6d\x7a\x0f\xee\x1e\x0a\x14\x45\x81\x6c\x7b\x09\x8c\x82\x12\xc7\x16\x1b\x8a\x54\x49\xca\xde\x20\xf0\x7f\x2f\x46\x56\xbc\x4e\xbc\x09\xda\x62\x2f\x06\xcd\xe1\x7b\x33\x6f\xbe\x74\xc9\x4c\xa8\xba\x06\x7d\xe6\x32\xa2\x36\x8f\x6c\xdd\xf9\x2a\xdb\xe0\x19\x7f\xba\x64\x70\x51\x05\xbf\xb6\x9b\x54\x58\xdf\x76\xf9\xde\x9a\x95\x28\x8e\x77\x09\x1d\x56\x19\xb8\xac\x6a\xed\x37\x78\x0a\xdd\xea\x58\xa0\xba\x64\xb9\xb6\x89\x8b\x18\x76\x0a\x65\xab\x23\xfa\x9c\x7e\xf7\xd9\x3a\x06\xb9\x0c\xe6\x11\x04\xe4\x08\x5c\x54\x5d\x24\x9b\x42\xb9\xd5\x8e\x71\x61\x70\xad\x08\x2a\x0d\xae\x75\xe7\xf2\x1f\xda\x75\x38\x83\xaa\xc6\xea\xa1\x0c\x5f\x40\x11\x5d\x0c\x2d\x83\xfc\xd8\x22\xf0\x8f\xec\x99\x01\x60\x74\xb4\xf5\xcf\xd1\xc0\x39\xdf\xf2\x60\xe1\x53\xb8\xfb\xf4\xcb\xa7\xe5\x67\x50\x07\x73\xd6\x9b\x5f\x75\x83\x57\x57\x8c\x10\x28\xd7\xd6\x1b\x06\xa1\x25\x55\xc0\xe5\xda\xba\x8c\xf1\x54\x67\xc4\xdc\x45\x5f\x0c\x42\x07\xbf\x83\x93\xbb\x3e\x3d\xe4\x7f\xcf\x0f\xba\xbe\x0a\x55\x06\xd7\x1f\x63\xd8\xc9\x88\x4d\xd8\xe2\xd2\xe9\x94\x18\xd4\x3a\x5d\x63\x8c\x21\x02\x9f\x92\x51\x1b\x73\x6e\x11\xa7\x65\x29\xbb\x9c\x83\xbf\xa7\x2c\xa8\xd4\x95\x8d\xcd\x2b\x38\x86\x61\x93\x2e\x1d\x1a\x10\x37\x4a\x9d\x82\x72\x94\x27\x7c\xd2\xa1\xdf\xe4\x9a\xef\xdf\x66\x8e\x98\xb0\x27\xae\x9c\xad\x1e\xde\x6a\x91\x57\xb4\x6f\x29\xfb\x1f\xf1\x8f\x6e\xde\x8b\xee\x2b\xee\x2c\x3c\x6a\xc3\xa6\xd7\xde\x04\xa3\xdd\xb2\x47\x53\xbf\x29\x80\xd9\xb1\x76\xdf\x96\x50\x64\x33\x74\xbd\xd7\x0d\xae\x7a\x5f\x23\x55\x5b\x63\xd0\x9f\x4e\xc1\x6b\xcc\x71\x2a\x50\x57\xf5\x9b\x33\xf1\x5e\x33\x6f\xcf\x3b\x78\xba\x3d\xce\x46\x35\x56\x30\x77\x76\x31\x4f\x39\x06\xbf\x59\xc0\xf8\xf9\x35\xc5\x09\x7c\x0c\xf3\xc9\x60\x9a\x16\x30\xde\xd2\x7f\x67\x17\xb0\xe7\xa2\x19\x1a\x5a\xf6\xd9\xb8\xee\x27\x90\xcb\x3a\x37\x8e\xc1\xb2\x9f\xe1\x54\x1c\x9a\x3d\x4d\xe7\x65\x2c\x26\x8b\x79\xe7\x16\x30\xae\x88\x83\x4e\x44\xd1\x63\x19\x17\xa3\xdb\xa1\x26\x27\xa9\xfd\x2f\x75\x31\x3a\x6b\xf5\xb4\xff\xf7\x65\xb0\xe6\xbb\x17\x81\x62\xb8\x7f\x99\xbe\x95\x7a\xa7\x32\xe7\x75\x19\xaa\x42\xa9\x90\x6d\x48\x99\xc1\x05\x08\xa2\x15\x27\x4e\x77\xd6\x9b\xb0\x93\x2e\x54\x9a\x6e\x64\x44\x17\xb4\xe9\x51\xc7\x2c\xbe\x5c\xb4\x7d\x02\x5b\x9d\xd2\x2e\x44\xb3\x92\xcf\xa7\xeb\x54\x87\x1d\x70\x19\x3c\x03\x3a\xca\x32\x1d\x6d\x45\x6d\x0d\x9e\x5e\x80\xf8\x86\x6e\xb8\x80\xf1\xcb\x6d\x65\x0d\xf5\xcc\x9f\x5b\x5a\xb1\xc0\x67\x76\xcd\xf0\x79\x25\xf4\x28\xda\xca\x19\xbf\x64\xc6\x67\xc3\x01\x25\x09\x64\x30\x40\xb8\x78\x79\x21\x32\xdf\xef\xf9\xec\x45\x41\x69\x90\x81\xcb\x9f\x74\xd6\x9f\xe9\xcc\x9e\x9c\xf6\x9b\x4e\x6f\x70\xfa\xd4\x45\x37\x85\x89\xd1\xa9\x2e\x83\x8e\x66\x42\x5c\x3d\x20\x4d\xec\xed\x8f\x5e\xfe\x95\x82\x87\xbd\x28\x7f\xd3\x1b\xeb\x75\xc6\xe9\xe8\x56\x94\x3f\xfb\x75\xa0\x83\xd6\x77\x21\x66\xeb\x37\xd3\xfb\x95\x30\x51\xef\x96\xda\xb9\x52\x57\x0f\xd3\x57\xe2\x75\x6b\x0f\x3b\x5e\xb7\x96\xf5\x9f\xa4\xa4\x74\x6b\x25\x1d\xd8\x53\x4b\x91\xc0\xb0\x9d\x61\xcf\xa5\x0f\x06\x13\xe3\xc2\xe9\x94\x95\xef\x9c\x9b\xd1\xe3\x2a\xb8\xae\xf1\xec\x46\x9c\x03\xfa\x0c\xbc\xee\xb7\x18\x76\xc2\x1e\xfc\x53\x77\xa9\x4b\xba\x19\xde\xc2\x26\x86\xae\x05\x3e\x23\x17\x23\xa5\x7c\xff\xe1\xa1\xdf\x21\xfd\x57\x57\xfd\xf3\xc4\x25\xfe\xcd\x2c\x97\x25\xae\x43\x44\xf6\x61\x9e\x63\x51\xd1\x5a\x55\x03\xc5\x62\x9e\x4d\x51\x05\x97\x5a\xed\x15\xfc\x00\x8b\x0f\x63\xa2\xa1\xa1\xcd\x66\x31\x9f\xe4\x48\x93\x7b\x50\xa2\x1b\xdc\x53\x79\xf6\x7c\xf6\x4f\x00\x00\x00\xff\xff\x96\x3a\x4d\x6b\xfc\x07\x00\x00")

func assetsJsManagerMinJsBytes() ([]byte, error) {
	return bindataRead(
		_assetsJsManagerMinJs,
		"assets/js/manager.min.js",
	)
}

func assetsJsManagerMinJs() (*asset, error) {
	bytes, err := assetsJsManagerMinJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/js/manager.min.js", size: 2044, mode: os.FileMode(420), modTime: time.Unix(1543014287, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesRuLc_messagesConfigMo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\x4f\x6b\xd4\x40\x18\xc6\x9f\xae\xf5\x12\x50\x50\x2f\x0a\x82\xaf\x07\xab\x22\x53\x27\x59\x0b\x65\x76\xb3\x55\xb7\x2d\x88\x5d\x2c\x25\x7a\xf2\x32\xee\x4e\xb3\xc1\xcd\x4c\x9c\x24\x45\xa1\xa0\xa8\x78\x15\x0f\x5e\x3c\x79\xf1\x5e\xf1\xdf\xe2\x9f\x7e\x86\xc9\x17\xf0\xb3\x48\xb2\x6a\xf5\x3d\x0c\xcf\xf3\x3e\xcf\xfc\x78\x7f\x1e\x9f\x7f\x0d\x00\xf3\x00\x4e\x03\xe8\x02\x38\x0c\xe0\x2e\x66\x93\x01\x38\x02\xe0\x01\x80\x63\x00\x1e\x03\x38\x01\xe0\x15\x80\xf1\x1c\xf0\x0e\xc0\x49\x00\x47\x5b\xc0\x29\x00\xe7\x5b\xc0\x19\x00\xd7\x5a\xc0\xdc\x6f\x6e\x3d\x2d\x00\x87\x6a\xd1\x37\x7a\x3b\x89\x4b\x2b\x8b\xc4\x68\x0c\x1b\x37\xbf\xaa\xee\x95\x31\xa5\x66\xa4\x0e\x36\x3b\x6a\x62\xb2\xd9\x6e\x4b\x65\xc6\x16\x6c\x90\xc7\xc9\x88\x5d\x2f\xe3\x9c\x45\x46\xd0\x48\xed\x5c\xbd\x9f\x8c\x65\x6a\x16\x6d\xe9\x6d\xde\x8a\x58\xdf\xaa\x06\xcb\x56\x65\xa1\x04\x05\xdc\x5f\x66\xbc\xcd\x82\x36\x05\x6d\xb1\xb4\x74\x89\xb7\x39\xf7\x36\x64\x5e\xb0\xc8\x4a\x9d\x4f\x64\x61\xac\xa0\x9b\x0d\x83\x06\xa5\x95\xa9\x19\x19\xea\xfe\x07\xee\x79\x1b\x52\xc7\xa5\x8c\x15\x8b\x94\x4c\x05\xfd\xf5\x82\xb6\xca\x3c\x4f\xa4\xf6\x06\x37\x06\x6b\xec\x8e\xb2\x79\x62\xb4\x20\x7f\x91\x7b\x7d\xa3\x0b\xa5\x0b\x16\x3d\xca\x94\xa0\x42\x3d\x2c\x2e\x67\x13\x99\xe8\x0e\x0d\xc7\xd2\xe6\xaa\x08\x6f\x47\xeb\x6c\xf9\xa0\x57\xdf\xb3\xad\x2c\x5b\xd3\x43\x33\x4a\x74\x2c\xc8\xdb\x9c\x94\x56\x4e\xd8\xba\xb1\x69\x2e\x48\x67\x8d\xcd\xc3\x76\x87\x66\x32\xd4\xe7\x7c\x1e\x86\x3e\x2d\x2c\x50\x2d\xf9\xd9\xd0\xf7\x69\x85\x38\x89\xc6\xf7\xc2\xe0\x4f\xd4\x0d\xaf\xd4\xf2\x42\x53\xeb\xfa\x9c\x76\x77\x67\x5f\x7a\x61\xc0\x2f\xd2\x0a\xf9\x24\x28\xe8\xc0\xbd\x71\xfb\xee\x47\xf5\xdc\x4d\xdd\x87\xea\x59\xf5\xc4\xed\x55\x2f\xdc\xb4\x7a\x09\xf7\xd6\x7d\x72\x9f\xdd\xd4\x7d\x27\xb7\x5f\x3d\x75\xdf\xdc\x9e\xfb\xe8\xbe\xba\xe9\xbf\x49\xdd\x77\x5f\x9a\xf7\x7d\xd3\xaa\xf3\x5f\x01\x00\x00\xff\xff\xa7\x17\xf2\x1f\x61\x02\x00\x00")

func localesRuLc_messagesConfigMoBytes() ([]byte, error) {
	return bindataRead(
		_localesRuLc_messagesConfigMo,
		"locales/ru/LC_MESSAGES/config.mo",
	)
}

func localesRuLc_messagesConfigMo() (*asset, error) {
	bytes, err := localesRuLc_messagesConfigMoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/ru/LC_MESSAGES/config.mo", size: 609, mode: os.FileMode(420), modTime: time.Unix(1543014294, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesRuLc_messagesManageMo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x4f\x4b\x6f\x1b\x55\x14\xfe\xa6\x2e\x2f\x97\x67\x85\x84\x84\x58\x9c\x0a\x11\x81\xd0\x94\x19\x9b\x4a\xd1\x24\x13\x53\x92\x14\x10\x0d\x54\xc1\x94\x2d\x57\xf6\x8d\x33\x62\x3c\x63\xdd\x3b\x13\x88\x14\x89\x38\x2e\x54\x60\x44\x11\x0f\x75\x05\xa8\x3b\x96\xa6\xc4\x74\x68\x6b\x77\xcf\xea\xdc\x35\x88\x9f\xc0\x2f\x40\x02\xcd\xa3\x0f\xee\xe2\x9e\xf3\x7d\xe7\x3b\xe7\x7c\xe7\xaf\xe3\x47\xbf\x03\x80\x87\x01\x3c\x03\x60\x0f\xc0\x13\x00\xfe\x44\xf9\xc8\x02\x1e\x03\x70\xc2\x02\x1e\x01\xe0\x5a\xc0\x7d\x00\x5a\x16\xf0\x34\x80\x75\x0b\x78\x00\xc0\xfb\x16\x70\x0c\xc0\xb6\x05\x1c\x05\x90\x54\xba\xdd\x0a\x7f\x5c\xe1\x0b\x16\xf0\x20\x80\xcf\xad\x52\xfb\x8d\x05\x2c\x00\xf8\xe7\x08\xf0\x24\x80\x67\x6b\xc0\xa3\x00\x16\x6b\xc0\xeb\x00\x5e\xab\x01\x4f\x01\xf8\xaa\x06\x3c\x0e\xe0\xa7\x2a\xfe\x56\xc5\xdf\x6b\xc0\x71\x00\x7f\x54\xf8\xef\x2a\xfe\x5b\x03\x2c\x94\xde\x6e\xbf\x7c\x7f\xad\xca\x8f\x55\xf1\xa1\xea\xf6\xdc\xd3\xfd\x28\xbd\xd6\xab\xda\x91\xfc\x3b\x1d\x86\xf1\x87\xb2\x4b\x3b\x22\x4c\xa5\xf6\x70\x5a\x49\xda\x8d\x53\xd2\xa9\x92\x2d\xac\x86\xb1\x96\x58\x8d\xa3\xad\x40\xf5\x49\x8b\x1d\x49\x9d\x1c\xf4\x52\x25\x92\x20\x8e\xb0\x26\xb7\x44\x1a\x26\x58\x93\xba\xa3\x82\x41\xc1\xbd\x25\xfa\x12\x9b\x52\xcb\x04\xef\x88\x1d\x89\xf3\xf9\x68\xbc\x27\x92\xce\xb6\x54\x1a\x9b\x72\x10\xab\xc4\xde\xd0\xbd\xa0\x6b\xbf\x9a\xf6\xb4\xdd\x8e\x3d\xea\xca\x9d\x57\x3e\x08\xb6\x45\x3f\x3e\xa9\xd2\xfa\xb9\xb7\xdb\xf6\xaa\x92\xc5\x0e\x7b\x4d\x24\xd2\xa3\x86\xe3\x2e\xda\x4e\xd3\x6e\x34\xa9\xd1\xf4\x4e\x9d\x7a\xd1\x69\x3a\x4e\xfd\xac\xd0\x89\xdd\x56\x22\xd2\xa1\x48\x62\xe5\xd1\x9b\xc5\x0c\xda\x48\x95\xe8\xc7\xdd\x98\x96\xff\x37\x78\xa5\x7e\x56\x44\xbd\x54\xf4\xa4\xdd\x96\xa2\xef\xd1\x1d\xec\xd1\x66\xaa\x75\x20\xa2\xfa\xc6\x1b\x1b\xeb\xf6\x79\xa9\x74\x10\x47\x1e\xb9\x27\x9d\xfa\x6a\x1c\x25\x32\x4a\xec\xf6\xee\x40\x7a\x94\xc8\x8f\x92\x97\x06\xa1\x08\xa2\x25\xea\x6c\x0b\xa5\x65\xe2\xbf\xdb\x3e\x63\x2f\xde\xd5\xe5\x7e\xb6\xa4\xb2\xd7\xa3\x4e\xdc\x0d\xa2\x9e\x47\xf5\x73\x61\xaa\x44\x68\x9f\x89\x55\x5f\x7b\x14\x0d\x0a\xa8\xfd\xe6\x12\x95\xa9\x1f\x3d\xe7\x3a\xbe\xef\xd2\xc2\x02\xe5\xa9\x73\xc2\x77\x5d\x6a\x91\x43\x5e\x81\x57\xfc\xc6\xed\xd2\xb2\xff\x72\x9e\x3e\x5f\xc8\x96\x5d\x87\xf6\xf6\xca\x96\x15\xbf\xe1\xbc\x40\x2d\x72\xc9\xa3\xc6\x12\xf8\x5b\x9e\xf3\x2d\x33\x32\x43\x73\xc0\x19\xdf\x34\x63\x9e\x12\x5f\xe3\x19\x4f\xcc\x45\x9e\xf2\x8c\x33\x73\xc9\x03\x7f\x6d\xc6\x64\x46\x7c\x95\xa7\x66\x3f\xa7\xcd\xb8\x05\xbe\xcc\x13\xbe\x6e\xf6\xcd\xd8\x1c\x98\x2f\xc0\x3f\xf2\x9c\x0f\xcd\x41\x25\xfa\x95\x0f\xcb\x7e\x9e\x92\x19\xf2\xdc\x7c\x62\xf6\x79\xc2\xb3\x3b\x24\x5f\xe7\x39\xcf\xcc\x05\xce\xf8\x17\x33\xca\x8b\xe6\x53\xce\x38\x2b\x06\xd9\x66\xc4\x37\x79\xce\x37\xcc\xc5\xa2\x29\x33\x5f\x82\x7f\xe0\x5b\x9c\x99\x61\x49\xf0\x14\xfc\x3d\x4f\xf8\x1a\x5f\xbd\x4b\x5c\xe1\x9f\xcd\x3e\xcf\xcd\x90\xb3\xd2\xd3\x95\x7b\x16\x57\xd4\xe5\x7b\x8f\x2b\x9b\x6e\xf0\x94\x0f\xcd\x25\xf3\x59\x41\xfc\x17\x00\x00\xff\xff\x0b\x24\x32\x8e\x10\x04\x00\x00")

func localesRuLc_messagesManageMoBytes() ([]byte, error) {
	return bindataRead(
		_localesRuLc_messagesManageMo,
		"locales/ru/LC_MESSAGES/manage.mo",
	)
}

func localesRuLc_messagesManageMo() (*asset, error) {
	bytes, err := localesRuLc_messagesManageMoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/ru/LC_MESSAGES/manage.mo", size: 1040, mode: os.FileMode(420), modTime: time.Unix(1543014294, 0)}
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
	"templates/views/manager.html": templatesViewsManagerHtml,
	"assets/js/manager.min.js": assetsJsManagerMinJs,
	"locales/ru/LC_MESSAGES/config.mo": localesRuLc_messagesConfigMo,
	"locales/ru/LC_MESSAGES/manage.mo": localesRuLc_messagesManageMo,
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
			"manager.min.js": &bintree{assetsJsManagerMinJs, map[string]*bintree{}},
		}},
	}},
	"locales": &bintree{nil, map[string]*bintree{
		"ru": &bintree{nil, map[string]*bintree{
			"LC_MESSAGES": &bintree{nil, map[string]*bintree{
				"config.mo": &bintree{localesRuLc_messagesConfigMo, map[string]*bintree{}},
				"manage.mo": &bintree{localesRuLc_messagesManageMo, map[string]*bintree{}},
			}},
		}},
	}},
	"templates": &bintree{nil, map[string]*bintree{
		"views": &bintree{nil, map[string]*bintree{
			"manager.html": &bintree{templatesViewsManagerHtml, map[string]*bintree{}},
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
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}
