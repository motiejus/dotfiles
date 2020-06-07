package page

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"strings"
	"time"

	"github.com/gomarkdown/markdown"
	"github.com/motiejus/dotfiles/joplin2site/internal/note"
	"gopkg.in/yaml.v2"
)

// Page contains everything that's necessary to render a page.
type Page struct {
	ID          string
	Title       string
	URL         string
	Body        string
	CreatedAt   time.Time
	PublishedAt time.Time
}

type userMeta struct {
	URL         string    `yaml:"url"`
	PublishedAt time.Time `yaml:"published_at"`
}

var (
	ErrMetaStart = errors.New("missing metadata opening tag")
	ErrMetaEnd   = errors.New("missing metadata closing tag")
)

const (
	_metaPrefix = "<!--\n"
	_metaSuffix = "\n-->\n"
)

// FromNote converts a Joplin Note to Page
func FromNote(n note.Note) (Page, error) {
	if !strings.HasPrefix(n.Body, _metaPrefix) {
		return Page{}, ErrMetaStart
	}
	endIdx := strings.Index(n.Body, _metaSuffix)
	if endIdx == -1 {
		return Page{}, ErrMetaEnd
	}
	metaS := n.Body[len(_metaPrefix):endIdx]
	body := n.Body[endIdx+len(_metaSuffix):]

	var meta userMeta
	if err := yaml.UnmarshalStrict([]byte(metaS), &meta); err != nil {
		return Page{}, fmt.Errorf("bad user's metadata: %w", err)
	}

	return Page{
		ID:          n.ID,
		Title:       n.Title,
		URL:         meta.URL,
		Body:        body,
		CreatedAt:   n.CreatedTime,
		PublishedAt: meta.PublishedAt,
	}, nil
}

type templateContext struct {
	tree *Tree
}

func (t *templateContext) indexFor(name string) ([]byte, error) {
	subpages, err := t.tree.SubPages(name)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err := _indexFor.Execute(&buf, subpages); err != nil {
		return nil, fmt.Errorf("failed to generate index: %w", err)
	}
	return buf.Bytes(), nil
}

// Render() renders the concrete page
func (p Page) Render(t *Tree) ([]byte, error) {
	tplName := p.Title + "-" + p.ID
	tpl, err := template.New(tplName).Parse(p.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse user's page: %w", err)
	}
	tctx := templateContext{tree: t}
	funcs := template.FuncMap{
		"indexFor": tctx.indexFor,
	}
	var buf bytes.Buffer
	if err := tpl.Funcs(funcs).Execute(&buf, nil); err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}
	html := markdown.ToHTML(buf.Bytes(), nil, nil)

	return html, nil
}
