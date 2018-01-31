package main

import (
	"testing"

	"github.com/spf13/afero"

	"github.com/gorook/rook/fs"
	"github.com/stretchr/testify/assert"
)

var configYml = []byte(`
baseURL: http://localhost:1414/
title: Test Site
`)

func TestNewApplication(t *testing.T) {
	filesys := fs.New(afero.NewMemMapFs(), afero.NewMemMapFs())
	a := newApplication(filesys)
	assert.NotNil(t, a)
}

func TestInit(t *testing.T) {
	assert := assert.New(t)
	t.Run("with config", func(t *testing.T) {
		memfs := afero.NewMemMapFs()
		filesys := fs.New(memfs, memfs)
		filesys.WriteFile("config.yml", configYml)
		a := newApplication(filesys)
		assert.Nil(a.init())
		assert.NotNil(a.config)
	})
	t.Run("without config", func(t *testing.T) {
		memfs := afero.NewMemMapFs()
		filesys := fs.New(memfs, memfs)
		a := newApplication(filesys)
		assert.Error(a.init())
	})
}

func TestLoadConfig(t *testing.T) {
	assert := assert.New(t)
	t.Run("with config", func(t *testing.T) {
		memfs := afero.NewMemMapFs()
		filesys := fs.New(memfs, memfs)
		filesys.WriteFile("config.yml", configYml)
		a := newApplication(filesys)
		assert.Nil(a.loadConfig())
		assert.NotNil(a.config)
	})
	t.Run("without config", func(t *testing.T) {
		memfs := afero.NewMemMapFs()
		filesys := fs.New(memfs, memfs)
		a := newApplication(filesys)
		assert.Error(a.loadConfig())
	})
	t.Run("with invalid config", func(t *testing.T) {
		memfs := afero.NewMemMapFs()
		filesys := fs.New(memfs, memfs)
		filesys.WriteFile("config.yml", []byte("some invalid config"))
		a := newApplication(filesys)
		assert.Error(a.loadConfig())
	})
}

func TestBuild(t *testing.T) {
	assert := assert.New(t)
	memfs := afero.NewMemMapFs()
	filesys := fs.New(memfs, memfs)
	filesys.WriteFile("config.yml", configYml)
	a := newApplication(filesys)
	assert.Nil(a.build())
}
