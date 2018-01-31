package main

import (
	"github.com/gorook/rook/config"
	"github.com/gorook/rook/fs"
	"github.com/gorook/rook/site"
	"github.com/gorook/rook/theme"
)

const (
	configFileName = "config.yml"
	contentDirName = "posts"
	themeDirName   = "_theme"
)

type application struct {
	fs     *fs.FS
	config *config.SiteConfig
	site   *site.Site
	theme  *theme.Theme
}

func newApplication(fs *fs.FS) *application {
	a := &application{
		fs: fs,
	}
	return a
}

func (a *application) init() error {
	var err error
	a.config, err = config.FromFile(a.fs, configFileName)
	if err != nil {
		return err
	}
	a.site, err = site.FromDir(a.fs, contentDirName)
	if err != nil {
		return err
	}
	a.theme, err = theme.FromDir(a.fs, themeDirName)
	return err
}

func (a *application) build() error {
	return nil
}
