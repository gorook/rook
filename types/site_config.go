package types

// SiteConfig contains all info from config.yml
type SiteConfig struct {
	BaseURL string `yaml:"baseURL"`
	Title   string `yaml:"title"`
}
