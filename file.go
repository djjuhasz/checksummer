package main

import (
	"hash"
	"os"
)

// FileInfo extends os.FileInfo with a HashFunc and Hash value
type FileInfo interface {
	os.FileInfo
	FullPath() string
	Hash() []byte
	HashFunc() hash.Hash
}

// File represents an OS file
type File struct {
	os.FileInfo
	fullPath string
	hash     []byte
	hashFunc hash.Hash
}

// FullPath returns the file path + file name
func (f *File) FullPath() string {
	return f.fullPath
}

// SetFullPath sets FullPath
func (f *File) SetFullPath(val string) {
	f.fullPath = val
}

// Hash returns the hash value
func (f *File) Hash() []byte {
	return f.hash
}

// SetHash sets hash
func (f *File) SetHash(val []byte) {
	f.hash = val
}

// HashFunc returns the Hashing function used
func (f *File) HashFunc() hash.Hash {
	return f.hashFunc
}

// SetHashFunc sets hashFunc
func (f *File) SetHashFunc(val hash.Hash) {
	f.hashFunc = val
}

// NewFile returns a File
func NewFile(fi os.FileInfo) *File {
	f := new(File)
	f.FileInfo = fi

	return f
}
