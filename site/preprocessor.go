package site

import (
	"fmt"

	"github.com/aymerick/raymond"
)

type preprocessor struct {
	baseURL string
}

func (p *preprocessor) link(page string) string {
	return p.baseURL + page
}

func (p *preprocessor) youtube(id string) raymond.SafeString {
	fmt.Println(id)
	s := `<iframe width="560" height="315" src="https://www.youtube-nocookie.com/embed/%s" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe>`
	return raymond.SafeString(fmt.Sprintf(s, id))
}

func (p *preprocessor) preprocess(content []byte) ([]byte, error) {
	tmpl, err := raymond.Parse(string(content))
	if err != nil {
		return nil, err
	}

	tmpl.RegisterHelper("link", p.link)
	tmpl.RegisterHelper("youtube", p.youtube)

	res, err := tmpl.Exec(nil)
	if err != nil {
		return nil, err
	}
	return []byte(res), nil
}
