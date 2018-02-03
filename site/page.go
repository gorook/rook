package site

import "time"

// Page is a parsed page
type Page struct {
	Content   []byte
	Summary   []byte
	Link      string
	Name      string
	Truncated bool
	Path      string

	Front *FrontMatter
}

// IndexPage is a colleciton of pages with pagination info
type IndexPage struct {
	PagerItem
	Path  string
	Pages []*Page
	Next  *PagerItem
	Prev  *PagerItem
	Pager []*PagerItem
}

// PagerItem is a pager entity
type PagerItem struct {
	Number int
	Link   string
}

// FrontMatter is a metadata for Page
type FrontMatter struct {
	Time  time.Time
	Title string
	Vars  map[string]interface{}
	Tags  []string
}
