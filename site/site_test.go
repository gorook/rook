package site

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"

	"github.com/gorook/rook/config"
	"github.com/gorook/rook/fs"
)

func TestFromDir(t *testing.T) {
	a := assert.New(t)
	conf := &config.SiteConfig{}
	t.Run("dir not exist", func(t *testing.T) {
		memfs := afero.NewMemMapFs()
		f := fs.New(memfs, memfs)
		_, err := FromDir(f, conf, "posts", ContentTypeBlog)
		a.Nil(err)
	})

	t.Run("dir is empty", func(t *testing.T) {
		memfs := afero.NewMemMapFs()
		f := fs.New(memfs, memfs)
		f.MkDirAll("posts")
		_, err := FromDir(f, conf, "posts", ContentTypeBlog)
		a.Nil(err)
	})

	t.Run("dir not empty", func(t *testing.T) {
		memfs := afero.NewMemMapFs()
		f := fs.New(memfs, memfs)
		f.MkDirAll("posts")
		f.WriteFile("posts/post1.md", []byte{})
		s, err := FromDir(f, conf, "posts", ContentTypeBlog)
		a.Nil(err)
		a.NotNil(s)
	})

	t.Run("recursive tree", func(t *testing.T) {
		memfs := afero.NewMemMapFs()
		f := fs.New(memfs, memfs)
		f.MkDirAll("posts/subdir")
		f.WriteFile("posts/post.md", []byte{})
		f.WriteFile("posts/subdir/subpost1.md", []byte{})
		f.WriteFile("posts/subdir/subpost2.md", []byte{})
		s, err := FromDir(f, conf, "posts", ContentTypeBlog)
		a.Nil(err)
		a.NotNil(s)
		a.Len(s.Pages, 3)
	})

	t.Run("tags from pages", func(t *testing.T) {
		memfs := afero.NewMemMapFs()
		f := fs.New(memfs, memfs)
		post1 := []byte(`
---
tags: ['tag1', 'tag2', 'tag3']
---
`)
		post2 := []byte(`
---
tags: ['tag2', 'tag3', 'tag4']
---
`)
		f.MkDirAll("posts")
		f.WriteFile("posts/post1.md", post1)
		f.WriteFile("posts/post2.md", post2)
		s, err := FromDir(f, conf, "posts", ContentTypeBlog)
		a.Nil(err, "err should be nil")
		if a.NotNil(s, "site should not be nil") {
			if a.NotNil(s.Tags, "tags should not be nil") {
				a.NotNil(s.Tags["tag1"])
				a.NotNil(s.Tags["tag2"])
				a.NotNil(s.Tags["tag3"])
				a.NotNil(s.Tags["tag4"])
			}
		}
	})
}

func TestCreateIndexPages(t *testing.T) {
	a := assert.New(t)
	s := &Site{Tags: make(TagSet), Pages: make([]*Page, 0)}
	for i := 0; i < 23; i++ {
		s.Pages = append(s.Pages, &Page{})
	}
	s.createIndexPages()

	t.Run("pagination", func(t *testing.T) {
		if a.NotNil(s.IndexPages) {
			if a.Len(s.IndexPages, 3) {
				a.Len(s.IndexPages[0].Pages, 10)
				a.Equal(s.IndexPages[0].Number, 1)
				a.Nil(s.IndexPages[0].Prev)
				if a.NotNil(s.IndexPages[0].Next) {
					a.Equal(s.IndexPages[0].Next.Number, 2)
				}

				a.Len(s.IndexPages[1].Pages, 10)
				a.Equal(s.IndexPages[1].Number, 2)
				if a.NotNil(s.IndexPages[1].Prev) {
					a.Equal(s.IndexPages[1].Prev.Number, 1)
				}
				if a.NotNil(s.IndexPages[1].Next) {
					a.Equal(s.IndexPages[1].Next.Number, 3)
				}

				a.Len(s.IndexPages[2].Pages, 3)
				a.Equal(s.IndexPages[2].Number, 3)
				if a.NotNil(s.IndexPages[2].Prev) {
					a.Equal(s.IndexPages[2].Prev.Number, 2)
				}
				a.Nil(s.IndexPages[2].Next)
			}
		}
	})
}
