// Code generated by go-bindata.
// sources:
// tmpl/base.html
// tmpl/edit.html
// tmpl/view.html
// public/css/style.css
// public/js/app.js
// DO NOT EDIT!

package main

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

var _tmplBaseHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\x90\xbd\x6e\xeb\x30\x0c\x85\x67\xeb\x29\x78\xb9\x5f\x73\xed\x20\x6b\xe9\xcf\xda\x02\x75\x86\x4e\x85\x20\xd3\xb5\x5d\xc5\x31\x44\xa2\x68\x20\xf8\xdd\x0b\x59\x43\x32\x91\x38\x87\xc4\xc7\xc3\x9c\x61\xe0\x71\x5e\x19\x70\x62\x3f\x70\x6a\x27\x3d\x47\x84\x7d\x37\xf6\xdf\xd3\xeb\x63\xff\xf1\xf6\x0c\x45\x72\xc6\x96\x02\xd1\xaf\x5f\x1d\xf2\xfa\x79\x7a\x47\x67\x1a\x5b\xb6\x9c\x69\x1a\xab\xb3\x46\x76\x90\x73\xdb\x1f\xed\xbe\x5b\xaa\x5a\x71\xcf\xac\x1e\xc2\xe4\x93\xb0\x76\x78\xea\x5f\xfe\x3f\x20\xd0\x61\xc5\x79\xfd\x86\xc4\xb1\x43\xd1\x6b\x64\x99\x98\x15\x61\x4a\x3c\x76\x48\xa2\x5e\xe7\x40\x41\x84\x0e\xb7\x0d\x22\x07\x97\x2a\x38\x67\xe0\x75\x28\xe7\x9a\xbb\x28\xe3\xe5\xa2\xf7\x51\x1a\x2b\x21\xcd\x9b\x82\x5e\x37\xee\x50\xf9\x57\x69\xf1\x3f\xbe\xaa\x08\x92\xc2\x8d\xb5\x08\xf9\x6d\x6b\x17\x41\x67\xa9\x4e\x38\x63\xa9\x3e\xe1\xc6\xfb\x0b\x00\x00\xff\xff\x01\x44\xbe\x69\x3a\x01\x00\x00")

func tmplBaseHtmlBytes() ([]byte, error) {
	return bindataRead(
		_tmplBaseHtml,
		"tmpl/base.html",
	)
}

func tmplBaseHtml() (*asset, error) {
	bytes, err := tmplBaseHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/base.html", size: 314, mode: os.FileMode(420), modTime: time.Unix(1483482480, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplEditHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x44\x8e\xc1\x6a\x04\x21\x10\x44\xcf\xf1\x2b\x9a\x86\x5c\xd7\xdd\x9c\x72\x50\x0f\x81\x9c\x13\xc8\xfe\x80\xbb\xf6\x66\x1a\x46\x1d\xb4\xc7\x64\x90\xf9\xf7\xe0\x61\xc8\xb1\xa0\xde\xab\x32\xd3\xc5\xbd\x07\x16\x4e\xdf\xd0\xfb\xe9\xca\x32\xd3\xbe\x1b\x3d\x5d\x9c\x52\xe6\x91\x4b\x04\x7f\x17\xce\xc9\xa2\xae\xbe\x91\xfe\x2f\x21\x44\x92\x29\x07\x8b\x9f\x1f\x5f\x57\x74\xea\xc9\x04\x6e\xce\x08\xfd\x8a\x2f\xe4\x21\xf9\x48\x16\x6f\x39\x6c\x08\x25\xff\x54\x8b\x2f\x67\x84\x7b\x9e\xab\xc5\xd7\x33\xba\xde\x97\xc2\x49\x1e\x80\xcf\x15\xe1\xf4\x96\xc3\x36\xa6\x0f\xde\x19\x3d\x7c\x87\x96\xd3\xb2\x0a\xc8\xb6\x90\xc5\xba\xde\x22\x0b\x42\xf3\xf3\x3a\xa2\x6f\x84\xa0\x0f\xc0\xe8\xf1\xdb\xa9\xbf\x00\x00\x00\xff\xff\x88\xa1\xe8\x59\xdc\x00\x00\x00")

