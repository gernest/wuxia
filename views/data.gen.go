// Code generated by go-bindata.
// sources:
// templates/footer.html
// templates/header.html
// templates/home.html
// DO NOT EDIT!

package views

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

var _footerHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb2\x49\xcb\xcf\x2f\x49\x2d\xb2\xe3\xb2\xd1\x47\xb0\x92\xf2\x53\x2a\x41\x74\x46\x49\x6e\x8e\x1d\x17\x17\x20\x00\x00\xff\xff\xca\xa1\xc0\x16\x24\x00\x00\x00")

func footerHtmlBytes() ([]byte, error) {
	return bindataRead(
		_footerHtml,
		"footer.html",
	)
}

func footerHtml() (*asset, error) {
	bytes, err := footerHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "footer.html", size: 36, mode: os.FileMode(420), modTime: time.Unix(1471687536, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _headerHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x90\xb1\x52\xc4\x20\x10\x86\xeb\xe4\x29\x90\xfe\x42\xeb\x8c\x90\xe6\xd4\x56\x8b\x58\x58\x22\xb7\x27\x1b\x03\x89\xec\x5e\xc6\xcc\xcd\xbd\xbb\x43\x88\x73\x6a\x05\x2c\xdf\xff\xed\x82\xbe\xb9\x7f\xda\x77\xaf\xcf\x0f\xc2\x73\x18\xda\x5a\x97\xa5\xd2\x1e\xec\xa1\xad\xab\x4a\x07\x60\x2b\x3c\xf3\xb4\x83\xcf\x13\xce\x46\xee\xc7\xc8\x10\x79\xd7\x2d\x13\x48\xe1\xca\xc9\x48\x86\x2f\x56\x39\x7d\x27\x9c\xb7\x89\x80\xcd\x4b\xf7\xb8\xbb\x95\x42\x5d\x3d\xd1\x06\x30\xf2\x00\xe4\x12\x4e\x8c\x63\xfc\x25\x58\xc1\x4c\x0e\x18\x3f\x44\x82\xc1\x48\xf2\x63\x62\x77\x62\x81\x2e\xa3\x3e\xc1\xd1\x48\x45\x6c\x19\x9d\xc2\xf0\xae\x8e\x76\xce\x57\x0d\xba\xf1\xa7\xcf\x9a\xfe\x4b\x3a\x22\x45\x10\x6c\x64\x74\x4d\xc0\xd8\x38\x22\xb9\x75\xe0\x65\x00\xf2\x00\xbc\xe5\xb3\xa2\x4c\x27\x78\x99\x60\x7b\x57\x6f\x67\x5b\xaa\x52\x50\x72\x57\x75\xff\xcf\xdc\x93\x6c\xb5\x2a\xe8\x3a\x0e\x23\x0f\xd0\x9e\xcf\x4d\x97\x37\x97\x8b\x56\xa5\x52\x57\x5a\x6d\x7f\xac\xdf\xc6\xc3\xd2\xd6\xdf\x01\x00\x00\xff\xff\xd0\xd1\x3f\x75\x8b\x01\x00\x00")

func headerHtmlBytes() ([]byte, error) {
	return bindataRead(
		_headerHtml,
		"header.html",
	)
}

func headerHtml() (*asset, error) {
	bytes, err := headerHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "header.html", size: 395, mode: os.FileMode(420), modTime: time.Unix(1471776808, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _homeHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xae\x56\x28\x49\xcd\x2d\xc8\x49\x2c\x49\x55\x50\xca\x48\x4d\x4c\x49\x2d\xd2\xcb\x28\xc9\xcd\x51\x52\xd0\xab\xad\xe5\x52\xb0\xc9\x30\xb4\x53\xc8\x48\xcd\xc9\xc9\x57\x28\xcf\x2f\xca\x49\x51\xb0\xd1\xcf\x30\xb4\xe3\x42\xd1\x95\x96\x9f\x5f\x82\xa2\x0b\x10\x00\x00\xff\xff\x12\xe9\x13\x7e\x54\x00\x00\x00")

func homeHtmlBytes() ([]byte, error) {
	return bindataRead(
		_homeHtml,
		"home.html",
	)
}

func homeHtml() (*asset, error) {
	bytes, err := homeHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "home.html", size: 84, mode: os.FileMode(420), modTime: time.Unix(1471687728, 0)}
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
	"footer.html": footerHtml,
	"header.html": headerHtml,
	"home.html": homeHtml,
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
	"footer.html": &bintree{footerHtml, map[string]*bintree{}},
	"header.html": &bintree{headerHtml, map[string]*bintree{}},
	"home.html": &bintree{homeHtml, map[string]*bintree{}},
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

