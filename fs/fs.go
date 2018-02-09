package fs

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

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

// HTTP returns http handler for file server
func (f *FS) HTTP(dir string) http.Handler {
	httpfs := afero.NewHttpFs(f.write)
	return http.FileServer(httpfs.Dir(dir))
}

// MkDirAll creates dir with 0777
func (f *FS) MkDirAll(path string) error {
	return f.write.MkdirAll(path, os.ModePerm)
}

// RemoveAll removes all from write file system
func (f *FS) RemoveAll(path string) error {
	return f.write.RemoveAll(path)
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

// Create creates new file
func (f *FS) Create(path string) (io.WriteCloser, error) {
	return f.write.Create(path)
}

// CopyTree make a deep copy of directory
func (f *FS) CopyTree(from, to string) error {
	return afero.Walk(f.read, from, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("walk error for %s: %v", path, err)
		}
		path = strings.TrimPrefix(path, from)
		if path == "" {
			return nil
		}
		if fi.IsDir() {
			err = f.MkDirAll(to + path)
			if err != nil {
				return fmt.Errorf("unable to create dir: %v", err)
			}
		} else {
			var in io.ReadCloser
			var out io.WriteCloser
			in, err = f.Open(from + path)
			if err != nil {
				return fmt.Errorf("unable to open file: %v", err)
			}
			defer func() { err = in.Close() }()

			out, err = f.Create(to + path)
			if err != nil {
				return fmt.Errorf("unable to create file: %v", err)
			}

			_, err = io.Copy(out, in)
			if err != nil {
				return fmt.Errorf("unable to copy file: %v", err)
			}

			return out.Close()
		}
		return err
	})
}
