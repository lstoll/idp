// Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// webauthn/webauthn.js
// webauthn/webauthn.tmpl.html
package webauthn

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

var _webauthnWebauthnJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x56\x5d\x8b\xe3\x36\x14\x7d\x76\x7e\xc5\x7d\xb3\xd3\x49\xed\x6c\x29\xa5\x64\x9a\x42\x9a\x5d\xe8\x07\xed\x14\x66\x4a\x1f\x4a\x29\xb2\x74\x6d\x6b\xd6\x91\x8c\x74\x3d\xd9\xb0\xe4\xbf\x17\xd9\xf2\x57\x66\xd2\x4d\xe8\xee\xdb\xce\x83\x61\xa4\x7b\xce\x3d\xf7\x23\x07\x25\x09\x14\x44\x95\x5d\x25\x49\x2e\xa9\xa8\xd3\x98\xeb\x5d\xf2\x56\xa3\x95\xf8\x6a\x99\xec\x31\x65\x35\x15\x2a\x49\x4b\x9d\x26\x3b\x66\x09\x4d\x7f\x18\x3f\xda\x59\x92\xc0\xaf\x5a\xc8\x4c\x72\x46\x52\x2b\xbb\x72\x27\x5f\xc0\xa6\x2c\xf5\x1e\x2a\x66\xad\x54\x39\x30\x48\xb5\x38\x00\x69\x30\x98\x4b\xc7\x31\x9b\xf1\x92\x59\x0b\x7f\x62\xba\x71\x54\xf0\x7e\x16\x24\x09\xbc\x46\xae\x05\xba\x78\x66\xf1\x9b\xaf\xc1\x92\x71\x78\xa9\x48\x03\x83\x3f\xa4\xa2\x6f\x37\xc6\xb0\x43\x3c\x0b\x2c\x31\x92\x1c\xfe\x11\x0d\xe4\x87\x3a\xcb\xd0\x44\x4f\xac\xac\x71\xee\xc8\x02\x83\x54\x1b\x35\xc6\x64\x46\xef\x22\x46\x3a\xf5\x61\x0b\xe0\xb0\xfe\x1e\x78\xcc\x0b\x66\xb6\x5a\xe0\x86\xa2\xe5\x7c\x7e\x3b\x0b\x8e\xb3\x46\xcd\x1b\xd5\xaa\x51\xd0\x10\xb4\x39\x3a\x31\x13\x85\x23\x3d\xa8\xfe\x53\x4f\x4a\x9a\x45\x0a\xf7\x23\x61\x3e\x2a\x36\x28\x6a\x8e\x51\x64\x17\x90\x1e\x08\xe7\x4e\x9c\x85\x1b\xb8\x6f\x53\x38\xf9\x5b\xaf\x34\x6a\x02\x16\x10\x86\x63\xbd\xdb\x02\xf9\x5b\x0b\xfb\x02\xa9\x40\x03\x54\x20\x38\x55\xb5\x85\x36\x39\x0a\xd8\x31\xe2\x05\xda\xf1\x5d\x2e\x9f\x50\x8d\x0a\xe0\x8e\xe5\xbe\xb9\x8b\xda\x90\x49\x01\x06\xad\x13\xe6\x4e\x02\x99\x41\x64\xd0\xc6\x9e\x69\xbd\x5e\xc3\x18\x31\xc6\xdc\xba\xff\x8f\xee\x43\x85\xd1\x7b\x70\x2d\x78\x63\x8c\x36\x23\x82\x07\x7c\x47\xae\x9c\xe0\xe8\x6b\xea\xb6\x25\x12\x8c\xd8\x44\x45\x86\xc4\x8b\x28\x1c\xf6\xb3\x0d\x35\xcd\x12\x26\x96\x98\xa1\x70\x01\xef\x67\x70\xf2\xb7\x43\x2a\xb4\x58\x41\xf8\xfb\xdd\xfd\x43\xb8\x78\x76\xef\xf6\x74\x05\x3f\xdf\xdf\xfd\x16\xb7\x93\x95\xd9\xa1\xcd\xde\xe8\x6f\xbe\x31\x15\xa8\xa2\x6e\x73\xe3\x49\xc3\xbe\x5a\x2e\xe7\xa3\x20\xdf\x2c\x57\xe2\xa3\xd5\x2a\x7a\xe1\xae\x6b\x94\x8d\xab\x3a\x2d\x25\xff\x05\x0f\x6e\x21\xcb\x12\x55\x8e\xb0\x86\x21\xcf\x64\xd3\xcf\x00\x9a\xf6\x9d\xd2\xd5\x16\x4d\x2c\xc5\xa5\x64\x3e\xdc\x53\x75\x33\x1e\xee\xf1\x1d\x2f\x6b\x81\x5b\x83\x02\x15\x49\x56\xf6\xe3\x0e\x32\x6d\x20\x7a\x62\x06\x24\xac\x61\x79\x0b\x12\xbe\x83\x0f\x81\x63\x27\x9c\x8a\x5b\x90\x37\x37\x3d\x51\xf0\x21\xd4\x5f\xf2\xef\x2b\x4a\x3a\x87\xf7\x35\xb6\x8b\xe9\xbf\xa7\x3b\xfb\x7c\x64\x8a\x3d\xc9\x9c\x91\x36\x31\x1f\x95\xc1\x0d\x32\x42\x17\x33\x9e\xf2\x10\x31\x1e\xf6\x05\x3b\x9c\x49\x25\x6d\xd1\x2c\x71\x23\xf1\x74\x73\x9b\xc3\x02\x99\x40\x63\x57\x7d\xdb\xc2\x0d\xe7\x58\x51\xb8\x82\x90\x55\x55\xe9\x5d\x39\x71\xcb\xd7\x61\x82\x70\xab\x15\xa1\xa2\x2f\x1f\x0e\x15\xbe\x18\xe9\x9b\xe2\x01\x2f\xfe\x26\xba\x84\x52\xac\x60\xa8\x31\x96\xa2\xcb\x62\xd8\xfe\x27\xb1\x1a\x8d\x67\x62\x8c\x23\x48\x13\x38\x5f\x0c\x73\xaf\xb4\xb2\x38\x94\x14\x30\x22\x6c\xcc\x49\xab\xbb\xf4\x11\x39\x5d\xc4\xea\x79\xe2\x67\xe8\x3e\x55\xc0\x4b\x89\x8a\x5e\x33\x62\xae\xb8\xab\x58\xa7\xd0\xb9\x67\xec\x3a\x16\xd0\xa1\xc2\x49\x5f\xdc\x81\xef\xaa\xcf\x7f\xbc\xd8\x51\x5e\xf5\x1e\x5f\xea\x5c\xaa\x0b\xcc\xb0\x89\xfb\xec\x82\xcf\xad\x8b\xb9\x67\xc9\x59\xe3\x2a\x91\xce\x1b\xd7\x29\xf4\x12\xdb\x3a\xc5\x5c\x67\x5a\x2f\xa3\x3f\xa6\x65\xe5\x48\xff\xc7\xaf\xda\x35\xfb\x6c\x54\x1f\xdf\x4f\x7a\x8b\x72\x8d\x76\xb1\xdc\xcd\xce\x5d\x5f\xe7\x7d\xa7\xe8\x81\xd8\xca\x5c\x31\xaa\x0d\x5e\x45\xd8\xa3\x06\x22\xf7\x5a\xf8\x91\x29\x51\x5e\xc7\x34\xc0\x7a\xaa\x4f\xe3\x9e\xdd\x8b\xfe\xf8\x6f\x00\x00\x00\xff\xff\x28\x15\x11\xe5\xea\x0c\x00\x00")

