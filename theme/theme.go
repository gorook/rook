package theme

import (
	"fmt"
	"time"

	"github.com/aymerick/raymond"
	strftime "github.com/jehiah/go-strftime"
	"github.com/yanzay/log"

	"github.com/gorook/rook/config"
	"github.com/gorook/rook/fs"
	"github.com/gorook/rook/site"
)

// Theme contains templates
type Theme struct {
	index *raymond.Template
	post  *raymond.Template
	page  *raymond.Template
	ctx   map[string]interface{}
}

// FromDir loads theme from dir and parses templates
func FromDir(f *fs.FS, dir string) (*Theme, error) {
	base, err := f.ReadFile("_theme/base.html")
	if err != nil {
		return nil, fmt.Errorf("unable to read template: %v", err)
	}
	index, err := f.ReadFile("_theme/index.html")
	if err != nil {
		return nil, fmt.Errorf("unable to read template: %v", err)
	}
	post, err := f.ReadFile("_theme/post.html")
	if err != nil {
		return nil, fmt.Errorf("unable to read template: %v", err)
	}
	page, err := f.ReadFile("_theme/page.html")
	if err != nil {
		return nil, fmt.Errorf("unable to read template: %v", err)
	}

	indexTemplate := raymond.MustParse(string(base))
	indexTemplate.RegisterPartial("content", string(index))

	postTemplate := raymond.MustParse(string(base))
	postTemplate.RegisterPartial("content", string(post))

	pageTemplate := raymond.MustParse(string(base))
	pageTemplate.RegisterPartial("content", string(page))

	formatHelper := func(format string, t time.Time) string {
		return strftime.Format(format, t)
	}

	indexTemplate.RegisterHelper("format", formatHelper)
	postTemplate.RegisterHelper("format", formatHelper)

	return &Theme{
		index: indexTemplate,
		post:  postTemplate,
		page:  pageTemplate,
	}, nil
}

// SetTags sets tags for site render context
func (t *Theme) SetTags(tags []string) {
	t.ctx["tags"] = tags
}

// SetConfig sets site config for render context
func (t *Theme) SetConfig(conf *config.SiteConfig) {
	t.ctx = siteContext(conf)
}

// Exec executes given template
func Exec(f *fs.FS, template string, data interface{}) string {
	content, err := f.ReadFile(template)
	if err != nil {
		log.Fatalf("unable to read template file: %v", err)
	}
	tpl := raymond.MustParse(string(content))
	return tpl.MustExec(data)
}

// RenderIndex renders index page template
func (t *Theme) RenderIndex(ip *site.IndexPage) string {
	posts := make([]map[string]interface{}, 0)
	for _, page := range ip.Pages {
		posts = append(posts, t.pageContext(page))
	}
	ctx := map[string]interface{}{
		"posts": posts,
		"site":  t.ctx,
		"pager": map[string]interface{}{
			"number": ip.Number,
			"link":   ip.Link,
			"prev":   ip.Prev,
			"next":   ip.Next,
			"all":    ip.Pager,
		},
		"now": now(),
	}
	res := t.index.MustExec(ctx)
	return res
}

// RenderPost renders single page template
func (t *Theme) RenderPost(page *site.Page) string {
	ctx := t.pageContext(page)
	res := t.post.MustExec(ctx)
	return res
}

// RenderPage renders single page template
func (t *Theme) RenderPage(page *site.Page) string {
	ctx := t.pageContext(page)
	res := t.page.MustExec(ctx)
	return res
}

func (t *Theme) pageContext(page *site.Page) map[string]interface{} {
	ctx := map[string]interface{}{
		"content":   raymond.SafeString(string(page.Content)),
		"summary":   raymond.SafeString(string(page.Summary)),
		"truncated": page.Truncated,
		"link":      page.Link,
		"site":      t.ctx,
		"now":       now(),
	}
	for k, v := range page.Front.Vars {
		ctx[k] = v
	}
	ctx["title"] = page.Front.Title
	ctx["date"] = page.Front.Time
	ctx["tags"] = page.Front.Tags

	return ctx
}

func siteContext(conf *config.SiteConfig) map[string]interface{} {
	return map[string]interface{}{
		"baseURL": conf.BaseURL,
		"title":   conf.Title,
		"params":  conf.Params,
	}
}

func now() map[string]string {
	return map[string]string{
		"year":  fmt.Sprint(time.Now().Year()),
		"month": fmt.Sprint(time.Now().Month()),
		"day":   fmt.Sprint(time.Now().Day()),
	}
}
