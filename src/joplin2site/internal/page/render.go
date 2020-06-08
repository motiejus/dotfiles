package page

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/gomarkdown/markdown"
	"github.com/motiejus/dotfiles/joplin2site/internal/note"
)

// Render() renders the concrete page
func (p *Page) Render(notes note.Notes) ([]byte, error) {
	tpl, err := template.New(p.ID).Parse(p.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse user's page: %w", err)
	}
	tctx := templateContext{notes: notes}
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

type templateContext struct {
	notes note.Notes
}

func (t *templateContext) indexFor(title string) ([]byte, error) {
	pages, err := SubPages(t.notes, title)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err := _indexFor.Execute(&buf, pages); err != nil {
		return nil, fmt.Errorf("failed to generate index: %w", err)
	}
	return buf.Bytes(), nil
}
