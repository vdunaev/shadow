// Code generated by go-bindata.
// sources:
// locales/ru/LC_MESSAGES/config.mo
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

var _localesRuLc_messagesConfigMo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x53\xdf\x8f\x13\x55\x18\x3d\xb4\xa5\x5d\xab\x88\xa0\x82\xbf\x73\x49\x04\x24\x38\x38\xed\x42\x42\x0a\x05\x23\x2c\x88\xb0\xb0\x61\x8b\x3c\x92\x6b\x3b\xdb\x9d\x30\x9d\xdb\xcc\x4c\x89\x26\x98\x2c\xbb\xc6\x48\x20\x41\x13\x25\xc6\x28\xb0\xca\x9b\xc6\x94\x75\x8b\x23\x4b\xbb\x4f\x3e\x99\x98\xef\xfe\x01\xea\xa3\x2f\x3e\xf8\xe4\x9b\x3f\x72\xe7\x4e\xa7\xdd\x8d\x0f\x3a\x2f\x73\xbf\xf3\x9d\x7b\xbe\x73\xcf\x9d\xf9\x79\x63\xe6\x23\x00\x78\x16\xc0\x33\x00\x32\x6b\x80\x1d\x00\x7e\x5b\x83\xe8\x79\x3f\x05\xac\x03\xf0\x41\x0a\x78\x18\xc0\xcd\x14\xf0\x10\x80\xaf\x53\xc0\x46\x00\x4b\x29\xe0\x09\x00\x3f\xa6\x80\x4d\x00\x7e\x8d\x79\x7f\xc5\xbc\x75\x69\xfd\x7e\x2a\x0d\xbc\x0a\x60\x7b\x1a\x78\x10\x40\x33\xad\x79\x17\xd3\xc0\x63\x00\xae\xc4\xf8\xad\x34\xf0\x08\x80\xaf\xd2\xc0\xa3\x00\xba\x31\x2e\xd3\x7a\xde\x4f\xb1\xde\xef\x69\x60\x03\x80\x3f\x63\x9d\xf5\x19\x8d\x3f\x97\x01\x9e\x06\xb0\x33\xa3\x7d\x1f\xcd\xe8\x7d\x67\x32\x5a\x67\x3a\xa3\xf9\x41\x8c\xcf\x65\x80\xe9\x35\xc0\xf5\xb8\xfe\x7b\xad\xee\x6f\xcc\xea\xf7\x96\xac\xf6\x67\x66\x81\x2d\x00\xc6\xb2\xc0\x2e\x00\xf5\x18\xbf\x9e\xd5\x73\xbf\xcc\xea\xfd\x61\x16\x38\x0e\xe0\x87\x18\xdf\x9c\xd3\xe7\x79\x3e\xa7\xf3\xdd\x93\x03\x32\x00\x4e\xe7\x80\xad\x00\xce\xe6\x74\x7e\x6f\xe7\x80\xbc\xd2\xcb\xe9\x1c\xe7\x73\x7a\x7f\x18\xe3\xdf\xe7\xb4\x1f\x8a\xfb\xbf\xe4\x00\x13\xc0\x1f\x31\x6f\xfb\x88\x3e\xf7\xee\x11\x20\x0b\x60\x62\x44\xcf\x3b\x3b\x02\x3c\xa9\xce\x3b\x02\xa8\x2b\x55\xeb\xb4\xbe\xda\xc8\xff\xe3\xf1\x5a\xed\x51\xbe\x36\xc4\xb5\xda\xab\x7c\x8f\x00\x78\x40\xdd\x1f\xf4\x7c\x75\x27\x6b\xa1\xb3\x55\x73\x55\xa6\xf9\xf8\xfb\xe9\x3f\x39\xac\x7c\x52\xf1\x5b\xe5\xb3\x1e\xda\xe7\xe6\xa1\xbe\x3a\xff\xa6\x7e\x51\x15\xee\x94\x5d\xcf\x1c\xe2\x8e\x63\x79\x49\x25\x5c\x5f\x38\x56\xbf\x3c\x6c\xbd\xd1\xaa\x0f\x8a\x0b\x96\x23\x9a\x0d\xcb\x0d\xfe\x05\x62\x4d\xee\xda\xd5\xa4\xd1\xf2\x78\x60\x0b\x97\x4d\x09\xaf\xc1\x13\xfe\x98\x5b\x15\xb5\xc1\xb4\x31\xcf\x13\x49\x71\x84\x07\xdc\x49\x0a\xdb\x72\x6a\x3e\xb3\xfb\x02\x6c\x4a\x01\xe7\x5c\xde\xb0\xca\xd1\xb2\x70\xee\x02\x77\x5a\xd6\x8b\x51\x51\x1c\x6a\x14\x75\x23\x11\x6a\x39\x89\xe8\x51\xcb\xb5\xbc\xc1\x8c\x63\xae\xd6\xb6\x85\x3b\x00\x5f\x9b\x3c\x75\xb2\xbf\x3e\x21\xea\xcc\x51\x27\xec\x03\xe3\xb6\xe3\xd8\xbe\x55\x15\x6e\xcd\x4f\x30\x51\x4b\x86\x9d\xe4\xae\x58\xd5\x9e\x18\x4e\x65\xc2\x13\xb5\x56\x55\x0d\xec\x23\x93\x2b\xd9\x93\xd3\xc2\x4b\xc2\x9a\x0c\x78\xf5\x7c\xe0\xf1\xaa\xc5\x9c\xd5\x4e\x26\x03\xcf\x76\x93\x9b\xa9\xd8\x0d\x6b\x55\xd2\x95\xb7\x9a\x89\xad\xb3\xdc\x73\x87\xd8\xbc\x56\x63\x5c\x07\x8a\xd3\x56\x53\x78\x81\x31\xee\xd7\xed\x9a\xf1\x4a\xab\xee\x1b\x15\x51\x62\x35\xeb\xc2\xcb\xe7\xed\x69\xde\x10\xbb\xbc\x56\x7e\xe2\x54\xc5\x38\xe4\x59\x51\x50\xc6\x61\x1e\x58\x25\x56\x34\x0b\x7b\x0d\x73\xd4\x28\x8e\xb2\xe2\x68\x69\xcf\x9e\x9d\xe6\xa8\x69\xe6\x4f\x70\x3f\x30\x2a\x1e\x77\x7d\x87\x07\xc2\x2b\xb1\xe3\x91\x06\x1b\x6f\x79\xbc\x21\x6a\x82\xed\x5f\x21\x7c\x20\x7f\x82\xbb\xf5\x16\xaf\x5b\x46\xc5\xe2\x8d\x12\x4b\xea\x12\x3b\xdd\xf2\x7d\x9b\xbb\xf9\xf1\x63\xe3\x63\xc6\xeb\x96\xe7\xdb\xc2\x2d\xb1\xc2\x2e\x33\x7f\x48\xb8\x81\xe5\x06\x86\x3a\x60\x89\x05\xd6\x9b\xc1\x4b\x4d\x87\xdb\xee\x3e\x56\x9d\xe6\x9e\x6f\x05\xe5\x33\x95\x23\xc6\xde\x01\x4f\xf9\x99\xb2\x3c\x23\xfa\xf6\x6c\xb7\x5e\x62\xf9\x09\xa7\xe5\x71\xc7\x38\x22\xbc\x86\x5f\x62\x6e\x33\x2a\xfd\xf2\xe8\x3e\xa6\x97\x65\x77\x6b\xc1\x2c\x97\x0b\x6c\xdb\x36\xa6\x96\xe6\x96\x72\xa1\xc0\x0e\x32\x93\x95\xa2\xfa\x40\xb9\xd8\x6f\xed\x2f\xef\x56\xcb\x17\x22\xda\xfe\x82\xc9\x2e\x5e\xd4\x5b\x0e\x94\x8b\xe6\x0e\x76\x90\x15\x58\x89\x15\xf7\x81\x3e\xa6\x2e\x85\xf2\x5d\x0a\xa9\x2d\x67\xa9\x27\x67\x40\x9f\x50\x8f\xba\xf2\x12\xf5\x68\x49\x5e\x05\xdd\x90\xb3\xb4\x44\x6d\x5a\xa4\x7b\xd4\x06\xdd\xa2\x36\x7d\x2b\x67\xa8\x4d\x77\xa8\x27\x67\x35\x78\x93\xda\x4a\x47\x15\x2c\xea\xad\x64\x84\xa0\xdb\x4a\x9b\xee\xab\x29\x8c\x42\xea\xca\x59\xea\xc8\x19\x5a\xa0\xb6\x12\x67\xb4\x20\x67\xa8\x43\xf7\xa9\xa3\x74\xb4\x87\x45\x0a\xe5\x0c\xf5\x68\x41\x5e\x56\xd2\xca\xc9\x7b\x14\xd2\x1d\x3d\xf2\x76\x64\xb8\xad\x3c\x52\x57\x5e\xa1\xef\x94\x0b\x65\xf9\x1a\xa3\x05\x26\xdf\x19\xcc\xa3\xce\xff\xff\x4d\xb5\x56\x5f\xf8\x86\xbc\x44\x5d\xe5\x24\x02\x3a\x51\x6a\x43\x13\x54\x7a\x3d\xea\xc6\xec\xe8\x57\xa5\x2f\xb4\x75\x75\x1e\x79\x95\xd1\x12\xf5\xe8\x9b\xfe\x79\x74\x58\xf2\x1a\xe8\x53\x0a\x69\x89\x96\x28\x94\x97\xa8\x43\xf7\xe4\x1c\x75\x69\x51\x5e\x51\x29\x77\xe8\x2e\x85\x74\x1f\xf4\x59\x44\xef\xad\x66\x0c\x45\x8e\xa1\x1f\x98\xe6\x57\xd2\xe6\xa9\x47\xf7\xd4\x65\xc8\xcb\x51\xb6\xda\xe3\xd0\x3f\x2c\xe7\xfe\x93\xd3\x79\x39\x1b\x41\x71\xf6\x43\x97\xb9\xea\xe6\x3e\xa7\x90\x96\x41\x37\x23\x74\x51\xce\xd1\x72\xb4\xba\x4b\x8b\x9a\xa1\xe2\xfb\x90\x7a\x74\x87\xda\xb4\x40\xa1\x9c\x55\x43\x97\xa3\xb4\x3b\xf8\x27\x00\x00\xff\xff\x94\x73\x58\x8e\x8e\x08\x00\x00")

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

	info := bindataFileInfo{name: "locales/ru/LC_MESSAGES/config.mo", size: 2190, mode: os.FileMode(420), modTime: time.Unix(1541531470, 0)}
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
	"locales/ru/LC_MESSAGES/config.mo": localesRuLc_messagesConfigMo,
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
	"locales": &bintree{nil, map[string]*bintree{
		"ru": &bintree{nil, map[string]*bintree{
			"LC_MESSAGES": &bintree{nil, map[string]*bintree{
				"config.mo": &bintree{localesRuLc_messagesConfigMo, map[string]*bintree{}},
			}},
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
