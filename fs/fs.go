package fs

import (
	"os"

	"github.com/spf13/afero"
)

// FS is a composite read and write file system
type FS struct {
	read  afero.Fs
	write afero.Fs
}

// New creates composite file system
func New(readFS afero.Fs, writeFS afero.Fs) *FS {
	return &FS{
		read:  readFS,
		write: writeFS,
	}
}

// MkDirAll creates dir with 0777
func (f *FS) MkDirAll(path string) error {
	return f.write.MkdirAll(path, os.ModePerm)
}

// ReadFile reads file from read file system
func (f *FS) ReadFile(path string) ([]byte, error) {
	return afero.ReadFile(f.read, path)
}

// WriteFile writes file to write file system with 0777
func (f *FS) WriteFile(path string, data []byte) error {
	return afero.WriteFile(f.write, path, data, os.ModePerm)
}
