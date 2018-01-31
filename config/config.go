package config

import (
	"fmt"

	"github.com/gorook/rook/fs"
	"gopkg.in/yaml.v2"
)

// SiteConfig contains all info from config.yml
type SiteConfig struct {
	BaseURL string            `yaml:"baseURL"`
	Title   string            `yaml:"title"`
	Params  map[string]string `yaml:"params"`
}

// FromFile loads config from given file
func FromFile(f *fs.FS, name string) (*SiteConfig, error) {
	content, err := f.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("unable to load config: %v", err)
	}
	siteConfig := &SiteConfig{}
	err = yaml.UnmarshalStrict(content, siteConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to parse config: %v", err)
	}
	return siteConfig, nil
}
