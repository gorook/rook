package site

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/aymerick/raymond"
	"github.com/gorook/rook/fs"
)

type preprocessor struct {
	baseURL string
	fs      *fs.FS
}

func (p *preprocessor) link(page string) string {
	return p.baseURL + page
}

func (p *preprocessor) youtube(id string) raymond.SafeString {
	s := `<iframe width="560" height="315" src="https://www.youtube-nocookie.com/embed/%s" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe>`
	return raymond.SafeString(fmt.Sprintf(s, id))
}

func (p *preprocessor) snippet(lang, path string) raymond.SafeString {
	f, err := p.fs.Open(path)
	if err != nil {
		log.Printf("unable to open file %s, skipping: %v", path, err)
		return ""
	}
	defer f.Close()
	content, err := ioutil.ReadAll(f)
	if err != nil {
		log.Printf("unable to read file %s, skipping: %v", path, err)
		return ""
	}
	return raymond.SafeString(fmt.Sprintf("```%s\n%s```", lang, string(content)))
}

func (p *preprocessor) preprocess(content []byte) ([]byte, error) {
	tmpl, err := raymond.Parse(string(content))
	if err != nil {
		return nil, err
	}

	tmpl.RegisterHelper("link", p.link)
	tmpl.RegisterHelper("youtube", p.youtube)
	tmpl.RegisterHelper("snippet", p.snippet)

	res, err := tmpl.Exec(nil)
	if err != nil {
		return nil, err
	}
	return []byte(res), nil
}
