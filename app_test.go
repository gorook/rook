package main

import (
	"testing"

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
		a.Nil(app.init())
		a.NotNil(app.config)
	})
	t.Run("without config", func(t *testing.T) {
		app := newApplication(appInMemory)
		a.Error(app.init())
	})
}

// func TestLoadConfig(t *testing.T) {
// 	a := assert.New(t)
// 	t.Run("with config", func(t *testing.T) {
// 		memfs := afero.NewMemMapFs()
// 		filesys := fs.New(memfs, memfs)
// 		filesys.WriteFile("config.yml", configYml)
// 		app := newApplication(filesys)
// 		a.Nil(app.loadConfig())
// 		a.NotNil(app.config)
// 	})
// 	t.Run("without config", func(t *testing.T) {
// 		memfs := afero.NewMemMapFs()
// 		filesys := fs.New(memfs, memfs)
// 		app := newApplication(filesys)
// 		a.Error(app.loadConfig())
// 	})
// 	t.Run("with invalid config", func(t *testing.T) {
// 		memfs := afero.NewMemMapFs()
// 		filesys := fs.New(memfs, memfs)
// 		filesys.WriteFile("config.yml", []byte("some invalid config"))
// 		app := newApplication(filesys)
// 		a.Error(app.loadConfig())
// 	})
// }

// func TestBuild(t *testing.T) {
// 	a := assert.New(t)
// 	memfs := afero.NewMemMapFs()
// 	filesys := fs.New(memfs, memfs)
// 	filesys.WriteFile("config.yml", configYml)
// 	app := newApplication(filesys)
// 	a.Nil(app.build())
// }
