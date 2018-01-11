// Code generated by go-bindata.
// sources:
// tmpl/model.html
// tmpl/x_helpers.html
// tmpl/x_helpers_test.html
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

var _tmplModelHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x58\x6b\x6f\xdb\xc6\x12\xfd\x6c\xfe\x8a\xb9\x82\x6f\x2e\xe5\x28\xb4\x03\x5c\xf4\x83\x0b\xa5\x70\x2d\x1a\x15\xaa\x58\x8e\x25\xf7\x81\xa2\x48\x56\xe4\xd0\xda\x9a\xda\x95\x77\x97\x92\x0d\x9a\xff\xbd\x98\xdd\xd5\x5b\x76\x19\x34\x2d\x5a\x01\x09\xf4\x98\x9d\x73\xe6\xcc\x8b\xeb\xb2\x4c\x31\xe3\x02\xa1\x31\x91\x29\xe6\x8d\xaa\x0a\xa6\x2c\xb9\x63\xb7\x08\x65\x09\xd1\x95\x7b\x7f\xc9\x26\x08\x55\x15\x04\xc7\x47\x6f\x9e\x7f\xbd\x0e\x9e\xe0\x5c\xa6\x08\xb7\x28\x50\x31\x83\x29\x8c\x1e\xc1\xfa\xbd\x45\x01\x4f\xc1\x13\xf8\x57\xa7\x0f\x97\xfd\x21\xc4\x9d\xee\x30\x5a\x7c\xf7\x14\xbc\x7e\xc1\xf7\xd1\x71\x10\x1c\x1f\x43\x37\x83\x47\x59\xc0\x9c\x09\x03\x46\x02\x3e\x18\x14\x29\x8c\x70\xcc\x66\x5c\x16\xea\x94\x6c\xae\x72\x64\x1a\x21\x51\xc8\x0c\x02\x83\xa4\xd0\x46\x4e\x20\xe3\x39\x02\x17\x60\xc6\x5c\x43\xca\x15\x26\x46\xaa\x47\x3a\xc0\x44\xba\xb0\x36\x63\x84\x09\x9a\xb1\x4c\x35\xbd\x57\x48\x28\x6c\x26\x79\x0a\x72\x86\x6a\xae\xb8\xe1\xe2\x96\x38\x28\x48\x28\xd4\x39\x37\x63\x7b\x6a\x15\xb3\x14\x18\x05\x01\x9f\x4c\xa5\x32\x10\x06\x8d\x6c\x62\x1a\x41\x23\x65\x86\x8d\x98\xc6\x63\x7d\x9f\x37\x02\x0a\xb8\x2c\x41\x31\x71\x8b\x70\x78\xd7\x82\xc3\xd9\x69\x1b\xa2\xf7\xa4\x55\xd4\xb5\x47\x35\x09\x4e\x76\x8d\xb2\x3c\xbc\xab\xaa\xc5\xa1\x37\x40\x21\x57\x55\xd0\xb4\x8a\x94\xa5\x3f\x45\x39\xaa\x2a\x50\x38\x55\xa8\x51\x18\x0d\x0c\x94\x9c\xbb\x90\x71\x65\x37\x64\xa3\x1c\xbd\xb1\xa1\xf7\x81\x79\x9c\xe2\x8e\x1f\x6d\x54\x91\x18\x28\x5f\xa6\x7a\xc1\x31\x4f\x97\x4c\xbd\xe9\xe1\x2c\xf2\x05\xe3\x3f\x0d\x09\xa0\xaa\xe0\xd3\x6f\x5a\x8a\x53\x8a\x67\x16\x9d\xcb\xbc\x98\x08\x07\xd6\xf8\xb4\x1d\x9b\xcc\x32\x8d\x06\xb8\x30\x41\xce\x27\xdc\xbd\xab\x5c\x09\x08\x8d\xca\x00\x03\x81\xf3\xdd\xe8\x6b\x05\x9c\x15\x22\x81\xb0\x2c\xa3\x6b\x4c\x90\xcf\x50\x55\x15\x1c\x6d\xb9\x6a\x7a\xa0\x30\x1d\xc1\x91\xbe\xcf\xa3\xce\xb7\x4d\x08\x73\xa6\x8d\xfb\xbe\xdb\x21\x4e\x5f\xfd\xbf\x05\xa8\x14\xfd\x93\xaa\xe9\xd5\x4a\xa4\xd0\x06\xb4\x99\x18\x68\x43\xa3\x7b\x39\x88\xaf\x87\xd0\xbd\x1c\xf6\xf7\x73\x0a\x97\xdf\x7a\x35\x9f\x80\x5b\x88\x8f\x99\xfd\x4c\x5c\x7e\x38\xeb\xdd\xc4\x83\x17\x4c\x67\x2c\x2f\x90\x4c\x5d\x99\x28\xd4\x8e\xd8\x69\x1b\xd2\x51\x14\x3f\x60\x12\x12\xa1\x96\xed\xea\xd5\x31\xa6\x6e\x29\x7b\x4d\x7b\x88\x67\xf6\xc8\x7f\xda\x20\x78\xee\x63\x71\xce\x4c\xa1\x04\x9c\x58\x8f\xf6\xdb\x2a\x58\xfb\x5e\xa1\x8e\x7a\x2b\x5d\xd2\xb0\xe9\x33\x75\x33\x4d\x6d\x07\x0a\xc0\x07\xae\x6d\xe3\xfc\x95\xf9\x72\x70\x6b\xf9\x6a\x01\x4f\x5d\x92\x9a\x2e\x41\x7b\xf3\x73\x73\xd5\x39\x1b\xc6\xfb\xe1\x07\xf1\x70\x21\x58\x61\xbd\x7b\x9d\xa9\x98\x7f\xfc\x2e\xbe\x8e\x09\xa1\x0d\xdf\x38\xd1\x3f\xbe\x2c\xb9\xf7\xe0\x25\x07\x62\xd7\x5c\xd7\x91\xc4\x5d\x08\x67\x4b\xdc\xe5\x48\xff\xa9\x52\x3f\x3e\xa6\xb4\x92\x51\x21\xf8\x7d\x81\x2e\x78\xc5\xb8\x1d\x10\x0a\x41\x48\x03\x99\x2c\x44\xda\x02\x49\x13\x6f\xce\x35\x02\x37\x9e\xad\x06\x6e\xa2\xda\xfa\xff\x1b\xfb\x05\xfa\x97\xd0\xb9\xb9\xea\x75\xcf\xa9\x0a\xbe\x8f\x7f\x86\x65\x41\xf8\xac\xad\xd9\xc3\x62\x0c\xff\x23\xfa\xeb\x82\x8b\xf4\x6f\xeb\x2e\x02\xfb\xec\xde\x1a\xc4\xbd\xf8\x7c\x08\x47\x70\x71\xdd\x7f\xbf\x9f\xc3\x6e\x1b\x11\x6b\x27\xeb\x87\x02\xd5\xe3\xb5\x9c\x7b\x69\xb7\xfa\x45\xc9\x79\x34\x48\x98\x08\xbd\xe4\x3a\x61\x62\x55\x02\x5e\xa3\x9e\x64\x29\xb0\x3c\x6f\x81\x54\xc0\x40\x17\x23\xda\x2b\x32\xdb\xa7\x95\x86\x4c\xc9\xc9\x97\x91\x8b\x70\x37\x5b\x81\x70\x7f\xf9\x75\xcb\x70\x4f\x37\x58\xed\x4e\x6b\x89\xd7\x08\x16\x75\xb5\x41\x27\x72\x2b\xb3\xdd\x86\x13\x78\xf5\x6a\xeb\x37\xbf\x59\xdf\xc1\xc9\x6e\x05\x6a\x34\x2d\xc8\x26\x26\x8a\x89\x50\x16\x36\x12\x26\x68\x3e\xdc\x53\x1e\xdc\xa3\x8e\x3f\x3e\x2a\x0c\x08\x09\x16\xa8\xd1\xf4\x35\xfb\x12\x99\x4d\x3c\x1b\xe3\xeb\xb6\xc5\x1a\x4c\x15\x17\x26\x0b\x1b\xd0\xeb\xbe\xef\x0e\xe1\xbf\x69\xa3\xb5\xcf\x45\x73\xad\x33\x76\x40\xf6\x46\xb5\x1f\xa5\x7f\x71\x41\x63\x7d\x0f\x8c\x73\xb2\x8e\x93\x62\x86\x0a\x28\xe3\x61\x73\xcd\xf1\x5e\xb5\xe1\xe4\x99\xdf\x3d\xb7\x85\x41\x15\x36\x17\x65\xbe\x31\x43\x6c\xb1\xdb\x4a\xaf\x33\x2d\x76\x48\x92\xbb\xe8\x3c\x97\x1a\xbd\x7f\x7a\x65\xd2\xff\x70\x89\x0f\x66\x23\x84\x35\x84\xb6\x33\xf1\xad\xb4\xdb\x49\x5f\xef\xe7\xb1\xc5\x65\xc5\xc7\x4a\x6f\x03\x66\xd3\x29\x8a\x34\xb4\x65\x75\xb4\x21\xca\x46\xc9\x78\x2f\xae\x61\x3b\x98\x63\x8d\x87\x86\x2f\xd7\xa9\x0e\xf0\x99\xd1\x16\x92\x34\x67\x59\x86\x09\x3d\xd9\xd7\x5b\x60\x9d\xb8\x17\x0f\xe3\xcf\x1b\x79\xa8\x8b\xdc\x3c\xb3\x51\x68\xec\x1d\x6c\x57\xc3\xc1\x81\x57\xed\xa0\x0a\x82\x83\xd5\xa6\x28\x72\x13\x5d\xaf\x71\x5e\x2e\x8b\x73\x59\xd0\x9d\x69\x8c\x20\x8a\xc9\x08\x15\xcd\xc0\xcf\x9a\x79\x7f\x28\xa4\x45\xd8\x9c\x79\x89\x05\xad\x27\x9b\x1f\x77\xe7\xfd\x9b\xcb\x61\x78\xd4\x7c\x69\xea\xbd\xb0\x25\x36\x7a\xa7\xbd\xda\x11\xaf\x2c\x95\xe7\x6a\x79\xa7\xa7\x36\x4a\x32\xa6\x42\xd4\x90\x8c\x31\xb9\xd3\xb6\xa7\x48\x2e\x6e\x70\xa2\x5d\x91\xa2\x48\x70\xb1\x68\x17\xf7\xbc\x16\xd0\xff\x74\x1b\x04\x6e\xfe\xa7\x81\xa7\x11\xf9\x3a\x13\x7e\x5d\xce\x79\x9e\x83\x14\xf9\x23\x8c\xd0\xc3\x51\x85\x65\xc0\x60\xf0\xa1\x07\x0a\x73\x7b\x99\xcc\x18\xcf\x0b\x85\x30\xb6\xbd\xa4\x23\x77\x03\xa2\x8d\xe6\x1e\xd7\x20\x61\x9a\x9e\x45\x18\x8c\xa4\xcc\xed\x55\x96\x42\xb3\xde\x9d\xd7\xa8\x5e\xf6\x5c\x90\xcf\xb5\x01\x3a\x09\x08\xa3\x6e\x22\xe3\x9f\xba\x83\xe1\x20\xf4\x9f\xde\xd6\xec\x07\xbf\x04\xde\xfa\x7b\xcc\x8c\xd1\x3d\xdb\x17\x51\xbd\xa7\x83\x2f\x94\x7b\x0f\xfb\x8e\x1e\xcb\x04\xcf\x7d\x29\x2c\x69\x7b\x2b\xfb\x57\x02\xd7\x24\x20\xd8\xa4\xee\xcc\x59\xba\x09\x9b\x74\xd1\xb6\x43\x2e\xf0\xb8\x8d\xfd\x35\xef\xf0\x07\x68\x7a\x76\xdd\x68\x34\x0e\xdb\xed\x66\xbb\x83\x6a\x82\x2f\x7c\x84\xcb\x9b\x75\x73\xc7\x0a\xca\x60\xff\x92\x73\x40\x9e\xea\x86\xc9\x8a\x61\xdf\x2d\xbc\x2d\x8a\x6e\x0d\xd6\xe7\xe8\xbc\x84\xab\xbf\x04\xd4\x60\xb9\x5c\xb5\x1e\xec\x19\x9e\x65\x89\x22\xad\xaa\x20\xf8\x3d\x00\x00\xff\xff\x3f\xa0\x37\x0d\xfe\x12\x00\x00")

