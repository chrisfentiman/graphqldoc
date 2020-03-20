// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// template/enum.tmpl
// template/interface.tmpl
// template/object.tmpl
// template/scalar.tmpl
// template/schema.graphql
// template/schema.tmpl
package parser

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templateEnumTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x91\x41\x8b\xdb\x30\x14\x84\xef\xfa\x15\x03\xb9\xb4\x90\xea\x4f\x34\x29\x04\x4a\x52\x9a\x52\xe8\x4d\xb2\x3d\x6e\x54\x6c\xc9\xe8\xc9\xa5\x41\xd6\x7f\x2f\xf6\xc6\x9b\x64\x0f\xfb\x0e\xd2\x63\x66\x24\x7d\xe2\x6d\xb0\xf7\x63\x2f\x6a\x59\x11\x39\x44\x0a\x7d\xc2\x10\x44\x5c\xd5\x11\xc2\x24\x08\x2d\xfe\xda\x6e\xa4\xa0\x0d\x11\x16\xad\x63\xd7\x68\xa5\xbe\x84\x08\xfe\xb3\xfd\xd0\x71\x8b\x74\x21\xcc\x41\x64\xa4\x41\xa8\xfe\xb0\x4e\xb8\x58\x59\xd3\xa8\x6d\xd7\xb1\x81\x91\x64\x13\x8d\xc6\x8f\x0b\xb1\xf4\x70\x02\xeb\x41\x3f\xf6\xf8\x20\x03\x6b\xd7\xba\x39\x7c\xdd\xce\xef\xa6\xeb\xb0\x5e\x7b\x5e\x4e\x7e\x44\xc5\xda\x8e\x42\xb8\x84\xde\x5e\x51\x11\xe6\xf4\x6d\x7f\x34\x08\x11\xe6\xf3\xd7\xd3\x79\xbf\x33\x5a\xe5\x1c\xad\xff\x4d\xe8\x52\xd4\x66\x83\x9c\xf5\xd1\xf6\x2c\x45\xe5\xac\x77\x94\x3a\xba\x21\xb9\xe0\x4b\x51\x6a\xc2\xcf\xf9\x77\xc0\x84\x07\x07\xcf\x35\x7b\x43\x64\x6d\x13\x9b\x55\xf9\x4e\x2b\xc1\x63\x52\x13\x3e\xad\x85\x87\xfe\x3d\xf1\x49\xc2\x74\xe7\x9d\x47\xb1\xf0\x48\x29\x13\xcc\x2b\xb8\xc1\x84\xb7\xec\x8b\xe4\x5a\xe8\x83\xdc\xe9\x4a\xf9\x45\xc9\x99\x9d\xb0\x94\x63\xc8\x99\xbe\xb9\x45\xf5\x9a\x72\xc1\xbf\xc0\xcf\x86\xba\x45\x6e\xdb\xff\x00\x00\x00\xff\xff\xbe\x3f\x7c\x02\x14\x02\x00\x00")

func templateEnumTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateEnumTmpl,
		"template/enum.tmpl",
	)
}

