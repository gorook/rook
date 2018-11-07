package site

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"github.com/russross/blackfriday"
	"github.com/yanzay/log"
	yaml "gopkg.in/yaml.v2"

	"github.com/gorook/rook/fs"
)

var readMore = []byte("<!--more-->")

func (s *Site) pageFromFile(fs *fs.FS, path string) (*Page, error) {
	f, err := fs.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open page file: %v", err)
	}
	page, err := s.parsePage(f, path)
	if err != nil {
		return nil, fmt.Errorf("unable to parse page %s: %v", path, err)
	}
	err = f.Close()
	if err != nil {
		return nil, fmt.Errorf("unable to close file: %v", err)
	}
	return page, nil
}

func (s *Site) parsePage(r io.Reader, path string) (*Page, error) {
	scanner := bufio.NewScanner(r)
	fm := parseFrontMatter(scanner)
	summary, content, err := s.parsePageContent(scanner)
	if err != nil {
		return nil, err
	}
	page := &Page{
		Path:  trimExtension(path),
		Front: &FrontMatter{Vars: make(map[string]interface{})},
	}
	err = yaml.Unmarshal(fm, page.Front)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal page frontmatter for page %s: %v", path, err)
	}
	page.Content = content

	if len(summary) > 0 {
		page.Summary = summary
	} else {
		page.Summary = content
	}

	page.Truncated = len(summary) < len(content)

	page.Path = strings.TrimSuffix(path, filepath.Ext(path))
	fmt.Println("page path", page.Path)
	page.Link = page.Path + "/"

	return page, nil
}

func parseFrontMatter(scanner *bufio.Scanner) []byte {
	buf := &bytes.Buffer{}
	frontMatter := false
	for scanner.Scan() {
		line := scanner.Text()
		if !frontMatter && line == "---" {
			frontMatter = true
			continue
		}
		if frontMatter && line == "---" {
			break
		}
		_, err := buf.WriteString(line + "\n")
		if err != nil {
			log.Fatalf("unable to write to buffer: %v", err)
		}
	}
	return buf.Bytes()
}

func (s *Site) parsePageContent(scanner *bufio.Scanner) ([]byte, []byte, error) {
	buf := &bytes.Buffer{}
	for scanner.Scan() {
		_, err := buf.WriteString(scanner.Text() + "\n")
		if err != nil {
			log.Fatalf("unable to write to buffer: %v", err)
		}
	}
	content, err := s.proc.preprocess(buf.Bytes())
	if err != nil {
		return nil, nil, err
	}
	content = blackfriday.Run(content)
	summary := bytes.SplitN(content, readMore, 2)[0]
	return summary, content, nil
}

func trimExtension(path string) string {
	return strings.TrimSuffix(path, ".md")
}
