package fs

import (
	"io"
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

// Open opens file for reading
func (f *FS) Open(path string) (io.ReadCloser, error) {
	return f.read.Open(path)
}

// TreeList returns file list of directory tree
func (f *FS) TreeList(dir string) ([]string, error) {
	var filelist []string
	err := afero.Walk(f.read, dir, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !fi.IsDir() {
			filelist = append(filelist, path)
		}
		return nil
	})
	return filelist, err
}