func templateEnumTmpl() (*asset, error) {
	bytes, err := templateEnumTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/enum.tmpl", size: 532, mode: os.FileMode(420), modTime: time.Unix(1584631096, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateInterfaceTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x50\xc1\x6e\xd4\x30\x10\xbd\xfb\x2b\x9e\xb4\x1c\x40\x02\x7f\x00\x12\x87\x4a\xa8\xd2\x4a\xd0\x43\xc5\x05\x21\xa4\x38\xce\x4b\x62\x9a\xd8\xc1\xe3\xa5\x54\x8e\xff\x1d\x25\xcd\x76\x37\x07\x3a\x87\xcc\xe4\xcd\x9b\x79\xf3\x7c\xc0\xd1\x27\xc6\xd6\x58\x8a\xba\x94\x10\xc6\x3f\x84\x11\x4c\x26\xd2\x27\x84\xfa\x17\x6d\x12\xb4\x31\x8c\x78\xec\x9d\xed\x11\x52\xcf\xf8\xd2\xb0\xc6\xc3\xf9\x9e\xd1\x25\xad\xd4\x6d\x88\xe0\x5f\x33\x4e\x03\xdf\xa3\xfa\x12\xec\x83\xa9\x07\x56\x70\x82\x95\xb7\xe9\xa0\xa6\x35\x27\x21\xea\x90\x7a\x54\x47\x91\x13\x2b\x18\xdf\xa0\xfa\xca\xd8\xf1\x9e\xbf\x4f\x94\x54\xed\x54\x6a\x62\x08\xf6\x81\x8d\xc6\xcd\xf5\xae\xde\x08\x5c\x12\x84\x47\x8f\xc1\x49\x42\x68\xe1\xcd\xc8\x06\xad\xe3\xd0\x08\x52\x6f\x12\x4c\x24\xa4\x37\x91\x0d\xea\x27\xb8\xe5\xc2\x91\x3e\x39\xdf\x9d\x45\xb4\xca\x39\x1a\xdf\x11\xba\x14\x75\x38\x20\x67\x7d\x67\x46\x96\xa2\x72\xd6\x9f\x29\x36\xba\x29\xb9\xe0\x4b\x51\x6a\xc6\x4d\xec\x4e\xcb\x06\x60\xc6\xb7\xa7\x89\x4b\xbe\x22\x61\x1f\x4b\x6f\x8a\xb4\x26\xb1\x39\x23\xf7\x34\x12\x3c\x66\x35\xe3\xc3\x39\xf0\x9f\xfa\x35\x70\x07\x61\xbe\xd8\xb8\x5d\xfd\x97\x32\xa3\x7a\xf1\x52\x61\x46\xce\x6f\x12\x3e\x7e\x42\xc7\xb4\x5e\xae\x97\x6f\x29\x3f\x16\x7c\xab\x7f\xbe\x5d\x7f\x9e\x67\xde\x2d\x23\xfb\x17\x58\xb7\xb8\x16\xfa\x28\x17\x63\xa5\x7c\xa7\xe4\xcc\x41\x58\xca\x5d\xc8\x99\xbe\xd9\xa8\xfa\xcc\x72\xc1\x3f\xfb\x5e\x1a\x6a\xa3\x6c\xe9\x5f\x00\x00\x00\xff\xff\xaa\x16\xe3\x7b\x95\x02\x00\x00")

func templateInterfaceTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateInterfaceTmpl,
		"template/interface.tmpl",
	)
}

func templateInterfaceTmpl() (*asset, error) {
	bytes, err := templateInterfaceTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/interface.tmpl", size: 661, mode: os.FileMode(420), modTime: time.Unix(1584631096, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateObjectTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8f\x41\x8b\xc2\x30\x10\x85\xef\xf9\x15\x0f\xba\x87\xdd\xc0\xe6\x07\x2c\xec\xad\x08\x5e\xf4\xe2\x4d\x84\xd6\x76\x2c\x11\xad\xa5\xc9\x45\x66\xe6\xbf\x4b\x63\xb5\xed\xc5\x39\x24\xef\x3d\xf2\xc8\x37\x19\xb6\xc7\x33\x55\x31\x18\xe6\xbe\x6c\x1b\x82\x53\x35\x59\x06\x66\xb7\x29\xaf\xa4\x6a\x98\x5d\x4e\xa1\xea\x7d\x17\xfd\xad\x55\x35\x46\xb0\xf2\x74\xa9\x01\xc1\xee\xde\x11\x96\x23\x98\x3d\x1f\x23\x23\xf8\x7d\x0d\x66\xfa\x53\x08\x99\x98\xd2\x7f\x41\x55\x50\xbc\xc1\x0a\x08\x98\xbf\x22\xfe\xfe\xd1\x50\x4c\x24\x6e\x38\x55\xf7\x43\x3e\xea\xc3\x77\x32\xcf\xce\x4f\xea\x2c\xf7\x01\xb3\x3f\xc1\xad\x43\x4e\x5d\x4f\x55\x19\xa9\x56\xb5\x76\x72\xd6\x32\x53\x5b\xab\x26\xa2\xa4\xc6\xeb\x11\x00\x00\xff\xff\x87\xd3\x8b\xcc\x3e\x01\x00\x00")

func templateObjectTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateObjectTmpl,
		"template/object.tmpl",
	)
}

