package main

import (
	"github.com/gorook/rook/config"
	"github.com/gorook/rook/fs"
	"github.com/gorook/rook/site"
	"github.com/gorook/rook/theme"
	"github.com/spf13/afero"
)

const (
	configFileName = "config.yml"
	contentDirName = "posts"
	themeDirName   = "_theme"
)

type appOption int

const (
	appDefault appOption = iota
	appRenderToMemory
	appInMemory
)

type application struct {
	fs       *fs.FS
	config   *config.SiteConfig
	site     *site.Site
	theme    *theme.Theme
	rendered map[string]string // {path: rendered page}
}

func newApplication(opt appOption) *application {
	var readfs, writefs afero.Fs
	switch opt {
	case appDefault:
		readfs = afero.NewOsFs()
		writefs = readfs
	case appRenderToMemory:
		readfs = afero.NewOsFs()
		writefs = afero.NewMemMapFs()
	case appInMemory:
		readfs = afero.NewMemMapFs()
		writefs = readfs
	}
	return &application{
		fs: fs.New(readfs, writefs),
	}
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
	a.renderAll()
	return a.saveAll()
}

func (a *application) renderAll() {
	a.rendered = make(map[string]string)
	for _, page := range a.site.Pages {
		a.rendered[page.Path] = a.theme.RenderPage(page)
	}
	for _, ipage := range a.site.IndexPages {
		a.rendered[ipage.Path] = a.theme.RenderIndex(ipage)
	}
	for _, tag := range a.site.TagPages {
		for _, tpage := range tag {
			a.rendered[tpage.Path] = a.theme.RenderIndex(tpage)
		}
	}
}

func (a *application) saveAll() error {
	for path, content := range a.rendered {
		dir := "public/" + path
		err := a.fs.MkDirAll(dir)
		if err != nil {
			return err
		}
		err = a.fs.WriteFile(dir+"/index.html", []byte(content))
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *application) startServer() error {
	return nil
}
