package theme

import (
	"log"
	"time"

	"github.com/aymerick/raymond"
	"github.com/gorook/rook/fs"
	"github.com/gorook/rook/site"
	"github.com/jehiah/go-strftime"
	"github.com/spf13/afero"
)

// Theme contains templates
type Theme struct {
	index *raymond.Template
	post  *raymond.Template
	ctx   map[string]interface{}
}

// FromDir loads theme from dir and parses templates
func FromDir(f *fs.FS, dir string) (*Theme, error) {
	base, err := f.ReadFile("_theme/base.html")
	if err != nil {
		log.Fatalf("unable to read template: %v", err)
	}
	index, err := f.ReadFile("_theme/index.html")
	if err != nil {
		log.Fatalf("unable to read template: %v", err)
	}
	post, err := f.ReadFile("_theme/post.html")
	if err != nil {
		log.Fatalf("unable to read template: %v", err)
	}

	indexTemplate := raymond.MustParse(string(base))
	indexTemplate.RegisterPartial("content", string(index))

	postTemplate := raymond.MustParse(string(base))
	postTemplate.RegisterPartial("content", string(post))

	formatHelper := func(format string, t time.Time) string {
		return strftime.Format(format, t)
	}

	indexTemplate.RegisterHelper("format", formatHelper)
	postTemplate.RegisterHelper("format", formatHelper)

	return nil, nil
}

// Exec executes given template
func Exec(fs afero.Fs, template string, data interface{}) string {
	content, err := afero.ReadFile(fs, template)
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
	}
	res := t.index.MustExec(ctx)
	return res
}

// RenderPage renders single page template
func (t *Theme) RenderPage(page *site.Page) string {
	ctx := t.pageContext(page)
	res := t.post.MustExec(ctx)
	return res
}

func (t *Theme) pageContext(page *site.Page) map[string]interface{} {
	ctx := map[string]interface{}{
		"content":   raymond.SafeString(string(page.Content)),
		"summary":   raymond.SafeString(string(page.Summary)),
		"truncated": page.Truncated,
		"link":      page.Link,
		"site":      t.ctx,
	}
	for k, v := range page.Front.Vars {
		ctx[k] = v
	}
	ctx["date"] = page.Front.Time
	ctx["tags"] = page.Front.Tags
	return ctx
}

// func (t *Theme) siteContext() map[string]interface{} {
// 	log.Printf("Site context updated. Tags: %q", sc.Tags)
// 	return map[string]interface{}{
// 		"baseURL": sc.BaseURL,
// 		"title":   sc.Title,
// 		"tags":    sc.Tags,
// 		"params":  sc.Params,
// 	}
// }
