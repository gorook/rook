package site

import (
	"fmt"
	"sort"
	"strings"

	"github.com/gorook/rook/config"
	"github.com/gorook/rook/fs"
	"github.com/yanzay/log"
)

const (
	pagesOnIndex = 10
)

// Site is collection of content pages
type Site struct {
	Pages      []*Page
	IndexPages []*IndexPage
	TagPages   map[string][]*IndexPage
	Tags       TagSet
}

// FromDir loads Site from given directory, recurcively
func FromDir(f *fs.FS, dir string) (*Site, error) {
	s := &Site{Tags: make(TagSet)}
	err := s.loadPages(f, dir)
	if err != nil {
		return nil, err
	}
	s.sort()
	s.createIndexPages()
	s.createTagPages()
	return s, nil
}

// Rebuild reloads page from file
func (s *Site) Rebuild(f *fs.FS, path string) error {
	page, err := pageFromFile(f, path)
	if err != nil {
		return err
	}
	for i, p := range s.Pages {
		if p.Path == page.Path {
			s.Pages[i] = page
			break
		}
	}
	s.sort()
	s.createIndexPages()
	s.createTagPages()
	return nil
}

// ByPath returns page by path
func (s *Site) ByPath(path string) *Page {
	path = strings.TrimSuffix(path, ".md")
	for _, page := range s.Pages {
		if page.Path == path {
			return page
		}
	}
	return nil
}

func (s *Site) loadPages(f *fs.FS, dir string) error {
	files, err := f.TreeList(dir)
	if err != nil {
		return err
	}
	s.Pages = make([]*Page, 0)
	for _, file := range files {
		page, err := pageFromFile(f, file)
		if err != nil {
			return err
		}
		s.Pages = append(s.Pages, page)
		s.Tags.Add(page.Front.Tags)
	}
	if len(s.Pages) == 0 {
		return fmt.Errorf("content directory is empty")
	}

	return nil
}

// PreprocessPages applies preprocessing rules to each page
func (s *Site) PreprocessPages(conf *config.SiteConfig) {
	var err error
	pp := &preprocessor{baseURL: conf.BaseURL}
	for _, page := range s.Pages {
		page.Content, err = pp.preprocess(page.Content)
		if err != nil {
			log.Warningf("Unable to preprocess page %s: %v", page.Path, err)
		}
	}
}

// PreprocessOne applies preprocessing rules to specific page
func (s *Site) PreprocessOne(conf *config.SiteConfig, page *Page) {
	var err error
	pp := &preprocessor{baseURL: conf.BaseURL}
	page.Content, err = pp.preprocess(page.Content)
	if err != nil {
		log.Warningf("Unable to preprocess page %s: %v", page.Path, err)
	}
}

func (s *Site) sort() {
	sort.Slice(s.Pages, func(i, j int) bool {
		return s.Pages[j].Front.Time.Before(s.Pages[i].Front.Time)
	})
}

func (s *Site) createIndexPages() {
	s.IndexPages = paginate(s.Pages, "")
}

func (s *Site) createTagPages() {
	tagged := make(map[string][]*Page)
	for _, page := range s.Pages {
		for _, tag := range page.Front.Tags {
			tagged[tag] = append(tagged[tag], page)
		}
	}

	s.TagPages = make(map[string][]*IndexPage)
	for tag, pages := range tagged {
		s.TagPages[tag] = paginate(pages, "tags/"+tag+"/")
	}
}

func paginate(pages []*Page, prefix string) []*IndexPage {
	indexPages := make([]*IndexPage, 0)
	currentIndexPage := &IndexPage{}
	currentPageNumber := 0
	var prevIndexPage *IndexPage

	for i, page := range pages {
		currentIndexPage.Pages = append(currentIndexPage.Pages, page)
		if len(currentIndexPage.Pages) == pagesOnIndex || i == len(pages)-1 {
			currentPageNumber++
			currentIndexPage.Number = currentPageNumber
			currentIndexPage.Path = prefix + pathToPage(currentPageNumber)
			currentIndexPage.Link = currentIndexPage.Path + "/"
			if prevIndexPage != nil {
				currentIndexPage.Prev = &PagerItem{
					Number: prevIndexPage.Number,
					Link:   prevIndexPage.Link,
				}
				prevIndexPage.Next = &PagerItem{
					Number: currentIndexPage.Number,
					Link:   currentIndexPage.Link,
				}
			}
			indexPages = append(indexPages, currentIndexPage)
			prevIndexPage = currentIndexPage
			currentIndexPage = &IndexPage{}
		}
	}
	return indexPages
}

func pathToPage(number int) string {
	if number == 1 {
		return ""
	}
	return fmt.Sprintf("page/%d/", number)
}
