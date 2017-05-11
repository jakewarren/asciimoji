// Code generated by go-bindata.
// sources:
// emojis.json
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

var _emojisJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x5a\x6f\x4f\x1b\x57\xd6\x7f\xdf\x4f\x71\x95\xa7\x52\x67\x22\x02\x91\xfa\xe2\x91\xf2\x28\x7a\xbe\x40\xa4\x48\xfb\xa2\x52\x15\xc2\x64\xb0\x2f\xf6\x2c\xe3\x19\x6b\x6c\xc7\xeb\xee\x6e\x35\xc6\x0d\x60\x42\x49\x9a\x7a\xb0\x49\x9c\x82\x31\x4e\x9b\x2c\x04\x97\x1a\x62\x88\x8d\xa5\xa9\x14\xa9\x60\x42\xd4\xad\xb4\xdb\xee\x9b\x99\x33\x5a\x55\xea\x9b\xf9\x0a\xab\x19\x13\x18\xda\x5e\xe7\x50\x55\x2b\x21\x73\x67\x6c\xfd\xce\xbd\xe7\x9e\x7f\xbf\x73\xef\xb5\xb7\x08\xf9\xf3\x5b\x84\x10\x72\x6e\x9c\x66\xd2\xaa\x16\x4e\x9c\xbb\x44\xae\xf9\x6f\x08\x39\x27\x86\xa4\xf0\x39\xff\xe1\xfa\x40\xef\x57\x34\xa6\xfe\x51\x3a\x77\x89\x9c\x83\x99\x09\x0e\x8a\x79\xd0\x3b\x50\xcc\xf3\x4e\x76\xc5\xfb\xdd\x5f\x07\xde\x80\x37\xa6\x89\x4c\x44\xce\x99\x9c\x24\x02\x71\x26\x27\x79\x14\x56\x22\x71\x6e\xe0\xf5\xc3\x68\x2a\x99\x64\x4d\x54\xef\xfc\x05\xf4\x0e\x0a\xd3\x83\xa1\xda\x98\x9c\x61\x60\xbd\x68\xfe\xbd\xfc\xe2\x19\x6a\x76\x4a\x84\xca\xac\x19\x95\xca\x1c\xcc\xde\x83\xf9\x36\xcc\xde\xe3\xa1\xf4\x19\x6a\x6a\xaa\x12\x3e\x59\xef\xc5\x8b\xff\xcb\x02\x2f\xcc\x72\xe4\xe0\xa1\x59\x27\x07\xe5\xbd\xf9\x83\x8a\x59\xe7\x2f\xbf\xdc\x1c\xea\x6e\x1e\x4c\x75\x3b\xdd\xce\x10\xe8\x0b\xdd\x8e\xff\x41\xba\x9d\xff\xfa\x9e\xfd\xde\x5a\x11\x95\x88\x96\x39\x51\x4b\x4c\x64\x9a\xab\xbe\x7c\x43\x30\x1b\xa0\x2f\xa3\x60\x35\x4d\x4d\x47\x29\x1b\x6d\xa5\xcc\xc1\x97\x3b\x64\x6f\xbe\xbb\x0e\x5f\xee\xf0\xb0\x72\x07\xb5\x87\x54\xd4\x18\x88\x7b\x86\xf9\xec\xa0\x62\x37\x6a\xe6\xb3\xbd\x02\x16\x6b\x4c\x96\xe2\x4c\x3c\x27\x37\x07\xfa\xb2\xdd\xa8\x81\xbe\xbc\x57\x70\x72\x73\xc4\x6d\x6e\x12\x28\x3c\x87\x42\x16\x0a\xcf\xb1\x32\xa2\xa9\x08\x5b\x44\xb6\x1a\x14\x91\xad\xa2\x40\xa5\x48\x94\x8a\x1a\xd3\x5d\x1f\xa2\x22\xc9\xa8\x2c\x86\xc6\x69\x86\x32\x50\x2e\x0a\xc2\xff\x60\x60\xd2\xd1\x80\xf5\x48\x61\x55\x49\xa6\x94\x30\xd5\x12\x49\x51\x61\xed\xbd\x6d\x2c\x70\x30\xa5\x83\xde\x84\xc9\x16\xe8\x5b\xbc\x6d\x14\x31\xa2\xc6\xc4\x10\x8d\x8b\x72\x8c\xe5\x49\x6e\xeb\x63\xd0\x9b\x76\xee\x3e\x2e\x52\xa9\xea\x68\x82\x05\x45\x06\xc9\xfb\x64\x90\xe0\x80\x34\x69\x94\xa5\x45\x8e\x80\xbe\x7c\x50\xd9\x2f\x11\xe1\x68\xc0\x3b\xb9\xb9\xaf\xeb\x18\xdc\x90\x46\x69\xfc\x44\xb7\xfe\x23\x65\x59\xff\x77\xb7\x38\xc8\xcf\x7b\xb9\x24\x3f\xff\xdd\x2d\xd4\xbc\x7d\xc0\x50\x2a\xc9\x9a\xfb\x8b\x27\xdc\x0f\x9f\xfe\x63\x83\x07\x7d\xee\xc5\x13\xd0\x6f\x83\x7e\x1b\x07\x1b\x30\x88\x84\x3a\xca\xd2\x0c\x18\x35\x77\xfb\x0e\x18\x35\xd4\x64\xc3\x62\x4c\xc9\xa8\x29\x16\x98\x6d\x14\xc9\xc1\x92\x59\x27\xd6\x17\x33\xbd\x01\xd6\xa8\xc2\xa2\x12\x62\x69\xc0\x36\x0c\x0e\x66\xef\xc2\xfc\x92\x00\xf3\x4b\x1e\x22\x81\x85\x27\xb0\xb0\x8a\x02\xa6\xa2\x9c\x96\x92\x51\x29\x79\xa2\x8e\x70\x5a\x62\xaa\xe3\x44\x0e\x0e\x3d\xae\xd1\x44\x82\x06\x32\x5a\x82\x19\x6c\x39\xb7\xb9\xe5\x36\xb7\xdd\xe6\x16\x12\x5b\x63\x05\x45\x28\xe5\x09\xe8\x1d\x02\xf7\x27\x51\x48\x2a\x2b\xf4\x71\x50\x34\xbc\xb0\x57\x34\xf6\x6e\xe3\x26\xa5\xca\xb2\xa8\x8d\x4a\xb2\x7c\xb2\xe2\xb7\x19\xe0\xd7\xba\x1b\xdd\x5b\x6f\x7b\x1f\x9c\xf7\xf1\x72\xfb\x65\xc7\xfb\xdf\xbd\xc5\x1f\x7f\x71\x1d\x27\x53\x49\x86\x44\x8d\x65\x1e\x60\xac\xc1\xf4\x63\x30\x9e\xba\xad\x69\x5f\xc3\x1d\xef\xb3\x95\x3f\x7e\x8f\x95\x91\x16\x15\x56\x30\x77\x72\x6d\xce\x6d\xeb\xaf\x0a\x66\x83\x77\x77\xf3\x38\xc4\xd4\xa8\x4c\x93\xe2\xa8\x4c\xfb\xa4\xb7\xe3\x34\xe6\x25\x35\x4f\xca\x8d\xd7\x32\xce\x9c\xe4\xc2\x29\x1a\xd8\x94\x50\x54\x94\x65\xaa\x44\xfa\x3a\x95\xb1\x66\x6d\xd4\xc1\x58\x23\x07\x8b\x7b\xf3\x60\x3c\x05\x7d\xb9\x9b\xe5\x61\x26\x67\xae\x5c\xe6\x40\x2f\x1c\xff\x5d\xc0\x4d\x20\xca\x34\xb3\xfc\xe7\x6e\xb3\x03\xf9\x47\x38\x23\x4b\x29\x8a\xca\x80\x32\xd7\x87\x87\x39\xb3\x2e\xa8\xfc\x90\xb9\x8e\x01\xa3\x4a\x38\xa5\xd1\x13\xc5\xc4\xa9\x96\xa0\x37\x69\xf0\xd5\x38\xa5\x71\x55\x89\xa8\x92\xc2\xf4\x93\x6f\x27\xa0\x58\x11\xa0\x58\xe1\x89\x6d\xa0\x12\x3b\xbd\x29\xb1\x2a\xc4\xc3\x69\xcf\x98\x60\x7a\xca\x6c\xf0\x87\xd3\x28\xb0\x3f\xd1\x50\x2a\x29\xa9\x0a\x22\x6e\xb9\xcd\xe7\x60\x3c\x02\x63\x05\x0a\x3a\x39\x63\x7c\x1f\x63\xd7\x33\xdc\xfe\x24\xec\xfc\x6d\x7f\xd2\xda\xfe\x84\x1f\x4e\x5d\xbc\xf8\x6e\xe8\x5d\x14\xa2\x14\x89\x32\x21\xad\xe6\x94\x97\x99\xbb\xba\xd0\xb3\x3c\xab\x39\x85\xc2\x94\xd5\x34\x33\x05\x73\x50\xee\x40\x71\xc9\xe7\x73\xb8\x20\xce\xa6\x49\x50\xd0\x2f\x43\xbe\xf2\xb2\xca\x71\x1c\x71\xb2\x2b\x50\x34\xbe\xff\xe8\xa0\x0c\x45\x03\x4b\x14\xc7\x34\x89\x2a\xe1\x3e\x01\x80\x03\x63\xdd\xac\xc3\xbc\xc7\x6f\xc0\x58\x3f\xe5\xf5\xbd\x42\xd7\x58\xe7\x8e\x7e\xe0\x8d\x71\x2b\xd2\xd4\x34\xd3\x54\xec\xdc\xa3\xfd\x12\xcc\x56\xf7\x4b\x76\x0e\xe7\x8f\x63\xa9\xd0\xb8\x3a\x36\x76\xe2\x30\x91\xe4\x18\xcb\x3f\xad\xb5\xcf\xad\x56\xcb\xda\x58\xfa\xf1\xc1\x8c\xb5\xb1\x74\x1c\x84\xad\x56\x1b\x2b\xca\x2b\x30\x06\x4e\x5e\xb0\xf9\xe1\xab\x45\x28\xdc\xe5\xac\x8d\x25\xc1\xda\xc0\x6d\x75\x84\x2a\x49\x99\xc6\x44\x25\x90\xaf\x25\x2d\xc0\xbb\x54\x45\x0d\xc9\xac\x98\xd9\x93\x54\xb7\xbe\xca\xa1\x64\x45\xd5\x04\xcb\xf2\xad\x56\xcb\xb3\x27\x02\x0f\x3f\x16\xe0\xe1\xc7\xc4\x6a\xb5\x91\x06\x15\x91\xc6\x02\x95\x8c\x57\x7d\x50\x66\xd6\xe2\xcc\x86\x93\x7b\x7e\x98\x77\x72\xcf\x79\x27\x5b\xfd\xb1\xf0\x25\x4e\x42\x2c\xc6\xd4\xc0\xeb\x79\x17\x0d\x01\x8a\xc6\x59\xe6\x2d\x4b\xc9\x24\xdb\x69\xcf\x3b\xb9\xe7\xa0\x77\xfc\x99\xe6\xe6\x60\xf6\xde\xf9\x4b\x6e\xbb\xe6\xee\x2e\x42\xf9\x73\x1c\xbc\x98\x48\xd0\x40\x23\x45\xa1\x1a\xb3\x0e\x83\xd9\xbb\xe4\xa0\x02\xf3\x4b\x7e\x6b\xa1\x37\xc4\x99\x4f\x6f\x11\x7e\x65\x36\x10\xac\xd3\xfa\x39\xb8\xbb\x9b\x87\x52\xde\xc9\x3d\x85\xfb\x93\x5e\x62\x27\xbf\x61\x75\x6a\x52\x62\x6e\x32\x94\x3e\xf3\x90\xa6\x75\x77\x77\x11\xdb\x6e\x88\x68\x94\x06\xcc\xc8\x7f\x94\x94\x48\x40\x7f\x51\x2a\xcb\x2c\x1f\xe7\x88\xd9\x80\x62\x81\x1c\xe6\xa1\x58\xb8\xc1\x13\x27\x37\xe7\x4c\xa0\xba\x49\x91\x14\x2b\x2c\x05\xb3\x16\x06\x29\x2a\x86\xd5\xd4\x38\x65\xc1\x59\xad\x96\x93\x5d\x39\x0e\x43\x3d\x3b\x25\xc7\xf1\xfc\x2a\x8f\xda\xef\xa8\x18\x15\x19\x02\xbe\x7f\xcc\x8d\x80\xde\x19\xe1\xff\xb5\x85\x43\x8a\xc7\x59\x59\xe6\xfb\xc7\x1c\xb1\xb6\x3f\x01\x63\x9b\xfc\x34\xd5\x00\x63\x1b\x8d\x99\xa4\x5a\x22\xa2\x2a\x8a\xe8\x0d\x03\x7b\xe7\x09\x8b\x88\x9a\xc6\x92\xe8\x15\x7f\xc4\xbe\xfb\x00\x4d\xcf\xfa\x36\x3a\x16\x6a\x38\x08\xcf\xa0\x8e\xe7\xa8\x46\x45\x29\xd0\xf6\xcc\x04\xa6\x9f\x16\x6f\x32\x89\xfc\x5e\x09\xf4\xce\x5e\xc9\x4b\x95\x28\xa1\x52\x24\x3a\x26\x31\xe1\x06\x85\x41\x7e\x68\x78\x98\x1b\x14\x06\x51\x68\xa9\x48\x40\xc9\xa9\x08\xb3\x6f\xe1\x64\x6b\x6e\xbb\x02\x45\xc3\xab\x42\xbc\x42\xc4\x70\xdb\x15\xde\xc9\xa2\x14\x25\x29\xb2\xca\x56\x80\x93\xad\x9a\x8d\xc3\xfc\x8d\x10\x0f\x0b\x15\x1c\x5c\x92\x6a\x0a\x65\x66\xa2\xd5\x12\x67\x3d\x5e\x3b\xbf\x3f\x63\x6f\x1a\xfb\xb7\x79\x6f\x0c\xfa\x4c\x57\xdf\xbf\x4d\xbe\x79\xf2\x4d\xc3\x6e\x3c\xb0\x1b\x53\x7b\xba\x37\x9a\xb2\x1b\x0f\x30\x22\xc7\xa5\x44\x3f\xd5\x10\xb7\x53\x25\xe6\x57\x6e\xa7\x8a\x55\x89\x4c\x15\x85\x65\xcb\x1c\x39\xa8\x1c\x35\x89\xfd\x11\xca\xb1\x7d\xc0\x37\x96\xa6\x7e\x07\xc3\xa3\x46\x3d\x60\x64\x75\xda\xc3\xee\x93\x10\x9c\xdc\xdc\xcf\xe6\xec\xe4\xe6\x78\xbf\xd8\xe3\xc8\x41\xd9\xd1\x27\xc8\x41\xc5\xfb\xc6\x1b\xe2\x97\x93\x88\x6a\xcc\x16\xa7\xc7\x9a\x84\x5f\x68\x4a\x40\x52\xa8\x1e\x7c\x52\x94\xc7\x69\xa0\x56\xfa\xc5\x8b\xfe\xad\x31\x28\xac\x42\xa1\xf1\xfa\x73\x25\x30\x19\x28\x94\x83\x5f\x9e\x61\x46\x9a\xca\x64\x6c\xb6\xf1\xe8\x17\xeb\x45\x32\x37\x1f\x3c\x2d\x7d\x20\x6a\x81\xb6\xce\xd1\x33\x63\x6d\x46\xfd\xe7\xc2\x88\xc7\x10\xa0\xa0\x7b\x7f\xa5\xc9\xf3\x97\x9c\xdc\x73\x77\x77\x11\x25\x5e\x65\xf1\xc6\x2b\x9c\x59\x27\x57\x89\x59\xbf\x82\xb3\x0a\x55\x1d\x67\x9a\xf7\x46\x45\xb0\x36\x2a\xd8\x7a\xa1\x4f\x44\x82\x85\x9a\x17\xe2\x70\x49\xc0\xc3\xe9\x7b\x78\x00\x0b\x35\xbb\x51\x83\x85\x1a\xee\xf0\x20\x26\x46\xa4\x10\x72\x93\x82\x3d\x1c\x98\xc9\x79\x14\xab\x34\xe9\xee\x2e\x0e\x92\xf3\xc4\x6d\xd7\x88\xdb\xae\xb8\xbb\x8b\x03\x68\xb1\x7d\x5c\x7c\xc8\x5c\x87\x62\xa5\xd7\xa2\x2b\x56\xf8\x21\x73\x9d\x7c\x78\xb6\x66\x4e\x8c\x52\x16\xb8\xdf\xff\x18\x31\xeb\xfc\x10\x0e\x88\xd5\x94\x39\xe2\x4d\x28\x10\x29\x91\xcc\xd0\x0c\x65\x45\xf5\x23\x6b\x42\x41\xa9\x4a\x82\xcd\x03\x3c\x7e\x61\xb5\xe6\xac\x9d\x2d\x62\x3d\x6b\x1c\x0d\x91\xd4\x51\x51\x63\x27\x96\x90\x49\xc5\x62\x99\x60\x8d\x2e\x4b\x21\x49\x4d\xf5\x49\x4b\xd5\xfd\xc9\x1f\x2a\xfb\x93\x87\xb8\x88\xab\x46\x13\x51\x76\x49\x4e\xcc\x1d\x77\x3b\x67\xee\xe0\x4e\x2e\xd4\x18\x2b\x88\x1d\xd1\x2c\x0c\x48\x5c\x95\xd8\x3c\xb0\x47\x11\x9c\xdc\xd3\x33\x50\x84\xb8\xaa\xb2\x2c\x67\xcf\xf0\xbb\x36\xd9\x7f\x4e\xfb\xcd\x1b\xd4\x22\x35\x31\xd2\xaf\x0f\xea\x25\xc5\xe3\x6a\xdd\xa3\x80\xff\x6e\x57\xce\xe4\x30\x1a\x15\x65\x39\x13\xac\x2f\x8f\xde\xfc\xba\xbc\xaf\x37\x84\xaf\x51\x99\x46\x93\x54\x66\xe9\x74\xba\xd5\x61\xb5\xda\xd6\xda\x17\x28\x4c\x55\x96\xc3\x52\x28\x50\xf2\xfa\x4f\x0c\x1b\xb8\x8f\xea\x36\x78\x98\x7d\x3c\x94\x83\x62\x41\x80\x62\x01\xb5\x57\xec\x23\x8c\x97\x9b\x9c\xd9\x70\xd7\xb2\x6e\x73\xd3\x5d\xcb\xde\xe0\x5f\xa1\x58\x9f\x87\x37\x10\x78\x08\xab\x4a\x84\x19\x02\x9c\x5c\xdb\xd3\xea\xce\xf4\xde\x7c\x77\xdd\xda\x99\xb6\x5a\x6d\x64\xc7\x3d\x11\x55\x53\x4c\x92\x62\xdc\xe7\x40\x5f\x86\xd9\x09\xd0\x97\x79\x30\x1e\xe2\x00\xdf\x54\x51\x39\xb9\x8f\xd0\x55\x54\x22\xca\x32\xc6\xcb\x23\xc2\xc8\x65\x14\x84\x3a\x1a\x2c\xb7\x98\x25\xf1\xff\x99\x8d\x5e\xe8\x7c\x55\xe8\xfd\xbf\x81\xdb\xf7\x94\x14\x92\xc2\x2c\x43\x64\x5e\xbe\xe0\xa0\x58\x71\x9b\x9b\x5e\xa6\xc3\x49\xd1\xe2\x9a\x94\xa0\xcc\xf6\x8c\xc7\x85\xbd\xf0\x42\x9c\xdc\xd3\x5e\x7b\xd8\xda\xfe\x04\x0b\xad\x51\x25\x1c\x2c\x4a\xa5\x88\x74\x93\xa6\xd8\xb9\x54\xe0\x2e\x08\x17\x78\x01\x95\x4b\x5f\x57\xbc\x0c\x13\xfb\x59\x81\xeb\xb6\x6b\xc2\x6f\x28\x6d\xdf\x74\x62\xe4\xd3\x07\x67\xa2\xfc\xaa\xe0\x4c\x94\xf9\xdf\x74\x1d\x22\x49\x45\x8d\x19\x28\xac\x8d\x9a\xbb\x7d\xc7\xda\xc0\x1d\x18\x24\xa3\xa2\x32\x1e\x68\x1b\xf9\xcf\xa7\x1a\xb8\x49\x96\x95\x0e\x0f\x73\x23\x17\x46\x70\x55\x4c\x52\xd2\xd8\xf6\x72\x59\x10\x04\xe1\x32\x6e\xba\xaa\x1c\x3e\x3d\xbb\xde\x8b\x04\xab\xd5\x05\xa5\x32\xf7\xd3\xbd\x0a\xcc\x7e\xfa\xd3\x3d\x9c\x79\xfb\xfb\x17\x66\xf7\xdf\x7d\x4b\xd0\xa1\xb0\x4a\x9c\xdc\x1c\xe7\xed\xe4\x05\x67\xa2\xec\x31\xc0\xb3\xa1\x1f\x2f\x21\xa5\xf4\x3d\x5f\x5c\x3d\x22\x21\x47\x02\xbd\xc4\xf7\x15\x56\x5a\x3a\x2a\x25\x92\xcc\x76\xb8\x57\x34\x8d\xbc\x7a\x36\xc2\xe3\x4f\xe3\xd3\x6a\xa0\x05\x94\x11\x59\xa6\xe1\xb6\x5a\x9c\xdb\xda\xbd\xea\xb6\x76\x79\xb7\x85\xba\x92\x94\x4e\x06\xce\x27\xd2\xaa\xa6\x49\xa7\x0e\xe5\xd3\x54\x4c\xf6\x39\xde\x83\x99\x05\xb7\xd5\x81\x99\x85\x77\x70\x7a\x61\xe7\x99\x99\x85\xc3\x3c\xcc\x2c\x60\x50\xd8\xeb\x1f\x1e\xe6\xc8\x49\xc1\x86\xf2\x90\x8c\xca\x6c\xd7\xbe\x6f\x6f\xde\xb3\x37\x4b\xe4\xaa\xbd\x55\xdf\xaf\xec\x35\xc9\x15\x7b\x6b\xc5\xde\x7c\x60\x6f\xe6\xfd\x77\xf6\x56\xd9\xde\x44\x65\xd8\x8c\x9a\x1a\x97\xc2\x01\x77\xef\x3d\xfe\xba\x58\x3b\x77\xdf\x6a\xb5\xfc\xa3\x42\x1a\x38\x63\xee\x9d\x1d\x5a\xad\xb6\x9d\xbb\x8f\x91\xf9\x01\x0d\xd8\x7a\x8c\x86\xa5\xa4\xc8\xbc\x2b\x03\x33\xdb\x1e\x15\x87\xd9\xea\xde\x22\xcc\x56\x79\x30\xd6\x61\x66\x1b\x25\x45\x95\xc2\xa3\x54\x63\x1a\xc8\x7b\x3c\xe1\xcc\xfa\xc0\xc0\xc0\x80\x59\xe7\x09\xf7\x1e\xca\x4c\x3e\x50\x63\xa3\x12\x6b\xae\xd7\xcc\x55\x73\xe7\x82\x59\xbf\x6e\xa2\x3c\x27\x21\x53\x1a\x0f\xd4\xb7\xbd\x88\x18\xf0\x25\x66\xc0\x71\x72\xbb\x9c\xdb\xa9\x5e\x74\x3b\x55\xe2\xe5\x0b\xac\x69\x1e\x63\x47\xc5\x78\x5c\x52\x28\xb3\xa7\xe7\xe8\x65\x4f\x00\xcc\xb7\xdd\x4e\xd5\xd1\xcb\x67\xac\x02\x02\x5e\x9b\x66\x13\x98\xbc\x00\xa5\x3c\xae\x69\xcf\xbc\x86\x00\x33\x0b\x83\x50\x42\x59\xba\x47\x57\x02\xda\xa5\xc1\xfb\x2e\xbd\xbb\x9a\x7d\x22\x57\x5b\xbf\x68\x36\xb0\x81\xeb\xb4\xaa\x53\x9a\xd6\x37\x2a\x8e\xa8\x23\x58\xe0\x90\xca\x6c\x24\x79\xa5\xb0\xe0\x15\xc2\xc4\xbf\x19\xe7\x0f\x8f\x1c\xd5\xbf\x47\x70\x01\xe6\x97\xc8\x59\xef\x42\x65\x82\xc7\x6f\x89\x54\x28\xc4\xb6\x18\xce\x6a\x4e\x79\x2b\x39\x33\x79\xe8\x5b\xb9\xc0\x6c\xb5\xfb\xb8\xfb\xb8\x9b\x15\x5e\x8f\x70\xd4\xf4\x94\x84\x5f\xb9\xd1\x35\x4a\xc5\x64\x9f\x84\xce\x7d\x3b\xe1\xc9\xd3\x7d\xa9\x59\x1c\xe3\xd7\xce\x40\x55\xdd\xc6\xee\x41\xc5\xda\xbe\xd3\x5d\xf7\xff\xb9\x8d\x0e\xb9\xfa\xce\x1f\xae\xbc\xff\xff\x18\x49\xbd\x7b\x75\x01\x51\x19\x2a\xb2\xdc\xe3\x43\x6e\xbf\x04\xf3\xbb\xfb\xa5\x0f\x71\x7e\x16\x8b\xb1\xee\x80\x9a\xab\x02\x2e\xaa\xfd\xae\x37\x7d\x42\xaa\x32\x96\xea\xc3\x2b\xcc\xba\x93\x2b\x21\x0f\x08\x4e\x62\xde\xa9\x30\xc8\x64\x5c\x67\xbe\x74\x12\x96\x12\x62\x3c\xae\xf5\x39\xe9\x39\xcb\xbd\x06\x4d\x0d\x8d\x33\x2f\x07\x0d\x0f\x0f\x0c\x0c\x71\x23\xc2\x08\xef\x8f\x70\x5c\x55\x1a\xeb\xdb\xf1\xe3\xcc\x55\xb7\xbd\x62\xae\xe2\x1c\x8c\x8a\xcc\x48\x64\x6f\x16\xec\x46\xcd\xde\x3c\xea\x4c\xbc\x75\xfd\x3f\x01\x00\x00\xff\xff\xf3\x42\x8b\x56\xca\x31\x00\x00")

func emojisJsonBytes() ([]byte, error) {
	return bindataRead(
		_emojisJson,
		"emojis.json",
	)
}

func emojisJson() (*asset, error) {
	bytes, err := emojisJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "emojis.json", size: 12746, mode: os.FileMode(420), modTime: time.Unix(1494456002, 0)}
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
	"emojis.json": emojisJson,
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
	"emojis.json": &bintree{emojisJson, map[string]*bintree{}},
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
