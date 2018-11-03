// Code generated by go-bindata.
// sources:
// templates/views/send.html
// locales/ru/LC_MESSAGES/config.mo
// locales/ru/LC_MESSAGES/mail.mo
// locales/ru/LC_MESSAGES/send.mo
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

var _templatesViewsSendHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x57\x5f\x6f\xe3\x36\x0c\x7f\xbf\x4f\x41\xe8\x61\xd8\x86\x39\x46\xd3\xe1\x30\x60\x76\x30\x60\x2f\x7b\xb8\x62\xc3\xee\xf6\x5c\x28\x16\xdd\xa8\xd0\x1f\x4f\xa4\x73\xd9\x19\xfd\xee\x83\x2c\x3b\xf1\x35\x8e\xef\xba\xa1\x40\x5e\x5a\x51\xa2\x48\xea\xc7\x1f\x19\xba\xeb\x40\x61\xad\x1d\x82\xa8\xbc\x63\x74\x2c\xe0\xe9\xe9\x4d\xa1\xf4\x1e\x2a\x23\x89\x4a\x11\xfc\x47\xb1\x79\x03\x00\x30\xdd\xad\xbc\xc9\xac\xca\x6e\xd6\x10\x57\x64\xc7\xd5\x81\xb2\x9b\xf5\xa0\xff\xfc\xce\xe1\xbe\x91\x0e\xcd\xe4\xf4\x5c\x83\x35\x1b\x7c\xa6\xd1\x6b\xed\xd6\x9b\xae\x03\x7d\xf3\x93\x03\x71\x27\xb5\x11\xb0\x82\xa7\xa7\x22\xdf\xad\x67\x94\xa7\x81\x1a\x94\xa1\xd6\x07\xb1\x29\x72\xa5\xf7\xcf\x7c\xcf\x6c\x7d\x16\xce\x08\xca\xb9\x8f\x18\x4c\x0d\x2b\x0c\xc1\x87\x08\xd9\x52\x0c\xd2\x60\x60\xe8\xff\x66\x4a\xba\x07\x0c\xa3\xa0\xc9\x6a\x22\xb9\x9d\x7d\x73\x6f\x66\xdb\x32\x7b\x07\xfc\x4f\x83\xa5\x48\x82\x38\xbd\xcd\x13\x0a\x50\x92\xe5\x68\x6a\x70\x26\x40\x06\x2d\xb3\x9d\x56\x0a\x5d\x29\x38\xb4\x28\x36\xdf\xb0\xb6\x48\x3f\x17\x79\x32\x33\xef\xb0\xeb\x96\x1e\x75\x8e\xd7\x70\x07\x9d\x8a\x17\x2e\xc1\x64\x91\x48\x3e\xe0\x4b\x80\xd2\xae\xf6\xd7\x0d\xd3\xd2\xa3\x5e\x0c\x54\x51\xfb\x60\xc7\x88\xe3\x3a\xdb\xf9\xa0\x3f\x79\xc7\xd2\x40\x2f\x1b\xb9\x45\x93\x19\xac\x59\x40\xf0\x06\x93\x9a\x00\x8b\xbc\xf3\xaa\x14\x8d\x27\x16\xa0\x55\x29\x08\x9d\xca\x6c\x5f\x24\xb2\x62\xed\x5d\x29\xf2\x28\xe6\xf1\x20\x17\xe0\xfc\x5e\x1a\xad\x24\xe3\x05\x34\x27\x29\xd1\x8c\x36\xf9\x7f\x08\xbe\x6d\x2e\xe0\xdf\xdf\xea\x03\x8c\xba\xa5\x60\x7f\x42\xdf\x3b\x0e\xde\xa4\xf0\x61\xe8\x1b\xb7\x63\xdb\xb8\x9d\xed\x1a\x17\x10\x4f\xd5\xff\xe1\xf7\x54\xfb\x50\x50\x23\xdd\xb1\x4d\xe1\xdf\xad\x0e\xa8\xc4\xe6\xfb\x22\x8f\x07\x0b\x81\xe6\x7d\x2c\x0b\x0a\xe7\x8d\xee\xed\x18\xf0\xdb\xaf\x0e\xb8\xd0\xae\x69\x79\xa0\x24\xe3\x81\xc5\x67\xe9\x1d\x70\x49\x19\x8b\x78\x39\x69\x31\xad\x1a\x23\x2b\xdc\x79\xa3\x30\x94\x02\x0f\xd2\x36\x06\x6f\x7e\x51\xde\x4a\xed\x56\x95\xb7\x3f\xc0\xb0\xb9\x9e\x6c\x0a\xd8\x4b\xd3\x62\x29\x04\x8c\x58\x4c\x51\x59\x40\x63\x96\xaa\x5f\x3a\xfa\xdf\x14\xa1\x76\xfb\x88\x15\xbf\x1e\x4f\xde\x8f\x0e\x56\x73\x15\x7a\x7a\xe3\xd5\x91\xe1\x88\x4c\x62\xc4\x51\x1c\xf3\x7b\x7d\xb9\x8c\xcf\x7a\xbd\x44\xfe\x9a\x7e\x84\x21\x79\xb9\x8a\x6c\x12\x1a\xac\x78\x2e\x85\x90\x8e\xd6\xb1\x47\x7f\xa4\x52\xfc\x78\x2c\xec\x3e\xfa\xbe\xd8\xe3\x6a\xd9\x41\xef\xc4\x37\xb1\x77\x8f\x69\x6f\x8c\xd4\x4e\x9c\x86\xa0\x3f\xa2\x0c\x89\x4a\x69\x14\x4a\xfa\x2f\x36\xbc\x63\x6b\x26\x76\x7f\xfb\x70\xf7\xee\x25\x16\x8b\x3c\xbd\xf8\xfa\x58\x39\xfc\x3e\xbf\x1e\x31\xef\x46\x07\x57\xc1\xc9\x48\x05\x19\x50\xce\xb2\x32\x20\xe9\x4f\x71\x80\xba\x1f\xd5\xce\x08\x7a\x84\x2b\x72\x74\x14\x36\x45\x3e\x5e\x78\xcd\x04\x1b\x77\x4f\xde\x68\x35\x3b\xa5\xcf\x5d\xf8\x3a\x32\x5c\x02\xd6\xaa\xcc\xd7\x35\x21\x67\xb7\x5f\x42\x75\x18\x2c\xc7\xb1\x4a\x0c\x2d\x9c\xda\xad\xd5\xa7\x26\xbe\x65\x07\x5b\x76\x19\xb5\x55\x85\x44\x93\x72\x7a\xdf\x5f\x4a\xe5\xb4\x34\x50\xfe\x37\x1c\x8b\x3c\x02\xb1\xf8\x51\x33\x11\x87\xe5\xf0\x6f\x32\x8c\x4e\x3e\x03\x1f\xa9\xff\x02\xec\x3a\x20\x96\xac\xab\xd8\x0e\xe0\xdb\xb4\xfe\xeb\xcf\x77\x20\x72\x25\x69\xb7\xf5\x32\xa8\x5c\x12\x21\x53\xbe\x47\xa7\x7c\xa0\x7c\x18\x2a\x7d\xc8\x1f\x27\xc2\xca\x6a\xb7\x8a\x56\x6b\x69\x08\xbf\x1b\x8c\x27\xcf\xff\x06\x00\x00\xff\xff\x26\xe1\x89\xf3\x7f\x0e\x00\x00")

func templatesViewsSendHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesViewsSendHtml,
		"templates/views/send.html",
	)
}

func templatesViewsSendHtml() (*asset, error) {
	bytes, err := templatesViewsSendHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/views/send.html", size: 3711, mode: os.FileMode(420), modTime: time.Unix(1540767205, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesRuLc_messagesConfigMo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x52\x4d\x6b\x13\x5b\x18\x7e\xa6\xd3\xfb\x95\x7b\x2f\xf7\x72\xe1\x42\x17\x8a\xa7\x0b\x8b\x45\xa6\xce\x24\x16\xca\xb4\xd3\x8a\xb5\x45\xb1\xa1\xa1\xa6\x22\xb8\x3a\x34\xa7\xe9\x60\x72\x4e\x38\x67\x62\x15\xba\x48\xb3\x71\x53\xa8\x1b\xc1\x5d\x57\x2e\x85\x50\x2c\x16\xdb\xa6\x7f\xe1\x3d\x7f\x40\xfc\x09\xba\x71\x2d\x93\x49\x52\x44\x7c\x37\xf3\x3c\xef\xf3\x31\xef\xc0\x7c\xfc\x6f\xf4\x15\x00\xe4\x00\x5c\x02\x50\x03\xf0\x37\x80\x53\x64\xf3\x19\xc0\x28\x80\x2f\x00\xfe\x04\xf0\xb5\xaf\x8f\x38\xc0\x18\x80\x7f\x1d\xe0\x7f\x00\x93\x0e\xf0\x0f\x80\x05\x27\xf3\xad\xf6\x9f\x8f\xfb\xfb\xaa\x03\x6c\x39\xc0\xb6\x03\x8c\x03\xf8\x34\x02\xfc\x9e\xf6\xb8\xc0\x5f\x69\xb7\x9b\xed\xc7\x5c\xe0\x32\x80\xb0\xbf\x7f\xe4\x66\x3e\xe1\x66\x77\x48\x17\xb8\x02\xc0\xb8\x80\x03\xc0\x05\xf0\x2b\x80\x5f\xfa\x77\xfe\x81\xcc\x9f\x7e\xcf\x08\x2e\xe6\xb7\x01\x28\xf2\xb8\x86\x0d\x25\x37\xe3\xea\xe8\x5d\x65\x92\x01\x5e\x11\x49\x22\xf4\x80\xa5\x2e\xb6\xa9\x55\x9d\xf1\x4a\x45\x0b\x63\x7e\x14\x24\xaf\x8b\xc1\xb6\xc4\x8d\xd9\x56\xba\x32\xe4\x4a\x0f\x8b\x1f\x14\xcb\xa5\x01\x5e\x37\x42\xf7\x72\x6b\xa2\xa1\x74\xe2\x15\x4d\x35\xae\x78\xb7\x9b\x55\xe3\x95\x55\xc8\x2a\xe2\xe9\xad\x27\xf1\x16\xaf\xab\x29\xdd\xcc\x95\x56\xcb\xde\xa2\x16\x3c\x89\x95\xf4\xee\xf0\x44\x84\x2c\xef\x07\x33\x9e\x5f\xf0\xf2\x05\x96\x2f\x84\xd3\xd3\xd7\xfd\x82\xef\xe7\x56\xb8\x49\xbc\xb2\xe6\xd2\xd4\x78\xa2\x74\xc8\xee\xf7\x3a\x58\xb1\xa9\x79\x5d\x55\x14\x9b\xfb\xae\x78\x3e\xb7\xc2\x65\xb5\xc9\xab\xc2\x2b\x0b\x5e\x0f\xd9\x90\x87\x6c\xad\x69\x4c\xcc\x65\xae\x78\xaf\xb8\xe4\x3d\x14\xda\xc4\x4a\x86\x2c\x98\xf2\x73\x8b\x4a\x26\x42\x26\x5e\xf9\x79\x43\x84\x2c\x11\xcf\x92\x1b\x8d\x1a\x8f\xe5\x2c\xdb\xd8\xe2\xda\x88\x24\x5a\x2f\x2f\x7b\x33\x17\xbe\xf4\x9e\x4d\xa1\xbd\x25\xb9\xa1\x2a\xb1\xac\x86\x2c\x57\xaa\x35\x35\xaf\x79\xcb\x4a\xd7\x4d\xc8\x64\xa3\x47\x4d\x54\x98\x65\x19\x8c\xe4\xd5\xc0\x8f\xa2\x80\x4d\x4c\xb0\x14\xfa\xe3\x51\x10\xb0\x05\xe6\xb3\xb0\xc7\xe7\xa3\xfc\x40\x9a\x8b\x6e\xa6\xf0\x5a\xcf\x36\x17\xf8\x6c\x67\x27\x8b\xcc\x47\x79\x7f\x92\x2d\xb0\x80\x85\x2c\x3f\x0b\x7a\x4b\x27\x74\x44\x1f\x6c\xdb\xb6\xa8\x4b\x67\x74\x46\x1d\xbb\xcf\xe8\x9c\xba\xf6\x85\x6d\x53\x07\xf4\x86\xba\x76\xd7\xb6\x41\x07\x74\x6c\x77\xed\x1e\x9d\x52\x17\xf4\x92\xde\xd9\x16\x1d\xd9\x5d\x46\x5d\xdb\xa6\x73\xdb\xa2\x0e\x1d\xd2\xb1\x6d\xd3\x11\x9d\xd8\x7d\xd0\x6b\x3a\x4d\x9b\x7e\xa2\x1e\x50\xa7\xf7\xc6\x13\xbb\x97\x92\xae\x6d\xd9\x36\x7a\xff\xc2\x20\x78\x9e\xa9\xf4\x9e\xba\x74\x48\x9d\x61\xf4\x5b\x00\x00\x00\xff\xff\xca\x69\x7f\x82\x93\x03\x00\x00")

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

	info := bindataFileInfo{name: "locales/ru/LC_MESSAGES/config.mo", size: 915, mode: os.FileMode(420), modTime: time.Unix(1541260570, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesRuLc_messagesMailMo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\xbd\x6e\x14\x31\x14\x85\xcf\x46\xa1\x99\x92\x9a\xe2\xa6\x20\x02\x81\x83\x67\x86\x48\x91\x77\xbd\x41\x84\x44\x42\x64\xc4\x12\x0d\xe9\xad\xac\x99\x1d\x31\x6b\xaf\x6c\x0f\x02\x29\x05\x4a\x43\x49\xc5\x63\x20\xa5\xa1\x01\x25\xfb\x0a\x9e\x17\xe0\x59\xd0\xce\xf2\x97\x53\x9d\x4f\xfe\x7c\x74\x7f\xde\xde\xfc\x02\x00\x1b\x00\xee\x00\x78\x08\xe0\x16\x80\x11\xd6\x99\x00\xd8\x04\xf0\x0a\xc0\x6c\x00\x9c\x02\xd8\x02\xb0\x1c\x00\x83\xdf\xce\x06\xfe\x4b\xa1\xea\x06\x27\x7a\x61\x5d\x60\x85\xaf\xea\x29\x7b\xda\x56\x9e\x95\x56\xd0\x54\xbf\x7b\xf2\xb6\x9e\xa9\xb9\xdd\x71\x6d\x32\x79\x59\xb2\x03\xa7\x55\xa8\xad\x61\xcf\x54\xd0\x82\x32\x9e\xee\x31\x9e\xb3\x2c\xa7\x2c\x17\xbb\xbb\x0f\x78\xce\x79\x72\xac\x7c\x60\xa5\x53\xc6\x37\x2a\x58\x27\xe8\x45\xbf\x41\x45\xeb\xd4\xdc\x4e\x2d\x8d\x6e\x0c\x8f\x93\x63\x65\xaa\x56\x55\x9a\x95\x5a\xcd\x05\xfd\x65\x41\x27\xad\xf7\xb5\x32\x49\xf1\xbc\x38\x64\xa7\xda\xf9\xda\x1a\x41\xe9\x0e\x4f\x0e\xac\x09\xda\x04\x56\x7e\x58\x68\x41\x41\xbf\x0f\x8f\x16\x8d\xaa\xcd\x90\xce\x66\xca\x79\x1d\xe4\xeb\xf2\x88\xed\xfd\xf3\x56\xf7\xbc\xd1\x8e\x1d\x9a\x33\x3b\xad\x4d\x25\x28\x99\x34\xad\x53\x0d\x3b\xb2\x6e\xee\x05\x99\x45\x8f\x5e\xe6\x43\x5a\x57\x69\xee\xa6\x5c\xca\x94\xb6\xb7\x69\x55\xf9\x96\x4c\x53\xda\x27\x4e\xa2\xe7\xb1\xcc\xfe\x3c\x8d\xe4\xe3\x55\xbd\xd7\x6b\xa3\x94\xd3\xf9\xf9\xfa\xcb\x58\x66\xfc\x3e\xed\x53\x4a\x82\xb2\x21\xe2\xd7\xf8\x23\x7e\x8b\xdf\xbb\x8b\xee\x63\xbc\x8e\x57\xf1\x2a\x5e\x76\x9f\x29\x2e\xe3\x75\xf7\xa9\xbb\x88\x97\xf8\x15\x00\x00\xff\xff\x8c\xc9\x8f\x42\xe1\x01\x00\x00")

func localesRuLc_messagesMailMoBytes() ([]byte, error) {
	return bindataRead(
		_localesRuLc_messagesMailMo,
		"locales/ru/LC_MESSAGES/mail.mo",
	)
}

func localesRuLc_messagesMailMo() (*asset, error) {
	bytes, err := localesRuLc_messagesMailMoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/ru/LC_MESSAGES/mail.mo", size: 481, mode: os.FileMode(420), modTime: time.Unix(1541260570, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesRuLc_messagesSendMo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x51\xcf\x4f\xd4\x5a\x14\xfe\x66\xe0\xbd\x07\x93\xf7\x92\x17\xa2\x31\x46\x17\xc7\x85\x04\x43\x8a\xed\x8c\x18\x52\x28\x18\x11\xa2\x91\x09\x04\xaa\xfb\xeb\xcc\x75\xa8\xce\xdc\x4e\x7a\x5b\x23\x09\x0b\x18\x36\x9a\xb8\xd4\x8d\x0b\x8d\xae\x59\x80\xc9\x24\x88\x50\x36\xfe\x01\xa7\xff\x80\xd1\x7f\xc0\x8d\x4b\x37\xa6\xed\x20\x9a\x78\x16\xf7\x7c\xdf\x39\xdf\xf9\x95\xfb\x69\xa8\xff\x05\x00\x0c\x02\x38\x0f\xa0\x0e\xe0\x3f\x00\xdb\xc8\xed\x33\x80\x7f\x01\x7c\x01\xd0\x0f\xe0\x2b\x80\x7f\x00\x7c\x03\x70\x0a\xc0\x77\x00\x25\x00\xff\x17\xf2\xfc\xb9\x42\x9e\xa7\x02\x50\x04\x30\x52\x00\x56\x0b\xc0\x68\x01\x38\x03\xe0\x79\x31\xd7\x6d\x17\x81\x21\x00\xbb\x45\xe0\x2a\x80\x8f\x45\xe0\x6c\xda\xaf\x17\x3f\xdd\x07\x0c\xa4\xf5\x3d\x6f\xf5\x01\x05\x00\x7f\xf7\xe6\xff\x05\xa0\x0f\xf9\x8c\xc1\xde\xae\xfd\x38\xb1\x81\x63\x30\xeb\xab\x50\xaa\x90\xc2\xb5\xb6\xc4\x4d\xb7\xba\x80\xaa\xd4\x5a\x34\xe4\xb1\x27\x2d\x55\x9d\x74\x54\xab\x49\xad\xb1\xd4\x14\x9e\xa2\x50\x3e\x0e\xb1\x22\x55\x1d\x2b\xd1\xbd\x07\xb2\x16\xc2\x5d\xc4\xb2\x6c\xfb\x41\x68\x54\x75\xc3\xab\x1b\xd7\xa3\x86\x36\x5c\xdf\xa6\xba\x7c\x74\xed\xa1\xb7\x2a\x5a\xfe\x58\x10\x95\x96\x16\x5d\x63\x36\x90\x22\xf4\x7c\x65\xdc\x10\xa1\xb4\xa9\x6c\x5a\x13\x86\x59\x31\xca\x15\x2a\x57\xec\xf1\xf1\x51\xb3\x62\x9a\xa5\x05\xa1\x43\xc3\x0d\x84\xd2\x4d\x11\xfa\x81\x4d\xb7\xb3\x1e\x54\x8d\x02\xd1\xf2\xeb\x3e\x4d\xfd\xd6\x78\xba\xb4\x20\x54\x23\x12\x0d\x69\xb8\x52\xb4\x6c\xfa\xc9\x6d\x5a\x8e\xb4\xf6\x84\x2a\x55\x6f\x55\xe7\x8c\xbb\x32\xd0\x9e\xaf\x6c\xb2\xc6\xcc\x52\xef\x76\xc3\x5d\x6b\x4b\x3b\x3b\xea\x72\x3b\xbd\x6f\x92\x6a\xab\x22\xd0\x32\x74\xee\xb8\xf3\xc6\xc4\x89\x2e\xdd\xe7\xbe\x0c\x8c\x39\x55\xf3\xeb\x9e\x6a\xd8\x54\x5a\x6a\x46\x81\x68\x1a\xf3\x7e\xd0\xd2\x36\xa9\x76\x46\xb5\x53\x99\xa4\x1c\x3a\xea\xa2\x65\x3a\x8e\x45\xc3\xc3\x94\x42\xf3\x82\x63\x59\x34\x43\x26\xd9\x19\x9f\x76\xca\xc7\xa9\x29\xe7\x4a\x0a\x47\x32\xd9\x94\x65\xd2\xfa\x7a\x5e\x32\xed\x94\xcd\x4b\x34\x43\x16\xd9\x54\x9e\x04\xbf\xe5\x3d\x3e\x22\xde\xe7\x98\x0f\x93\x0e\x77\xb3\x77\x27\xff\x3d\x7e\xc3\x31\xc7\xbc\x9b\x3c\x4d\xe3\xbc\xc7\xdd\x3f\x84\x28\xd9\x4a\x36\xf9\x88\xbb\xc9\x13\x3e\xe4\x98\x38\x4e\x3a\x7c\x94\x6c\xf0\x0e\xbf\xe3\x0f\x99\x2a\x06\xbf\x4e\x36\x38\x4e\x36\x93\x0e\xc7\xfc\x9e\xb2\x41\xfb\x29\x05\xbf\xfa\x45\xbe\x97\x74\x92\x67\xe9\x4e\x5d\x3e\xe0\x1d\xf0\x4b\x8e\xf9\x20\xd9\xc2\x8f\x00\x00\x00\xff\xff\x38\x92\x03\x7b\x3a\x03\x00\x00")

func localesRuLc_messagesSendMoBytes() ([]byte, error) {
	return bindataRead(
		_localesRuLc_messagesSendMo,
		"locales/ru/LC_MESSAGES/send.mo",
	)
}

func localesRuLc_messagesSendMo() (*asset, error) {
	bytes, err := localesRuLc_messagesSendMoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/ru/LC_MESSAGES/send.mo", size: 826, mode: os.FileMode(420), modTime: time.Unix(1541260570, 0)}
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
	"templates/views/send.html":        templatesViewsSendHtml,
	"locales/ru/LC_MESSAGES/config.mo": localesRuLc_messagesConfigMo,
	"locales/ru/LC_MESSAGES/mail.mo":   localesRuLc_messagesMailMo,
	"locales/ru/LC_MESSAGES/send.mo":   localesRuLc_messagesSendMo,
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
				"mail.mo":   &bintree{localesRuLc_messagesMailMo, map[string]*bintree{}},
				"send.mo":   &bintree{localesRuLc_messagesSendMo, map[string]*bintree{}},
			}},
		}},
	}},
	"templates": &bintree{nil, map[string]*bintree{
		"views": &bintree{nil, map[string]*bintree{
			"send.html": &bintree{templatesViewsSendHtml, map[string]*bintree{}},
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
