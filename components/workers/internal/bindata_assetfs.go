// Code generated by go-bindata.
// sources:
// templates/views/manager.html
// assets/js/manager.min.js
// locales/ru/LC_MESSAGES/config.mo
// locales/ru/LC_MESSAGES/manager.mo
// locales/ru/LC_MESSAGES/workers.mo
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

var _templatesViewsManagerHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\xdd\x6e\xdb\x3a\x12\xbe\xcf\x53\x10\x2c\xb0\x48\x2e\x14\x35\x6d\x17\x58\xb4\xb6\x17\x41\x77\x17\x2d\x90\x2d\x16\x69\x16\xc5\xb9\x0a\x28\x71\x64\xd1\xa1\x48\x1d\x72\x24\xdb\x30\xf2\xee\x07\xd4\x8f\x2d\xd9\x94\x1b\xbb\x4d\x91\x06\xb9\x0a\x4d\x72\x7e\xf8\xcd\x70\xe6\x93\xa2\xd5\x8a\x70\x48\x84\x02\x42\x63\xad\x10\x14\x52\x72\x7f\x7f\x32\xe2\xa2\x24\xb1\x64\xd6\x8e\x69\xce\xa6\x10\xa0\x40\x09\x74\x72\x42\x08\x21\xdd\xc5\x6a\xfe\x56\x42\x82\xcd\x62\xb5\x21\x7d\x3b\x59\xad\x88\xb8\xf8\x87\x22\xf4\x9b\x36\x77\x60\x2c\x25\xe7\xe4\xfe\x7e\x14\xa6\x6f\x1b\x25\x21\x17\xe5\xe4\x64\x40\xa1\x11\xd3\xb4\xa7\xb1\xb3\x23\xd1\x26\x0b\xa6\x46\x17\x39\xc9\x0b\x29\x83\xed\xbd\xdb\xfb\x85\xca\x0b\xac\x05\xb6\x76\x55\x3b\x25\x8b\x40\xee\xce\x57\x6b\x95\x28\xc1\x65\x0e\x63\x1a\xa7\x10\xdf\x45\x7a\x41\x5b\xbd\x33\x1b\xd8\xb9\xc0\x38\xa5\x44\xf0\x31\x65\x05\x6a\x03\x89\x01\x9b\x52\x12\x4e\xc8\x1a\x80\xcb\xee\x82\x03\x61\xd7\x87\xd0\xe3\x44\x03\x90\xe7\x67\x33\x6c\x11\xec\x9e\x35\x96\xc0\x4c\x22\x16\x74\xe2\x5b\x35\x7a\x4e\x50\xe7\xb7\x28\x24\x58\x4f\x30\x99\x12\x19\x43\xe0\x24\x91\x22\xff\xac\xfe\x20\xb1\x96\x81\x9c\x06\xef\xaa\x41\xc6\x9b\x81\xcd\x9a\xc1\xc2\x06\x17\x6f\x06\xa2\xe4\x6c\x04\x16\x19\xda\x7d\xa1\x89\xb5\xa2\x93\x91\x58\x87\x96\x91\x84\x05\x1c\x58\xe2\xfe\xfe\xbd\x3a\x86\x98\x6c\x21\xb1\xad\x24\xd6\x85\xcb\x5a\x17\x03\x29\x2c\x82\x02\x63\x83\x7a\x72\xf2\xda\x27\xdb\xcd\xce\xab\x56\xc2\x85\x86\xd6\xfe\xf6\xf2\xd4\x8f\xfd\x13\x45\xae\xb0\x60\xec\x71\xd0\xcd\xeb\x6b\xfa\x60\xe0\x3a\xd7\xfa\xb7\x87\x0d\x99\xbd\x3b\x12\xb6\x4a\xf4\xc1\xa0\xdd\xb8\xdd\x87\x41\xd6\xbb\xc7\xbd\x1c\xf7\x5c\xe0\xc5\x6d\xce\x14\xc8\x01\x84\x16\xb7\xdd\x2a\xbe\x71\xf0\x8d\xf7\x06\xd6\x48\x10\xff\x4d\xa9\xfc\x7e\xb3\xa5\xa8\x90\xad\x16\xc5\x4a\xa2\x58\x19\x31\x53\x57\x67\x52\xb9\x75\x8b\x5a\x4b\x57\x41\x7d\x55\x58\x4c\x46\x6c\x83\xaf\x94\x2c\xb7\x10\x48\xa1\xee\x76\xc3\x15\xa7\x50\x1a\xad\x02\x57\xcf\xeb\x68\xb1\xc9\x28\x94\xc2\xab\xb6\x15\xe5\x46\xe7\x5c\xcf\x95\xc7\x78\xb5\x93\x91\xd4\x40\x32\xa6\xaf\xe8\xb6\x44\x80\x7a\x3a\x95\x40\x09\x67\xc8\x9a\x1f\x1d\x7d\xc4\x68\xf7\x3b\x2a\x10\xb5\xa2\x84\x19\xc1\x02\x58\xe4\x4c\x71\xe0\xce\x63\x69\x61\xf7\x08\x73\x03\x2a\x4e\x37\xee\xfb\x7d\xda\x00\xba\xf6\x25\x03\x55\xb4\x26\xab\xb1\x5f\xb4\xc5\x74\x70\xb1\x77\xe6\x19\x2b\x99\x8d\x8d\xc8\xf1\x7d\xa9\x05\x3f\x7d\x7d\xb6\x75\xd8\x4c\x73\x26\xdb\x39\x66\xa6\x80\x63\xfa\xaa\x3b\x59\x8d\x6b\x92\x30\xa6\xeb\x94\xf9\xa8\x55\x22\x4c\x46\x0c\x64\xba\x04\xc2\xa4\x24\xb2\x9f\x45\x3d\xf1\x48\xf3\x65\x47\xfa\x9b\x90\x92\x44\xe0\x95\x26\xb0\x88\x21\x47\xc7\x5d\x58\x21\x71\xbd\x70\x4e\x2e\x0d\x90\xa5\x2e\x88\x2d\x0c\xfc\xd3\x63\x24\x66\x52\x46\x2c\xbe\xeb\x5c\xa4\xeb\xca\xc0\xe9\xd9\x87\x3d\x60\xae\x31\x5b\x07\x72\x2a\x97\x79\xea\xea\x09\x59\x8f\x02\x34\xcc\xa6\xdb\x17\xe7\x7a\xed\xff\x40\xfb\xef\xe9\x1f\x4a\x86\x7a\x71\x28\xa4\xa3\xb0\xf0\x70\x98\xdd\xfd\xbb\xfb\xf6\xb2\x87\x8d\x58\xff\x67\xaf\xa8\xb4\xbc\x71\x58\x2f\xb2\x48\x42\x60\xc0\xe6\x5a\x59\x51\x6e\x57\xa0\x6a\x7b\xb5\xa7\x27\x40\x6a\xb1\x54\x97\x60\x9a\xb1\x45\x23\x72\xe0\x43\x77\x18\x53\x60\x7c\x68\xcd\xec\xc1\x15\xd3\x4d\x8d\xfe\xfc\xaf\xb6\xc0\x61\xfa\x40\x91\x2f\x2c\x83\x83\x85\xfe\x5d\x82\x42\x7b\xb0\xd8\x7f\x84\x81\xa3\xa4\x2c\x92\x44\x18\xe0\x07\xcb\x5e\xb1\xa3\x45\x2f\x63\x14\x5a\x7d\xd7\xdd\x51\x38\x14\x1d\x27\xe3\x8d\xe9\x28\xac\x32\xe2\x47\x28\x73\x87\xf5\x3c\x62\x2b\xad\x28\xd9\x76\x49\xd8\x7e\x26\x7a\xe9\xa4\xcf\xb8\x93\x76\xd9\xb5\x4d\xf5\xfc\x87\xbb\x0c\x2c\x21\xd0\x39\xa8\xed\xac\xfa\x9a\xea\x79\xd5\x26\xb1\xa5\x99\x8f\xd3\x6c\x7e\x2e\x24\xa9\xe0\xbe\x86\x70\x38\x24\xb1\xd4\x55\x82\xf4\x30\xf9\x24\x38\xfc\x42\x4c\xd6\x59\x26\x4a\xc1\xc1\xd0\x81\xdb\xd4\x91\x78\x7a\x14\x6d\xde\x2d\x4e\x03\xdc\xa9\xd9\xf3\x0c\x98\x13\x79\x92\x81\x6b\x2f\x07\xe3\x43\x64\xa7\x67\x6c\x2f\xb4\xb9\x2c\x76\x1a\xd0\x25\xe7\x4d\x9c\x5f\x38\xe9\xf3\xe4\xa4\x1f\x0d\x30\x3c\x82\xb2\x7d\x45\x86\xc5\xe1\x04\xf3\x4a\xc7\x77\x47\x58\xbb\x61\xf6\xee\xf9\xb1\xca\xba\xd7\x3c\x1e\xa7\x6c\xf4\xf7\xaf\xf4\xcd\xa6\xc1\xbd\x30\xca\xe7\xcd\x28\x1f\xa9\xf1\x77\x28\xd2\x40\xdb\xaf\x76\x3c\x83\xa6\xff\xd2\x9a\x7e\xbb\xd7\x25\xff\x33\x42\x1b\x81\xcb\x83\x05\xaf\x21\x07\x76\xc4\x8b\x96\x5a\x8e\x08\x85\x60\x4a\x26\x0f\x6f\x6d\x22\x03\x5d\xe0\xaf\x6c\xdc\xe6\xe9\x37\xfc\x4b\x44\xc8\xf2\x23\xc2\x71\x29\xa5\x9e\x13\xeb\x0e\x79\xe4\xdb\x2f\x7b\x24\x40\xd5\xfb\xaf\x63\x85\x9f\x34\x57\x69\xca\x47\x55\xe5\x49\xc2\x38\xf4\x9f\xcc\xdd\xc3\x87\xab\x24\x42\x71\x58\x8c\x69\x70\xd1\x36\x37\x2e\x98\xd4\xd3\xa6\x9f\xa6\x82\x73\x50\x63\x8a\xa6\xf0\x7d\x20\x50\xb7\x90\x46\xc2\x5f\x13\x9b\x2e\xf3\xdd\xba\x58\xef\x73\x68\xb8\x07\xb4\x5d\x40\xea\x36\xdf\xfc\xbb\xbe\xed\xf9\xeb\x5a\xad\x6d\x4b\x13\xb8\xb0\x99\x58\x2b\xf4\x1e\xe3\x6f\x28\x32\xb0\x1f\x46\x61\xad\xc6\x63\x2c\x7d\xd7\x77\xab\xe1\x6e\x9e\x27\xac\x0d\x29\x7b\xb7\x37\x58\xfe\xe3\x46\x9a\x2f\x7d\x87\x4d\xb4\xc9\x06\x32\xc6\xfb\xad\xc4\x5e\x32\xc2\x22\x90\x24\xd1\xa6\x17\xfb\xe6\xbf\x9a\x1b\xe2\xa7\xd0\x68\x19\x54\x9b\x3b\x07\xfd\x52\x64\x11\x18\xa2\x60\xde\x9e\xf7\x7d\x7b\xe0\x3d\x9f\x56\x54\x76\xbb\x9f\x57\xa8\x4a\x0d\x25\x99\x50\x63\x7a\x41\x7b\x07\x68\x4c\xef\x64\x67\xeb\x61\xc9\x64\x01\x4e\x6a\xe8\x0a\xed\xe0\x5c\x4f\xef\x62\xf8\xa0\x90\x24\x5a\xe3\xe1\x19\x18\xa1\x22\x11\xaa\xa0\xf9\x97\x95\x3f\x17\x3b\x6d\xa0\x4e\xd8\x1a\xc8\xc1\x2c\xec\x19\xb4\x45\x94\x09\xdc\x31\x68\x8b\x38\x06\x6b\xbf\x67\xf0\x92\xf3\x7d\xe6\x1e\x5a\x59\x56\x2b\x02\x8a\x3b\xde\x76\xd2\xf9\xb6\xc8\x5d\x5a\xda\x92\xb9\xd5\xca\x95\x53\x14\xf1\xa7\x9b\xff\x5e\x91\xd3\x7a\xfc\xff\xeb\x2b\x42\x43\xce\x6c\x1a\x69\x66\x78\xc8\xac\x05\xb4\x61\x09\x8a\x6b\x63\x43\xe7\x7b\x55\xef\xec\xb9\x02\x0c\x22\x1b\xc6\xb6\x9e\xbd\xa9\x67\x23\xad\xd1\xa2\x61\xf9\x79\x26\xd4\x79\xec\xce\x5b\xf1\xfb\x33\x67\xd5\xef\xd4\xcc\xfe\x44\x97\xc2\x99\x0d\x67\x7f\x16\x60\x96\xe7\x1d\xaf\x9c\x2f\xb3\x9e\x2b\x3f\x0f\x80\xd9\x9e\xf3\x1f\x60\xf3\x34\x37\x42\x21\xa1\x21\x25\xe7\x1f\x75\x96\x6b\x05\x0a\x1d\x45\x23\xb4\xf5\x61\x66\xc3\x8c\x29\x36\x05\xe3\x14\x9f\x11\x57\x23\xfb\xb8\xfe\x15\x00\x00\xff\xff\x3a\xad\xbe\x27\x47\x26\x00\x00")

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

	info := bindataFileInfo{name: "templates/views/manager.html", size: 9799, mode: os.FileMode(420), modTime: time.Unix(1543013830, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsJsManagerMinJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x59\x6d\x6f\xe3\x2a\xf6\xff\x2a\xbe\xcc\xfc\x6b\xa3\x38\xce\xf4\xfe\xf7\xc5\xca\xb1\x5d\x8d\x66\x74\xa5\x91\x46\xbb\x57\xd3\x4a\xfb\xa2\xad\x2a\x62\x4e\x12\x6e\x09\x44\x80\x9b\x89\x22\x7f\xf7\x15\xe0\xd8\x4e\x9a\xf4\x61\xfa\xb0\xf7\x8d\x85\xe1\x70\xf8\x01\xe7\xfc\x38\x07\x3e\x46\x54\x96\xd5\x02\x84\xc1\x89\x02\x42\xd7\xd1\xb4\x12\xa5\x61\x52\x44\x78\xf3\x31\x42\x1f\x56\x52\xdd\x82\xd2\x43\x3d\x97\x2b\x84\x93\x92\xb3\xf2\xf6\x88\x4c\x90\x18\xa2\x6f\x9d\x64\x3a\x27\x3a\x62\xc9\x8c\xaf\x97\x73\x56\x4a\x31\x84\x35\x0c\xe5\x12\x04\x6e\x75\xe0\x1a\xc7\xfd\x01\xe6\x8c\xc2\x0b\x07\x28\xb9\xd4\x70\x7c\x04\x42\x69\x30\xa9\x8c\x91\xe2\xd2\xac\x97\x90\xeb\x6a\xb2\x60\xe6\xfa\xe0\xa0\x09\xf9\x8b\xfc\x8c\x36\x56\x2e\x45\x7f\xfe\xfb\xfc\x02\xc5\x95\xe2\x29\x1a\x35\xca\x46\x67\xc4\xc9\xe6\x3d\xe5\x28\xa6\xc4\x90\x74\x53\xca\x4a\x98\x74\x6f\xe4\xa1\xab\x45\x38\xb9\x23\x3c\xc2\x75\xac\xab\xb2\x04\xad\xd3\x6a\x49\x89\x81\x1a\xd7\x78\x7c\x47\x54\x60\xc8\x84\xc3\x77\xa6\x0d\x08\x50\x3a\xb7\x4a\xf8\xf6\xcf\x37\x22\x9c\x48\x11\x21\xaa\xc8\x2a\xa1\x06\xc5\x2d\x6a\x88\x35\x18\xc3\xc4\x4c\xe3\xcd\xb6\x94\xfc\xa5\xa5\x38\x39\xd9\x51\xd3\x22\x31\xf0\xd3\x44\x3b\x92\x89\x82\x52\x2a\xaa\x2f\xa4\x21\x1c\xd7\x38\xf9\x4a\x0c\xb9\xb0\xa3\x46\x1b\x4e\xc4\xac\x22\x33\x48\x37\x7e\x21\x28\xd1\xf3\x89\x24\x8a\x8e\xec\xac\x1d\x34\x3d\x62\xa7\xff\x14\x4e\x13\xaa\x63\xbb\x82\x5b\xe1\xfd\x55\xd3\x86\x18\x7d\x02\xc2\x30\xb3\xce\x5b\x64\x7e\x01\xcf\x55\x99\x22\x5b\x40\x75\x5c\x4a\x5e\x2d\x84\x4e\x2f\x37\x6e\x69\x11\xa3\xa8\x8e\x9b\xb2\x20\x0b\xe8\xfe\xe0\x0e\x84\xd1\x28\x56\x20\x28\xa8\xb4\x5d\x15\xdb\x1a\xdb\x6d\x8c\x95\x5c\xe1\x8d\x5d\xe3\x52\x0a\x03\xc2\xe4\x08\x8d\xa7\x52\x45\xb6\xca\xf5\xfe\x46\x03\x26\x02\xdb\x01\x2b\xb9\x4a\xb8\x2c\x6f\x81\x9e\x35\xd2\x83\x3c\xcc\xf4\x92\x88\xa0\xe4\x44\xeb\x1c\x71\x32\x01\x1e\xb8\xef\x90\x89\xa9\x44\x45\x38\xb0\x5d\x2f\x1b\x55\xd7\x03\x94\x8d\x6c\x87\x22\x40\x69\x4f\x07\x09\xe6\x0a\xa6\x39\xfa\x80\x02\xc3\x0c\x87\x1c\xfd\x80\x85\xbc\x63\x62\x16\x6c\xd7\x01\x1d\x1d\xc3\x81\x1b\x1a\x39\x9b\xd9\x8e\x0b\x49\x09\xdf\xd6\x11\x35\x03\x93\xa3\x0f\xfd\x4a\x57\x1e\x36\xc3\x7c\x91\x62\xca\xd4\x22\x50\x76\x38\x68\x07\x0b\x3e\x84\x03\x3b\x5b\x46\x07\x28\x98\xca\x66\x29\x02\xb4\x37\x99\x70\x47\x65\x49\x38\x9f\x90\xf2\x36\x47\xed\xde\xb9\x59\x40\x74\x15\x76\xea\xc2\x38\x08\xd1\xa0\x51\x31\x40\x21\x1e\x5f\xa1\x62\x5f\x31\x0a\x7e\x66\x23\x52\x04\x68\xac\xc0\x54\x4a\x6c\x77\xa7\x6e\x77\x76\xca\x14\x68\xd4\xff\xd5\xe6\xc6\x56\xd2\x1b\x62\x0e\x6e\x38\xe0\x4d\xa3\xcc\xfe\x9c\xd9\xcf\x85\x3c\x37\x8a\x89\x99\x6f\x4e\x11\xea\xf4\x73\xf2\x5a\xfa\xa4\xa2\xa0\xac\x23\xa4\xbf\x9d\x7a\x2e\x10\x15\xe7\x8f\x99\x64\xa3\xba\x67\x72\x08\xa5\x61\x46\xd9\xdd\xd6\x0e\x26\x46\x0c\x67\x4a\x56\xcb\xa0\x2d\x0d\x7f\x6a\x54\x64\x9e\xcd\x02\xc7\x66\xc8\xff\xa0\x5e\x27\x27\x4e\x89\x98\x81\x72\x45\x4b\x94\x6f\x6c\x43\xa1\xb3\x21\xc2\x79\xd0\x38\xe4\xb3\xcd\x26\xbc\x0a\xf1\x18\x15\x19\xdb\x4e\xa4\xe5\xf8\xa0\x63\x7b\xa3\x88\x9e\xef\x7a\x50\x07\x47\xef\x63\x28\xb2\x11\x2b\xb2\x91\x5f\xa0\x22\x1b\x51\x76\x57\x84\x75\x7d\x1d\xbb\x0d\x4b\x2f\x2f\x4f\x63\x44\x74\x89\xae\xe3\xcb\xdf\x9b\xd2\x75\x8d\x63\xc7\x69\xff\xf1\xcc\x95\xf7\x8f\xa1\x17\xf1\xf0\xf6\x4c\xf8\xbb\xb1\x70\xd3\xf8\x3c\x0e\x2e\x15\x10\x03\xf4\x09\x5e\xb3\xeb\x30\x9d\xf7\x59\x0c\x55\xcf\xbd\xbd\x0b\xdc\x57\x38\xe5\x64\xd6\x2a\xb4\x3f\x67\xe8\xbb\x17\x4d\xd1\x1f\x0a\xa0\x73\x68\xe7\x73\x14\xa6\xa4\xe2\xe6\x8b\x27\x13\xeb\x9f\xcf\x72\xcf\xdd\x83\xe2\x49\xce\x18\x8e\xbb\xd9\x12\x17\xa8\x9c\x9c\x44\x3d\xe2\x7f\x8a\xaf\x36\x41\x81\x2b\x97\x4c\x95\x1c\x82\x36\xe2\x69\x9d\x54\xdf\xe6\xe8\x2a\x0c\x06\x01\x0b\x06\xc1\x55\xf8\x98\xab\x6c\x23\xaf\xd6\x5b\xce\xe7\x72\xe5\xd4\x5e\x85\x3a\xa0\x60\x08\xe3\xfb\x2e\x12\xe2\xf8\x99\xc8\xdf\x82\x65\x6e\x19\xe7\x81\x37\x4b\xcb\x30\x6e\x5d\x2d\x41\x1c\xe1\x94\xc6\x80\x7b\x8c\xd2\xf6\xf8\x65\x4a\xf1\x3a\x1f\x27\x90\xdf\x5b\x02\xf9\xd4\x11\xc8\xf8\x1e\x67\x04\x66\x22\xe9\xba\x61\x0e\x17\x74\xa2\xb8\x59\xd1\x2e\xb4\xed\x53\x09\xde\x40\xb2\x54\x8e\xc6\xbe\x7a\x8b\x8e\x7c\x9c\x38\xc9\x3f\x46\x66\xce\x34\x4e\xa6\x4c\xd0\x08\x31\x84\xed\x69\x92\xf7\x39\x2b\x51\x72\x15\x6d\xc5\x5c\x6c\xac\x4d\x84\x8c\x42\x18\x8f\xd9\x34\x9a\x24\x73\xa2\xbf\xd8\x05\x89\xd0\x01\x73\xc1\xde\x07\x9c\xbd\x59\x6a\xb6\x8b\x19\x61\x07\x73\x3c\x49\xfc\x09\xf0\x40\xef\x84\x50\x7a\xb0\xd9\x01\xf1\x68\x93\x72\xce\x38\x8d\xc2\xcc\x2f\xce\x8a\x51\x33\xcf\xd1\xe9\xa7\x4f\xff\x87\x8a\xcc\xa8\x22\x33\xb4\xc8\x2a\xde\x86\x43\x4c\x1b\xef\x70\xa8\xc8\x38\xbb\x5f\x3d\x64\x06\x16\xa8\xd8\x89\xd3\x96\x15\xe7\x43\xc5\x66\x73\x13\x58\xa6\x1d\x2e\x2a\x03\x34\xd0\x0b\xc2\x39\x2a\x32\x58\x14\xe1\xc0\x4e\xc9\x99\x49\x36\x82\x45\xd1\x44\x6d\x99\x36\x4a\x8a\x59\xf1\xed\x6b\x36\x6a\x8a\xd9\x44\x05\xa3\x22\x1b\x71\xf6\xfa\xc3\xdb\x60\xf6\x30\x80\x7f\x91\x05\xbc\x0b\x84\xa5\x62\x52\x31\xb3\x3e\x0c\xe3\xcf\xa6\xf5\x5d\xa0\x28\x58\x02\x31\xfa\x30\x92\x1f\xbe\xf1\xad\x81\xd0\x4a\x11\xeb\x84\x17\xf2\x07\x10\x6a\x0d\xb4\x39\xc0\x7a\x08\x6f\x98\x30\xa0\xee\x08\xc7\x0f\x41\xfd\xd6\x08\xfd\x6f\x11\x1b\xb6\x00\x59\x99\x23\x48\x2f\x7c\xeb\x9b\x43\xec\x47\x02\x0e\x56\x13\x41\xdc\x90\x63\xc8\xbe\x78\x81\xb7\x46\xe6\xd1\x68\x43\x94\x47\x73\x76\x1f\x6a\xd7\x68\x03\xfe\x23\x70\xcf\xbd\xd0\xbb\xb8\x89\x8f\x9d\x8e\xe2\x30\xd5\x9b\x3b\x89\x83\x41\x8c\x81\xc5\xf2\x98\xbb\x7e\x6e\x5a\xdf\xdf\xb4\x08\xe7\x72\x75\xe3\x76\xed\xb8\x79\x7d\xb6\x42\x81\x13\x7a\x1f\x13\xf3\xf9\xeb\x83\x86\xb6\x2f\xf2\x80\xb9\xfd\x61\x45\x03\xfd\x3e\x46\xe7\xd1\xb9\x7c\xf9\x41\xfc\x7b\x12\x1e\x3e\x3a\x04\xff\x3b\x79\x04\xfd\xa8\xe2\x45\x36\xb2\x71\xc0\xc8\x06\x04\x23\x17\x29\x14\x08\x27\x36\x52\x8a\x70\x0d\x5c\x43\xf0\x60\x38\xd2\xc4\x1b\x47\xe3\x11\x1f\xae\x74\xe1\x48\x32\x67\x14\xa2\x9d\xfb\xb8\x0b\xa2\x6f\x7d\x06\x68\xa7\xf7\xc2\xfc\xcf\xa9\xf8\xdb\x65\x7f\x0e\xd5\x4b\xee\xdf\xb6\xd1\x43\x57\xd3\x1c\xe2\xfb\x15\xed\x99\x79\x3f\xc7\x13\xba\x4b\x19\x8f\x9d\x65\x42\xf7\x92\xc7\xe6\x4c\x7b\x05\x4d\xdd\x31\xf4\xb2\x5c\x56\x3d\x59\xc9\x63\xd7\x52\xaf\x9f\x18\xa3\x2d\x51\x77\x4a\x77\x29\xf2\x05\x53\xdf\xa7\xac\xd7\xba\x97\x7b\x25\x8d\xcf\x4a\xfd\xbd\xc6\x37\xbb\x81\xf3\x29\xbd\x67\xac\xd7\xbc\x8d\xb3\x6a\x9f\x92\x27\x3b\x57\x7f\xd5\x2c\xd9\x6a\x7c\x4e\x8e\xfc\xff\xbd\x4b\x36\xff\xfe\x91\xf7\xde\x5e\x76\x9f\x40\xdc\x43\x4c\xa2\x80\x4b\x42\xa3\xdd\x3b\xb9\x43\x4d\x8e\xaa\x77\x1b\xea\x98\x54\x46\x2a\x98\x2a\xd0\xf3\xdc\xee\xbb\x4b\xcc\x7b\x95\x87\x5e\x80\x6c\xd6\x9c\x94\x73\x70\x17\xb2\xb6\x53\x9e\xe7\xbd\x2e\x27\x27\x91\x47\x1e\xe1\x1d\xf5\x2b\x26\xa8\x5c\x25\x1a\xda\xa8\xbf\x91\x8b\x4f\xe1\x1f\x18\x3b\xb3\xfb\x6d\x5f\x53\xd3\xa9\xe4\x40\x54\xdb\xad\x27\x82\xef\xcd\x00\xd7\x38\x6e\x7a\xed\x5d\xa5\x76\x2b\xb9\x6d\xf8\x46\xf1\xa6\x94\x42\x4b\x0e\x09\x97\xb3\x88\xa8\x99\x7b\xf0\xd3\x38\x7e\xfa\x33\x57\xf7\x7a\xd4\x18\x6e\xf3\xd6\xc5\x68\xda\x8d\x13\xfb\xfb\xd6\xf4\xb3\x52\x64\x9d\x90\xe5\x92\xaf\x23\xe7\x67\xdd\x90\x89\xe6\xac\x84\xe8\xb4\xf7\x02\xf6\xc4\xad\xaf\x6b\x5c\x6f\xe7\xbc\x73\xd5\xd3\xcd\x98\xd1\x5f\x79\xb9\xbb\x37\x21\x46\x8f\x83\x7b\xa6\xe9\xf5\x30\xf7\xdc\xee\x57\x11\xfb\xd0\xe1\x8d\xf1\xd6\x78\xfc\xdf\x00\x00\x00\xff\xff\xdb\xcb\x0a\x72\x15\x1e\x00\x00")

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

	info := bindataFileInfo{name: "assets/js/manager.min.js", size: 7701, mode: os.FileMode(420), modTime: time.Unix(1543014288, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesRuLc_messagesConfigMo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8e\xbd\x6e\x13\x4f\x14\xc5\x4f\xfc\xcf\xbf\xb1\xa0\xa1\xa6\xb8\x14\x44\x20\x34\x61\xd7\x26\x52\xd8\x78\x13\x44\x3e\x24\x44\xac\x84\xc8\x50\xd1\x8c\xec\xb1\xbd\x8a\x3d\x63\x66\x67\xf9\x90\x22\x91\x8f\x22\x5d\x68\xa0\xa1\xe1\x15\x92\x10\x4b\x16\x81\x8d\xc4\x13\xdc\x79\x01\x9e\x05\xad\xf3\x81\x28\x98\xea\x9c\x73\xef\xef\xdc\xf9\x75\x63\xf2\x13\x00\x4c\x02\xb8\x09\xa0\x06\xe0\x7f\x00\x2f\x71\xfe\x06\x17\xf9\x2b\x00\x55\x00\x07\x00\xae\x03\xf8\x09\xa0\x3b\x01\x30\x80\x65\x00\x0f\x4b\xc0\x33\x00\xef\x4b\xc0\x35\x00\xbe\x04\x4c\x00\x28\x5d\xf4\x14\xfd\xff\x15\xa2\x69\x74\x3b\xe9\x4c\x2e\xa9\xb6\xcc\x7a\x8e\xde\x18\xbb\xa9\x6c\x4a\x4d\x93\x69\x77\x35\xcc\xac\x74\x89\xd1\xd4\x36\x96\x5c\xd2\xdc\x54\x96\x12\x4d\xad\x24\x1d\x48\xd7\xec\x2a\x4b\xa6\x7d\x49\x5e\x32\x6b\xae\x5b\xb8\x0d\x35\x30\xd6\x89\x7a\xda\x49\x5a\xe2\x71\xd6\x49\x45\xc3\x44\xd4\x52\xaf\x1f\x6d\x26\x5d\xd9\x37\xd3\x36\x2b\xaf\xaf\x35\xc4\xa2\x55\xe3\x13\x62\x49\x3a\x15\x51\x25\x08\x67\x45\x50\x15\x95\x2a\x55\xaa\xd1\xcc\xcc\xbd\xa0\x1a\x04\xe5\x55\x99\x3a\xd1\xb0\x52\xa7\x3d\xe9\x8c\x8d\xe8\xe9\xb8\x83\xea\x99\x95\x7d\xd3\x32\x54\xfb\xab\x78\xbe\xbc\x2a\x75\x27\x93\x1d\x25\x1a\x4a\xf6\x23\xba\xf2\x11\x6d\x64\x69\x9a\x48\x5d\xae\x3f\xa9\x2f\x8b\x17\xca\xa6\x89\xd1\x11\x85\xd3\x41\x79\xd1\x68\xa7\xb4\x13\x8d\x77\x03\x15\x91\x53\x6f\xdd\xfd\x41\x4f\x26\x7a\x8e\x9a\x5d\x69\x53\xe5\xe2\xe7\x8d\x15\x31\xfb\x67\xaf\xf8\x4f\x5b\x59\xb1\xac\x9b\xa6\x95\xe8\x4e\x44\xe5\xf5\x5e\x66\x65\x4f\xac\x18\xdb\x4f\x23\xd2\x83\xb1\x4d\xe3\xea\x1c\x9d\xcb\x58\xdf\x0e\x83\x38\x0e\x69\x6a\x8a\x0a\x19\xdc\x8a\xc3\x90\x16\x28\xa0\x68\xec\xe7\xe3\xca\xe5\xa8\x16\x3f\x28\xe4\x9d\xf1\x5a\x2d\x0c\x68\x6b\xeb\x1c\x99\x8f\x2b\xc1\x5d\x5a\xa0\x90\x22\xaa\xcc\x81\x3f\x73\xce\xa7\x3c\xf2\xfb\x3c\xf4\x3b\x7e\x97\x8f\x39\x27\xce\xf9\xc8\x6f\xf3\x21\x1f\x71\xee\x77\xfd\x3e\x8f\xf8\x1b\xe7\x7c\x4c\x7c\xc6\xb9\xf0\x7b\xfc\xbd\x80\xfc\x3e\x1f\xf2\x0f\x1e\xf9\x03\xf0\x17\x1e\xfa\x6d\x1e\x71\xce\x27\xc4\x27\x7c\xea\x3f\x90\xdf\x1d\x63\x45\x7e\x48\x05\x7b\xc2\x23\xbf\xc3\x67\x3c\x1c\x57\x16\xf9\xf0\x9f\x97\xc0\x1f\xfd\xb6\xdf\xe3\xaf\x9c\xf3\x10\xbf\x03\x00\x00\xff\xff\xcb\x67\x1c\x8f\xde\x02\x00\x00")

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

	info := bindataFileInfo{name: "locales/ru/LC_MESSAGES/config.mo", size: 734, mode: os.FileMode(420), modTime: time.Unix(1543014295, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesRuLc_messagesManagerMo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x54\x5f\x6c\x5b\x57\x19\xff\x75\xd8\xc6\x71\x12\x18\x2b\x8c\xf2\xff\xec\x4f\xbb\x8d\xe1\xcc\x4e\x98\x54\xb9\x75\x42\xc8\x5a\x5a\xd1\x94\x34\x71\xe8\x2b\x77\xf1\x49\x72\x15\xfb\xde\x70\xef\x75\xd2\x8a\x3d\xc4\x8e\xa6\x16\x0c\x2b\xeb\x98\x06\x48\xdb\x68\x27\x24\x84\x90\x70\xa3\xb8\x75\xd2\xc4\x01\x09\x84\x78\xe2\x3b\xe2\x9d\x07\xc4\x1b\x7b\x41\x3c\x31\x24\x84\xbe\x73\x8e\x13\xa7\x6d\x10\x9b\xa5\xdc\x73\xbe\xf3\xfd\xfb\x7d\xbf\xef\xfb\xf2\x97\x47\x62\xaf\x03\xc0\x20\x80\xcf\x02\xf8\xd5\x21\x60\x0c\xc0\xc9\x0f\x41\xff\xce\xc4\x80\x0f\x03\x38\x1b\x03\xf8\xe9\x42\x0c\x48\x01\x98\x8e\x01\xbd\x00\xbe\x65\xcf\x52\x0c\x48\x02\xb8\x64\xe5\x6a\x0c\x88\x03\xf8\x41\xcc\xc4\x7d\x2d\x06\x1c\x01\xd0\x88\x01\x9f\x06\xf0\x47\x1b\xf7\x6f\x31\x20\x01\xe0\x5d\x6b\xff\x4f\xeb\xff\xaf\x18\xd0\x0f\x20\x1e\x07\x3e\x02\xe0\x70\x1c\x78\x08\xc0\xe3\x71\x93\xff\x68\x1c\xe8\x03\x90\x89\x03\x3d\x00\xf2\x71\x13\xe7\x6b\x71\x20\x06\xe0\x7c\x1c\x38\x0c\x60\x2a\x6e\x70\xcd\x59\xbf\xc5\x38\xf0\x51\x00\xdf\x89\x9b\xfc\x75\x1b\xff\x35\x2b\xdf\xb4\x71\x7e\x69\xe3\xdc\x8a\x1b\x5c\xb7\xad\x7e\x2b\x6e\xf8\xf9\x9d\x8d\x93\x4c\x18\xbc\x47\x12\x06\xef\x93\x56\x1e\x4c\x18\x5c\xc3\x09\x83\xf3\x4c\xc2\xd8\x4f\x5b\x59\x26\x4c\xde\xc0\xda\x55\xad\x7f\xdd\xbe\xff\x24\x01\x3c\x0c\xe0\x17\x36\xde\x46\x02\x98\x3f\x04\xfc\xde\xbe\x0f\x24\xcd\x39\x92\x04\x8e\x02\x98\x48\x02\xcf\x00\x58\x4a\x02\x5f\x00\xf0\xb3\xa4\x89\xf3\x9b\xa4\xe1\xff\x4f\x56\x7e\x37\x09\x9c\x05\xf0\xef\x24\x30\x02\xe0\x54\x0f\x30\x0e\xe0\xe5\x1e\xa3\xff\xb3\x3d\xff\xda\x63\x70\xfe\xa3\x07\xf8\x24\x80\xff\xf4\x00\x9f\x02\xf0\x68\x0a\x10\xcc\x7b\xca\xf4\xe3\x42\x0a\xf8\x1c\x80\x42\xca\xe4\x0d\x52\xc0\x23\x00\xae\xa7\x80\x8f\x03\xf8\x79\xca\xf0\xb9\x99\x02\x8e\x03\xf8\xad\xd5\xbf\x97\x02\x3e\x01\xa0\xbf\xd7\xf8\x3d\xd1\x6b\xf2\x8e\xf5\x02\x4f\x02\x98\xec\x35\xf5\x2d\xf7\x1a\x1c\x57\xed\xf9\xba\x3d\x6f\x5a\xfb\x5f\xf7\x02\xd7\x01\xb4\x7b\x4d\xbe\x0b\x7d\xa6\xcf\xb3\x7d\x66\xde\xbe\xdd\x07\x3c\x0a\xe0\x7b\x7d\xc6\xfe\x86\x95\x1b\x7d\xc0\x13\x00\xfe\x60\xdf\xff\xde\x67\xea\x78\xcf\xca\x47\xfa\x8d\xdd\x53\xf6\x1c\xee\x37\xfa\xe9\x7e\xa3\x0f\xfb\x81\x43\x66\x45\xf0\x1c\x4c\x0f\x98\x9f\xa7\x61\xe6\xad\xfb\xf7\x79\x98\x19\x62\xdc\x0f\xd9\x37\xe6\xe4\x71\x18\x8c\x59\x98\x9a\x07\xba\x7c\xb8\x97\x99\x2e\x99\x31\x30\x5f\xcc\xfd\xc3\xf6\xed\xb0\xdd\x59\xae\xe3\x18\x80\xc7\x00\x7c\xc6\xea\x78\x9f\x79\xae\xbe\x08\xb3\x6f\xcc\x0d\xe3\x67\xec\x5f\xb2\x36\x3c\x57\x76\xcd\xf5\xec\x71\x9f\x79\xe6\x3f\x06\x20\x6d\xdf\x9f\xb2\x27\x73\xfa\x2c\xcc\x0e\x74\x7e\xdc\x47\x9e\x69\xee\x1f\x46\x67\x22\xd7\xf7\x42\x8c\x16\x8b\xfc\x27\x96\xfd\x60\x41\x06\x5d\xd7\x10\xa3\xa5\x92\xbf\x2c\xc2\xc8\x09\x22\x8c\x46\x91\x2c\x2f\x46\x21\x46\x2b\x91\x1f\xc8\xd9\x40\x86\xf3\x18\x2b\xf9\xa1\xc4\x98\xef\xcd\xba\x41\x59\x04\xb2\xec\x2f\x49\xe1\x94\x4a\xa2\xe4\x86\x91\xf4\x38\xc6\x03\x94\x91\x13\x2e\x3c\x50\xd1\xc9\x3b\x16\x48\x27\x92\x45\x9c\x5a\x92\x5e\x14\xe2\xb4\x1b\x48\xfd\x0d\x23\x31\xeb\x06\xb2\x68\xef\x1a\x98\x2c\xe2\x8c\x5b\xec\x0e\x7c\xf6\x05\x9c\x73\x76\x4d\xf5\xb5\x63\x79\x6e\x17\xd6\x39\x7f\x66\x41\x16\x71\xde\x29\x4b\x9c\xaf\x94\x5f\x94\x81\xf0\xe4\x72\x07\x41\x0e\x13\x81\xeb\x07\x6e\x74\x19\x93\xbb\xf0\x30\x29\x17\xa5\x13\x09\xd7\x8b\x64\xb0\xe4\x74\xe4\x10\x53\xf3\xfe\x72\x57\xfe\x29\x9b\x6c\x2a\x72\xa2\x4a\x88\x82\x13\x2e\xe8\x4f\x88\x82\x5b\x96\x7e\x25\xc2\x45\xb7\x54\x12\x2f\xca\x07\x12\x26\xe4\xa5\x19\xb9\x18\x89\xa2\x9c\x75\x2a\xa5\x68\x57\x31\x20\x46\x03\x29\x2e\xfb\x15\x11\x56\x02\x39\x82\x30\x72\xa2\x30\xb6\x57\x90\x91\x4d\x1e\x73\xbf\x68\xd9\x64\x54\xb1\x31\xc7\x9b\x91\x25\x73\x3f\xed\xb8\xf6\x36\x11\xf8\x33\x32\xb4\x26\xa6\x9c\x8b\x8e\x1b\x19\x79\xaa\x32\xb3\xa7\x9c\xf6\x8a\x72\xd6\xf5\x64\xd1\x88\xda\xca\x90\xd5\x09\x6d\xa5\x4e\x48\x2b\xee\xf9\xd9\x07\xed\x39\x29\x17\xfd\x20\x4a\x8f\x87\x73\x6e\x31\xfd\xd5\xca\x5c\x98\x2e\xf8\x39\x51\x94\x4b\x5f\x59\x70\xe7\x9d\xb2\x3f\x10\x54\x52\x13\xdf\x28\xa4\xf5\x20\xb8\xbe\x97\x7e\xc1\x89\x64\x4e\x0c\x66\xb2\xc7\xd3\x99\xa1\xf4\xe0\x90\x18\x1c\xca\x3d\xff\xfc\xb3\x99\xa1\x4c\x26\xc5\x1d\x4e\x17\x02\xc7\x0b\x4b\x4e\xe4\x07\x39\xf1\x75\x1d\x43\x8c\x57\x02\xa7\xec\x17\x7d\x71\x72\x5f\xe0\xe1\xd4\x39\xc7\x9b\xab\x38\x73\x32\x5d\x90\x4e\x39\x27\x76\xe5\x9c\x98\xac\x84\xa1\xeb\x78\xa9\xf1\xb3\xe3\xa7\xd2\xdf\x94\x41\xe8\xfa\x5e\x4e\x64\x07\x32\xa9\x31\xdf\x8b\xa4\x17\xa5\x0b\x97\x17\x65\x4e\x44\xf2\x52\xf4\xdc\x62\xc9\x71\xbd\x13\x62\x66\xde\x09\x42\x19\xe5\xa7\x0b\xa7\xd3\xc7\xf7\xec\x18\xcf\xac\x0c\xd2\xa7\xbc\x19\xbf\xe8\x7a\x73\x39\x91\x9a\x28\x55\x02\xa7\x94\x3e\xed\x07\xe5\x30\x27\xbc\x45\x2d\x86\xf9\xa1\x13\xc2\x5c\xf3\xde\xd1\x6c\x26\x9f\xcf\x8a\x63\xc7\x04\x5f\x33\x8f\xe5\xb3\x59\x31\x22\x32\x22\xa7\xe5\xe1\xfc\x60\x47\x75\x32\xff\x65\xbe\x3e\xad\xcd\x4e\x66\x33\xe2\xa5\x97\x8c\xcb\x70\x7e\x30\xf3\x8c\x18\x11\x59\x91\x13\x83\x27\x40\x3f\xa2\x26\x6d\xa8\xaa\xaa\xd1\x1a\xb5\xd4\x35\x7e\x68\xd3\x2d\x6a\x68\xb1\xa6\xbe\x7f\xdf\x83\x60\x51\xad\x50\x83\x6e\x51\x5b\xd5\xd4\x15\x6a\xd1\xe6\xff\x67\x45\x6d\x5a\x03\xbd\x4a\x77\xa9\x45\xb7\xa9\x41\x1b\xea\x2a\xb5\x68\x43\xd0\x1d\x6a\xd0\x8e\x5a\x55\x55\x8e\xf4\x36\xb5\x69\x47\xd5\x55\x8d\x36\xa9\x05\xfa\x21\xad\xa9\x1a\xb5\x75\xf8\x6d\x8e\x40\x77\xa9\x49\xdb\xd4\xa2\x26\xe8\x0d\x6a\xd0\xa6\x5a\x61\x6b\x86\xca\xae\xeb\xba\x94\xa6\x5a\xa1\xdb\xb4\xde\xb1\x14\x6a\x95\xd6\xa9\xd1\x71\x55\xd7\x04\xad\xa9\x2a\x35\xd5\xcb\x42\x55\xe9\xae\x5a\x55\x57\xa9\xa1\x6a\xd4\xd4\x16\x1b\x1f\x2c\x92\x2e\x63\x9d\x1a\xea\xca\x07\xf4\x3f\x90\xb2\x1b\xd4\xa6\x3b\xda\x6d\x9b\xda\x46\xbc\xa5\x19\xd2\x1d\xbb\xae\xea\x74\x87\x0d\x55\x9d\x13\x73\xc2\x35\x55\x67\x5e\xd7\x3a\x9a\x7b\xde\xf7\xf1\x7d\x63\x8f\x41\x8b\xa5\xab\x12\x6a\xf1\xbf\x4b\x2e\x87\x79\xa2\x26\xad\xeb\x32\xee\x8d\x7d\x9f\xf6\x9e\x0c\xfb\x19\x6e\x99\x21\x68\x73\x83\xd5\x8a\xee\xe9\x26\x35\x40\x3f\xa6\x2d\xae\xe7\xa7\xd4\x66\x23\x75\x85\x9a\x76\x30\xdb\xc2\xb4\x5e\xd5\xff\x27\x4d\x6f\xab\x15\x6a\x51\x9b\xbf\x9c\x49\xd5\x40\xef\x58\xb2\x5b\x5d\xe5\x71\xa2\x6d\x6d\xb0\x42\x6b\xac\x15\xb4\xa3\xa3\xd7\xb4\x6f\xc3\x74\xcf\x8a\x86\x53\x86\xda\xe0\xa2\x0e\x64\x49\x8f\x22\x57\xfc\x5d\xdd\xd8\x06\xe8\x86\xaa\x69\xfb\x55\x55\x35\x5a\x63\xda\xe8\x16\x5a\xa0\x9b\xbc\x08\xb4\x45\x0d\xb5\xca\x80\x5f\xe5\xf9\xe0\xeb\xbe\x41\x51\xf5\x4e\xd2\xfb\xc6\xb5\xa5\x71\x08\x6a\x31\xd7\x74\x57\xbd\xc2\xbc\x99\x59\xa3\xad\x07\x4e\xb7\x2e\x37\xad\x56\x69\x8b\x89\x66\x48\x7a\x12\x5f\x19\x10\x3c\x4a\x9c\xd7\x4c\xad\xce\x3b\x72\x7f\xfb\xf4\x82\xbc\xb1\x37\xeb\x6f\x1d\xd4\x8f\xdd\xed\xd4\x34\x37\x55\x4d\x55\xb9\xbf\x6f\x52\xd3\xd4\x66\xd8\xe0\xf1\xdd\xd1\x2d\xdf\x56\xd7\xf6\xac\xde\xa2\xdb\xd4\xd2\x0c\x34\x55\x6d\x7f\x87\xec\xe6\x80\xde\x51\x55\xda\xa1\xa6\xba\x6a\xf6\xe2\x4d\x6e\x0a\xff\xf3\x58\xd1\xa3\xd8\xb4\x4b\xb6\xcd\x43\x7f\x4f\xc0\x03\xc1\x1d\x80\xe6\xfd\x84\xfe\x6f\x00\x00\x00\xff\xff\x76\x75\x5a\xb4\x82\x0d\x00\x00")

func localesRuLc_messagesManagerMoBytes() ([]byte, error) {
	return bindataRead(
		_localesRuLc_messagesManagerMo,
		"locales/ru/LC_MESSAGES/manager.mo",
	)
}

func localesRuLc_messagesManagerMo() (*asset, error) {
	bytes, err := localesRuLc_messagesManagerMoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/ru/LC_MESSAGES/manager.mo", size: 3458, mode: os.FileMode(420), modTime: time.Unix(1543014295, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesRuLc_messagesWorkersMo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\x4f\x6b\x14\x31\x18\xc6\x9f\x2d\x8a\x30\x47\x0f\x9e\x3c\xbc\x1e\x2c\x8a\xa6\x26\x33\x16\x4a\x76\xb3\x15\x6b\x0b\x62\x17\xd7\x32\x2a\x1e\x43\x37\xce\x0e\xdd\x49\x86\x24\x23\x0a\x3d\x88\x17\x3f\x81\x57\x3f\x83\xbd\x2d\x1e\x3c\xf8\x09\xb2\x5f\xc0\xcf\x22\x3b\xeb\xdf\xe7\xf4\xfc\xc8\x2f\x0f\xef\x8f\xab\x97\x3e\x01\xc0\x16\x80\xeb\x00\xee\x02\xb8\x0c\x60\x84\x4d\xa6\x00\xae\x00\x78\x06\x60\x3e\x00\x5e\x01\xb8\x06\xe0\xfb\x00\x18\xfc\x72\xb6\xf0\x4f\x5e\x3a\x7f\x66\x7c\xc0\x89\x69\x9d\x8f\x6c\x12\xaa\x7a\xc6\x1e\x76\x55\x60\xa5\x93\x34\x33\x6f\x1e\x9c\xd5\x73\xdd\xb8\x1d\xdf\x65\xd3\xa7\x25\x3b\xf0\x46\xc7\xda\x59\xf6\x48\x47\x23\x29\xe7\x62\x8f\xf1\x82\xe5\x05\xe5\x85\xdc\xdd\xbd\xc3\x0b\xce\xb3\x63\x1d\x22\x2b\xbd\xb6\x61\xa1\xa3\xf3\x92\x9e\xf4\x1b\x34\xe9\xbc\x6e\xdc\xcc\xd1\xe8\xbf\xe1\x71\x76\xac\x6d\xd5\xe9\xca\xb0\xd2\xe8\x46\xd2\x1f\x96\x74\xd2\x85\x50\x6b\x9b\x4d\x1e\x4f\x0e\xd9\x0b\xe3\x43\xed\xac\x24\xb1\xc3\xb3\x03\x67\xa3\xb1\x91\x95\xef\x5a\x23\x29\x9a\xb7\xf1\x5e\xbb\xd0\xb5\x1d\xd2\xe9\x5c\xfb\x60\xa2\x7a\x5e\x1e\xb1\xbd\xbf\xde\xfa\x9e\xd7\xc6\xb3\x43\x7b\xea\x66\xb5\xad\x24\x65\xd3\x45\xe7\xf5\x82\x1d\x39\xdf\x04\x49\xb6\xed\x31\xa8\x62\x48\x9b\xaa\xec\x4d\xc1\x95\x12\xb4\xbd\x4d\xeb\xca\x6f\x28\x21\x68\x9f\x38\xc9\x9e\xc7\x2a\xff\xfd\x34\x52\xf7\xd7\xf5\x56\xaf\x8d\x04\xa7\xf3\xf3\xcd\x97\xb1\xca\xf9\x6d\xda\x27\x41\x92\xf2\x21\xd2\xe7\x74\xb1\x7a\x9f\xbe\xa4\x8b\xf4\x6d\xf5\x61\xf5\x31\x2d\xd3\xd7\xb4\xc4\xcf\x00\x00\x00\xff\xff\x4e\x6a\x2e\xd8\xd9\x01\x00\x00")

func localesRuLc_messagesWorkersMoBytes() ([]byte, error) {
	return bindataRead(
		_localesRuLc_messagesWorkersMo,
		"locales/ru/LC_MESSAGES/workers.mo",
	)
}

func localesRuLc_messagesWorkersMo() (*asset, error) {
	bytes, err := localesRuLc_messagesWorkersMoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/ru/LC_MESSAGES/workers.mo", size: 473, mode: os.FileMode(420), modTime: time.Unix(1543014295, 0)}
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
	"locales/ru/LC_MESSAGES/manager.mo": localesRuLc_messagesManagerMo,
	"locales/ru/LC_MESSAGES/workers.mo": localesRuLc_messagesWorkersMo,
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
				"manager.mo": &bintree{localesRuLc_messagesManagerMo, map[string]*bintree{}},
				"workers.mo": &bintree{localesRuLc_messagesWorkersMo, map[string]*bintree{}},
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
