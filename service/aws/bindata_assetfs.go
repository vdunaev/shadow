// Code generated by go-bindata.
// sources:
// templates/index.tpl.html
// DO NOT EDIT!

package aws

import (
	"github.com/elazarl/go-bindata-assetfs"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
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
	name string
	size int64
	mode os.FileMode
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

var _templatesIndexTplHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x56\xc1\x8a\xdb\x30\x10\xbd\xe7\x2b\x84\x76\x8f\x75\xcc\x5e\x8b\x62\xc8\xa1\xd7\x52\xda\xfe\xc0\xc4\x52\x12\x81\x56\x32\xd2\x38\x65\x59\xf6\xdf\x3b\x96\xed\xdd\x71\x62\xa7\xa5\x10\xba\x1b\x88\x63\x8f\xde\xbc\x79\x92\x9e\x26\x7e\x7e\x16\xd6\xd7\xae\xd5\x46\xc8\x7d\x0c\x1e\x8d\xd7\x52\x48\x07\x4f\xa1\xc5\xf2\x68\x40\x9b\xb8\xc6\xc6\xad\x8f\xf8\xe8\xa4\x78\x79\x59\xad\x94\xb6\x27\x51\x3b\x48\x69\x23\x1b\x38\x98\xa2\x47\xc9\x6a\x25\xe8\xa3\x8e\x0f\xd5\xf6\x57\x52\x25\xfd\xae\x54\x49\xd8\x6a\x48\x89\xc1\x99\x8d\x44\xd8\x35\xe0\x8d\x1b\xe1\xad\x1b\xb9\x3c\x9c\x04\x7d\x0b\x42\x24\xf9\x86\x76\x36\xe1\x00\xce\x09\xce\x0e\x63\x4d\x34\xc9\x78\x04\xb4\xc1\xcb\x91\x04\x6a\xb4\x27\x23\x2b\x05\xe2\x18\xcd\x7e\x23\xef\xa0\x69\x9c\xad\x33\x8a\x68\x21\x5a\x28\x6a\x9a\x27\x71\x74\xf0\xc9\xe0\x6b\x4d\x29\x34\x20\x14\x18\x0e\x87\x31\x52\x6d\x19\x54\x95\x50\xa9\xd2\xd9\x3f\xca\x62\x3a\x30\x34\xb6\xbe\x54\x30\x86\xaf\xd6\xfe\x99\x41\xff\x52\x35\xb5\xbb\x54\x47\xdb\xcc\x4f\xff\x6c\xf4\xaa\x86\x1f\x1c\x3b\x95\xa2\xca\xd6\xd1\x2e\xe7\x5b\x66\x0e\x4a\xcb\xb5\x48\x17\xdf\xc0\x19\x2f\xf0\x8c\x2e\x22\x86\x6d\x14\x56\x9f\xed\xd1\x1b\x4f\xe6\xea\xfc\x61\x58\x36\x3d\xe4\x6b\x91\x90\xa4\x1a\x3d\x3c\x91\x0a\x6d\x7c\x32\xfa\x2c\xbf\xe7\xe8\xfc\x7b\x19\x1f\xc6\xaa\x3b\x55\xd2\x75\x71\x98\xd9\x42\x6c\xbf\x7f\x9d\x07\x77\xd1\xd9\x22\x0a\x77\x41\x3f\x5d\xc6\xe9\x58\x46\xf0\x07\x23\xee\xed\x27\x71\xcf\x56\x40\x7c\xde\x88\x35\xb7\x62\x77\x24\x2f\x59\xe3\x92\x60\x5d\x11\x35\x68\x4d\xc4\xe2\x81\x72\x49\xd9\xe2\xdc\x33\x96\x17\x5f\x7f\x73\x80\xfb\x10\x1f\x99\x00\x9a\xf3\x22\x0d\x45\x67\x84\x10\x29\x35\x99\x59\xdd\xe5\xcc\x72\x50\xb0\xdb\x43\x66\xa0\xa1\xa9\xbc\x3e\xff\x85\xa1\x7a\x27\x0d\x67\xed\xdd\x79\x28\x1f\xef\x9b\xb9\x27\xcf\x3a\xfb\xa6\x6f\x23\x37\x77\x4c\x2e\xd8\x17\xfb\x48\xee\x98\x36\xc3\x77\x67\x12\xde\x7f\x97\xbd\x32\xa2\xbf\x78\xdd\x04\xeb\xf1\x26\x8e\xe2\x2b\x95\x8d\x35\xf9\x6f\xb8\xb9\xbf\x78\xf9\x49\xe9\x6b\x6e\x5b\x24\x18\x57\xea\xff\xf8\x94\xdd\x8e\xce\xbd\xfe\x4a\xb6\x0f\x01\xcf\x5e\xc9\x7e\x07\x00\x00\xff\xff\xc4\x66\xa2\x9e\xc4\x09\x00\x00")

func templatesIndexTplHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesIndexTplHtml,
		"templates/index.tpl.html",
	)
}

func templatesIndexTplHtml() (*asset, error) {
	bytes, err := templatesIndexTplHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/index.tpl.html", size: 2500, mode: os.FileMode(420), modTime: time.Unix(1435004621, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"templates/index.tpl.html": templatesIndexTplHtml,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"templates": &bintree{nil, map[string]*bintree{
		"index.tpl.html": &bintree{templatesIndexTplHtml, map[string]*bintree{
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
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, Prefix: k}
	}
	panic("unreachable")
}
