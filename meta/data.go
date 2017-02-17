// Code generated by go-bindata.
// sources:
// meta/config.yaml
// DO NOT EDIT!

package meta

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

var _metaConfigYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x92\xc1\x6e\xdb\x30\x0c\x40\xef\xfa\x0a\x22\x39\xcf\x46\x76\xcc\xcd\xcb\x8a\xcd\x40\x07\x14\x48\xb7\x3b\x2d\xb1\x36\x51\x99\x12\x24\xca\x43\xff\x7e\x70\x8a\x7a\x2d\x9a\x06\x75\x4e\x86\xc5\xf7\x48\x4a\xe4\x16\x0e\x41\x1e\xb8\x2f\x09\x95\x83\x64\x78\x08\x09\x9a\xbb\xf6\x46\x5c\x0c\x2c\x9a\xab\xf9\xef\x48\x69\xa2\x04\x7f\xd9\x7b\xb3\x05\xb4\x0a\x41\x80\xc5\xf1\xc4\xae\xa0\x87\x91\x74\x08\x2e\x43\x89\x41\x40\x07\x02\xfb\x26\xa9\xd9\x42\x4c\x61\x62\x47\xae\x32\x46\x9f\x22\xed\x61\xd3\x8a\x52\xb2\x14\x35\xa4\xe7\x16\x36\x66\x42\xcf\xee\xa4\xbc\x0a\xee\x0d\x40\x48\x8e\xd2\x1e\x76\x06\x40\xc2\x9f\x85\x9a\x43\x00\x5f\x60\x53\x63\x8c\xd9\x06\x47\xd5\x40\xe8\x75\xa8\x7e\x9e\x3e\xf5\x51\x51\x4b\xde\x18\x2c\x3a\x90\x28\xdb\x0b\xc9\xbf\x1a\x80\x22\xaf\x48\x72\x2f\x8f\xf0\xe9\x3a\xef\xb0\x39\x5d\x35\xed\x3a\x52\xdc\x55\xcd\x9b\x2e\xea\xdb\xd0\xb3\xac\x77\x42\xd1\x95\xd2\x7d\x78\xa4\x73\x85\x1e\x4b\x47\x49\x48\x29\x2f\xe6\xc1\x97\xac\x94\xf2\x7c\xa3\xa4\x25\xbe\x4c\x66\x95\xdb\x4a\x56\x14\x4b\xdf\x9e\xda\xbb\x33\xea\x88\xec\x59\x7a\xcf\x59\x17\xf7\xd7\xf3\xd9\x2d\x67\xad\x8f\x24\xee\x66\x86\xae\x71\x4b\x97\x6d\xe2\x8e\xae\x70\x7f\x4b\xbe\x60\x0b\x8e\x94\x23\x5a\x5a\xdc\x7b\xc2\x31\xd7\x87\x44\xa8\x6b\x84\x1f\x74\x6e\x7e\x1f\xd1\x6d\x6e\x26\x64\x8f\x9d\x3f\xd5\x58\xb0\xc6\x8d\x2c\xcd\x27\x97\xd5\xf2\xff\xd5\xe8\x49\xf4\x42\xd7\xef\xd1\xef\xe4\x69\x46\xff\x05\x00\x00\xff\xff\x08\xe9\x11\x69\x2a\x04\x00\x00")

func metaConfigYamlBytes() ([]byte, error) {
	return bindataRead(
		_metaConfigYaml,
		"meta/config.yaml",
	)
}

func metaConfigYaml() (*asset, error) {
	bytes, err := metaConfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "meta/config.yaml", size: 1066, mode: os.FileMode(420), modTime: time.Unix(1453795200, 0)}
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
	"meta/config.yaml": metaConfigYaml,
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
	"meta": &bintree{nil, map[string]*bintree{
		"config.yaml": &bintree{metaConfigYaml, map[string]*bintree{}},
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