func templateObjectTmpl() (*asset, error) {
	bytes, err := templateObjectTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/object.tmpl", size: 318, mode: os.FileMode(420), modTime: time.Unix(1584631096, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateScalarTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xce\x31\x6a\xc4\x30\x10\x85\xe1\x5e\xa7\x78\xb0\xad\xd1\x01\xd2\x25\x2c\x09\x86\x10\x12\xb6\x48\xab\x89\x3d\x6b\x0f\xc8\x23\x33\x92\x16\x8c\xd0\xdd\xc3\x6e\x52\xbd\xbf\xf8\x8a\x77\xc2\x65\xa2\x48\x96\xdd\xff\x82\x8c\xb1\x9b\x6c\x52\xe4\xc6\xb8\x51\xac\x9c\x9f\x10\x46\x2d\x61\x40\x78\x8d\x89\x1e\x71\x29\x26\xba\xdc\xeb\x25\xa5\xc8\xa4\x61\x40\x32\x84\xf1\x1c\xbc\x73\xdf\x2b\x2b\x26\x8a\x51\x74\x41\x59\x19\x6f\x46\xfb\xfa\xf5\x8e\xe7\xcf\x71\xc0\x91\x2a\xb6\x9a\x0b\xf2\xce\x93\x5c\x0f\x28\xe7\xc2\x33\x72\xfd\xb9\x0a\xc7\x19\x55\x8b\xc4\x07\x33\x2e\xd5\x14\x49\xe3\x81\xfc\xf7\xd0\x3b\xd7\x9a\x91\x2e\x0c\xdf\xfb\xe9\x84\xd6\xfc\x07\x6d\xdc\xbb\x6b\xcd\x9f\x39\x4f\x26\x7b\x91\xa4\xbd\xdf\x25\xeb\xdc\xfb\x6f\x00\x00\x00\xff\xff\xb6\xdd\x72\x6f\xe7\x00\x00\x00")

func templateScalarTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateScalarTmpl,
		"template/scalar.tmpl",
	)
}

func templateScalarTmpl() (*asset, error) {
	bytes, err := templateScalarTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/scalar.tmpl", size: 231, mode: os.FileMode(420), modTime: time.Unix(1584631096, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateSchemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x93\xcd\x6e\xeb\x20\x10\x85\xf7\x7e\x8a\x91\xee\xe6\x76\x93\x07\xe8\xba\x8a\xd4\x55\xa5\x28\xea\xd6\x22\x30\xa4\xa8\x18\x28\x3f\x95\xa2\xc8\xef\x5e\x41\xb0\x3d\x38\x6a\x6b\x75\x93\x98\xe1\x83\x73\xce\x00\x1f\x09\xfd\x05\xae\x1d\x40\xdf\x07\xfe\x86\x03\x2b\x03\x80\x32\x71\xbc\x38\xac\x63\x80\xdd\x6e\xb7\x4f\x5a\xe7\x5a\xa9\x8c\xe5\x77\x48\x91\x45\x65\xcd\x06\x34\xa4\x53\xe0\x5e\xb9\x8d\x78\xbc\x38\x0c\xbf\x30\x42\x79\xe4\x51\x7d\x12\xd0\xb0\x01\xeb\xa7\xc0\x59\xaf\x56\x98\x3f\x2f\x64\xd9\xf4\xd9\xb8\x14\x5f\x99\x4e\xd3\xa2\xb1\xfe\xff\x03\x6b\x5e\x1c\x7a\x46\x96\xe7\xda\xde\xb3\xf3\x80\x26\xd2\x92\x42\x2d\x66\x57\x63\x37\x76\x9d\xac\x14\x4c\xa6\xc1\x1a\xe8\xfb\x39\xf6\xbb\x32\x79\x45\x35\xdb\x1a\x95\x79\xbb\xf0\x5f\x19\xae\x93\xc0\x27\x74\x1e\x39\x8b\x28\x1e\x21\xfa\x84\x0f\xd5\xff\x9c\x73\x9d\xb2\xc9\x78\x9f\x70\x69\x2e\x85\xb2\xb1\x03\x4a\x42\xa8\xb0\x28\x57\x9d\xdb\x50\x59\x73\x40\x16\x8a\x5a\x46\x55\xde\xbf\xb4\x60\x92\x5d\x8b\xde\xa8\x88\x5e\x32\x8e\x04\x5a\x44\x33\x81\x26\x0d\x65\xc5\xdf\xa3\x6f\xf6\xec\x6c\x08\xea\xa4\xf1\x48\xee\xd8\xca\x10\x3d\xc4\x25\xcd\xed\x18\xc9\xf8\xfa\xdd\x29\x92\x0e\xdf\x45\x15\x28\x59\xd2\xb5\x3f\x54\xa8\x62\x3f\x5e\x16\x2b\xc9\xe3\xa9\x53\xa4\x27\xcd\x34\x01\x9a\x97\xb1\x82\x1a\xac\x01\xc7\xe6\x5a\x7f\x05\x00\x00\xff\xff\x79\x7b\x42\xc7\x2e\x04\x00\x00")

func templateSchemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_templateSchemaGraphql,
		"template/schema.graphql",
	)
}

func templateSchemaGraphql() (*asset, error) {
	bytes, err := templateSchemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/schema.graphql", size: 1070, mode: os.FileMode(420), modTime: time.Unix(1584670571, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateSchemaTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x52\xcd\x6a\xf3\x30\x10\xbc\xfb\x29\x06\x3e\x1f\xf2\x89\x54\xf4\x1c\xe8\x21\x10\x0a\xbd\xf4\x50\x7a\x0b\x05\x99\x78\xe3\x08\x1c\x25\xd8\x2a\xb4\xac\xf6\xdd\x8b\x64\xd9\x89\x49\x68\x75\xb0\xf6\x47\x3b\x3b\x3b\xeb\x7f\x60\xd6\xaf\xd5\x91\x44\x0a\x66\xbd\xa1\x7e\xd7\xd9\xb3\xb7\x27\x27\x52\x14\xcc\x5d\xe5\x1a\x82\x7e\xb6\xd4\xd6\xbd\x08\x33\xca\x96\x5c\xe3\x0f\x58\x3d\xa1\x25\x07\xbd\xee\x9a\x1e\x31\x53\xfa\x03\xbd\x7f\x9f\x29\x66\x1a\xf2\xc9\xd4\xf1\x9b\xa0\xed\x1e\xfa\xa5\xdf\xd0\xb9\xa3\x5d\xe5\xa9\x16\x31\x17\xc7\x30\x93\xab\x45\x94\x9a\xd8\x28\x05\x66\xd8\x3d\x1c\x4d\x2d\x1f\x21\xb2\x18\x39\x95\xd6\xd5\xf4\xb5\x44\x49\x2d\x1d\xc9\xf9\xd8\x36\x91\x89\x5c\xec\x3e\xe7\x45\x96\xc8\xe0\x86\x79\x7c\x9b\x7b\x98\x15\x22\xed\x6b\xc2\xd3\x8b\x81\xf8\x36\xe6\xb3\xfd\xb1\x48\xce\x50\xfa\x3f\x83\x8e\x37\x56\xd8\x5e\x24\x98\x55\xe4\x50\xae\xbb\x2b\x85\x4a\xd2\x0f\xbe\x3d\xb9\x37\xaa\xfa\xb8\x00\x95\xc1\x6f\x37\x73\x57\x9a\x22\x60\xdd\x35\x9f\x49\x0c\x04\xa4\x79\xe6\x27\xe0\x0a\x26\x87\x8a\x80\x87\xf1\xe0\xca\xfe\x2d\x88\x70\xf9\x35\x06\xc9\x03\xcc\xb4\x3a\x83\x70\xa3\xeb\x9f\x7a\xa6\x9a\xf9\x94\xa9\x4b\x12\x60\x7e\xfd\x04\x00\x00\xff\xff\x1b\xab\x48\x4c\xb4\x02\x00\x00")

func templateSchemaTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateSchemaTmpl,
		"template/schema.tmpl",
	)
}

func templateSchemaTmpl() (*asset, error) {
	bytes, err := templateSchemaTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/schema.tmpl", size: 692, mode: os.FileMode(420), modTime: time.Unix(1584631096, 0)}
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
	"template/enum.tmpl":      templateEnumTmpl,
	"template/interface.tmpl": templateInterfaceTmpl,
	"template/object.tmpl":    templateObjectTmpl,
	"template/scalar.tmpl":    templateScalarTmpl,
	"template/schema.graphql": templateSchemaGraphql,
	"template/schema.tmpl":    templateSchemaTmpl,
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
	"template": &bintree{nil, map[string]*bintree{
		"enum.tmpl":      &bintree{templateEnumTmpl, map[string]*bintree{}},
		"interface.tmpl": &bintree{templateInterfaceTmpl, map[string]*bintree{}},
		"object.tmpl":    &bintree{templateObjectTmpl, map[string]*bintree{}},
		"scalar.tmpl":    &bintree{templateScalarTmpl, map[string]*bintree{}},
		"schema.graphql": &bintree{templateSchemaGraphql, map[string]*bintree{}},
		"schema.tmpl":    &bintree{templateSchemaTmpl, map[string]*bintree{}},
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