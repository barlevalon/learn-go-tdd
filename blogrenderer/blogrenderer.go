package blogrenderer

import (
	"embed"
	"html/template"
	"io"
	"learn-go-with-tests/blogposts"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

type PostRenderer struct {
	templ    *template.Template
	mdParser *parser.Parser
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}
	return &PostRenderer{templ: templ}, nil
}

//go:embed templates/*
var postTemplates embed.FS

func (r *PostRenderer) Render(w io.Writer, p blogposts.Post) error {
	return r.templ.ExecuteTemplate(w, "blog.gohtml", newPostVM(p, r))
}

type postViewModel struct {
	blogposts.Post
	HTMLBody template.HTML
}

func newPostVM(p blogposts.Post, r *PostRenderer) postViewModel {
	vm := postViewModel{Post: p}
	vm.HTMLBody = template.HTML(markdown.ToHTML([]byte(p.Body), r.mdParser, nil))
	return vm
}