func tmplModelHtmlBytes() ([]byte, error) {
	return bindataRead(
		_tmplModelHtml,
		"tmpl/model.html",
	)
}

func tmplModelHtml() (*asset, error) {
	bytes, err := tmplModelHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/model.html", size: 4862, mode: os.FileMode(420), modTime: time.Unix(1515658612, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplX_helpersHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x58\x5b\x8f\xdb\x36\x13\x7d\xe7\xaf\x98\xe8\xe1\x83\xe4\x38\x72\x3e\x24\x29\x50\x17\x7e\xc9\x0d\xdd\xa2\xd9\x4d\xb3\x6e\x5e\x8a\x02\xa1\x2d\xca\xcb\x54\xa2\xbc\x24\x9d\xc5\xc2\xeb\xff\x5e\xf0\x26\x52\x92\x2d\xd9\x7b\xe9\x3e\x78\x6d\x6a\xe6\xf0\xcc\xf0\x0c\xc5\xe1\x76\x9b\x91\x9c\x32\x02\xd1\x15\x29\xd6\x84\x8b\x68\xb7\x43\x6b\xbc\xfc\x07\xaf\x08\x6c\xb7\x90\x7e\x36\xdf\xcf\x71\x49\x60\xb7\x43\x68\x32\x7a\x71\xf8\xef\x39\xba\x83\x77\x55\x46\x60\x45\x18\xe1\x58\x92\x0c\x16\xb7\x50\x56\x19\x29\x56\x84\xc1\x1d\xba\x03\xfb\xf7\xfe\x02\xce\x2f\xe6\xf0\xe1\xfd\xd9\x3c\x75\x63\x77\xe8\x79\x0f\xf6\x68\x82\x10\x2d\xd7\x15\x97\x10\xa3\x28\xc3\x12\x2f\xb0\x20\x13\x71\x5d\x44\xcd\x9f\x93\x8c\xd3\x1f\x84\x47\x28\x22\x6c\x59\x65\x94\xad\x26\xdf\x45\xc5\x22\x14\x71\x92\x17\x64\x29\x23\x14\x09\xc9\x29\x5b\x89\x08\x45\x92\x96\x24\x42\x28\x5a\x51\x79\xb5\x59\xa4\xcb\xaa\x9c\xac\xaa\x17\xe2\xba\x78\x61\x60\x26\xe5\xad\x9e\x22\x41\x68\x32\x81\x4b\x99\xcd\x69\x49\x60\xcd\xab\x1f\x34\x23\x02\x32\x92\xe3\x4d\x21\xe1\xf2\x8f\xdf\x61\x7e\xf6\xe9\x03\xe4\x15\x2f\xb1\x44\xcb\x8a\x09\x59\x9b\xcf\x20\xfa\xff\x9b\xe9\xcb\xd7\xd3\x97\x6f\x22\x8d\x43\xca\xb5\xbc\xd5\x8f\x70\x51\x54\x37\x1e\x47\xd1\x11\x20\x2b\x58\x10\x50\x18\x34\x23\x9c\x64\xca\x85\x6d\x8a\x42\xa1\x03\x65\x82\x70\x49\x2b\x06\x94\xc9\x0a\xe4\x15\x01\x17\x7e\x8a\x7e\x60\x1e\x80\xcf\x34\x5e\xaa\xbe\x6f\xbb\x6b\xa7\x56\x6b\x7e\xbb\x56\x1c\x28\x16\x44\x74\x16\x40\xa5\x7c\x32\x81\xf3\x4d\x51\x7c\x2c\x2a\x2c\x7f\x7a\x5d\x9b\x8a\xeb\x22\x0d\xc6\x91\x54\x38\xa1\x61\xdb\xc0\x01\x5d\xea\xcc\x77\x70\xcc\xb0\x87\xb1\x66\xad\xc7\x0e\xe4\x6d\x55\x15\x1d\x08\x35\xe8\x01\xb4\x49\xe3\x91\x73\x3e\x63\xfb\x02\xd1\xa3\xde\xdd\x18\x35\x1f\x3a\x00\xbb\x6e\x4d\x7f\x35\xe8\xdd\xb5\x89\x56\x8e\x7f\xa8\xdc\xbf\xe0\x9b\xdf\x2e\x2f\xce\x6b\x6f\x25\xcc\xf4\x0b\xbe\xf9\x44\x84\xc0\x2b\x0b\xe0\x8c\xda\x0f\x87\x6b\x2f\xc8\x1c\x2d\xd7\x05\x29\x09\x93\x58\x49\xa5\xbb\xb6\xdd\xe2\x9a\x4c\xe0\x13\xe6\xe2\x0a\x17\x7a\x76\x25\xb5\x20\xf5\xf9\x86\x2d\x21\x66\xc1\x50\x12\x9a\xc7\x09\xc4\x7f\xfd\xbd\xb8\x95\x64\x0c\x84\xf3\x8a\x27\xb0\xd5\x62\xc4\x30\x32\xc5\x86\x68\x0e\x2c\xfd\x8a\x0b\x9a\xc1\x16\x61\x98\xc1\xff\x58\x6a\xc1\x77\x88\x13\xb9\xe1\xcc\x84\x6c\x61\x63\x9c\xa0\x9d\xa6\xf5\x27\x2b\x87\x89\x8d\x42\x66\x0d\x8f\x78\x01\x86\x5a\x62\xa8\xc1\x16\x11\xce\x61\x3a\x33\xd3\xd5\xb6\xf1\x62\xec\x39\x25\xc8\x91\x9d\x29\x2f\x98\xcd\x80\xd1\xc2\xf1\x24\x9c\x5b\x6e\x5f\x71\xb1\x21\xc7\x24\x4b\x1b\xaa\x34\x99\x7d\x25\xd5\xbf\x83\x64\xd1\x1c\x9e\xf9\xfc\xd8\x79\x18\x2d\xc6\x7a\xde\x3a\x43\x8e\x9f\x1b\xd6\xbb\xd2\x12\xb3\xa3\xd2\xa2\x0c\x63\xc1\x97\x6a\xe7\x20\x3c\xc7\x4b\xb2\xdd\xf9\xa4\x98\xe5\x6a\x95\x1c\xcd\xc1\x26\x0b\xa7\xce\x3d\xf9\x45\x8f\x3d\xd3\x29\xf1\x64\x4d\x52\x1c\x41\xd0\x1e\x35\x88\xdd\x7b\x53\xb5\xe1\x5c\xe4\x1a\xc5\x03\xf8\x54\x4b\xbe\x21\x41\xb0\x36\xc4\x3e\xdd\xd7\xc2\x77\x3b\xcf\x69\xca\xef\x93\xbe\xdb\xbb\xc2\xe5\xb4\x63\x47\x8b\x3f\xb7\x18\xfb\xd4\xef\xf0\x07\xe4\xdf\x94\x58\x2f\xa9\x47\x13\x99\x45\x0c\x55\xb6\xbf\x0a\xdb\x74\x46\x0d\x3e\x0f\xa9\x43\x07\x72\x64\x21\x36\x8a\xa0\x9f\xd5\x29\x65\xf0\xd1\xaf\xdf\xa9\x75\xe0\x04\xa9\x5c\x02\x98\xa7\xaa\x04\x57\x07\xe6\xd5\x75\x52\x15\xf4\xd4\x80\x79\xf5\x85\x62\xd3\x23\x47\xeb\x9f\xb2\x43\xea\x37\xc8\x27\x69\xff\x20\x99\x47\xd3\xbd\xc6\x1b\x56\x7d\x93\xc8\x28\x60\xf2\x10\xc5\x1b\x88\xfb\xe8\xfd\x30\x9f\x53\xb4\x7e\xc6\xee\xa9\x74\x23\x39\xe5\x50\x43\x3c\x91\xca\x9d\xc8\xf5\xf1\xee\x14\x8d\xf7\x48\x5c\x9f\x0e\x43\x51\xa9\x81\xa3\x05\xbe\x50\xde\xfb\xf4\xad\x61\x4f\x92\xf7\x01\x22\x8f\x26\x6e\x05\x37\xac\xed\x06\x8b\x91\xa7\xf1\x10\x65\x6b\x84\xfb\x08\xfb\x10\x99\x53\x64\xfd\xd6\xae\xd0\xa9\xaa\xd6\x1a\x53\xf6\x0e\xe0\x89\x35\xad\x9b\x86\x47\xd2\xb4\xee\x38\x42\x29\xa9\x81\xa3\x35\x5d\x77\x8d\x7b\x85\xad\x1f\x9c\x24\xec\x03\x6c\x1e\x4d\xd8\x0a\x6e\x58\xd8\x0d\x16\x23\x4f\x63\x48\xd8\x42\xa9\xc6\x34\x31\xf1\x22\x41\x02\xdc\x2f\x91\xce\x39\x2d\x63\x31\x86\x6f\xd1\xb7\x04\x49\x5a\x8e\x9d\xc8\x74\x06\x3f\x63\x2e\x48\xac\xbf\x7e\xf9\xf8\xee\xd5\xab\x57\x3f\x8f\x41\x24\x4e\x8a\x7b\x94\x93\xe3\x42\x90\x56\x45\x98\xe8\x4c\x2b\xdf\x16\x59\x4b\x62\xad\xda\x39\x14\x6f\x7f\xed\x28\x14\x22\x81\x32\x2a\x29\x2e\x40\x48\x2c\xcd\x42\x8a\xcd\x42\x90\xeb\x0d\x61\x12\xc4\x12\x33\x91\x76\x88\x5b\xf9\xb4\xfa\xde\x7b\x54\x9e\x8d\x18\xd7\x1a\x7c\x9a\xca\x53\x85\xe7\x9a\xed\x13\xea\xee\xf8\x66\xd9\x82\x0f\xd7\x1d\xae\x77\x4e\xdf\xef\xc7\x2c\x71\x91\xe0\xb4\x01\x70\x4c\xeb\x59\x4f\x3d\x54\x64\x76\x0a\xab\x6f\x96\x0c\xd7\x51\xbb\xc3\xac\xa7\x1a\x2a\x24\x23\x8f\xf6\xb5\xc6\x9e\x17\x06\x4e\xd0\x52\x65\xc4\x22\xab\x6d\x65\xc4\x60\x06\xcb\x01\xc5\x1f\x64\xf6\x5f\x34\xbe\xdf\x05\x0b\x39\x9b\xe0\x63\xd7\x03\x27\x36\x84\xef\x82\x0d\xeb\x53\x09\xf3\x57\x7d\x1f\x0b\x2a\x96\xc3\x8a\xb4\x52\x9c\x57\xc1\xcd\x8f\x41\x17\x80\x81\x91\x9b\x4e\x5a\x42\xd3\x58\xb8\xcb\x99\x24\xbc\x3a\xd2\x1b\xaf\xb0\xef\x67\x1f\x65\xe8\xd7\x48\xd5\x56\x97\xdf\xd4\x6c\x02\xbb\xc4\x17\xe0\x61\x0f\xf3\x6f\x0a\x23\x31\x06\xeb\xad\x6a\x77\xe7\xa4\x6d\x58\x9a\x63\x65\x37\x9e\xe0\xb4\x1b\x18\xc6\xd4\x76\x1b\x49\xd0\x06\xe9\x58\xe8\xde\x58\x8c\x53\xe3\xec\xdb\x1b\xc9\x3e\x7b\xfd\x39\x85\x11\xed\x0b\xc3\xf5\x81\xdd\x40\x1a\x6d\x6a\xc3\x58\x05\x93\xbb\x7e\x35\x04\xe9\x09\xc8\xb9\xb6\x5a\xd7\xde\xa0\x0e\xf9\xd8\xff\x43\xa1\xe9\x23\xd2\x92\x13\x2c\x49\x18\x97\x3f\xb5\x79\xb3\x78\x61\xce\xca\x89\x3f\xbf\xeb\x58\x16\x7b\x63\xd1\x1e\xe1\x01\xae\x37\x8a\xae\xb5\xfa\x98\xc2\x68\xd1\x47\x5e\xbf\x65\xba\xe4\xfd\x6b\xd3\x9b\xc5\xd2\xdf\xa4\x27\xfe\xb0\xa6\x03\x90\x2a\x00\x7f\xe7\xde\x08\x43\xfb\x36\xdf\x87\xbd\x81\xec\xb3\x57\x1f\x53\x90\xdd\x48\xb6\x5b\xc2\xb2\xdd\x0e\xfd\x1b\x00\x00\xff\xff\xd1\x82\xba\x29\xc4\x19\x00\x00")

