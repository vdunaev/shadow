// Code generated by go-bindata.
// sources:
// templates/views/list.html
// DO NOT EDIT!

package metrics

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

var _templatesViewsListHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x57\x5b\x6f\xe3\x36\x13\x7d\xcf\xaf\x20\x08\xe3\xdb\x5d\x20\xba\x24\xfb\x6d\x0b\xb8\xb2\x8a\xa2\xb7\x00\xed\x16\x28\x62\xec\x43\x5f\x0a\x4a\x1c\x5b\x74\x29\x52\x21\x29\xc7\x86\xa1\xff\x5e\x50\x94\x14\xd9\x91\x2f\x6b\x3b\x0f\xf5\x83\x20\x0e\xc9\x33\x67\xce\x8c\xc6\xe4\x66\x83\x28\xcc\x98\x00\x84\x53\x29\x0c\x08\x83\x51\x55\xdd\xdc\x6c\x36\x88\xcd\x90\x0f\x4a\x49\x65\x0d\x11\x65\x4b\x94\x72\xa2\xf5\x04\x13\x0e\xca\xa0\xfa\xe9\x51\x22\xe6\xa0\xda\x01\xd3\x39\xd3\x9a\x24\x1c\x70\x7c\x83\x10\x42\x51\x52\x1a\x23\x05\x32\xeb\x02\x26\xd8\x0d\x70\x8b\x93\x72\xa9\x01\x23\x4a\x0c\x69\xb7\x36\xe0\x18\x11\xc5\x88\x97\x31\x4a\x41\x4c\xb0\x51\x25\xe0\xf8\x7f\x86\xe5\xa0\xbf\x8b\x02\x07\xe3\x1c\x6c\x36\x0d\x49\xff\xe7\x8e\x6a\x40\xd9\x32\xb6\x11\x80\xa0\x75\x30\x7d\xf2\x4a\x3e\xb7\xdc\x7a\xd6\xd5\xdf\x05\x11\xc0\x9b\x99\xd7\xb3\x86\x99\x2e\xa6\x6e\x45\x76\x1f\x7f\x06\xa3\x58\xaa\xa3\x20\xbb\xdf\x99\xec\x6d\x4f\x39\x10\x35\x63\x2b\x1c\x37\xd4\xba\x45\x3b\xc3\x2d\x97\x6d\x3a\x76\x70\x8b\x76\x85\x81\x95\xf1\xf2\xd2\x00\x45\x33\x29\x8c\x77\xf7\x11\xe5\x5e\xe2\x7d\x0c\x77\x76\xd8\xdf\x63\x26\x9f\xb5\xd5\x8a\x83\x40\x7e\xee\x48\xa3\xaa\x42\xcd\x6b\x93\xed\xb2\xa0\xc4\xe2\x55\x15\xe2\x44\x1b\xd4\x8e\x89\x41\x91\x4e\x15\x2b\x4c\x93\x48\x52\x14\x9c\xa5\xc4\x30\x29\x82\x05\x59\x12\x37\x89\x63\x2a\xd3\x32\x07\x61\xfc\x67\xc5\x0c\xbc\xb7\xbb\xa7\xf2\xd1\x28\x26\xe6\xef\xdf\xd9\x54\x35\x88\xfe\x2f\x52\xe5\xc4\x20\x7c\x1f\x86\xdf\x78\xe1\x9d\x17\xde\x4f\xef\x3e\x8d\xc3\xff\x8f\xc3\x4f\x7f\x85\xdf\x8e\xc3\xd0\x56\xe1\xbb\x0f\x1f\xa2\xc0\x41\xc7\x2f\xd9\xdc\x52\x23\x28\xe2\x9b\x6d\x8b\xb1\xd5\xd7\x69\x54\x0f\xea\xa7\xa7\x8d\x62\x05\xd0\xba\xdc\x6a\x0b\x46\x8c\x4e\x70\xa3\xc0\x80\x68\x91\xc9\x80\xd0\xd7\x76\x37\xa7\x86\x27\x9a\x8d\xf1\x1f\x24\x87\x28\x30\xd9\xe1\x55\x0f\xc0\x8b\xa3\xab\xda\x60\x72\xea\xa5\x92\x7b\x77\x38\x9e\xae\x8b\x13\xc0\x7f\x27\x09\x70\x7d\x7c\xdd\x17\xc2\xcb\x03\x70\x51\x30\x14\xab\x5d\x3f\xa8\x4e\x64\x12\x49\xd7\xc3\x50\x9b\x0d\x52\xb6\x5f\xa0\x51\x7e\x8b\x46\x4e\x79\x34\x9e\xf4\x2b\xf2\x1c\xb5\xa9\xad\x8e\x06\xce\xff\x09\x5c\xc5\x30\x29\x7c\x9b\x04\x54\x55\x51\x60\xf6\xa4\xf1\xc8\x7e\x9b\x9e\x4b\xf6\xdb\x3c\x9d\xb4\x7f\xef\xe4\xb6\x6a\xfc\x16\x8d\xb8\xcd\xaa\x15\x6d\xc8\xa1\x4b\xf9\x3e\x19\x3b\x8f\xba\x20\xa2\xad\x2a\x87\x57\x3f\x3d\x5d\xa6\x29\x68\x8d\xeb\x78\x6a\x93\xff\x1b\xac\x51\x55\x8d\xd1\x8b\xa5\x2e\x97\x3a\x2a\x0b\x73\x94\xfa\xc0\x57\xbb\xc5\xe5\x62\x6d\x46\xb6\x25\xed\x13\xc4\x66\xc0\x77\x0d\xe8\x98\x2a\xae\x07\xc2\x53\x03\x88\x53\x59\x0a\x03\x0a\x9f\xad\x26\x13\x33\x89\xe3\xa5\x95\xcb\xe9\xd7\xf0\xfb\xd1\x01\x7f\xb5\x90\x5c\xc3\x36\xc3\x39\x29\xe7\x70\x7d\x7e\xbf\x5a\xd8\xcb\xd9\x65\x4c\x1b\x39\x57\x24\xbf\x94\x61\x9d\x88\x2d\x86\x0f\x2d\xb4\xff\x48\xf2\x82\x43\xad\xe8\x89\x5c\x8f\x38\xcb\x99\x38\xe8\xea\x33\x13\x57\x72\x44\x56\x87\x1d\x91\xd5\x75\x1c\x2d\xed\x51\x4a\xa4\x70\xd0\xdb\x97\x66\x51\xcf\xe5\x89\x3d\xe9\xe9\x16\x8d\xea\x1a\xea\x7f\x82\x2f\xf0\x7f\x96\x44\x18\xc6\xe1\xfc\xa6\xe4\xa2\xb0\xdc\x9f\xba\x4e\xb4\xbc\x6a\x0f\x42\xfb\x4a\xd8\x1e\x38\x2f\x6e\x00\xaf\xcb\x77\x6a\x61\xdf\xbe\x74\xfb\x6e\xde\xac\x6c\xb7\x9c\xbc\x69\xc9\xf6\x3d\x5d\xb9\x5c\x1d\xf4\x7f\xa9\x54\xcf\xfd\x53\x1d\x3e\xcc\xed\x47\x8c\x82\x81\xe3\x5c\x14\xd4\xe7\xe7\xc1\x7b\x4c\xf3\xfa\xfa\x0e\xd6\xbb\x68\xda\x83\x63\xfd\x55\x45\x9c\x89\x7f\x50\xa6\x60\x36\xc1\x01\x25\x3a\x4b\x24\x51\x34\x20\x5a\x83\xd1\xc1\x12\x04\x95\x4a\x07\xdd\x81\x5d\xfb\x02\x8c\x97\xe8\x20\xd5\xce\x3a\x75\xd6\x44\x4a\xa3\x8d\x22\x85\x9f\x33\xe1\xa7\x5a\x7f\xbf\x9c\xd8\xcb\xc6\x0f\x2f\xb7\x14\x3f\x29\x19\xb7\x44\x30\x52\xc0\x27\x58\x9b\x35\x07\x9d\x01\xd8\xfb\xd5\x59\x34\x66\x6c\x05\xd4\x46\x02\xaa\xa5\x54\x9b\x1e\x6a\xd3\xc5\x9c\x86\x95\x5b\x68\xa7\x5b\x73\x19\xd3\x2a\x3d\x9d\x71\xb0\xd0\xc1\xe2\xa9\x04\xb5\xf6\x7b\xda\x59\x76\x8b\x83\xe4\xe2\xee\x0e\x76\x9e\x5f\xab\xce\xe2\x40\xbe\xde\xd8\x7b\x2f\x4f\x3b\x34\xfa\xe9\xfa\x2a\x22\x5d\x6e\xfe\x0d\x00\x00\xff\xff\x72\x65\x9f\xbf\x3d\x11\x00\x00")

func templatesViewsListHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesViewsListHtml,
		"templates/views/list.html",
	)
}

func templatesViewsListHtml() (*asset, error) {
	bytes, err := templatesViewsListHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/views/list.html", size: 4413, mode: os.FileMode(420), modTime: time.Unix(1502810433, 0)}
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
	"templates/views/list.html": templatesViewsListHtml,
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
	"templates": &bintree{nil, map[string]*bintree{
		"views": &bintree{nil, map[string]*bintree{
			"list.html": &bintree{templatesViewsListHtml, map[string]*bintree{}},
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
