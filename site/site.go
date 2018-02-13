package site

import (
	"fmt"
	"sort"
	"strings"

	"github.com/gorook/rook/config"
	"github.com/gorook/rook/fs"
)

const (
	pagesOnIndex = 10
)

// ContentType blog or docs
type ContentType int

// List of content types
const (
	ContentTypeBlog ContentType = iota
	ContentTypeDocs
)

// Site is collection of content pages
type Site struct {
	Pages      []*Page
	IndexPages []*IndexPage
	TagPages   map[string][]*IndexPage
	Tags       TagSet

	proc *preprocessor
}

// FromDir loads Site from given directory, recurcively
func FromDir(f *fs.FS, conf *config.SiteConfig, dir string, contentType ContentType) (*Site, error) {
	s := &Site{
		Tags: make(TagSet),
		proc: &preprocessor{baseURL: conf.BaseURL},
	}
	exists, err := f.DirExists(dir)
	if err != nil {
		return nil, err
	}
	if !exists {
		return s, nil
	}
	err = s.loadPages(f, dir)
	if err != nil {
		return nil, err
	}
	switch contentType {
	case ContentTypeBlog:
		s.sort()
		s.createIndexPages()
		s.createTagPages()
	case ContentTypeDocs:
		s.createDocsIndex()
	default:
		return nil, fmt.Errorf("unknown content type: %d", contentType)
	}
	return s, nil
}

// Rebuild reloads page from file
func (s *Site) Rebuild(f *fs.FS, path string) error {
	page, err := s.pageFromFile(f, path)
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

// AllTags returns all tags sorted by popularity desc
func (s *Site) AllTags() []string {
	tags := s.Tags.All()
	sort.Slice(tags, func(i, j int) bool {
		var inum, jnum int
		for _, ip := range s.TagPages[tags[i]] {
			inum += len(ip.Pages)
		}
		for _, ip := range s.TagPages[tags[j]] {
			jnum += len(ip.Pages)
		}
		return inum > jnum
	})
	return tags
}

func (s *Site) loadPages(f *fs.FS, dir string) error {
	files, err := f.TreeList(dir)
	if err != nil {
		return err
	}
	s.Pages = make([]*Page, 0)
	for _, file := range files {
		page, err := s.pageFromFile(f, file)
		if err != nil {
			return err
		}
		s.Pages = append(s.Pages, page)
		s.Tags.Add(page.Front.Tags)
	}

	return nil
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

func (s *Site) createDocsIndex() {
	page := s.ByPath("docs/index")
	page.Path = ""
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
