package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorook/rook/config"
	"github.com/gorook/rook/fs"
	"github.com/gorook/rook/site"
	"github.com/gorook/rook/theme"
	"github.com/rjeczalik/notify"
	"github.com/spf13/afero"
	"github.com/yanzay/log"
)

const (
	configFileName = "config.yml"
	contentDirName = "posts"
	themeDirName   = "_theme"
	staticDirName  = "static"
	publicDirName  = "public"
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
	if err != nil {
		return err
	}
	a.theme.SetConfig(a.config)
	a.theme.SetTags(a.site.Tags.All())
	return nil
}

func (a *application) build() error {
	a.renderAll()
	err := a.saveAll()
	if err != nil {
		return err
	}
	err = a.copyStatic()
	return err
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

func (a *application) renderChanged(path string) {
	page := a.site.Pages[path]
	a.rendered[page.Path] = a.theme.RenderPage(page)

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
		dir := publicDirName + "/" + path
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

func (a *application) copyStatic() error {
	return a.fs.CopyTree(themeDirName+"/"+staticDirName+"/", publicDirName+"/"+staticDirName+"/")
}

func (a *application) startServer() error {
	newWatcher("posts/...", a.contentChanged)
	newWatcher("_theme/...", a.themeChanged)
	newWatcher("config.yml", a.configChanged)

	err := a.build()
	if err != nil {
		return err
	}

	handler := a.fs.HTTP(publicDirName)
	log.Infof("Listening on %s", a.config.BaseURL)
	return http.ListenAndServe(":1414", handler)
}

func (a *application) contentChanged(e notify.EventInfo) {
	start := time.Now()
	wd, err := os.Getwd()
	if err != nil {
		log.Errorf("unable to get current dir: %v", err)
		return
	}
	path := strings.TrimPrefix(e.Path(), wd+"/")
	log.Infof("Content changed: '%s'. Rebuilding...", path)
	err = a.site.Rebuild(a.fs, path)
	if err != nil {
		log.Errorf("unable to rebuild site: %v", err)
	}
	a.theme.SetTags(a.site.Tags.All())
	a.renderChanged(path)
	err = a.saveAll()
	if err != nil {
		log.Errorf("unable to save site: %v", err)
	}
	log.Infof("Done in %s", time.Since(start))
}

func (a *application) themeChanged(e notify.EventInfo) {
	fmt.Println(e.Path())
}

func (a *application) configChanged(e notify.EventInfo) {
	fmt.Println(e.Path())
}
