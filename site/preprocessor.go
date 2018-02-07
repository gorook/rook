package site

import (
	"github.com/aymerick/raymond"
)

type preprocessor struct {
	baseURL string
}

func (p *preprocessor) link(page string) string {
	return p.baseURL + page
}

func (p *preprocessor) preprocess(content []byte) ([]byte, error) {
	tmpl, err := raymond.Parse(string(content))
	if err != nil {
		return nil, err
	}

	tmpl.RegisterHelper("link", p.link)

	res, err := tmpl.Exec(nil)
	if err != nil {
		return nil, err
	}
	return []byte(res), nil
}