func tmplEditHtmlBytes() ([]byte, error) {
	return bindataRead(
		_tmplEditHtml,
		"tmpl/edit.html",
	)
}

func tmplEditHtml() (*asset, error) {
	bytes, err := tmplEditHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/edit.html", size: 220, mode: os.FileMode(420), modTime: time.Unix(1483038164, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplViewHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x8e\xb1\x0a\x42\x31\x0c\x45\xf7\x7e\x45\x08\xb8\xbe\xf2\xf6\x98\xc1\x6f\x70\x13\x87\x42\x53\x5a\xe8\xb3\xa5\x06\x41\x42\xff\x5d\x3a\xa9\xeb\xbd\x87\xc3\x31\x03\x95\xa3\xd7\xa0\x02\x98\x25\x44\x19\x5b\xd6\xa3\x22\xcc\xe9\x28\xef\x6c\xb6\x5d\x8b\x56\x99\x93\x7c\xde\xd9\x39\xea\x7c\xa3\x00\x79\x48\x3a\xa3\x97\x58\xd4\x7f\x19\xe4\x35\x90\x0f\x7c\x27\xdf\x17\x1d\xcb\x8b\xcd\xfa\x28\x0f\x4d\x80\xa7\x27\xc2\x76\x69\xf1\xbd\x74\xeb\x72\xee\xaf\x20\xb5\xa6\xbf\x05\x9f\x00\x00\x00\xff\xff\x0e\xcc\x7c\x2b\x9f\x00\x00\x00")

func tmplViewHtmlBytes() ([]byte, error) {
	return bindataRead(
		_tmplViewHtml,
		"tmpl/view.html",
	)
}

func tmplViewHtml() (*asset, error) {
	bytes, err := tmplViewHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/view.html", size: 159, mode: os.FileMode(420), modTime: time.Unix(1483480730, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _publicCssStyleCss = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\x52\xa8\xe6\xe2\x2c\x48\x4c\x49\xc9\xcc\x4b\x57\x30\xb0\xe6\xe2\x4c\xca\x2f\x4a\x49\x2d\x52\x00\xb1\x6b\xb9\x00\x01\x00\x00\xff\xff\x75\x90\xf1\xc5\x1e\x00\x00\x00")

func publicCssStyleCssBytes() ([]byte, error) {
	return bindataRead(
		_publicCssStyleCss,
		"public/css/style.css",
	)
}

func publicCssStyleCss() (*asset, error) {
	bytes, err := publicCssStyleCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "public/css/style.css", size: 30, mode: os.FileMode(420), modTime: time.Unix(1483462631, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _publicJsAppJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func publicJsAppJsBytes() ([]byte, error) {
	return bindataRead(
		_publicJsAppJs,
		"public/js/app.js",
	)
}

func publicJsAppJs() (*asset, error) {
	bytes, err := publicJsAppJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "public/js/app.js", size: 0, mode: os.FileMode(420), modTime: time.Unix(1483482490, 0)}
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
	"tmpl/base.html": tmplBaseHtml,
	"tmpl/edit.html": tmplEditHtml,
	"tmpl/view.html": tmplViewHtml,
	"public/css/style.css": publicCssStyleCss,
	"public/js/app.js": publicJsAppJs,
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
	"public": &bintree{nil, map[string]*bintree{
		"css": &bintree{nil, map[string]*bintree{
			"style.css": &bintree{publicCssStyleCss, map[string]*bintree{}},
		}},
		"js": &bintree{nil, map[string]*bintree{
			"app.js": &bintree{publicJsAppJs, map[string]*bintree{}},
		}},
	}},
	"tmpl": &bintree{nil, map[string]*bintree{
		"base.html": &bintree{tmplBaseHtml, map[string]*bintree{}},
		"edit.html": &bintree{tmplEditHtml, map[string]*bintree{}},
		"view.html": &bintree{tmplViewHtml, map[string]*bintree{}},
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

