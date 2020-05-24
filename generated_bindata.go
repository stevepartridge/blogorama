// Code generated for package blog by go-bindata DO NOT EDIT. (@generated)
// sources:
// _sql/mysql/1000_initial_users.down.sql
// _sql/mysql/1000_initial_users.up.sql
// _sql/mysql/1001_initial_posts.down.sql
// _sql/mysql/1001_initial_posts.up.sql
package blog

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

var __sqlMysql1000_initial_usersDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x48\x28\x2d\x4e\x2d\x2a\x4e\xb0\xe6\xc2\x23\x1b\x9f\x91\x59\x5c\x92\x5f\x54\x09\x57\x15\xe4\xe9\xee\xee\x1a\x84\x53\x5d\x7c\x49\x51\x66\x7a\x7a\x6a\x51\x7c\x69\x41\x4a\x62\x49\x2a\xc9\xda\x52\x52\x73\x52\x41\xda\x00\x01\x00\x00\xff\xff\x67\xaa\x0f\x8f\xb1\x00\x00\x00")

func _sqlMysql1000_initial_usersDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__sqlMysql1000_initial_usersDownSql,
		"_sql/mysql/1000_initial_users.down.sql",
	)
}

func _sqlMysql1000_initial_usersDownSql() (*asset, error) {
	bytes, err := _sqlMysql1000_initial_usersDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_sql/mysql/1000_initial_users.down.sql", size: 177, mode: os.FileMode(420), modTime: time.Unix(1590089791, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __sqlMysql1000_initial_usersUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x55\x51\x6f\x9b\x3a\x14\x7e\xe7\x57\x9c\xb7\x18\xdd\xa6\xba\xb9\x57\xaa\x26\x45\x79\x20\xc1\x49\xd1\x12\x93\x19\xa3\xae\x4f\xc0\x8a\x97\x5a\x0b\x4e\x45\x9c\x49\xf9\xf7\x93\xc1\x80\x47\xd2\x68\xed\x36\xa9\x9a\x86\x22\xc5\x3e\xdf\x77\x8e\x0f\x87\xef\x83\xe1\x10\x86\xf6\xe5\x54\x3f\x38\xec\x79\xb9\xd7\x6b\x9f\x86\x6b\x60\xde\x74\x89\x21\x98\x03\xfe\x18\x44\x2c\x82\xb4\x82\xd3\xb1\x33\xa3\xd8\x63\xd8\xe0\x26\x0a\xc8\x01\x48\x45\x9e\x82\x90\x0a\x8d\x46\x2e\x1c\xe4\x5e\x6c\x24\xcf\x81\x84\x0c\x48\xbc\x5c\x82\x17\xb3\x30\x09\xc8\x8c\xe2\x15\x26\xec\x4a\x27\x64\x4f\x22\xf9\xc2\x8f\x29\x3c\x3c\x66\x25\xfa\xff\xc6\x6d\xd9\x15\x2c\xb3\x82\x1b\x6c\xf4\xdf\xbb\x1e\x98\x3d\x28\xf1\x95\xa7\xa0\x84\x3c\x56\x87\x9e\x3b\x33\xe7\x9f\xb3\xc3\x56\xc1\xe0\xdf\x41\x95\x74\x78\xca\x33\xc5\xf3\xe4\xd3\xf1\x52\xa7\xcf\x65\x65\x4a\x1f\x57\xf0\xbd\xca\x8a\x27\x74\xe3\xb6\xc4\x59\x4c\x29\x26\x2c\x61\xc1\x0a\x47\xcc\x5b\xad\x35\xb8\x93\x50\x27\x9e\x85\xab\xc2\x0f\x25\x7f\x45\x3b\x4d\x96\x6e\x47\xd7\xd7\x2d\xa1\xde\xec\xd6\x34\x58\x79\xf4\x1e\xde\xe3\x7b\x40\xfa\xc1\x54\x07\xc6\x24\xf8\x10\xe3\x2a\xd8\xcd\x1e\xb5\xcb\x8a\x53\x81\xf5\xe4\x51\xfd\xdf\x85\x9b\x99\xa3\x66\xd5\x41\xf6\x64\x91\xbd\x3b\xa5\xe8\xbe\x91\xbd\xeb\x28\xf6\x40\x90\xbd\x3b\xa5\xd4\x55\xac\x9d\xeb\xb8\x80\xc9\x22\x20\x78\x12\x48\xb9\xf3\xa7\xe0\xe3\xb9\x17\x2f\x19\xcc\x6e\x3d\x1a\x61\x36\xd9\x66\x4a\xc8\xd1\xd8\xb9\x24\xf0\xe4\x51\xec\xd5\xae\x3c\x9e\x17\x7a\x8b\xbe\x52\xf0\xba\x48\x72\x31\xeb\x2d\xfb\xa2\x99\xe7\x39\x5f\x3c\x27\xc4\xdf\xa8\xf1\xd4\x3c\x8c\xa4\xe0\xea\x71\x97\xa7\xc0\xe5\xa1\x40\x83\xba\xab\xc1\xd5\x20\xe7\x5b\xae\xf8\xc0\x3d\x73\x03\x0d\xc7\xae\xf2\x72\x3b\xd5\xaa\x6e\x1e\x29\x6a\x97\x96\x61\xfe\x38\x93\x35\x14\xd4\x9f\xfe\x19\xe4\x65\xa6\x1c\x0e\xe1\x9f\x42\x6c\xca\x4c\x89\x9d\x84\x29\xdf\x08\x19\xa9\x4c\xf1\x82\x4b\x65\x1c\x4b\x83\xc5\x02\xd3\x67\x3d\x9b\xa8\x52\x6c\x36\xbc\x4c\xea\xdb\xb6\x2c\x6c\x12\x2f\xd3\x61\x8a\xe7\x21\xc5\x10\xaf\x7d\x9d\x14\x92\xf6\xeb\x36\x0f\x29\x60\x6f\x76\x0b\x34\xbc\x83\x29\x5e\x04\xc4\x01\x00\x08\x48\x84\x29\x83\x80\xb0\xf0\xe4\xfd\x10\x61\x56\x71\x3a\x7d\x34\xd7\x04\xc2\xa5\x7f\x2d\xf2\xab\x1a\x6f\x25\x62\x5d\x13\xd8\x6d\xf3\x6b\x83\x18\x5e\xad\x96\x13\x92\x0e\x37\x95\x8c\x70\xfa\x65\xaa\xb0\xe1\xd8\x0a\xb2\x38\x5d\xb8\xc7\xd3\x02\xb0\xba\xee\xc2\x86\x67\x6b\xc9\xe2\x75\xe1\x1e\xaf\x57\xaf\x0b\x1b\x5e\xdf\xd2\x30\xf9\xce\xac\x3d\xbb\x4e\x80\x84\x77\xe8\xc6\x1d\x3b\x98\xf8\xe3\x71\x4f\x42\x58\xe6\x9d\x80\x7e\xa5\xbc\xea\x17\xcb\x0f\xcb\xcb\xd0\x1b\x79\xf9\x78\x89\xff\xca\xeb\xed\xc8\xcb\x7c\x26\x7e\x4a\x5e\xdf\x02\x00\x00\xff\xff\xa3\x97\xc1\x38\x52\x0b\x00\x00")

