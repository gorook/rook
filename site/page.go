package site

import (
	"fmt"
	"time"
)

// Page is a parsed page
type Page struct {
	Content   []byte
	Summary   []byte
	Link      string
	Name      string
	Truncated bool
	Path      string
	Draft     bool

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

// UnmarshalYAML implements yaml.Unmarshaler interface
func (fm *FrontMatter) UnmarshalYAML(unmarshal func(interface{}) error) error {
	m := make(map[string]interface{})
	err := unmarshal(m)
	if err != nil {
		return err
	}
	for key, val := range m {
		err = fm.applyValue(key, val)
		if err != nil {
			return err
		}
	}
	return nil
}

func (fm *FrontMatter) applyValue(key string, val interface{}) error {
	switch key {
	case "tags":
		tags, ok := val.([]interface{})
		if !ok {
			return fmt.Errorf("can't parse tags")
		}
		for _, tag := range tags {
			fm.Tags = append(fm.Tags, tag.(string))
		}
	case "date":
		str, ok := val.(string)
		if !ok {
			return fmt.Errorf("unable to parse date")
		}
		var err error
		fm.Time, err = time.Parse("2006-01-02 15:04:05", str)
		if err != nil {
			return err
		}
	case "title":
		str, ok := val.(string)
		if !ok {
			return fmt.Errorf("unable to parse title")
		}
		fm.Title = str
	default:
		fm.Vars[key] = val
	}
	return nil
}
