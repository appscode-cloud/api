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

var _metaConfigYaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x92\x41\x8f\x9b\x30\x10\x85\xef\xfc\x8a\x11\x7b\x2e\xa8\x3d\xe6\x86\xd2\x55\x1b\x69\x5b\x55\xa5\xed\xdd\xd8\x93\x30\x8a\x19\x5b\xf6\x98\x2a\xff\xbe\x26\x11\x29\x11\x12\xd9\x13\x03\xef\x9b\xe7\xc7\x78\x5e\x60\xef\xf8\x48\xa7\x14\x94\x90\xe3\x08\x47\x17\xa0\xf9\x71\x78\x65\xe3\x1d\xb1\xc4\x6a\x7a\x6b\x31\x8c\x18\xe0\x2f\x59\x5b\xbc\x80\xd2\x02\x8e\x81\xd8\xd0\x48\x26\x29\x0b\x03\x4a\xef\x4c\x84\xe4\xf3\x77\xe9\x11\xf4\x83\x69\xee\xf1\xc1\x65\x16\x4d\x55\x14\x72\xf1\xb8\x83\xf2\xc0\x82\x41\xa3\x17\x17\x6e\x11\xca\x42\xa5\xdc\xcb\x42\xfa\xda\xb6\x00\x76\x05\x80\x0b\x06\xc3\x0e\x3e\xe6\x32\xf1\x82\x44\x33\x67\x9d\x28\x80\x0f\x50\xd6\x3d\x2a\x2b\x7d\xf5\xf5\xfa\xa8\x5b\x51\x92\x62\x79\x57\xa7\xe6\xaa\x79\x38\xab\x7e\x73\x27\xe2\xa7\x88\x4b\xb2\xcd\xfc\x72\x67\x5c\xd8\x9c\x53\x87\x81\x51\x30\x56\x7b\x9b\x62\xfe\x9f\x38\xa5\x09\x92\x7c\xab\x03\xf9\x85\xdb\xa0\xc8\x12\x9f\x2c\x45\xa9\xbe\xdd\xea\xb7\x5c\xd7\x2d\xb2\x79\x9d\xc4\x77\xa0\xa9\x8b\xd9\xb5\xc3\xe7\xe8\x6f\x8e\x6b\x98\xd5\x80\xd1\x2b\x8d\xd5\xf7\xb9\xaa\xf7\x01\xf3\x88\xb7\x99\x2f\x28\xdb\xc0\x21\x36\x63\x3e\x5c\x75\xf6\xea\x14\x2f\x79\x12\x43\xb3\x7d\x85\x9a\x72\xe0\x69\x62\xab\x08\xff\x95\xcf\x68\x71\xa9\x74\x4a\x9f\x93\xcf\xa3\xa6\x6c\x5c\xff\xc4\x79\xb5\xd7\xc8\x6d\xa1\x57\xd6\x8f\xea\x6c\x3f\x2a\x4b\x66\x63\x25\x3f\xe5\x92\xdd\x9f\x3b\xf5\x64\x0f\xff\x05\x00\x00\xff\xff\x76\xb2\x29\xf0\x73\x03\x00\x00")

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

	info := bindataFileInfo{name: "meta/config.yaml", size: 883, mode: os.FileMode(436), modTime: time.Unix(1469499308, 0)}
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

