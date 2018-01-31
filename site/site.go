package site

import (
	"fmt"

	"github.com/gorook/rook/fs"
)

type Site struct {
	Pages []*Page
}

func FromDir(f *fs.FS, dir string) (*Site, error) {
	files, err := f.TreeList(dir)
	if err != nil {
		return nil, err
	}
	fmt.Println(files)
	return nil, nil
}