func tmplX_helpersHtmlBytes() ([]byte, error) {
	return bindataRead(
		_tmplX_helpersHtml,
		"tmpl/x_helpers.html",
	)
}

func tmplX_helpersHtml() (*asset, error) {
	bytes, err := tmplX_helpersHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/x_helpers.html", size: 6596, mode: os.FileMode(420), modTime: time.Unix(1515658630, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplX_helpers_testHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\x7d\x6f\xdb\x36\x1a\xff\x9f\x9f\x82\x21\xb0\xc0\xca\x54\xc5\x56\x8a\x3b\x40\xbb\x1c\xb0\x36\x29\xb0\x03\x9a\x1e\x56\x6f\x7f\xdc\x30\x2c\x92\x4d\xa7\xba\xe9\x25\x95\xe8\xe4\x02\x57\xdf\xfd\xc0\x87\xd4\x0b\x25\xca\x8e\x64\x2e\x71\x80\x14\x41\xea\x50\x0f\x9f\xd7\x1f\x1f\xfe\x64\x51\x9b\xcd\x92\xae\xc2\x84\x62\xf2\x85\x46\xb7\x34\xcb\x19\xcd\x19\x29\x0a\x74\xeb\x2f\xfe\xf4\x6f\x28\xde\x6c\xb0\xf3\x6f\xf1\xf9\xca\x8f\x29\x2e\x0a\x74\x7a\xf2\xa6\xff\xdf\xf7\xe8\x1b\x7e\x9f\x2e\x29\xbe\xa1\x09\xcd\x7c\x46\x97\x38\x78\xc0\x71\xba\xa4\xd1\x0d\x4d\xf0\x37\xf4\x0d\xcb\x7f\x17\x9f\xf0\xd5\xa7\x39\xbe\xbc\xf8\x69\xee\x94\x63\xdf\xd0\xf7\x5b\x74\x9f\x9c\x22\x14\xc6\xb7\x69\xc6\xf0\x04\x91\xa5\xcf\xfc\xc0\xcf\xe9\x69\xfe\x35\x3a\x5d\x66\xe1\x1d\xcd\x08\x22\x34\x59\xa4\xcb\x30\xb9\x39\xfd\x6f\x9e\x26\x04\x91\x8c\xae\x22\xba\x60\x04\x11\x1e\x59\x98\xdc\xf0\x4f\x61\x4c\x09\xb2\x10\x5a\xad\x93\x05\x9e\xd3\x9c\x7d\x66\xd9\x7a\xc1\x2e\xe3\x80\x2e\xf9\xe4\x09\xc3\x27\x52\xdc\x99\x5b\x78\x83\x58\x18\x63\xef\x1c\xf3\x89\xce\x85\xcf\xe8\xc4\x9d\xce\xfe\x6e\xe3\x19\xfc\x4c\x1b\x3f\x20\xf1\xcb\xfc\xbd\x85\xe8\xff\x6e\xe9\x82\x87\xef\x9d\xe3\xdf\x7e\x0f\x1e\x18\x9d\x5c\x6f\x88\x4f\xbc\x99\x7b\x66\x93\x80\x78\x2c\x5b\x53\x9b\x2c\x60\xc0\x81\xc1\x25\xf1\x48\xce\x32\xee\xa4\x4d\x28\xf1\x08\xb7\xf2\x66\x3a\x7b\x33\x9d\xcd\xa7\x53\x0f\x7e\xfe\x43\x6c\xb2\x22\xde\x6f\x33\xdb\xb5\xcf\x7e\x2f\xae\x2d\xc4\x1e\x6e\x29\xa6\xdc\x75\x9c\x43\x18\x78\x83\x7e\xc4\x57\xeb\x28\xfa\x29\x61\x7f\x7b\x8b\x31\xbe\xe6\xb9\xf0\x88\x6f\xa7\x71\xc8\x68\x7c\xcb\x1e\xc8\x35\x7a\x07\x22\xef\xd2\x34\xc2\xb5\x48\xa0\x88\xbc\x07\x91\x0f\x51\xea\x73\x3d\x52\x64\xa1\x88\x5c\x80\xc8\x67\x70\xba\xd2\xb2\x54\x44\x2e\x41\x64\x1e\xc6\xb4\x61\x88\x2a\x22\x1f\xf0\xcf\xfe\xfd\xbf\x3e\x7f\xba\xc2\x0d\x91\x95\x22\x52\x20\x0a\x25\x80\x40\x37\xe8\x47\xaf\x8e\x70\xf3\xab\x1f\x85\x4b\x0f\x43\x42\x31\x0c\x79\x78\xe6\x9e\x15\x36\x7a\xe7\x55\x51\xaa\x52\x7c\x44\x7c\x2e\x6c\xf4\xde\x6b\x06\xaa\x0a\xca\x41\x50\xe8\x08\xa5\x17\x5e\x23\x68\x55\x5a\x8c\x79\xb8\x2c\x63\x61\xa3\x4b\xaf\x8a\x5f\x95\xe5\x23\x1e\x07\x4c\x61\xa3\x0f\x5e\x99\x81\xc9\xb5\x2c\xed\xb5\x65\xa3\x02\x05\x36\xa6\x59\xc6\x03\xe7\x49\x71\x3e\xfa\x59\xfe\xc5\x8f\x26\x34\xb6\x50\xb8\x82\x4b\x47\xe7\x38\x09\x23\x8e\x51\xe7\x83\xcf\xf8\xb5\x2c\xb3\x50\xc1\x2f\x1f\x49\xf0\x3b\x17\x94\xde\x5e\x7e\x5d\xf3\x8b\x12\x95\x36\x0e\xac\xc6\x1c\x92\xa4\x0c\xb3\x2f\x14\xe7\x7c\x8d\x73\x3f\x8e\x48\xa9\x64\x22\x62\x99\x04\x16\x3e\x3f\xc7\xf2\x8f\x52\x8f\xd5\xab\x45\x28\x40\x77\x7e\x86\x69\xec\x8a\xba\x95\x3e\x97\xe1\xfc\x92\xc4\x65\x40\x95\x5f\xc7\x34\x76\xad\x1f\xc6\x84\x16\xbb\x36\xa6\x71\xc7\xa1\x45\x9a\x65\xbc\x03\xf0\x99\x45\x63\xcd\xd7\x35\xfc\xa3\xf2\x03\x2a\xd0\x59\xfc\x34\x67\xb9\x58\xc6\xd5\x02\x4b\x7c\x01\x67\x91\x0e\x94\xc8\xfe\x55\xeb\x44\x79\xba\xce\x16\x14\xcb\xb5\x8f\xee\xfd\x84\x5d\x66\x19\x0e\xd2\x34\x42\xc5\x06\x09\x15\x1e\x9f\x44\xee\x38\x2c\x88\x2d\xa7\x78\xb8\xea\x17\xbc\x29\x47\x29\xe1\x50\x90\xf3\x3d\xbc\xf2\xa3\x9c\xda\xa8\xb0\x15\x15\x61\xd2\xab\x64\x43\xfe\xa4\x0f\xc4\xe3\x56\xd6\x94\x14\x8a\x32\x40\x62\x5b\x97\x58\x72\x1d\x4d\x9b\x42\x37\xb1\x40\xab\x34\xc3\x7f\xd8\x98\x31\x9e\xa2\xcc\x4f\x6e\x28\x16\x19\xe3\x75\xf8\x79\x9d\x4c\x18\x73\xb8\x72\x1b\xf3\xd4\x77\xb2\x5b\x43\x82\xcb\x39\xad\x52\x30\x47\x78\x61\xfd\x80\x27\x35\x24\x2c\xfe\x3f\x63\x4e\x99\x54\x6e\xe9\x32\xcb\xd2\x6c\x35\x21\x75\x09\x5a\xba\x2c\x6e\x27\xcd\xf0\x39\xfe\xee\xce\xc6\xe5\xd4\xef\xee\x08\x2c\x31\xbb\xa1\x0f\x90\xb2\x05\x2e\xbf\xf2\x4c\x1a\x82\x09\x37\xc9\x87\xc4\x06\xe6\x80\xea\x5e\xb0\x54\x48\x49\x94\x16\x84\x64\x5f\x91\x55\xa9\x9a\x90\x40\x0f\x94\x89\x6b\xf4\xda\x76\x26\x52\x62\x3b\xbc\x1a\xd8\xd2\x9b\xad\xa7\x54\x56\x92\x30\xd2\xea\xdc\x1b\x2d\x37\x29\xb3\x15\xbc\x88\x40\xa0\x1b\x0e\x07\x88\x9c\x3c\x00\x18\x19\x65\xeb\x2c\xe9\xed\x41\xe0\x9e\x94\xb7\x76\x59\xad\xed\x09\x63\xca\xe4\x1d\x10\xfc\xbc\xf0\x93\x7d\x10\x78\xd2\x82\x60\x85\xb6\x3c\x5b\x80\x40\x98\x30\x9a\xad\xfc\x05\xdd\x14\x7d\x08\x3c\x6e\x62\xa1\x03\x3a\x15\x93\x85\xcd\x35\x03\x36\x30\x21\xdb\xf1\xc6\x7b\xbe\x68\x56\x03\xcc\xd4\x7a\x2a\x3b\x7f\x19\x08\x5b\x2d\x4b\xd4\x82\x39\x79\xb6\x18\xd5\xa6\x60\xfe\xb0\xee\x14\xae\x2a\xf4\x87\x4b\xbe\x2d\x43\x74\xf8\xf8\x18\x0b\x3f\x9a\x5b\x67\x69\x31\xff\x92\xae\xa3\x25\x16\x08\xc6\xc9\x3a\x8a\x48\x57\x95\xa2\x41\x44\x27\xb8\x5d\x53\x93\x6c\x08\xb2\x48\xbb\xa0\xfa\xd1\xcc\xd6\xda\x46\x2c\x1f\xdb\xb1\xb7\xee\x01\xd6\xaa\x93\x0d\xdb\x8a\x85\x45\x91\xdc\x2d\x66\xb5\xb8\xed\x98\xe4\x5a\x76\xd9\x53\x7a\xf3\xb1\xa6\x39\x6f\x89\xa8\x2f\x18\xf3\x4d\xba\x09\x81\x91\xad\xfa\xe3\xc8\x9d\xdc\x44\xc3\x56\x6d\x8f\x6a\xdb\xfc\x7e\xc3\x2c\xcd\x7c\x07\xdd\xda\x00\xc9\x84\xc2\xbf\x52\xcc\x81\x14\x93\xa7\xdf\x3c\xc1\x04\x98\x18\xa1\x97\x00\x8f\xbd\xc8\x25\xdc\x34\xab\x6d\x44\xdc\x35\x6b\x7a\x8a\xc2\x2a\xf9\xe5\xa1\x94\x52\x31\xf6\x02\x09\x25\xc0\xe1\x89\xe9\xa4\x62\x73\x7c\x57\x32\x42\x25\x2b\xb8\x8d\x26\x92\x02\x01\x0a\xc2\xda\x5b\x58\xc5\xec\xc4\xc8\x50\x0e\xd9\x6f\x41\x43\x1e\xe5\xd0\x21\xd3\x47\xa8\xff\x61\x93\x47\xf8\x7a\x71\x04\x75\x04\x60\x9a\x24\x8e\x4a\x3b\x1c\x4a\x1b\x35\x9d\x50\x47\xa8\x1e\xb3\x91\xea\x09\xe2\xae\xee\xb7\x1f\x23\x34\xe0\xfe\x61\x51\x42\x00\xfe\xb3\x10\x42\x8d\xe5\x51\x8d\x77\x1e\xc6\xd4\x2c\x1d\xe4\x1a\xcd\x7c\xe7\x08\x0f\x1d\x66\xb3\x37\xee\xdb\xc6\x43\x87\x57\x7e\x78\x10\xfc\x10\x70\xd3\xc3\x0f\x1b\x8f\xa8\xae\xd2\xfb\x89\x35\x0a\x41\x7b\x31\x46\x78\xc6\xa1\x36\x1a\xf1\x90\x83\x7b\xb5\x8d\x30\x86\xf1\x50\xbe\xa8\x98\x7a\x81\x7c\x91\xfb\xff\xd4\x7c\x51\xb1\x39\xbe\x6d\xe9\xf9\xe2\x3e\xe8\x3b\x51\xe0\x37\x9a\x41\x0a\x4c\x34\x11\xd7\x4f\x20\xf9\xd5\xa1\xfc\xb1\x57\xbf\x86\x3e\xd6\x89\x38\x6c\x0e\x09\x98\x38\x6c\x0e\x09\x4f\x8e\x47\x70\x48\x00\xab\x49\x0e\xa9\x34\xc8\xa1\x1c\xb2\x0d\x1e\xf5\x1c\xc1\xcc\xc6\xee\xdb\x9e\x93\x04\x5d\x18\x77\xbf\xca\x1b\xb3\x67\xeb\xa9\xe8\xae\xc6\xba\x1f\x15\xd5\xec\x10\xda\x78\xa6\x53\x38\xf5\xa0\x1e\x7c\x78\x11\xdc\x14\x16\xd4\xb3\x70\x53\x8d\xe5\x51\x4d\x1e\x0e\x50\x98\x25\xa7\xa0\xd2\x08\x3b\x9d\xb9\x67\xaf\x54\x74\x20\x15\x85\xec\x9b\xe7\xa2\x02\x27\x46\xbe\xac\x14\xf8\xd8\x8b\x7b\x8a\x93\x40\x6a\x6b\xa9\x8f\x02\x6d\x21\x9f\x21\x17\x9a\xcc\xdc\x33\x6b\x28\x07\x55\x4d\xbe\x40\x12\x2a\x80\xf1\xc4\x2c\x54\x35\xba\x47\x87\x32\xf2\xbd\x65\x8d\xbc\xd1\xb4\x53\xc2\xa0\x09\xb6\x5e\xde\x59\x83\x6d\x38\xfd\xec\x83\xdb\x63\x95\x1f\x0c\xe7\x14\x08\x38\x6c\xd2\x29\x8e\x4e\x8e\x60\x9d\x02\x9b\x26\x69\xa7\xda\x1c\x87\xf2\xce\x47\xa0\xb3\x43\xbf\x76\x6f\xb2\x7a\xee\xb8\xb3\x21\xee\x47\x1e\x75\x2d\x5e\xa7\x7d\xfa\x22\x88\xa2\x58\x05\xcf\xc2\x14\x75\xa6\x1f\xd9\x88\x15\xac\xcb\xb3\xb1\x66\xb9\xa2\x54\x6a\x8a\x2d\x3a\xaf\x8c\x71\x38\x63\x94\x35\x30\xcf\x19\x4b\xc4\x18\x61\x8d\x25\x52\xf6\xe2\x8d\xe5\xa1\xef\xea\x40\x8e\x2c\x4d\xfb\xdc\xf7\x16\x06\xb9\x12\xa2\x13\x29\x39\x98\x47\xb6\x5d\x78\x81\x4c\xb2\x04\xcc\x13\x73\xc9\xb6\xd9\x51\x6c\xb2\xc4\xa4\x11\x3e\xd9\xc4\xe4\x68\x46\x59\x01\xa2\x0b\xc2\x16\x4a\x95\x27\xd6\x35\x0c\x47\xb0\xcb\x7e\x10\xea\x0c\x38\x87\xcf\x32\x4b\x6c\x1c\x36\xcf\x2c\x5f\xae\x19\xc1\x34\x4b\xdc\x9a\xe4\x9a\xed\x96\x3a\x94\x6d\x0e\x42\xae\x8e\x77\x3e\x62\xbb\xd6\x73\xcf\x47\x34\xd1\xfd\xd8\x67\xdb\xc0\x4b\xe7\x9f\xe5\xfa\x78\x16\x06\xaa\x37\x3e\xb8\x7d\xcf\xd3\xf2\x51\x6b\x27\x57\x01\x24\x29\x5b\x53\x14\xc0\xc7\x86\xe8\x71\x00\xb9\x3a\x0a\x02\xb9\x2c\x9b\xde\x55\xaf\xed\x41\xfd\xc1\x07\xe1\x4c\x29\x5d\xbe\x80\x14\x04\xdd\xd3\x2d\xd5\x64\xf1\x86\x57\x73\x2e\x97\xad\x5e\x88\x0a\x5c\x7c\x02\x6b\x29\x08\xdc\x96\x73\x81\x0b\xce\x05\x81\xbb\xcd\xb9\x24\x65\x1a\x07\x5d\xc5\x43\xfe\x77\xbf\x87\x02\x96\xea\xec\xd2\xc7\xa2\x9d\x61\xb8\x61\xd0\xa7\xb8\xfe\xd2\x41\x49\xb4\x98\x61\x24\xd3\x41\x20\xef\xc6\x8f\xce\x79\x37\xd1\xeb\xe0\x2d\x46\xd1\x00\x53\x94\x84\x83\xa7\x6a\xc6\x85\x97\x66\x53\x5e\xf9\x3a\xd5\xeb\x99\xb6\xe6\x57\x8e\x76\x81\x2d\x57\x89\x3e\xf1\xed\xdd\x58\x49\x7f\x39\xd3\x54\x01\xca\x6d\x4a\x94\xc0\xd9\x56\x06\xa7\x53\x0a\x39\x59\x29\x86\xf4\x5e\x2d\x47\xe9\xb5\xd9\x82\x34\x7c\x7f\x64\x49\x1a\x0e\x77\x96\x82\x38\x12\xae\x2f\x09\xf9\x7a\x4f\x89\x52\x07\x29\x6d\xaa\x0c\xf2\x5d\x8c\x23\x69\x4a\xaf\xe6\xeb\x7d\xab\xf5\x88\x59\x4a\xfa\xe5\xde\xaf\x64\x5f\xfa\x6a\x36\xf9\x0d\x8f\x7b\xdc\xfd\x07\xdc\x2e\x4b\x36\xf2\xcf\x96\xba\xd2\x75\xbe\xbd\x16\x8d\xdf\xad\xa2\xcc\xc3\xb8\x7b\x37\x29\x1f\xc3\xb7\x8e\x25\xf0\xd1\x66\x85\xc4\xd4\x30\x16\x57\x5a\x45\x02\x15\x43\x0a\x55\x7b\x28\x0a\x06\xcf\xae\x8f\xc0\x81\x7e\x75\x7c\xd3\xab\x74\xb1\x30\xb6\xcb\x99\xa5\xbe\x2a\x14\x19\x09\x3c\x45\x2d\xca\x50\x7a\x23\xd9\x1d\x88\x64\x39\xe6\x42\x51\xb1\xa7\x8b\x45\xfc\xde\x6c\x68\xb2\x2c\x0a\xf4\xff\x00\x00\x00\xff\xff\x2b\x74\xdd\x7a\x7f\x40\x00\x00")

func tmplX_helpers_testHtmlBytes() ([]byte, error) {
	return bindataRead(
		_tmplX_helpers_testHtml,
		"tmpl/x_helpers_test.html",
	)
}

func tmplX_helpers_testHtml() (*asset, error) {
	bytes, err := tmplX_helpers_testHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/x_helpers_test.html", size: 16511, mode: os.FileMode(420), modTime: time.Unix(1512482539, 0)}
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
	"tmpl/model.html": tmplModelHtml,
	"tmpl/x_helpers.html": tmplX_helpersHtml,
	"tmpl/x_helpers_test.html": tmplX_helpers_testHtml,
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
	"tmpl": &bintree{nil, map[string]*bintree{
		"model.html": &bintree{tmplModelHtml, map[string]*bintree{}},
		"x_helpers.html": &bintree{tmplX_helpersHtml, map[string]*bintree{}},
		"x_helpers_test.html": &bintree{tmplX_helpers_testHtml, map[string]*bintree{}},
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

