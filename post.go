package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	posts []*Post
)

type Post struct {
	Title   string
	Summary string
	Tags    []string
	Slug    string
	Time    string
	Path    string
	Raw     []byte
	Html    string
}

func init() {
	log.Println("Loading Posts...")
	posts = loadPosts()
}

func (p *Post) HTML() string {
	if p.Html == "" {
		buf := bytes.Buffer{}
		if err := md.Convert(p.Raw, &buf); err != nil {
			return err.Error()
		}
		p.Html = buf.String()
	}
	return p.Html
}

func GetPostByPath(path string) *Post {
	if strings.HasPrefix(path, "/post/") {
		path = strings.Replace(path, "/", "", 1)
	}

	for _, p := range posts {
		if p.Path == path {
			return p
		}
	}
	return nil
}

func loadPosts() []*Post {
	_posts := []*Post{}
	filepath.Walk("./posts", func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(path, ".md") {
			// filepath.Ext(path)
			raw, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			log.Println("load file:", path)
			_posts = append(_posts, ParseToPost(raw, "/"+path))
		}
		return nil
	})
	return _posts
}

func GetPostsAll() []*Post {
	return posts
}

func GetPostsByTag(tag, p string) []*Post {
	postsOfTag := []*Post{}
	for _, v := range posts {
		for _, t := range v.Tags {
			if t == tag {
				postsOfTag = append(postsOfTag, v)
				break
			}
		}
	}
	return postsOfTag
}

func GetPostBySlug(slug string) *Post {
	for _, v := range posts {
		if v.Slug == slug {
			return v
		}
	}
	return nil
}

func ParseToPost(raw []byte, path string) *Post {
	content := string(raw)
	post := Post{}
	splits := strings.Split(content, "\n")
	for _, line := range splits {
		line = strings.TrimSpace(line)
		if !strings.Contains(content, ":") || strings.HasPrefix(line, "<!--") {
			continue
		}

		if strings.HasPrefix(line, "-->") {
			break
		}

		pair := strings.Split(line, ":")
		value := strings.TrimSpace(pair[1])
		switch strings.ToLower(pair[0]) {
		case "title":
			{
				post.Title = value
				break
			}
		case "summary":
			{
				post.Summary = value
				break
			}
		case "tag":
			{
				post.Tags = strings.Split(value, ",")
				break
			}
		case "slug":
			{
				post.Slug = value
				break
			}
		case "time":
			{
				post.Time = value
				break
			}
		default:
			{
				log.Println("ignore line", line)
			}
		}

		fmt.Println(line)
	}

	post.Raw = raw
	return &post
}
