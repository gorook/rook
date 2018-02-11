package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/rjeczalik/notify"
	"github.com/spf13/afero"
	"github.com/yanzay/log"

	"github.com/gorook/rook/assets/newsite"
	"github.com/gorook/rook/config"
	"github.com/gorook/rook/fs"
	"github.com/gorook/rook/site"
	"github.com/gorook/rook/theme"
)

const (
	configFileName = "config.yml"
	contentDirName = "posts"
	themeDirName   = "_theme"
	staticDirName  = "static"
	publicDirName  = "public"

	dateTimeFormat = "%Y-%m-%d %H:%M:%S"
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
	local    bool
}

func newApplication(opt appOption) *application {
	var readfs, writefs afero.Fs
	var local bool
	switch opt {
	case appDefault:
		readfs = afero.NewOsFs()
		writefs = readfs
	case appRenderToMemory:
		readfs = afero.NewOsFs()
		writefs = afero.NewMemMapFs()
		local = true
	case appInMemory:
		readfs = afero.NewMemMapFs()
		writefs = readfs
	}
	return &application{
		fs:    fs.New(readfs, writefs),
		local: local,
	}
}

func (a *application) clean() error {
	return a.fs.RemoveAll(publicDirName)
}

func (a *application) init(addr string) error {
	var err error
	a.config, err = config.FromFile(a.fs, configFileName)
	if err != nil {
		return err
	}
	if a.local {
		a.config.BaseURL = fmt.Sprintf("http://%s/", addr)
	}
	a.site, err = site.FromDir(a.fs, a.config, contentDirName)
	if err != nil {
		return err
	}
	a.theme, err = theme.FromDir(a.fs, themeDirName)
	return err
}

func (a *application) prepare() {
	a.theme.SetConfig(a.config)
	a.theme.SetTags(a.site.AllTags())
}

func (a *application) build() error {
	err := a.clean()
	if err != nil {
		return err
	}
	a.prepare()
	a.renderAll()
	err = a.saveAll()
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
	page := a.site.ByPath(path)
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
	siteStatic := fmt.Sprintf("%s/", staticDirName)
	themeStatic := fmt.Sprintf("%s/%s/", themeDirName, staticDirName)
	publicStatic := fmt.Sprintf("%s/%s/", publicDirName, staticDirName)

	err := a.fs.CopyTree(siteStatic, publicDirName+"/")
	if err != nil {
		return err
	}
	return a.fs.CopyTree(themeStatic, publicStatic)
}

func (a *application) startServer(addr string) error {
	newWatcher("posts/...", a.contentChanged)
	newWatcher("_theme/...", a.themeChanged)
	newWatcher("config.yml", a.configChanged)

	err := a.build()
	if err != nil {
		return err
	}

	handler := a.fs.HTTP(publicDirName)
	log.Infof("Listening on %s", a.config.BaseURL)
	return http.ListenAndServe(addr, handler)
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
		return
	}
	a.theme.SetTags(a.site.AllTags())
	a.renderChanged(path)
	err = a.saveAll()
	if err != nil {
		log.Errorf("unable to save site: %v", err)
		return
	}
	log.Infof("Done in %s", time.Since(start))
}

func (a *application) themeChanged(e notify.EventInfo) {
	start := time.Now()
	log.Infof("Theme changed. Rebuilding...")
	var err error
	a.theme, err = theme.FromDir(a.fs, themeDirName)
	if err != nil {
		log.Errorf("unable to load theme: %v", err)
		return
	}
	err = a.build()
	if err != nil {
		log.Errorf("unable to build site: %v", err)
		return
	}
	log.Infof("Done in %s", time.Since(start))
}

func (a *application) configChanged(e notify.EventInfo) {
	start := time.Now()
	err := a.init("")
	if err != nil {
		log.Errorf("unable to init application: %v", err)
		return
	}
	err = a.build()
	if err != nil {
		log.Errorf("unable to build site: %v", err)
		return
	}
	log.Infof("Done in %s", time.Since(start))
}

func (a *application) createPost(name string) error {
	data := map[string]string{
		"date": time.Now().Format(dateTimeFormat),
	}
	content := theme.Exec(a.fs, "_theme/post.md", data)
	err := a.fs.WriteFile(contentDirName+"/"+name, []byte(content))
	if err != nil {
		return fmt.Errorf("unable to write file: %v", err)
	}
	return nil
}

func (a *application) createSite(name string) error {
	err := newsite.RestoreAssets(name, "")
	if err != nil {
		return fmt.Errorf("unable to create new site: %v", err)
	}
	return nil
}