func _sqlMysql1000_initial_usersUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__sqlMysql1000_initial_usersUpSql,
		"_sql/mysql/1000_initial_users.up.sql",
	)
}

func _sqlMysql1000_initial_usersUpSql() (*asset, error) {
	bytes, err := _sqlMysql1000_initial_usersUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_sql/mysql/1000_initial_users.up.sql", size: 2898, mode: os.FileMode(420), modTime: time.Unix(1590262588, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __sqlMysql1001_initial_postsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x48\x28\xc8\x2f\x2e\x29\x4e\xb0\xe6\xc2\x23\x1b\x9f\x91\x59\x5c\x92\x5f\x54\x09\x57\x15\xe4\xe9\xee\xee\x1a\x84\x53\x5d\x7c\x49\x51\x66\x7a\x7a\x6a\x51\x7c\x69\x41\x4a\x62\x49\x2a\xc9\xda\x52\x52\x73\x52\x41\xda\x00\x01\x00\x00\xff\xff\x98\x2d\x4e\xcc\xb1\x00\x00\x00")

func _sqlMysql1001_initial_postsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__sqlMysql1001_initial_postsDownSql,
		"_sql/mysql/1001_initial_posts.down.sql",
	)
}

func _sqlMysql1001_initial_postsDownSql() (*asset, error) {
	bytes, err := _sqlMysql1001_initial_postsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_sql/mysql/1001_initial_posts.down.sql", size: 177, mode: os.FileMode(420), modTime: time.Unix(1590090229, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __sqlMysql1001_initial_postsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x55\x41\x6f\xda\x30\x14\xbe\xe7\x57\xbc\x1b\x8e\x56\xaa\xb1\x43\x35\x09\x71\x08\xc4\xd0\x68\xe0\x20\xc7\xa8\xed\x29\xc9\x1a\x8f\x5a\x02\x07\x05\x33\xa9\xff\x7e\x72\xe2\x24\x6e\x9a\xa2\x8d\x6d\x52\x35\x2d\xe2\x60\xbf\xef\x7b\x7e\xcf\x2f\xdf\x47\x86\x43\x18\xda\x8f\x53\xfe\xe0\x90\x1f\xd5\x51\xaf\x7d\x1a\xae\x81\x79\xd3\x25\x86\x60\x0e\xf8\x3e\x88\x58\x04\x49\x09\x27\x63\x67\x46\xb1\xc7\xb0\xc1\x4d\x14\x90\x03\x90\x88\x2c\x01\x21\x15\x1a\x8d\x5c\x38\xc9\xa3\xd8\x4a\x9e\x01\x09\x19\x90\xcd\x72\x09\xde\x86\x85\x71\x40\x66\x14\xaf\x30\x61\x57\x3a\x41\x09\xb5\xe3\x09\x3c\x3e\xa5\x05\x1a\x7d\xfa\xec\x36\xe4\x12\x7d\xcc\xa5\xe2\x52\x25\xc0\xf0\x7d\xc5\x3f\x14\xe2\x7b\xaa\x78\x02\x4a\xc8\xe7\xb2\x52\x5f\xa1\x8c\x7f\x4b\x4f\x3b\x05\x83\x8f\x83\x32\xeb\x74\xc8\x52\xc5\xb3\xf8\xeb\xf3\xb9\xf6\xde\xca\x4a\x95\x2e\xb7\xe7\x47\x95\xee\x0f\xe8\xc6\x6d\x88\xb3\x0d\xa5\x98\xb0\x98\x05\x2b\x1c\x31\x6f\xb5\xd6\x60\x2e\xa1\x4a\xec\x85\xab\x6b\x15\xfc\x82\x76\xea\x2c\xdd\x8e\x3e\x5f\xb7\xa4\x0b\xda\x13\x5b\xd3\x60\xe5\xd1\x07\xf8\x82\x1f\x00\xe9\xb7\x51\x16\xd4\xbb\x7a\xd2\xc8\x2c\x5a\xa0\x19\x29\x6a\x96\x2d\x68\x4f\x0e\xd9\xbb\xd7\x14\xdd\x17\xb2\x77\x2d\xc5\xbe\x30\xb2\x77\xaf\x29\xd5\x29\xd6\xce\x75\x5c\xc0\x64\x11\x10\x3c\x09\xa4\xcc\xfd\x29\xf8\x78\xee\x6d\x96\x0c\x66\xb7\x1e\x8d\x30\x9b\xec\x52\x25\xe4\x68\xec\x9c\x53\x6d\xfc\x24\x8e\x2a\x2f\x9e\xfb\xd5\xdb\xa0\x17\xaa\x58\x1f\x12\x9f\xcd\x7a\x77\x62\xaf\x87\xd8\x27\xf6\xb7\xd4\xf5\x17\x85\x9b\x98\x37\x10\xef\xb9\x7a\xca\xb3\x04\xb8\x3c\xed\xd1\xa0\xea\x6a\x70\x35\xc8\xf8\x8e\x2b\x3e\x70\x7b\x2e\x50\x73\xec\x53\x2e\xf4\x48\xf3\x1e\x51\xb3\xfc\x07\x0c\x54\x53\x50\x77\xc8\x3d\xc8\xaf\x19\x6e\x38\x84\x0f\x7b\xb1\x2d\x52\x25\x72\x09\x53\xbe\x15\x32\x52\xa9\xe2\x7b\x2e\x95\x71\x23\x0d\x16\x0b\x4c\xdf\xf4\x63\xac\x0a\xb1\xdd\xf2\x22\xae\xae\x6d\xd9\xd3\x24\x9e\xa7\xc3\x14\xcf\x43\x8a\x61\xb3\xf6\x75\x52\x48\x9a\xcf\xd1\x3c\xa4\x80\xbd\xd9\x2d\xd0\xf0\x0e\xa6\x78\x11\x10\x07\x00\x20\x20\x11\xa6\x0c\x02\xc2\xc2\x57\xde\x8f\x30\x2b\x39\xad\x0c\xea\x67\x02\xe1\xd2\xbf\x16\xd9\x55\x85\x1b\x25\x80\x85\xe7\xbb\xec\xba\x0c\x1b\x4a\x63\xe6\x97\x14\x13\x36\xa4\x46\x38\x2f\x49\x26\x6c\x48\xb6\x80\x2c\x52\x1b\xee\xf0\x52\x95\xd8\x4d\xb7\xe1\xba\x33\x4b\x4a\x16\xaf\x0d\x77\x78\x9d\xf3\xda\xb0\xe1\x75\x8d\x0b\x93\x17\x96\xec\x98\x72\x02\x24\xbc\x43\x37\xee\xd8\xc1\xc4\x1f\x8f\x3b\x0a\xc2\x32\x6b\xf5\xf3\x27\xd5\x55\xfd\x7d\xfc\xb4\xba\x0c\xbd\x56\x97\x8f\x97\xf8\xbf\xba\xde\x85\xba\x26\x50\x7f\x0a\x7e\x4b\x5c\x3f\x02\x00\x00\xff\xff\xf4\x9a\x36\xc1\x00\x0b\x00\x00")

