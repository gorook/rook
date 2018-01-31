package main

import (
	"fmt"

	"github.com/gorook/rook/fs"
	"github.com/gorook/rook/types"
	"gopkg.in/yaml.v2"
)

const (
	configFileName = "config.yml"
)

type application struct {
	fs     *fs.FS
	config *types.SiteConfig
}

func newApplication(fs *fs.FS) *application {
	a := &application{
		fs: fs,
	}
	return a
}

func (a *application) init() error {
	err := a.loadConfig()
	return err
}

func (a *application) build() error {
	return nil
}

func (a *application) loadConfig() error {
	content, err := a.fs.ReadFile(configFileName)
	if err != nil {
		return fmt.Errorf("unable to load config: %v", err)
	}
	siteConfig := &types.SiteConfig{}
	err = yaml.UnmarshalStrict(content, siteConfig)
	if err != nil {
		return fmt.Errorf("unable to parse config: %v", err)
	}
	a.config = siteConfig
	return nil
}
