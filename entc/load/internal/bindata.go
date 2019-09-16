// Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// template/main.tmpl
// schema.go
package internal

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

var _templateMainTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x51\x5d\x6b\xdb\x30\x14\x7d\xb6\x7e\xc5\x99\xe9\xa8\x5d\x52\xa5\xed\xdb\x06\x79\x28\x6d\x06\x19\x5b\x3b\x48\x61\x0f\x5d\x29\x8a\x7d\x9d\x88\x3a\x92\x77\xa5\x94\x05\xa1\xff\x3e\x24\x27\x61\x7b\xb2\xa5\x73\xee\xf9\xd0\x0d\x61\x7a\x21\xee\xec\xb0\x67\xbd\xde\x78\xdc\x5c\x5d\x7f\xba\x1c\x98\x1c\x19\x8f\x2f\xaa\xa1\x95\xb5\x6f\x58\x98\x46\xe2\xb6\xef\x91\x49\x0e\x09\xe7\x77\x6a\xa5\x78\xda\x68\x07\x67\x77\xdc\x10\x1a\xdb\x12\xb4\x43\xaf\x1b\x32\x8e\x5a\xec\x4c\x4b\x0c\xbf\x21\xdc\x0e\xaa\xd9\x10\x6e\xe4\xd5\x11\x45\x67\x77\xa6\x15\xda\x64\xfc\xdb\xe2\x6e\xfe\xb0\x9c\xa3\xd3\x3d\xe1\x70\xc7\xd6\x7a\xb4\x9a\xa9\xf1\x96\xf7\xb0\x1d\xfc\x3f\x66\x9e\x89\xa4\xb8\x98\xc6\x28\x44\x08\x68\xa9\xd3\x86\x50\x6e\x95\x36\x25\x62\x14\xd3\x29\xee\x52\x9e\x35\x19\x62\xe5\xa9\xc5\x6a\x8f\x73\x32\xbe\x39\x5d\x9d\x4b\xdc\x3f\xe2\xe1\xf1\x09\xf3\xfb\xc5\x93\x14\x83\x6a\xde\xd4\x9a\x90\x34\x84\xd0\xdb\xc1\xb2\x47\x25\x8a\xd2\xba\x52\x14\xe5\x6a\xef\x29\xfd\x84\x00\x4f\xdb\xa1\x57\x9e\x50\x8e\x2c\x97\x2d\x33\x34\xb0\x36\xbe\x43\xf9\xf1\x77\x09\xf9\xe3\xa0\x18\xa3\xa8\x73\xcc\xb3\x95\x72\x84\xcf\x33\xe4\xef\x11\x4f\xb3\xef\x8a\xe1\x9a\x0d\x6d\x95\xc3\x0c\xcf\x2f\x64\xbc\x5c\x18\x4f\xdc\xa9\x86\x42\x96\x66\x65\xd6\x84\xb3\xd7\x09\xce\x8c\xda\x66\x19\xf9\xa0\xb6\xe4\x92\x7e\x51\x84\x70\x79\xd0\x8f\x51\xa6\xc3\x29\x8a\x0b\xb1\x3c\xcc\xc4\x38\xc9\x5a\x64\x5a\x5c\xc6\x28\xa2\x10\xdd\xce\x34\xb9\x73\x55\x23\x88\x22\x05\xe9\xb5\x21\x87\xe7\x97\xe7\x97\x54\x5a\x14\x9d\x65\xbc\x4e\x0e\xf9\x92\xef\x18\xe5\x98\x37\x88\xa2\x58\x4d\x40\xcc\x09\xfb\xae\xd8\x6d\x54\xbf\xcc\x60\x35\x72\x6a\x51\x14\xba\xcb\x8c\x0f\x33\x18\xdd\xe7\x99\xa2\x53\xba\xaf\x88\x39\xc1\xa9\xc2\xe8\x3b\x83\x1a\x06\x32\x6d\x95\x8f\x13\xac\x6a\x91\x50\xeb\xe4\xd2\xb7\x76\xe7\xe5\x4f\xd6\x9e\xaa\xbc\x0f\xf9\xd5\x6a\x73\x24\x8e\x71\xab\xf2\x97\x29\xeb\xba\x3e\x75\x3b\xba\x24\x7b\xcb\xb9\xe4\xa8\x45\xcc\xa3\xd6\xd2\xb3\x36\xeb\xc4\x91\xf3\xc4\xa9\xea\x3a\x73\xe6\x7f\xb4\xaf\xae\xb3\xd2\x7f\x5b\x1f\x4b\x8d\x4b\x3f\x3c\x66\x8c\xe2\x6f\x00\x00\x00\xff\xff\xe4\x6e\x0c\x4d\x4b\x03\x00\x00")

func templateMainTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateMainTmpl,
		"template/main.tmpl",
	)
}

func templateMainTmpl() (*asset, error) {
	bytes, err := templateMainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/main.tmpl", size: 843, mode: os.FileMode(420), modTime: time.Unix(1567330508, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\xdd\x6f\xdb\x38\x12\x7f\x96\xfe\x8a\x59\x03\x1b\x48\x81\x4f\xee\xed\xdb\x69\xe1\x87\x45\x37\x0b\xe4\xf6\x2e\x2d\x9a\xde\xbd\x04\x41\x4a\x8b\x43\x9b\x8d\x44\xaa\x24\xed\xc4\x0d\xf2\xbf\x1f\xf8\xa5\x2f\xdb\x41\xaf\xdd\x04\x08\x6c\xce\x17\x39\xbf\xf9\x71\x38\x5e\x2c\xe0\xad\x6c\xf7\x8a\xaf\x37\x06\x7e\x79\xf3\xf7\x7f\xfc\xad\x55\xa8\x51\x18\xf8\x83\x54\xb8\x92\xf2\x1e\x2e\x45\x55\xc0\x6f\x75\x0d\xce\x48\x83\xd5\xab\x1d\xd2\x22\x5d\x2c\xe0\xe3\x86\x6b\xd0\x72\xab\x2a\x84\x4a\x52\x04\xae\xa1\xe6\x15\x0a\x8d\x14\xb6\x82\xa2\x02\xb3\x41\xf8\xad\x25\xd5\x06\xe1\x97\xe2\x4d\xd4\x02\x93\x5b\x41\x6d\x08\x2e\x9c\xc9\xbf\x2e\xdf\x5e\x5c\x5d\x5f\x00\xe3\x35\x46\x99\x92\xd2\x00\xe5\x0a\x2b\x23\xd5\x1e\x24\x03\x33\xd8\xcf\x28\xc4\x22\x4d\x5b\x52\xdd\x93\x35\x42\x2d\x09\x4d\x53\xde\xb4\x52\x19\xc8\xd2\x64\x86\xa2\x92\x94\x8b\xf5\xe2\xb3\x96\x62\x96\x26\x33\xd6\x18\xfb\xa1\x90\xd5\x58\x99\x59\x9a\x26\xb3\x35\x37\x9b\xed\xaa\xa8\x64\xb3\x60\x21\x61\x2e\xaa\xed\x8a\x18\xa9\x16\x28\xcc\x42\x57\x1b\x6c\xc8\x02\xe9\x1a\xbf\xc9\x61\xf6\x7f\x04\x65\x1c\x6b\x3a\x4b\xf3\xd4\xc2\x70\xed\x64\xa0\x30\x14\x40\x03\x11\x80\xc2\x14\x41\x61\x36\xc4\xc0\x03\xd1\x2e\x4f\xa4\xc0\x94\x6c\x80\x40\x25\x9b\xb6\xe6\x16\x6c\x8d\x0a\x02\x16\x45\x6a\xf6\x2d\xc6\x90\xda\xa8\x6d\x65\xe0\x29\x4d\xae\x48\x83\x10\xff\xb4\x51\x5c\xac\xbb\xe5\x27\x8b\x52\x39\x13\xa4\xc1\xb9\x6c\xb8\xc1\xa6\x35\xfb\xd9\xa7\x34\xb9\xa0\x6b\xd4\xd1\xea\xe6\xf6\xdc\xae\x27\x4e\x16\x1d\x3d\xf6\xfa\xc3\xe6\xa6\x7b\x2f\xb7\x1e\x7b\xb9\xf4\x27\x6e\x97\x82\xe2\x63\xdc\xee\xe6\xf6\xdc\xad\xc7\x6e\xdc\x9b\x8c\xfd\xae\x5d\x8e\x61\xd3\x9b\xdb\xf3\xc1\x3a\xfa\x79\x18\xee\x8e\xec\xfa\xec\x0a\xe0\x8d\x0f\xf1\xf7\xf2\xef\x80\xdf\x3b\x9e\x40\x7f\x08\xff\x49\xe8\x3f\xda\x30\xdd\x9f\x3b\x79\xe1\x64\xc1\xc3\x6e\x33\xf1\x20\x7d\x45\x8f\xed\x61\xc8\x7a\x02\x1c\xff\x3a\xd8\xe2\x9c\x0b\x33\x46\x5b\xf3\xaf\x93\x2d\xfe\x23\xf8\x97\x6d\xe7\xb3\x92\xb2\x1e\x7b\x6c\x9d\x7e\xec\x73\xc5\xeb\x9a\xac\x6a\x3c\xe5\x23\x82\x7e\xec\xf5\xae\x35\x5c\x0a\x52\x9f\xf2\x92\x41\x3f\xf6\xfa\x1d\x19\xd9\xd6\xe6\xe4\xf9\xa8\xd7\x4f\x92\x6a\x29\x31\x18\x5d\x8f\x24\xe5\xf4\x77\x47\x7d\x2f\x9b\x66\x6b\xba\xec\x0e\x7d\x79\xd4\x8f\xdd\xfe\x4b\x6a\x4e\x6d\x4b\x70\x6c\xef\x81\x8f\x6e\xbb\x4e\x3f\xa5\xba\x54\x64\x8d\x7f\xe2\xfe\x78\x89\xb5\xd7\xdf\xdd\xe3\xfe\x08\xcb\x87\x17\x63\xc2\xf5\x47\x83\xca\xa2\x1d\x18\xeb\xe8\x06\x14\x19\x17\x48\x63\x47\xf6\x9d\x2b\xf6\x97\x41\xac\x9e\xe6\x1d\x05\xc3\xd1\x4e\x31\xaf\xbb\x0e\x63\xbb\x17\x6e\xc1\x24\xe0\x01\xf9\xdf\xca\xa6\xb1\x4f\xd7\xc4\xb0\xf2\xe2\xb1\xed\xfb\xfb\xf5\x7b\x62\x36\x53\xdb\xf6\x7e\x7d\xd7\x12\xb3\x99\xb4\xc0\x66\x85\xd4\xde\xfa\x50\xdc\xd8\xf6\x82\xf8\x08\xcc\xae\x49\x1e\xf6\x12\x27\xfe\x8e\x56\xe2\xfc\x8e\x74\x92\xbf\x0c\xba\x6f\x2d\xda\x07\x64\x7e\xf3\xb1\x9d\x42\x76\x77\xb8\xfb\x07\x64\xa1\xad\xf8\x37\xa3\x37\x3e\xd1\x51\xc6\xf0\x1e\x6b\x24\x97\x62\x87\x4a\xe3\xd4\x94\x7b\xf1\x74\xfb\x2f\x5b\xae\x0e\xaa\xa6\x82\xf8\x48\xd5\xfc\x6b\x73\x58\x36\x2f\xff\x8e\xba\x79\xc7\xbe\x70\x21\xd3\xae\x47\xbc\x90\x69\x78\x76\x6f\x6e\xc7\x48\x9f\x7e\x6a\xa7\x96\x27\x1f\xba\x2b\x7c\x70\xf5\xa8\x14\x12\x83\x2e\xc9\x90\x91\x0d\xee\xd3\x72\xdf\x28\xea\x4a\xf1\xd6\x48\x55\xa4\x6c\x2b\xaa\xe8\x99\x21\x85\x73\x6b\x51\xfc\xde\x59\xe4\xa1\xc8\x4f\x69\x22\x10\xca\x25\x9c\xd9\xe5\x53\x9a\x58\x6a\x95\x9e\x06\x48\x8b\x8f\x64\x3d\xb7\xb2\x7d\x8b\x65\x27\xb3\x6c\x4c\x13\xc7\xea\x4e\x68\x17\x56\xe8\x11\x2b\xbd\xd0\x2f\xac\x38\xf0\xa0\x74\xe2\xb0\xb0\xf2\x58\xf3\xd2\xca\xe3\xc2\x2b\x58\x88\xef\x14\x2c\xc4\x7f\x4e\x13\xce\x40\x21\xb3\x47\xf6\x9a\x5f\xdd\xf2\xa7\x25\x08\x5e\xdb\x74\x12\x81\x56\x0c\xcb\x2e\x7d\x85\x2c\x77\xae\x0a\xcd\x56\x09\x10\x18\x90\xfd\x37\x51\x7a\x43\xea\x30\x77\xb9\xf9\x13\x5d\xdb\x1c\xcc\x71\x5c\x18\x54\x76\x2c\xb4\xdf\x24\x10\xf8\xe7\xf5\xbb\x2b\xeb\xec\xe8\x55\x11\x01\x2b\x8b\xbc\x75\xa5\xde\xc4\x06\x08\xce\x72\xf5\x19\x2b\x13\x3e\x42\x51\x46\x9b\x66\x3a\xee\x6d\x59\x1b\x76\xca\x21\x5b\xc1\xcd\xed\x6a\x6f\x70\x0e\xa8\x94\xfd\xb7\x15\x7b\x4a\x13\xed\x4a\xe5\x7d\x9f\x3c\x40\x5c\xf8\x89\x3b\x0b\x73\xb2\xab\xcf\x3b\x16\x22\xe7\xb9\x2b\x4d\x96\x3f\xa7\x49\x60\x98\x0b\x59\x2e\x41\x13\x86\x9e\x8b\xd1\xd6\x81\x6b\xb5\x03\x34\x23\x66\xbc\x9e\x03\x6b\x4c\x71\x61\xcf\xc2\xb2\x59\x38\xf8\xcf\x5f\x4a\xf8\x79\x37\x9b\x83\xf6\x14\xb0\xee\x1e\x6c\x26\x15\xdc\xcd\xc1\x55\x4a\x11\x61\x99\xea\x89\x6f\xa3\x32\x6a\xc5\x6c\x40\xc8\x2c\x4f\x93\x44\x3b\xeb\x33\x77\x2a\x6b\x36\xe0\x98\x1f\xa9\x7a\xa2\x0d\x38\x19\x55\x91\x98\x03\x0a\x77\x2a\xcf\xe3\x01\x3d\xa3\xa6\xe7\x68\x37\xf5\x94\xfd\x66\x71\xce\xb1\xea\x38\xde\xf4\xea\x28\x71\xea\x6e\xac\x28\xa3\xba\x93\x38\x7d\x3f\x07\x94\x41\xdf\x4b\x9c\x41\x3f\x60\x38\x83\x1a\x45\xc6\x68\xd1\x4b\x73\x67\x15\x66\x9e\xb2\xcf\x20\x4e\x41\xbe\x68\x3e\xcd\xe1\x78\x54\xba\x34\x47\x03\x53\x6f\xfa\x9c\x26\xb6\xe8\xf6\x34\x76\xb2\xfc\x69\x09\x6f\x5c\x81\x12\xcd\xbc\x64\x09\x67\x41\x19\xac\x75\x11\x1a\xd8\x12\x48\xdb\xa2\xa0\x59\x94\xcc\x41\x87\x8b\xe6\xbb\xde\x90\x68\xae\x3d\xbe\x26\xcf\xb0\xe7\x99\xdb\xdd\x05\xd5\x85\x6f\xcb\x83\xa3\x5e\xf8\xa3\x75\xad\x71\x44\xc1\xdc\x87\x8c\x3f\x59\x86\x09\x84\x5f\x3a\xaf\x99\x02\xa7\x8f\x7d\x12\xe1\x0c\x2e\x70\x50\x70\xfa\x78\x70\x61\x8a\xf8\x0b\x6c\x90\xe2\x65\x3c\xfe\x99\xfb\xe6\xca\xe9\xd2\x2e\xc1\xc5\xf0\x10\x58\xa9\xaf\x5b\xe9\xa4\xa1\x86\xc3\x5b\x62\xc5\xfd\xfd\x78\x1e\x35\x51\xfb\x68\x15\xa1\x97\x65\x3a\x0f\x1d\xb5\xef\x29\xf0\xa0\x48\xab\x87\x43\x68\x90\x37\x68\x36\x92\xc2\x03\x37\x1b\x50\x58\xc9\x1d\x2a\x30\x12\x50\xe8\xad\x42\x10\x12\x5a\x22\x78\xa5\xed\x04\xdb\xf8\xf0\x5c\xac\x43\xef\x3c\x68\x59\x07\x8d\x93\xc5\xc7\xb5\xfb\x29\x38\x6d\xa1\x14\x19\x2a\xb0\xe1\xb2\xdc\xa3\xcb\x60\xe7\x70\xf7\x87\xc9\xf2\x5f\x61\x37\x2c\x6b\x62\xfd\x97\x47\x2a\x1a\x33\xf2\x07\x0e\xc5\xdd\xd9\xb2\x84\x56\x0b\x2e\x88\xbf\x37\xcf\xb6\x5e\x01\xbb\x91\x7b\x96\xcf\x9d\x55\x0f\xa0\xe7\xec\x01\x7e\x5e\xfc\xa3\xf0\x0d\x2f\xe2\x01\x7a\xfe\xe6\x78\xf0\xac\xe1\x2b\x62\xe7\xb3\x39\x02\x1d\x86\x1b\xfb\x12\x72\x3e\x89\x03\xe0\xe2\x5d\x38\x80\x2e\x2a\x7e\x14\xbc\x71\x13\x38\x80\x2f\xde\x59\x0f\xa0\x33\x7e\x45\x04\x63\x52\x47\x30\xe4\x5d\x53\x78\x09\xc5\x98\x4d\x8f\xa3\x4b\xb4\x1b\x28\x0c\x0c\x47\x8a\x7c\xb4\xb2\x67\xb3\x6d\xcb\x14\x7f\x72\x41\xb3\x1c\x96\xcb\x4e\xff\xde\x28\x77\x74\x03\x4b\x30\xc5\x45\x8d\x4d\x36\x6a\x1d\x26\x7d\x4e\xff\x17\x00\x00\xff\xff\x9d\x69\x83\xb8\xc9\x14\x00\x00")

func schemaGoBytes() ([]byte, error) {
	return bindataRead(
		_schemaGo,
		"schema.go",
	)
}

func schemaGo() (*asset, error) {
	bytes, err := schemaGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.go", size: 5321, mode: os.FileMode(420), modTime: time.Unix(1568625880, 0)}
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
	"template/main.tmpl": templateMainTmpl,
	"schema.go":          schemaGo,
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
	"schema.go": &bintree{schemaGo, map[string]*bintree{}},
	"template": &bintree{nil, map[string]*bintree{
		"main.tmpl": &bintree{templateMainTmpl, map[string]*bintree{}},
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
