package main

import (
	"strings"
	"testing"

	"github.com/gorook/rook/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

var configYml = []byte(`
baseURL: http://localhost:1414/
title: Test Site
params:
  param1: value1
`)

func TestNewApplication(t *testing.T) {
	app := newApplication(appInMemory)
	assert.NotNil(t, app)
}

func TestInit(t *testing.T) {
	a := assert.New(t)
	t.Run("with config", func(t *testing.T) {
		app := newApplication(appInMemory)
		app.fs.WriteFile("config.yml", configYml)
		app.fs.MkDirAll("posts")
		app.fs.WriteFile("posts/post.md", []byte{})
		app.fs.MkDirAll("_theme")
		app.fs.WriteFile("_theme/base.html", []byte{})
		app.fs.WriteFile("_theme/index.html", []byte{})
		app.fs.WriteFile("_theme/post.html", []byte{})
		a.Nil(app.init())
		a.NotNil(app.config)
	})
	t.Run("without config", func(t *testing.T) {
		app := newApplication(appInMemory)
		a.Error(app.init())
	})
}

func TestBuild(t *testing.T) {
	a := assert.New(t)
	app := newApplication(appInMemory)
	loadAssets(t, app.fs, "test-assets/simple/")
	a.Nil(app.init())
	a.Nil(app.build())
}

func loadAssets(t *testing.T, appfs *fs.FS, dir string) {
	t.Helper()
	osfs := afero.NewOsFs()
	f := fs.New(osfs, osfs)
	list, err := f.TreeList(dir)
	if err != nil {
		panic(err)
	}
	for _, path := range list {
		cont, err := f.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		newpath := strings.TrimPrefix(path, dir)
		err = appfs.WriteFile(newpath, cont)
		if err != nil {
			t.Fatal(err)
		}
	}
}
