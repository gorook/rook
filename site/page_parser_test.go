package site

import (
	"bytes"
	"testing"

	"github.com/jehiah/go-strftime"
	"github.com/stretchr/testify/assert"
)

func TestParsePage(t *testing.T) {
	a := assert.New(t)
	content := `
---
title: Post
tags: ["tag1", "tag2"]
date: 2016-08-10 15:25:51
custom: value
---
# Heading
content

<!--more-->

more content
`
	buf := bytes.NewBufferString(content)
	page, err := parsePage(buf, "posts/post.md")
	a.Nil(err)
	if a.NotNil(page) {
		a.Equal("Post", page.Front.Title)
		a.Equal([]string{"tag1", "tag2"}, page.Front.Tags)
		a.Equal("2016-08-10 15:25:51", strftime.Format("%Y-%m-%d %H:%M:%S", page.Front.Time))
		a.Equal("value", page.Front.Vars["custom"])

		a.Contains(string(page.Summary), "<h1>Heading</h1>")
		a.Contains(string(page.Summary), "<p>content</p>")
		a.NotContains(string(page.Summary), "<p>more content</p>")

		a.Contains(string(page.Content), "<h1>Heading</h1>")
		a.Contains(string(page.Content), "<p>content</p>")
		a.Contains(string(page.Content), "<p>more content</p>")
	}
}
