package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"

	"github.com/gorook/rook/fs"
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
		a.Nil(app.init(""))
		a.NotNil(app.config)
	})
	t.Run("without config", func(t *testing.T) {
		app := newApplication(appInMemory)
		a.Error(app.init(""))
	})
}

func TestBuild(t *testing.T) {
	a := assert.New(t)
	app := newApplication(appInMemory)
	loadAssets(t, app.fs, "test-assets/simple/")
	a.Nil(app.init(""))
	a.Nil(app.build())
}

func loadAssets(t *testing.T, appfs *fs.FS, dir string) {
	// t.Helper()
	osfs := afero.NewOsFs()
	from := dir
	to := ""
	err := afero.Walk(osfs, dir, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("walk error for %s: %v", path, err)
		}
		path = strings.TrimPrefix(path, from)
		if path == "" {
			return nil
		}
		if fi.IsDir() {
			err = appfs.MkDirAll(to + path)
			if err != nil {
				return fmt.Errorf("unable to create dir: %v", err)
			}
		} else {
			var in io.ReadCloser
			var out io.WriteCloser
			in, err = osfs.Open(from + path)
			if err != nil {
				return fmt.Errorf("unable to open file: %v", err)
			}
			defer func() { err = in.Close() }()

			out, err = appfs.Create(to + path)
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
	if err != nil {
		t.Fatalf("unable to load assets: %v", err)
	}
}
