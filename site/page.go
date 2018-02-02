package site

import "time"

type Page struct {
	Content   []byte
	Summary   []byte
	Link      string
	Name      string
	Truncated bool

	Front *FrontMatter
}

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

type FrontMatter struct {
	Time  time.Time
	Title string
	Vars  map[string]interface{}
	Tags  []string
}
