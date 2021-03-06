// Code generated by go-bindata.
// sources:
// data/newsite/_theme/base.html
// data/newsite/_theme/index.html
// data/newsite/_theme/post.html
// data/newsite/config.yml
// DO NOT EDIT!

package newsite

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

var __themeBaseHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb2\xc9\x28\xc9\xcd\xb1\xe3\x52\x50\x50\x50\xb0\xc9\x48\x4d\x4c\x81\x30\xc1\xdc\x92\xcc\x92\x9c\x54\xbb\xea\x6a\x30\x5d\x5b\x6b\xa3\x0f\x11\x80\xa8\xd5\x47\x28\xb6\x49\xca\x4f\xa9\x44\xe8\xab\xae\xb6\x53\x48\xce\xcf\x2b\x49\xcd\x2b\xa9\xad\x85\x2a\x86\xa8\xb0\xd1\x87\x58\x06\x08\x00\x00\xff\xff\x59\xaf\x86\xe9\x74\x00\x00\x00")

func _themeBaseHtmlBytes() ([]byte, error) {
	return bindataRead(
		__themeBaseHtml,
		"_theme/base.html",
	)
}

func _themeBaseHtml() (*asset, error) {
	bytes, err := _themeBaseHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_theme/base.html", size: 116, mode: os.FileMode(420), modTime: time.Unix(1517852780, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __themeIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2c\x8d\x41\x0a\xc3\x30\x0c\x04\xef\x7e\x85\x48\x1f\x60\x72\x57\xf5\x17\x93\xaa\xc8\x34\x38\xa6\x11\xbd\x2c\xfb\xf7\x92\xd2\xf3\x0c\x33\xc0\xcd\xdb\x16\x32\x8f\x33\x4f\xb2\xe8\xa3\x7f\xac\x88\x88\x68\xac\xa6\x4d\xe2\xed\xcf\xfb\x02\xec\x7d\xbc\xc8\xc5\x80\xec\xb9\x3b\xa9\xb5\x99\xd6\x58\xff\xf6\x34\x60\x3b\x46\xfa\xc8\x8b\x4d\x2b\x5a\x7f\x2d\xa0\x5e\x07\xb2\x7c\x03\x00\x00\xff\xff\x4c\x0a\x76\x4f\x6c\x00\x00\x00")

func _themeIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		__themeIndexHtml,
		"_theme/index.html",
	)
}

func _themeIndexHtml() (*asset, error) {
	bytes, err := _themeIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_theme/index.html", size: 108, mode: os.FileMode(420), modTime: time.Unix(1517852780, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __themePostHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb2\xc9\x30\xb4\xab\xae\x2e\xc9\x2c\xc9\x49\xad\xad\xb5\xd1\xcf\x30\xb4\xe3\xb2\x29\xb0\xab\xae\x4e\xce\xcf\x2b\x49\xcd\x2b\x01\x89\x15\xd8\x71\x01\x02\x00\x00\xff\xff\x0f\x87\xf2\x9b\x26\x00\x00\x00")

func _themePostHtmlBytes() ([]byte, error) {
	return bindataRead(
		__themePostHtml,
		"_theme/post.html",
	)
}

func _themePostHtml() (*asset, error) {
	bytes, err := _themePostHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_theme/post.html", size: 38, mode: os.FileMode(420), modTime: time.Unix(1517852780, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4a\x4a\x2c\x4e\x0d\x0d\xf2\xb1\x52\xc8\x28\x29\x29\xb0\xd2\xd7\xcf\xc9\x4f\x4e\xcc\xc9\xc8\x2f\x2e\xb1\x32\x34\x31\x34\xd1\xe7\x2a\xc9\x2c\xc9\x49\xb5\x52\xf0\x4b\x2d\x57\x08\xca\xcf\xcf\x56\x70\xca\xc9\x4f\xe7\x02\x04\x00\x00\xff\xff\x90\x8a\x66\x42\x35\x00\x00\x00")

func configYmlBytes() ([]byte, error) {
	return bindataRead(
		_configYml,
		"config.yml",
	)
}

func configYml() (*asset, error) {
	bytes, err := configYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config.yml", size: 53, mode: os.FileMode(420), modTime: time.Unix(1517852780, 0)}
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
	"_theme/base.html": _themeBaseHtml,
	"_theme/index.html": _themeIndexHtml,
	"_theme/post.html": _themePostHtml,
	"config.yml": configYml,
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
	"_theme": &bintree{nil, map[string]*bintree{
		"base.html": &bintree{_themeBaseHtml, map[string]*bintree{}},
		"index.html": &bintree{_themeIndexHtml, map[string]*bintree{}},
		"post.html": &bintree{_themePostHtml, map[string]*bintree{}},
	}},
	"config.yml": &bintree{configYml, map[string]*bintree{}},
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