func _sqlMysql1001_initial_postsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__sqlMysql1001_initial_postsUpSql,
		"_sql/mysql/1001_initial_posts.up.sql",
	)
}

func _sqlMysql1001_initial_postsUpSql() (*asset, error) {
	bytes, err := _sqlMysql1001_initial_postsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_sql/mysql/1001_initial_posts.up.sql", size: 2816, mode: os.FileMode(420), modTime: time.Unix(1590351207, 0)}
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
	"_sql/mysql/1000_initial_users.down.sql": _sqlMysql1000_initial_usersDownSql,
	"_sql/mysql/1000_initial_users.up.sql":   _sqlMysql1000_initial_usersUpSql,
	"_sql/mysql/1001_initial_posts.down.sql": _sqlMysql1001_initial_postsDownSql,
	"_sql/mysql/1001_initial_posts.up.sql":   _sqlMysql1001_initial_postsUpSql,
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
	"_sql": &bintree{nil, map[string]*bintree{
		"mysql": &bintree{nil, map[string]*bintree{
			"1000_initial_users.down.sql": &bintree{_sqlMysql1000_initial_usersDownSql, map[string]*bintree{}},
			"1000_initial_users.up.sql":   &bintree{_sqlMysql1000_initial_usersUpSql, map[string]*bintree{}},
			"1001_initial_posts.down.sql": &bintree{_sqlMysql1001_initial_postsDownSql, map[string]*bintree{}},
			"1001_initial_posts.up.sql":   &bintree{_sqlMysql1001_initial_postsUpSql, map[string]*bintree{}},
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