func webauthnWebauthnJsBytes() ([]byte, error) {
	return bindataRead(
		_webauthnWebauthnJs,
		"webauthn/webauthn.js",
	)
}

func webauthnWebauthnJs() (*asset, error) {
	bytes, err := webauthnWebauthnJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "webauthn/webauthn.js", size: 3306, mode: os.FileMode(420), modTime: time.Unix(1546748549, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _webauthnWebauthnTmplHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\x4d\x8f\xdb\x36\x13\xbe\xef\xaf\x98\xf0\x3d\xac\x8d\x37\x92\x1a\x34\x97\x26\x92\x81\x14\x69\xd0\x00\x5b\x24\x08\x12\xe4\x58\xd0\xe2\xd8\xe2\x86\x22\x05\x72\x64\xc7\x5d\xec\x7f\x2f\xa8\x6f\xd9\xf2\xd7\x36\x28\xaa\x93\x48\xcf\xc7\x33\x33\x7c\x86\x63\xc5\x19\xe5\x6a\x71\x73\x13\x67\xc8\xc5\xe2\x06\x00\x20\x26\x49\x0a\x17\x5f\x71\xc9\x4b\xca\x74\x1c\xd5\xeb\xfa\x37\x25\xf5\x37\xc8\x2c\xae\x12\x16\x45\xa9\xd0\x61\x5e\xca\xd4\xb9\x30\x35\x79\x94\x97\x32\xf8\x29\xfc\x25\x7c\xf9\x22\x4a\x9d\xf3\xcb\x30\x97\x3a\x4c\x9d\x63\x60\x51\x25\xcc\xd1\x4e\xa1\xcb\x10\x89\x01\xed\x0a\x4c\x18\xe1\x77\x8a\x2a\x81\xa8\x71\xe0\x52\x2b\x0b\x02\x67\xd3\xd3\x1e\xee\x7b\x07\xf7\x8e\x2d\xe2\xa8\x56\x9c\xb2\xb2\x6d\x22\x99\x14\xf4\x90\xea\x77\xff\x84\x99\x14\x08\x0f\xdd\xba\x7d\x84\x74\x85\xe2\xbb\x57\xa0\x8d\xc6\xd7\xa3\x9f\x1f\xbb\x55\x1c\x35\xd6\xe2\xa8\xce\xe6\x4d\xbc\x34\x62\xd7\x78\xf2\x5b\x68\x21\x55\xdc\xb9\x84\xf9\x48\x78\x51\x2c\xb9\x05\xff\x1a\xfc\xf5\x82\xf5\x28\xe2\x67\x41\x00\x6f\xea\x5f\x7f\xff\xfc\xc7\x1d\xac\x0d\x3a\xc8\xd0\x22\x04\x41\x63\x2e\xaa\xed\x35\x2b\x21\x37\x20\x45\xc2\x52\xa3\x09\x35\x05\x5b\xcb\x8b\x02\xed\xd0\xa6\x17\x19\x38\x6f\xbc\x07\x19\xca\x75\x46\x3e\x2f\x42\x6e\x8e\x8b\x7b\xc3\x5c\x6a\xb4\xc1\x4a\x95\x52\xb0\xc5\xcd\x28\x09\x15\xe2\xb7\xb8\xe2\xa5\x22\x20\xbb\x03\xae\x05\xf8\xa4\x43\xc1\xd7\x3d\xea\x91\x75\x8f\x57\x99\xb5\xd4\x29\xb7\x82\x0d\x9d\x59\xb3\x65\x8b\x83\x1a\x1c\x42\x52\x41\x2e\x82\x97\x30\x58\x98\xd5\xca\x21\x05\x2f\x27\xd4\xa7\x4c\x14\x5c\xa3\xaa\x0b\xe0\x4f\x62\x90\xa2\xa6\x51\xd6\x0e\x0c\x14\x8b\x3b\x0f\x19\xb6\x92\x32\xd8\x99\xd2\x02\x7e\x97\x8e\xa4\x5e\x57\xe1\xa2\x26\x99\x72\x32\x36\x8e\x8a\x13\x56\x56\xc6\xe6\x7d\xfc\xef\x8c\xcd\x4f\xf8\xac\x34\xa4\x2e\x4a\x1a\x90\x86\x81\xe6\x39\x26\xac\x74\x68\xfd\x1b\xeb\xcd\x7d\xe9\xb6\x9a\x48\xbd\xb7\xaa\x80\xd6\x28\x06\x85\xe2\x29\x66\x46\x09\xb4\x09\xeb\x64\x4f\x7a\xf7\xcf\x86\xab\x12\x13\xf6\xf0\x00\x61\xab\x04\x8f\x8f\x1d\x6f\x8f\x02\x5f\x96\x44\x46\x37\xc8\x5d\xb9\xcc\x25\x75\xc0\x96\xa4\x61\x49\x3a\x70\x65\x9a\xa2\x73\x90\x53\xf0\x33\xab\x13\x1c\x47\xb5\xe2\x89\x24\x46\x3e\xae\x93\xa5\x8a\x79\x95\x15\x8b\x6b\xe9\x08\xad\xef\x5d\xac\x69\x5e\xff\x63\x8b\x4f\xcd\x36\x7c\xc3\x5d\x1c\xf1\xc5\xd1\x92\xed\x51\xe3\xc4\x76\xb3\x75\x48\x8d\xce\x57\x55\xf9\xa3\x74\x68\x91\x4e\x31\x02\x7c\x6b\xfa\xcf\xd2\xa2\x55\xaf\xfa\xa7\x42\x2e\xd8\x28\xa0\x3b\xc3\x85\xd4\xeb\x3e\xe9\x52\xaf\xc3\x30\x84\x8f\x0a\xb9\x43\x20\x5e\xd4\x64\x1a\x71\x28\xbc\x90\x44\xad\x93\x1f\xc7\xa3\x1e\xf6\x5a\xea\xab\x78\x74\x96\x0f\x43\x00\x05\x77\x6e\x6b\x7c\xa9\x6b\x10\xfd\x7a\x08\xe2\x63\xb7\x3b\x85\xe3\x2c\x71\x9f\x88\xf3\x7a\xde\xb6\xa5\xfd\x67\xd4\x7d\x3a\xd9\x06\x22\xc3\xd7\x95\x31\xd4\xde\x93\x70\xea\x66\x3b\x75\xe0\x47\xb6\x5b\x8b\xa3\x51\x63\x30\xcf\xdc\xf3\x0d\xaf\x77\x07\x16\x14\x12\x74\x15\x45\xed\xd9\x00\x09\xac\xb8\x72\x83\x79\xc2\x0b\x55\x3d\xfc\x40\x62\x24\xb2\x85\x04\x34\x6e\xe1\x2b\x2e\xdf\xf8\xc1\x66\x36\x1f\x08\x08\x93\x96\x39\x6a\x0a\xd7\x48\xbf\x29\xf4\xaf\xbf\xee\xde\x8b\xd9\x98\x26\xf3\xd0\xe8\xba\xaa\xde\x47\xa9\x53\x92\x46\xc3\x0c\xe7\x7b\xc3\x0f\x86\x85\xc5\x0d\x6a\x6a\x2e\x77\xef\x6a\xf8\xbb\x5c\xc1\x6c\x2f\xac\x39\x58\xa4\xd2\xea\xb1\xe0\x61\xec\x64\xcb\xbd\x51\xea\x2c\xf4\xb6\x8d\xcc\xc3\xaa\x80\x77\xd2\x51\x68\x31\x37\x1b\x9c\x55\x9d\x87\xed\xa1\xdb\x86\xad\xe6\xec\x01\x7a\x8e\xbf\x3a\xea\xe9\x76\xc4\xfc\xdb\x79\x58\x5d\x7b\xcf\xa1\xa7\xe6\x05\xba\x2d\x61\x5b\x75\x78\x9c\x1f\x9c\xe7\xd0\xf7\xb9\x99\x45\x07\xc9\x02\xb8\x42\x4b\xb3\xdb\xcf\x99\x74\xe3\x0e\x08\x19\x77\xb0\x44\xd4\x5d\xfa\x50\xdc\xce\x27\xac\xa5\x9c\xd2\x6c\x86\xd6\x7a\x73\x87\xe3\xab\x7f\x52\xa3\x9d\x51\x18\xa2\xb5\xc6\x7a\xd1\x43\x33\xfe\x69\xb0\xbc\xe3\x52\xa1\x00\x32\x9d\xe7\x57\x70\x0b\xff\x07\xaf\xf7\xfa\x40\xf1\x68\x80\xb3\xf9\x71\x40\x67\xd9\x30\x7c\x9e\x74\x34\xb8\x10\xd3\xe7\xa2\xc6\x3c\xde\x8b\x22\xd8\x66\xa3\x4c\x43\xd3\xdf\x56\xa5\x52\xbb\xe7\xc0\x89\x30\x2f\x68\x58\x21\x69\x74\x67\xe3\xf1\x1a\x0e\x56\xa3\x88\xe7\x60\xaa\x64\xfa\xed\x34\x05\x8f\x1a\xeb\x87\xe7\x8b\x82\x3e\x0b\xea\xc0\xd4\x31\x6a\x0d\x43\xed\x80\x37\xe9\x79\x33\xca\xce\x6c\x3f\x18\xdf\x2f\x86\x1d\x6e\xba\x59\xec\xf5\xc0\xc3\x4e\x11\x45\x67\xd2\x72\x7d\xa7\xa8\xd4\x2e\x6c\x13\x7b\x83\xf6\xa5\x3c\xb7\xe8\xc2\x7b\xe7\xb3\x72\x46\x70\x2b\xb5\x30\x1e\x52\x9d\xc5\xd0\xcf\xac\x90\x54\xfa\x16\x85\xb4\x98\xd2\x9f\x64\xfe\xdd\x3e\x50\x85\xfc\xc3\x9b\xc0\xf4\x6d\x37\x25\xf9\xa4\x92\x5f\xdb\x01\x3e\x7f\x78\xfb\x01\x28\x43\x58\x49\xcd\x15\x38\xc2\xc2\x8f\xec\x40\xbe\x35\x6b\x44\xe1\x7c\x2a\x96\x08\x6d\x19\x3c\x6c\x32\x95\x4a\x7d\x90\x51\xc0\x97\x4f\x77\x61\xcf\x95\x0b\xba\x42\xff\x2f\xf0\x47\x5d\xcb\x47\xc8\xf8\x7a\x0a\x56\x14\xc1\xfb\x15\x6c\x11\x32\xbe\x41\xe0\x50\x58\x74\x88\x02\x05\xb4\x4c\x78\x5e\xff\x95\x97\x79\x8e\x42\x72\x42\xb5\xeb\x94\x3d\x9f\xaf\x64\xc9\xb3\x04\x18\xdb\x0b\xe8\x02\xbc\x50\x7f\x61\x69\x3e\xde\xc4\x51\xfd\x69\xe5\x26\x8e\xaa\x2f\x58\x7f\x07\x00\x00\xff\xff\x84\xd6\x6a\xf5\xc8\x12\x00\x00")

func webauthnWebauthnTmplHtmlBytes() ([]byte, error) {
	return bindataRead(
		_webauthnWebauthnTmplHtml,
		"webauthn/webauthn.tmpl.html",
	)
}

func webauthnWebauthnTmplHtml() (*asset, error) {
	bytes, err := webauthnWebauthnTmplHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "webauthn/webauthn.tmpl.html", size: 4808, mode: os.FileMode(420), modTime: time.Unix(1546796870, 0)}
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
	"webauthn/webauthn.js":        webauthnWebauthnJs,
	"webauthn/webauthn.tmpl.html": webauthnWebauthnTmplHtml,
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
	"webauthn": &bintree{nil, map[string]*bintree{
		"webauthn.js":        &bintree{webauthnWebauthnJs, map[string]*bintree{}},
		"webauthn.tmpl.html": &bintree{webauthnWebauthnTmplHtml, map[string]*bintree{}},
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
