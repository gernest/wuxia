// Code generated by go-bindata.
// sources:
// scripts/up.ql
// DO NOT EDIT!

package migration

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

var _scriptsUpQl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x92\x4f\x4f\x02\x31\x10\xc5\xef\xfd\x14\x73\x5c\x12\x12\xee\x12\x0f\x0b\x56\xd3\x44\x4a\x80\x92\x70\x23\x85\x8e\xa6\xb2\xb4\x6b\xa7\xa8\xfb\xed\xcd\xfe\xd3\xdd\xac\x11\x0f\xde\x3a\x2f\xfb\xf2\x7e\x6f\x66\x27\x13\x78\xb3\x37\x4f\x36\xc3\x58\xe4\x78\x4b\xaf\x19\x63\x33\xfe\x20\x24\xa8\x75\x2a\x37\xe9\x5c\x89\xa5\x9c\x32\x80\xf9\x9a\xa7\x8a\x83\x4a\x67\x8f\x1c\xc4\x3d\xc8\xa5\x02\xbe\x13\x1b\xb5\x01\x42\x22\xeb\x1d\x25\x0c\x00\xe0\x84\x05\x50\x0c\xd6\x3d\x8f\xab\xd9\xe8\xa8\xe1\x90\xf9\x43\x3d\x1e\x03\xea\x88\x66\xef\x1d\x44\x7b\xc6\x5a\xbc\xe4\x66\x28\xe2\x47\x6e\x03\x52\x2b\x8e\xae\x50\x44\x4d\xa7\x06\xe1\x72\xb1\xa6\xcf\xe0\x1d\xc2\xc1\xfb\xac\x89\x23\x0c\x7b\x6b\xc0\xba\x58\x0b\x79\xf0\x2f\x78\x8c\x3d\xad\x05\xd5\xf1\x07\xd0\x46\xbc\xc6\x54\x06\xb5\x4c\x84\xc1\xe9\x33\xf6\xb8\x72\x4d\xf4\xee\x83\xe9\xec\x07\xcf\xda\x66\xbd\x8f\xfe\x03\xa4\x29\xd8\x61\xe9\x75\x1d\x80\xfd\x39\xf3\x3b\x76\x2b\xc5\x6a\xcb\x41\xc8\x3b\xbe\x83\xed\x8a\xef\xab\xf2\xe0\x5d\xbd\x05\x48\xda\x0d\x8c\xab\x8a\x5d\xe2\x81\xb5\xba\x65\x69\xad\x1f\x49\x79\xd0\x5f\x0d\xed\x2f\x58\x7a\xbe\xde\xc9\x09\x8b\xd1\x94\xcd\x97\x8b\x85\x50\x53\xc6\xd8\x67\x00\x00\x00\xff\xff\xa0\x76\x4c\xa1\xed\x02\x00\x00")

func scriptsUpQlBytes() ([]byte, error) {
	return bindataRead(
		_scriptsUpQl,
		"scripts/up.ql",
	)
}

func scriptsUpQl() (*asset, error) {
	bytes, err := scriptsUpQlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/up.ql", size: 749, mode: os.FileMode(420), modTime: time.Unix(1469858211, 0)}
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
	"scripts/up.ql": scriptsUpQl,
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
	"scripts": &bintree{nil, map[string]*bintree{
		"up.ql": &bintree{scriptsUpQl, map[string]*bintree{}},
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

