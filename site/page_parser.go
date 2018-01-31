package site

import (
	"bufio"
	"bytes"
	"log"
	"path/filepath"
	"strings"

	"github.com/gorook/rook/fs"
	blackfriday "gopkg.in/russross/blackfriday.v2"
	yaml "gopkg.in/yaml.v2"
)

var readMore = []byte("<!--more-->")

func pageFromFile(fs *fs.FS, dir, name string) *Page {
	filePath := filepath.Join(dir, name)
	f, err := fs.Open(filePath)
	if err != nil {
		log.Fatalf("unable to render page: %v", err)
	}
	scanner := bufio.NewScanner(f)
	fm := parseFrontMatter(scanner)
	summary, content := parsePageContent(scanner)
	err = f.Close()
	if err != nil {
		log.Fatalf("unable to close page file: %v", err)
	}
	page := &Page{
		Front: &FrontMatter{},
	}
	err = yaml.Unmarshal(fm, page.Front)
	if err != nil {
		log.Fatalf("unable to unmarshal page frontmatter: %v", err)
	}
	page.Content = content

	if len(summary) > 0 {
		page.Summary = summary
	} else {
		page.Summary = content
	}

	page.Truncated = len(summary) < len(content)

	name = strings.TrimSuffix(name, filepath.Ext(name))
	page.Link = name + "/"
	page.Name = name

	// Handling date
	// date, ok := page.Vars["date"].(string)
	// if ok {
	// 	page.Time, err = time.Parse("2006-01-02 15:04:05", date)
	// 	if err != nil {
	// 		log.Printf("[WARN] unable to parse time: %v", err)
	// 		page.Time = time.Now()
	// 	}
	// } else {
	// 	page.Time = time.Now()
	// }

	// Handling tags
	// log.Println("parsing tags", page.Vars["tags"])
	// tags, ok := page.Vars["tags"].([]interface{})
	// if ok {
	// 	log.Println("ok")
	// 	for _, tag := range tags {
	// 		page.Tags = append(page.Tags, tag.(string))
	// 	}
	// }
	// if !ok {
	// 	fmt.Printf("%T", page.Vars["tags"])
	// }

	return page
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

func parsePageContent(scanner *bufio.Scanner) ([]byte, []byte) {
	buf := &bytes.Buffer{}
	for scanner.Scan() {
		_, err := buf.WriteString(scanner.Text() + "\n")
		if err != nil {
			log.Fatalf("unable to write to buffer: %v", err)
		}
	}
	content := blackfriday.Run(buf.Bytes())
	summary := bytes.SplitN(content, readMore, 2)[0]
	return summary, content
}
