package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/http/httputil"
	"regexp"
	"strings"
	"time"

	ext "github.com/OhYee/goldmark-fenced_codeblock_extension"
	uml "github.com/OhYee/goldmark-plantuml"
	toc "github.com/abhinav/goldmark-toc"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

var (
	prefix   = regexp.MustCompile("^(/static|/posts)+(.*)")
	md       goldmark.Markdown
	handlers []*Handler
)

func init() {
	log.Println("init handlers")
	md = goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			&toc.Extender{},
			ext.NewExt(
				uml.RenderMap(50, "plantuml"),
				ext.RenderMap{
					Languages:      []string{"*"},
					RenderFunction: ext.GetFencedCodeBlockRendererFunc(highlighting.NewHTMLRenderer()),
				},
			),
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(),
	)

	// assets handler
	handlers = append(handlers, &Handler{
		Regex: regexp.MustCompile("((.*).(txt|jpg|png|js|ico|css|map)$)"),
		Do: func(w http.ResponseWriter, r *http.Request) error {
			path := r.URL.Path
			if !prefix.MatchString(r.URL.Path) {
				path = "/static" + path
			}
			http.ServeFile(w, r, fmt.Sprintf(".%s", path))
			return nil
		},
	})
	// markdown handler
	handlers = append(handlers, &Handler{
		Regex: regexp.MustCompile("^(/post/)(.*)+"),
		Do: func(w http.ResponseWriter, r *http.Request) error {
			slug := strings.Replace(r.URL.Path, "/post/", "", 1)
			post := GetPostBySlug(slug)
			if post == nil {
				w.WriteHeader(404)
				return render.ExecuteTemplate(w, "404", nil)
			}
			return render.ExecuteTemplate(w, "markdown", template.HTML(post.HTML()))
		},
	})
	// index handler
	handlers = append(handlers, &Handler{
		Regex: regexp.MustCompile("^(/|)+$"),
		Do: func(w http.ResponseWriter, r *http.Request) error {
			return render.ExecuteTemplate(w, "index", GetPostsAll())
		},
	})
	// list handler
	handlers = append(handlers, &Handler{
		Regex: regexp.MustCompile("^/tag/(.*)+"),
		Do: func(w http.ResponseWriter, r *http.Request) error {
			tag := strings.Replace(r.URL.Path, "/tag/", "", 1)
			p := r.URL.Query().Get("p")
			return render.ExecuteTemplate(w, "tag", GetPostsByTag(tag, p))
		},
	})
	// mgr handler
	handlers = append(handlers, &Handler{
		Regex: regexp.MustCompile("^/status(.*)+"),
		Do: func(w http.ResponseWriter, r *http.Request) error {
			action := r.URL.Query().Get("action")

			switch action {
			case "reload":
				{
					posts = loadPosts()
					w.WriteHeader(200)
					w.Write([]byte("{ \"msg\": \"reload ok\"}"))
				}
			}

			return nil
		},
	})
	// dump handler
	handlers = append(handlers, &Handler{
		Regex: nil,
		Do: func(w http.ResponseWriter, r *http.Request) error {
			if dump, err := httputil.DumpRequest(r, true); err == nil {
				w.Write(dump)
			} else {
				w.Write([]byte(err.Error()))
			}

			return nil
		},
	})

}

type Handler struct {
	Regex *regexp.Regexp
	Do    func(w http.ResponseWriter, r *http.Request) error
}

func handle(w http.ResponseWriter, r *http.Request) {
	start := time.Now().UnixMilli()
	var message string
	defer func() {
		log.Println(message, " with ", time.Now().UnixMilli()-start, " ms")
	}()
	for _, h := range handlers {

		if h.Regex.MatchString(r.URL.Path) {
			message = fmt.Sprintf("[ %s ] => [ %s ]", r.URL.Path, h.Regex.String())
			if err := h.Do(w, r); err != nil {
				message = message + " throw err " + err.Error()
			}
			break
		}
		if h.Regex == nil {
			message = fmt.Sprintf("%s match %s", r.URL.Path, "*")
			h.Do(w, r)
			break
		}
	}
}
