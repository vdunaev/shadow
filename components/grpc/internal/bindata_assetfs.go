// Code generated by go-bindata.
// sources:
// templates/views/manager.html
// assets/js/manager.min.js
// locales/ru/LC_MESSAGES/config.mo
// locales/ru/LC_MESSAGES/grpc.mo
// locales/ru/LC_MESSAGES/manager.mo
// DO NOT EDIT!

package internal

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/elazarl/go-bindata-assetfs"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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

var _templatesViewsManagerHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x1b\x5d\x93\xdb\xb8\xed\x3d\xbf\x82\xc3\xdb\xeb\x24\x6d\x64\x6f\x76\x93\x4e\xeb\xd8\x9b\xe9\x5c\xdb\x69\x3a\x49\xae\x93\xbb\x3e\xa5\x69\x86\x96\x60\x8b\x5b\x9a\xd4\x91\x94\xbd\x5b\x8f\xff\x7b\x87\xd4\x17\xe5\x95\x64\x5b\xa2\x93\x6d\xa7\x7a\x48\xbc\x02\x09\x80\x00\x08\x02\x20\xb4\xdd\xa2\x08\x16\x94\x03\xc2\xa1\xe0\x1a\xb8\xc6\x68\xb7\x7b\x32\x8d\xe8\x1a\x85\x8c\x28\x35\xc3\x09\x59\x42\xa0\xa9\x66\x80\x6f\x9e\x20\x84\x90\x0b\xb4\xef\xbf\x30\x58\xe8\x1c\x68\x07\xc4\xd7\x37\xcb\x8f\x7f\xfb\x61\x3a\x8e\xaf\xf3\x29\xe3\x88\xae\x6f\x9e\xe4\xff\xd5\xd0\x87\x0c\x88\x5c\xd0\x3b\x7c\xd3\x04\x95\x62\xd3\x40\x35\x14\x2c\x58\x45\xc1\x8b\xab\x1c\xb6\xdd\x22\x49\xf8\x12\xd0\x85\x7a\x8e\x2e\x14\xc8\x35\x0d\x01\x4d\x66\x68\x94\xff\x56\x66\x51\x05\x7b\xdb\x6d\x39\xe6\x6d\x64\x46\x49\x48\x18\x09\xa1\x7c\x3b\xfa\x40\x56\x80\xf0\x08\x23\xfc\xc5\x8a\xa3\x5a\x99\xc3\xc4\xdd\x97\x84\x70\x60\xce\xba\x1f\x8e\x70\xc5\x56\x1b\x15\x5f\xdd\x38\x6c\x64\x04\x77\xbb\xe9\x38\xbe\x6a\x18\xdc\x25\xad\xda\xc0\x86\x57\x35\x76\x0a\x15\xdf\x3c\xe9\x24\x42\xc2\x50\xc8\x88\x0a\x8e\x11\x8d\x66\x98\x04\x75\x89\xed\x76\x18\x49\xc1\x60\x86\x35\x99\x33\xaa\x34\x46\x44\x52\x12\xac\x52\xa6\xa9\x02\x06\xa1\x79\x6f\xc0\x32\x6d\x5a\x7c\x5d\x63\xab\xe7\xe8\x62\x05\x3a\x16\x56\x15\xa5\x44\xde\xdb\x57\x35\xbd\x75\xb1\xdc\xa4\x8a\x3d\x7a\x39\x95\x3d\x95\x67\x2f\xdb\x35\xde\x80\x88\x2e\xca\x69\x6f\x79\x92\xea\x9f\xef\x13\xe8\xe2\xd3\xf2\x4a\x6a\x9c\x06\x31\x90\x88\xf2\xa5\x23\xc8\x4c\xd6\xf1\x03\x59\x07\x35\xd6\x8d\xec\x23\xa2\x49\xa0\xc5\x72\x69\x66\x86\x82\x31\x92\x28\xc8\x5f\x27\x44\x02\xd7\x33\xfc\x5d\x93\xd2\x62\x09\x8b\x19\xfe\x2e\x3c\x82\x86\x55\x28\xdc\x25\x84\x47\x10\xcd\xf0\x82\x30\x43\xc2\xbe\x35\x46\x24\x05\x33\xc6\x78\x18\x51\xbb\x46\x4a\xc9\xc4\x2f\xeb\xa2\xc9\xf7\x4c\x85\xc9\xd9\x1d\x2f\xbb\xf1\x4d\xc7\xe4\xc0\x00\x63\x32\x46\xce\xc7\xc8\xa0\xc6\x55\x21\x67\x54\x09\xbc\x54\x5d\x66\x7c\x99\x74\x18\x99\x03\x63\x10\xcd\xef\x8f\x52\xe6\x11\xf2\xd9\xb7\xf2\x60\x2e\xa2\xfb\x23\x26\xee\x4f\x6e\xf6\x57\xc7\xcf\x6e\xf3\x65\x9d\x18\x0e\x69\xac\x71\x12\x2d\x88\x2e\x08\x5a\x90\x80\x48\x29\x36\x81\xa4\xcb\x58\x9b\x3f\x17\x1b\xe3\xfc\xe8\xe9\x78\x1d\x8b\x2a\xf7\x6d\x61\x5b\x7d\x90\xed\xfb\x81\x9f\xb4\x04\xb2\xea\x83\xcc\x3c\x53\x95\x10\x5e\x2c\xdc\x5a\x11\xb2\xff\x06\x1b\x22\xb9\x71\x15\x66\x47\xd0\x17\xbf\xe3\x08\x2b\x4b\x08\xa3\x0b\xbb\x29\xcc\xbc\x5e\xc2\x00\x1e\x9d\xca\x6c\xb6\x05\x4f\x9b\x72\xec\xe9\xd5\x4d\xf7\xb4\xe1\x2d\xa7\xde\x49\x7c\x2f\x84\x5c\xa1\x4c\xbf\x33\x9c\x08\x7b\xcc\x85\x9a\x0a\x3e\xc3\x6f\xf2\x1f\x21\x61\xac\xf4\x13\xe6\x0f\x64\x26\x05\xb1\x90\xf4\xdf\x82\x6b\xc2\x72\xaf\x2c\x41\xa5\x4c\xcf\xb0\x3c\xc2\xeb\x64\x5e\xc5\xe0\xd9\xf3\xf4\x6b\xc2\x68\x44\xb4\x90\x27\x2e\xc4\x2e\x86\x1a\x03\x45\xfa\x3e\x81\x19\x8e\x69\x14\x01\xc7\x88\x93\x15\xcc\xf0\x97\x9c\x1d\x8c\xd6\x84\xa5\x30\xc3\x0d\x71\x89\x67\x8a\xd9\x7a\x6b\x04\xeb\x9e\xbe\x29\x40\x39\xf4\xb4\x1c\xcc\xa3\x3f\x53\x60\xdd\x71\x44\xeb\x12\xdc\x38\xd7\x44\x34\x46\x8f\x89\xe0\x8a\xae\x4f\xf5\x83\x25\x4a\x8b\xa7\x86\x14\x65\xa8\x63\xb1\x06\x99\xff\x56\x5a\xd2\x04\xa2\x9e\x34\x32\x3a\x26\xc0\x18\x32\x5f\xf6\x9f\x9c\x33\x80\x36\x34\xd2\xf1\x0c\xbf\xb8\xfc\xde\xf1\x5d\x56\x1d\x85\xeb\xd2\xb1\x37\x2a\xaf\x6a\x54\x8c\xe6\xcf\x40\xa4\xbe\x94\x77\xc6\x3d\x9f\x81\xca\x75\x8d\xca\x1f\xac\xab\xf1\x41\x66\x3a\x1e\xa2\x55\x43\x7b\xa0\x4d\x99\xc0\xa5\xff\xfc\x2a\x65\x58\x3c\x47\x17\x0b\x63\x47\x36\x63\xf0\xb8\xe3\x1d\x52\x17\x6b\x22\xf7\x52\x05\x4b\xf2\x41\xa6\xd0\x97\x86\x87\x2d\x16\xd9\x20\xd9\x61\xcb\x5a\xc8\x00\x15\x15\x68\x07\x21\x40\x95\x33\xce\x58\x7b\xab\xde\x83\x52\x64\xd9\x2b\xce\x6a\x64\x91\xe4\xa9\xcc\x2d\x59\x13\x15\x4a\x9a\xe8\xc9\x5a\xd0\xe8\xe9\xe5\xb3\xee\xec\x48\x13\xb9\x04\x93\x1d\x1d\x3e\x88\x83\xca\x06\x76\xbb\xe0\x56\x09\x8e\x1d\x69\xe7\x29\xdf\xc1\x7c\xe3\xe8\x15\x25\x12\x6c\x5e\xd2\x8b\x31\xa7\x2c\x92\xe5\x28\x22\xd5\x96\x5b\x49\x36\x05\xc7\x7f\xfd\xe9\xc7\x0f\x96\xe3\x44\x82\x17\xfd\x02\x53\xde\x14\xfa\x40\xb0\x5e\x18\x3c\x3d\xbc\xdd\x7f\xfc\xec\xa6\xd6\xc8\x5e\xa5\x61\x08\x4a\xb9\x86\x65\xcf\x94\x2a\xac\xf7\xc2\x40\x55\xd6\xb1\x27\xc9\x70\xed\xbb\x81\x91\x0d\xf3\x82\xa5\x14\x69\x52\x06\xbb\xbf\xa4\x54\x42\x64\x03\x3b\xd7\x0b\x7c\xcc\x01\x68\xb7\xd3\x32\x85\xca\x86\x6c\x75\xa1\xd4\x58\x89\x26\x01\xa2\x1b\xd1\x64\x80\x43\x68\x3c\xed\xcc\x5a\x5d\xcc\x08\x30\xb0\x4b\xce\x98\xe2\x42\x37\xae\x0f\xc5\x34\xf2\xce\x4a\x46\x52\x48\xf4\x14\x7e\xa9\xed\x17\x1c\x89\x74\xce\x00\x3f\x7b\x08\xf9\xf4\xb9\x1d\xb6\x60\x82\xe8\xe6\x49\x05\xc8\x97\xc3\x76\x93\x01\x9e\xae\xe6\x20\x8b\x64\xe0\xc1\x09\x56\x7a\x33\x9b\x4a\xe5\xc5\x26\x8c\x12\xa2\x35\x48\x3e\xc3\xff\xfc\x14\xfc\xe6\xf3\x9b\x4f\x97\xc1\xef\x3f\xff\xfa\xe9\x3f\x46\xd9\x8f\x67\x6f\x2e\x6a\xe9\x44\x86\xf1\x8f\xb0\x20\x29\xd3\x16\xe9\xd8\x9b\x0e\xac\xb5\xb5\x28\x82\x72\x7d\x7d\xd5\x2c\xd2\x56\x10\xe5\xfa\xb7\x2f\x5b\xe7\x34\x83\x54\x07\xa1\x76\x98\xea\x20\xd5\x0e\x53\x0b\x7a\x07\x51\x2b\xb5\x0e\x68\x06\x6b\xa5\x58\x41\x1f\xb9\x99\x3d\x0e\xdb\x4a\x3b\x74\xde\x0e\x4b\x3b\x74\xde\x0e\xeb\x54\x79\x07\xb0\x53\xe1\x8f\x5c\xdf\x8f\x4c\xdd\x73\x21\x58\xb3\x1c\x73\xc8\x39\x84\x18\xc6\x10\xfe\x6b\x2e\xee\x0e\x8b\xf1\x56\x05\x6a\x43\x75\x18\xa3\xba\x40\x6b\x87\x75\x25\x35\x8b\x18\xa2\x2a\x2e\xfb\x3a\x42\x54\x5a\x52\xbe\x6c\xf1\x3f\xad\xb0\xf9\xbd\x06\xd5\x22\xfb\x1c\x74\x0e\xe1\x6b\xb8\xd3\xa7\xda\xef\xd7\x34\x54\x4f\x6b\x46\x0f\xf3\xc3\x3f\xf1\xb4\x77\xdd\xbc\xe9\x99\x66\xb7\x90\x27\xc9\x12\x65\x73\xae\x3c\x05\x6b\xc5\x53\x55\x2d\xe0\x39\xba\x00\xb3\xd0\xc9\xac\xe0\xc7\xf7\xba\xcd\x33\x15\x89\x89\x52\x5d\xcb\x30\x54\x47\x1f\xac\x7f\x34\xab\xcf\x85\x6f\xdf\xbe\x55\x95\xc1\xe4\x12\x30\x21\x77\xf1\x0b\x97\xfb\xf5\xa6\x42\x54\x54\x1a\x32\x42\xde\xc5\xe5\x21\x6f\x73\x9f\xe9\x38\x5b\x8c\x3f\x3e\xcf\xb0\x1f\xa6\x66\xef\x13\x09\xe4\xc4\xed\x2f\xc5\x46\xcd\xf0\xb5\x9b\x41\x96\x89\x7e\x81\xd2\xef\xc2\xfd\x29\xc7\x2f\xb6\xa2\x79\xc4\x07\xae\x2a\xd1\x6a\xcc\xef\x9e\x35\xe6\xa2\xde\xb3\x4d\x27\xb3\x0e\xe6\xda\x47\xde\x5e\x52\x99\xa7\x5a\x0b\x9e\x1f\x3b\xd9\x1f\xa5\x81\xcd\x35\x47\x73\xcd\x83\xc8\x78\x2d\x99\x5f\x79\x05\x12\x56\x62\x0d\xdf\x20\xe9\x2d\x79\x76\x8b\x28\xd9\xbd\xb0\x96\x44\xc5\x01\x61\x1a\xdf\xf4\xbd\x05\x6d\xa5\x36\xce\xa4\xf2\x95\x65\x4e\xf9\x42\x14\x12\x27\xd1\x90\x0b\xa0\x46\x16\x1e\x8a\x30\x61\xa9\x7a\xec\xd2\x3b\xf1\xee\xb7\xeb\xf1\xe7\x72\x3c\x30\x35\xac\xb6\x37\xec\x32\x67\xb8\x20\xa6\xe3\x01\xd7\x39\xd3\xb1\xbd\xe8\xec\x71\xab\xdc\x4f\xec\xd5\x72\xfb\x5d\x02\x1f\x57\x99\xaf\x75\x02\xe4\x77\xfe\x18\x29\x7d\xcf\x60\x86\x23\xaa\x12\x46\xee\x27\x5c\xf0\xde\x97\xc7\xb1\xec\x1b\xd2\x77\x7a\x9f\x90\x09\x05\xf8\xe6\x57\x9a\xae\x40\xbd\x1e\xb6\x79\x6b\xbd\x9b\xd9\x6d\x39\x20\xc2\x40\xea\x93\x7b\x3e\x4a\x94\x3d\x4f\x76\x97\x15\xc6\xbf\x28\xc1\x68\x54\x35\x99\x0e\xc1\x66\x63\xb0\xac\xe4\x3d\x5c\x4a\x26\x4c\x0b\x42\xe0\x1a\xfa\x34\x76\x94\x18\xdd\x44\x52\xa5\xf3\x15\xd5\x0f\x4e\x97\xe2\xba\xc1\x49\x0d\xb2\x2b\xe6\x1f\x6c\x1b\xcb\xc5\x90\x9c\x71\x98\x6e\x4f\x9b\x62\xc4\x7f\x96\xd6\xa1\x53\x6c\xe3\x7f\xa8\xb3\x8e\xc1\xc2\x5b\x63\xdd\x8f\xa9\xf6\xdb\x59\x57\x21\xfc\x7f\x63\xdd\xde\x94\xff\xd6\xc6\x3a\xbf\x6d\x55\x5f\xa3\xa5\x6a\x48\x3b\xd5\xa0\x3e\x0f\xb7\xf7\xe8\xd5\x59\xda\xa8\xdc\xbe\xa3\x57\xe7\x68\xa1\x6a\xed\xd1\xf2\xd2\x3e\xd5\x3f\x0c\x1e\xd4\xcf\x34\xa4\x97\xe9\x40\x1f\x93\xe3\x41\x87\x35\x32\x9d\xbb\x89\x69\xa0\x61\x7b\x6f\x5e\x1a\xdc\xb8\x74\xbe\xa6\xa5\xe1\x0d\x4b\x9b\x47\xd4\xb0\x54\x36\x2b\xf5\x62\xea\xeb\x36\x2b\x79\xac\xd5\xfa\x6d\x52\xf2\x92\x81\x7f\xcb\xdd\x32\xa0\xab\x69\x58\x95\x76\x7f\x9b\x92\xc4\xcb\x16\x3d\xb8\x9c\x15\x49\x7c\x14\xca\xbe\xa5\xe2\xfb\x9f\x96\xc3\xb8\xee\x5d\x2c\xea\x55\x28\x3a\x5f\x08\x7c\xe4\xd0\x23\x86\x1d\x18\x72\xa4\xdb\x3a\xe6\xab\xc2\x96\xe3\xa5\xf5\x3c\x7a\x44\x5f\xec\x1d\x36\xba\x0e\x31\xb6\x4f\x3e\xfc\xe5\xac\x5b\x08\xa8\xe3\xa9\x7f\xc8\xec\xd4\x15\x9d\x0f\xa9\x8d\x16\xb2\xaf\xa8\x6d\xed\xcf\xb9\xdd\x1f\x87\xc6\x91\x58\x44\xa3\xb2\x3a\xb6\x2d\xc9\x6e\x62\xaa\x21\x50\x09\x09\x61\x82\x12\x09\xc1\x46\x92\xe4\x75\x09\x36\xe9\xcc\x82\x89\x4d\x70\x3f\x41\x24\xd5\xa2\x82\xac\xc8\x5d\x10\x03\x5d\xc6\x7a\x82\xae\x2f\x61\x55\x41\x6c\xe4\x3d\x41\x2f\x2e\x2f\xbf\xcf\x5e\xe6\x25\xd0\xd1\xde\xed\x12\xca\x0b\x84\x5b\x07\xa7\x5c\x52\x9e\x7d\x76\x37\x41\x97\xaf\xf7\x01\x73\xa1\xb5\x58\x59\x48\x0d\xb1\x53\xfc\x74\xb0\x25\x42\x51\x4d\x05\x9f\x20\x09\x8c\x68\xba\x86\x3a\x3b\xb6\x06\xd9\x38\x9e\xcc\x95\x60\xa9\x86\x8a\x81\x9c\xa5\x57\xc9\x5d\xf5\x4e\x8b\x64\x82\xae\xca\x57\x05\xda\xec\x2e\xc5\xc1\x6b\xab\x6d\x84\xd1\xa5\xe1\xc4\xe0\x69\x1c\x3f\x72\x3b\x3d\x5b\x66\x33\x58\xb4\x4c\x76\x44\xeb\x2a\xb7\x4d\x13\x0d\xd3\xac\x46\xaa\xa9\x6b\x90\x9a\x86\x84\x15\xa4\xb5\x48\x0a\x14\xd3\xb1\xb5\xb1\x36\x53\xbc\x55\xd6\x10\xb7\x5b\xa4\x34\xd1\x34\xfc\xcb\xcf\xef\xdf\xa1\xa7\xd9\xef\xbf\x7f\x7c\x87\xf0\x38\x22\x2a\x9e\x0b\x22\xa3\x31\x51\x0a\xb4\x1a\xaf\x81\x47\x42\xaa\x71\xf9\xb9\xd9\xf8\xd6\xf9\x63\xb4\xa2\x7c\x64\xb0\xda\x66\xdb\x67\x07\x90\x2f\x65\x12\x16\x78\x6f\xd5\x78\x45\x38\x59\x82\xb4\xf3\xb5\x4c\xcb\xe9\x19\xe3\xff\x09\x00\x00\xff\xff\x68\xc6\xc7\x5c\x8b\x40\x00\x00"

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

	info := bindataFileInfo{name: "templates/views/manager.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsJsManagerMinJs = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\x4d\x6f\xdb\x38\x10\xfd\x2b\x5a\xae\x61\x90\xb0\x4c\xec\x2e\xf6\x24\x2d\x11\x2c\xd2\x53\x81\xb6\x87\xa2\xa7\x24\x07\x46\x1a\x59\x8c\x69\x92\x1d\x52\x4a\x5c\x47\xff\xbd\xa0\x3e\xfc\x15\xbb\x0d\x74\x10\x66\x38\xf3\x38\xf3\xde\x70\x66\xb4\x6a\x4c\x11\x94\x35\x94\xed\x66\x94\x70\xd9\x1b\xc9\x63\x13\x82\x35\xa3\xb5\x44\xd8\xd8\x16\x08\xe3\x85\x56\xc5\xfa\x38\xa5\x95\x98\x80\x98\xd1\x50\x2b\xcf\x52\x14\xc0\x0b\x6d\x3d\xf8\x40\x09\x57\xc6\x35\x61\xb9\x42\xdb\x38\xc2\x52\x27\x90\x3b\x89\x60\x02\x65\xb9\xaa\x28\xf2\x52\x06\x49\x09\x82\x03\x19\xa0\x24\x6c\x40\x43\x2f\x1c\xaf\x94\x29\xcf\x11\xf2\xbf\xff\x43\xcf\x35\x98\x55\xa8\x6f\x28\xf2\xa1\x28\xca\xd2\x7f\x84\x10\xfb\x93\xf9\xfc\x00\xfc\xbd\x51\x18\x81\xa3\x6f\x44\x3c\x6f\x48\x96\xe5\xad\x96\xde\x53\x52\xab\x12\x08\x63\x19\x7d\x77\x6c\xfa\xf6\xa6\xd7\xd7\xf3\xec\xbe\x85\x4b\x17\x75\xa0\x3d\x24\x70\x09\xf5\x14\x41\x96\x25\x61\x63\xb7\xbf\x8e\xbc\x76\x57\xc7\xd2\xab\xd2\x0e\xe8\xef\xd0\x75\xd2\xee\xaa\xc0\xad\xb8\x52\xcf\xa0\xf6\x15\x52\x2f\xf7\xf5\x66\x34\x54\x45\x5b\x5e\x4b\x7f\xca\x61\x7b\x29\x3d\x8f\xc4\xf6\x0d\x14\x02\x63\xb5\x06\xe8\x1f\x7f\xb1\xfc\x30\x7e\x5c\x3a\x07\xa6\xa4\xc5\x28\xc2\x45\x98\xf4\xad\x34\x23\x8f\x85\xd4\x7a\x89\xe0\x1b\x1d\x26\x32\x7b\x4e\x2e\xf1\x38\x32\x78\xb8\x3a\x22\xd1\x11\xa9\xb2\xb8\xe9\xd1\x08\xe3\xbe\x79\xdc\xa8\x70\x9c\x0a\x2d\x98\xc0\x1d\xf6\xff\x0f\x50\xc9\x46\xc7\xa7\x73\xaa\x4c\x7d\xa4\x4c\x3a\xd4\x24\x66\x94\xfc\x49\x16\xb0\x27\x31\x3a\x09\xeb\x8f\x9d\x35\x1e\xc4\xe0\x9a\x14\x99\xdc\x84\xa5\xc1\xae\x56\x1a\xc4\xbe\x08\x40\xb4\xc8\x76\x53\xc4\x09\x4d\x47\x95\x22\x84\x06\x4d\xd2\x47\xdf\x10\xa9\x01\xc3\xd2\x37\x45\x01\xde\x93\x6c\xb4\x4b\x69\x56\x80\xa4\x3b\x9a\xce\xdf\x20\x8c\x19\xd9\x19\x60\xc7\xba\xbc\xe6\xcf\x52\x85\x4f\x40\x77\x50\x55\x50\x84\x8c\x54\xb2\x80\x47\x6b\xd7\x24\x32\xcb\xe5\x93\x7c\xa1\xbb\xb0\x75\x90\x01\x97\x21\x20\x25\x1b\x08\xb5\x2d\x09\x4b\x1b\xd4\x7b\xe7\x30\x8c\x84\xa5\x91\xaa\x0c\xb8\x07\x54\x52\xab\x1f\x71\xab\x14\x76\xe3\x34\x04\xc8\x4e\xca\xec\x89\xf3\xb5\x7d\xa6\x2c\xdd\x57\x31\x8d\x47\xda\xd7\x7f\x96\x30\x30\x57\x87\x8d\xa6\xe4\xff\x27\xf9\x92\xc4\x6d\x01\x3e\x24\x95\x54\x3a\xce\xf6\xc8\x7a\x9c\xd1\x2e\x1d\xbb\x3c\x60\xe0\xf0\x16\x9d\x20\x64\x78\x47\x35\xc8\x12\xd0\x0f\xee\xd1\x10\x77\x0f\x79\x65\x91\x46\xd7\x1a\xb6\x89\x32\xc9\x21\x70\x3a\x68\xa5\x3e\x39\xb8\x5b\xc3\xf6\x81\x4d\xc6\xf8\x1f\xb7\xe8\x83\x58\xc3\x96\x07\xfb\xcd\x39\xc0\x5b\xe9\x81\xb2\x05\xc9\x12\xb2\x68\xa5\xce\x9d\x98\x82\x9f\xac\x32\x94\xdc\x1b\xc2\x16\xe4\xde\xdc\x1b\xd2\xc5\xb5\x1c\x39\xba\x71\x0b\xf1\xf1\xeb\x97\xcf\xdc\x07\x54\x66\xa5\xaa\x2d\xed\x4d\x27\xd1\x03\x9d\xa2\x58\x6a\x1a\xad\xd3\x7f\x59\x86\xbc\xe7\x6e\x3e\xa7\x6e\x21\x46\xe3\x30\xb1\x03\x7d\x6e\x4f\xd5\x14\xd0\x75\x2c\x7e\xf9\xcf\x00\x00\x00\xff\xff\x8a\x1d\x3b\x9a\xc6\x06\x00\x00"

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

	info := bindataFileInfo{name: "assets/js/manager.min.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesRuLc_messagesConfigMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\x4f\x6b\x33\x55\x14\xc6\x9f\x34\xf1\x55\xa2\x22\xbc\xe2\xce\xc5\x79\x05\x5f\x5e\x91\xa9\x33\x89\x95\x32\xed\xb4\x62\x6c\xb1\xd8\xd0\x50\xa2\x22\xb8\xb9\x4d\x6e\x26\x83\x93\x3b\xe1\xce\x4d\x69\xa5\x8b\x34\x15\x5d\x58\xba\x10\xdd\xd5\x85\x0b\xf7\xc1\x12\x88\xfd\x93\x82\x2b\xdd\x9d\xfb\x05\xfc\x04\x7e\x00\x37\x82\x4c\xd2\xb4\x78\x16\x33\xcf\x73\xce\xef\x3c\xf7\xfc\xf5\xb8\xf0\x23\x00\xbc\x00\xe0\x75\x00\x5f\x00\x78\x11\xc0\x39\x66\xf5\x07\x80\x37\x00\xfc\x79\xd7\xff\x1b\xc0\xcb\x00\xfe\x01\xf0\x04\xc0\xa3\xdc\xcc\x3f\xcb\x01\x8f\x01\xbc\x97\x03\x5e\x01\xb0\x95\x03\xda\x39\xe0\xf3\x1c\xe0\x03\xf8\x7d\x61\xf6\xc6\xbf\x0b\xc0\x4b\xd9\x5e\x1e\xa8\x00\x78\x35\x3f\xf3\x9f\xe5\x81\xd7\x00\x34\xf3\x40\x01\xc0\x57\x79\x20\x87\x87\x7a\x74\xf7\xcf\x03\x78\x1e\xc0\x02\x66\x5c\x96\xf9\x5c\x36\x68\x24\xaa\x15\x85\x85\x0d\x25\xf6\x62\xd9\x24\x2d\xc3\x28\x35\x52\x93\x96\xad\x58\x36\x4c\x94\xa8\x39\xf2\x51\x92\x9a\xb9\xde\x9e\x42\x73\x57\x15\x07\x14\xcb\x7d\x19\x53\xd2\xa2\xae\xd0\x69\xa4\x42\x32\x87\x5d\x99\xce\x89\x1d\xd3\x96\xfa\xde\xd5\x12\x6d\x48\xf5\x3a\x7b\xf7\x11\x8e\x11\x7b\x85\x50\x77\x1b\xd8\x95\xdd\x44\x1b\xa7\x9a\x86\x51\xd3\xf9\xa0\x17\xa6\x4e\x3d\xf1\xa9\x29\xf7\xdf\xff\x32\x6a\x8b\x4e\xb2\xa8\x7b\xc5\xda\x4e\xdd\xa9\x68\x29\xb2\xeb\x9c\x0f\x85\x91\x3e\x95\x5c\x6f\xd9\x71\xcb\x4e\xa9\x4c\xa5\xb2\xbf\xb4\xf4\xb6\x5b\x76\xdd\xe2\xb6\x48\x8d\x53\xd7\x42\xa5\xb1\x30\x89\xf6\xe9\xe3\x69\x06\x55\x7b\x5a\x74\x92\x66\x42\xab\xff\x0b\x5e\x2b\x6e\x0b\x15\xf6\x44\x28\x9d\xba\x14\x1d\x9f\xee\xbd\x4f\xbb\xbd\x34\x8d\x84\x2a\x56\xb7\xaa\x1b\xce\xa7\x52\xa7\x51\xa2\x7c\xf2\x16\xdd\x62\x25\x51\x46\x2a\xe3\xd4\x0f\xbb\xd2\x27\x23\x0f\xcc\x3b\xdd\x58\x44\x6a\x85\x1a\x6d\xa1\x53\x69\x82\x4f\xea\x9b\xce\xf2\x03\x97\xdd\xd3\x92\xda\xd9\x50\x8d\xa4\x19\xa9\xd0\xa7\x62\x2d\xee\x69\x11\x3b\x9b\x89\xee\xa4\x3e\xa9\xee\xd4\xa6\x41\x79\x85\x66\x32\x50\x6f\x7a\x6e\x10\x78\xf4\xf4\x29\x65\xd2\x7d\x12\x78\x1e\xad\x93\x4b\xfe\xd4\xaf\x05\xa5\xf9\x68\x35\x78\x37\x93\xcf\xa6\xd8\xaa\xe7\xd2\xd1\xd1\x6c\x65\x2d\x28\xb9\x6f\xd1\x3a\x79\xe4\x53\x69\x05\xfc\x3d\x5f\xf2\x95\x3d\xb3\xdf\xf2\xd8\x0e\xec\x29\xd9\x3e\x8f\xf8\x82\xc7\xf6\xd8\x0e\x6c\x9f\x87\xf6\x1b\x1e\xdb\xb3\x69\xdb\x7e\xcd\x57\x3c\xe2\x4b\x7b\xcc\x63\x1e\x83\x7f\xe1\x49\x46\x81\x7f\xe6\x91\xed\xf3\xaf\xd9\x17\x7c\xce\xc3\x3b\xe4\x9a\x87\x7c\x65\x4f\xf9\xc6\x7e\xc7\xbf\x91\x3d\xb1\x7d\x9e\x64\x14\xdf\xd8\x53\xe2\x5b\x1e\xda\xfe\x94\xbb\xe1\x0b\x1e\x92\x1d\xf0\x98\x6f\x33\x02\xfc\x83\xed\xdb\x13\xbe\xe0\x09\x8f\xc0\x3f\xf1\x84\xaf\xb3\xec\x6c\x67\x62\xfb\x76\xc0\x43\x84\xbb\xb5\x0a\xfe\x0b\x00\x00\xff\xff\xda\x37\x4c\x26\x7f\x03\x00\x00"

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

	info := bindataFileInfo{name: "locales/ru/LC_MESSAGES/config.mo", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesRuLc_messagesGrpcMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x90\xcd\x6e\xd4\x30\x14\x85\xcf\x54\x65\x93\x25\x6b\x16\x97\x05\x15\x08\x5c\xec\x84\x4a\x95\x27\x9e\x22\x86\x56\x42\x34\x22\x8c\x42\xf7\xd6\xc4\x64\x22\x12\x3b\xb2\x1d\x04\xd2\xbc\x06\xcf\xc3\x9b\xf0\x2c\x68\x92\xe1\xa7\x67\x75\x3e\xdd\x73\x8f\xae\xee\xaf\x87\xa7\x3f\x00\xe0\x04\xc0\x23\x00\x2f\x00\x3c\x00\x90\x63\x56\x09\xe0\x14\xc0\x47\x00\xbb\x05\x70\x77\xe4\x9f\x0b\x60\x71\xcc\x9c\xe0\x3f\x35\x9b\x72\x8d\x8d\x19\x9c\x8f\xac\x08\x4d\x5b\xb3\x37\x63\x13\x58\xe5\x24\xd5\xe6\xeb\xeb\x2f\xed\x4e\xf7\xee\xdc\x8f\x49\xf9\xa1\x62\x6b\x6f\x74\x6c\x9d\x65\x6f\x75\x34\x92\x52\x2e\x2e\x19\xcf\x58\x9a\x51\x9a\xc9\x8b\x8b\xe7\x3c\xe3\x3c\xb9\xd5\x21\xb2\xca\x6b\x1b\x3a\x1d\x9d\x97\xf4\x7e\xea\xa0\x62\xf4\xba\x77\xb5\xa3\xfc\x5e\xf1\x2a\xb9\xd5\xb6\x19\x75\x63\x58\x65\x74\x2f\xe9\x2f\x4b\xda\x8c\x21\xb4\xda\x26\xc5\xbb\xe2\x9a\xdd\x19\x1f\x5a\x67\x25\x89\x73\x9e\xac\x9d\x8d\xc6\x46\x56\x7d\x1f\x8c\xa4\x68\xbe\xc5\x97\x43\xa7\x5b\xbb\xa4\xed\x4e\xfb\x60\xa2\xfa\x54\xdd\xb0\xcb\x7f\xb9\xc3\x3d\x9f\x8d\x67\xd7\x76\xeb\xea\xd6\x36\x92\x92\xb2\x1b\xbd\xee\xd8\x8d\xf3\x7d\x90\x64\x87\x09\x83\xca\x96\x34\x5b\x65\x9f\x08\xae\x94\xa0\xb3\x33\x3a\x58\xfe\x58\x09\x41\x57\xc4\x49\x4e\xbc\x52\xe9\x9f\x51\xae\x5e\x1d\xec\xd3\x29\x96\x0b\x4e\xfb\xfd\xbc\xb2\x52\x29\x7f\x46\x57\x24\x48\x52\xba\x9c\xbf\xfd\x3b\x00\x00\xff\xff\x73\x83\x3e\x08\xc4\x01\x00\x00"

func localesRuLc_messagesGrpcMoBytes() ([]byte, error) {
	return bindataRead(
		_localesRuLc_messagesGrpcMo,
		"locales/ru/LC_MESSAGES/grpc.mo",
	)
}

func localesRuLc_messagesGrpcMo() (*asset, error) {
	bytes, err := localesRuLc_messagesGrpcMoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/ru/LC_MESSAGES/grpc.mo", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesRuLc_messagesManagerMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8f\x4f\x6b\x13\x51\x14\xc5\x4f\xc7\xfa\x87\x11\x44\x5c\xbb\xb8\x2e\x2c\x8a\xbc\x3a\x33\xb1\x50\x5e\x3b\xa9\x1a\x1b\x10\x13\x2c\x65\xd4\xf5\x33\x79\x4d\x06\x27\xef\x85\x37\x33\xa2\xd2\x4d\xb3\xb5\x2b\xd1\x9d\x20\xf8\x09\x4a\x21\x98\x56\x48\xdd\xb8\x72\xf3\x66\xe3\x4e\x3f\x8b\x4c\x32\x5a\xbc\x9b\x7b\xce\x9d\xdf\x9c\x7b\xdf\xef\x2b\x8b\x1f\x00\xe0\x3c\x80\xab\x00\x22\x00\x17\x01\xec\x63\x5e\x13\x00\xe7\x00\x1c\x01\x58\x04\xf0\x0d\xc0\x59\x00\xdf\xab\xfe\xa3\x9a\x17\x15\xf7\x13\x40\x7f\x01\xf8\x05\xe0\x32\x80\xa6\x03\x5c\x02\xf0\xcc\x01\x2e\x00\xd8\x71\x00\x17\x80\x76\xe6\xfc\x9b\xca\xef\x39\xc0\x42\xb5\xd3\xa9\xfa\x22\x4e\xab\x64\xcf\x54\x77\x96\x7b\x71\xaf\x93\xc5\x5a\xa1\x21\x92\x04\xcd\x58\x26\x5d\xb4\xc4\x73\x99\x20\x7a\x3d\x94\x48\x33\x23\xc5\x00\xdb\x72\xa8\x4d\xc6\xda\x69\x2f\xee\xb2\xfb\x79\x2f\x65\x91\xe6\xd4\x95\x2f\xef\xbe\x88\xfb\x62\xa0\x97\x4d\xee\x6e\x3d\x8e\x58\xc3\x48\x51\xa6\xb1\x07\x22\x93\x9c\x02\xcf\x5f\x65\x5e\x8d\x05\x35\x0a\x6a\x7c\x65\xe5\x96\x57\xf3\x3c\xb7\x25\xd2\x8c\x45\x46\xa8\x34\x11\x99\x36\x9c\x1e\xcd\x32\xa8\x9d\x1b\x31\xd0\x5d\x4d\xeb\xff\x05\xd7\xdd\x96\x50\xbd\x5c\xf4\x24\x8b\xa4\x18\x70\xfa\xe7\x39\x6d\xe7\x69\x1a\x0b\xe5\xb6\x1f\xb6\x37\xd9\x53\x69\xd2\x58\x2b\x4e\xfe\xb2\xe7\x36\xb4\xca\xa4\xca\x58\xf9\x0a\x4e\x99\x7c\x95\xdd\x1e\x26\x22\x56\x6b\xd4\xe9\x0b\x93\xca\x2c\x7c\x12\x35\xd9\xea\x29\x57\xde\xb3\x23\x0d\xdb\x54\x1d\xdd\x8d\x55\x8f\x93\xbb\x95\xe4\x46\x24\xac\xa9\xcd\x20\xe5\xa4\x86\x33\x9b\x86\xb5\x35\x9a\xcb\x50\x5d\xf7\xbd\x30\xf4\x69\x69\x89\x4a\xe9\x5d\x0b\x7d\x9f\x36\xc8\x23\x3e\xf3\xf5\x30\xf8\xfb\x69\x3d\xbc\x53\xca\x1b\x33\x6c\xdd\xf7\x68\x77\x77\xfe\x4b\x3d\x0c\xbc\x9b\xb4\x41\x3e\x71\x0a\xd6\x60\xdf\xdb\xb1\x3d\x2a\xf6\x8a\x91\x3d\xb4\x13\x3b\x86\x7d\x57\xbc\xb5\x5f\xec\xa1\x3d\x28\x46\xc5\x3e\xec\x27\x3b\xb5\x5f\xcb\xf9\x47\x3b\x2e\x46\xf6\xd8\x1e\xc0\x7e\xb6\x13\x7b\x02\x7b\x62\xa7\xc5\xc8\x4e\xed\x31\xfe\x04\x00\x00\xff\xff\x12\x91\x5c\x79\x8c\x02\x00\x00"

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

	info := bindataFileInfo{name: "locales/ru/LC_MESSAGES/manager.mo", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"templates/views/manager.html":      templatesViewsManagerHtml,
	"assets/js/manager.min.js":          assetsJsManagerMinJs,
	"locales/ru/LC_MESSAGES/config.mo":  localesRuLc_messagesConfigMo,
	"locales/ru/LC_MESSAGES/grpc.mo":    localesRuLc_messagesGrpcMo,
	"locales/ru/LC_MESSAGES/manager.mo": localesRuLc_messagesManagerMo,
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
				"config.mo":  &bintree{localesRuLc_messagesConfigMo, map[string]*bintree{}},
				"grpc.mo":    &bintree{localesRuLc_messagesGrpcMo, map[string]*bintree{}},
				"manager.mo": &bintree{localesRuLc_messagesManagerMo, map[string]*bintree{}},
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
