package page

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/motiejus/dotfiles/joplin2site/internal/note"
	"gopkg.in/yaml.v2"
)

// Page contains everything that's necessary to render a page.
type Page struct {
	ID          string
	Title       string
	URL         string
	BodyHTML    string
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
		return Page{}, fmt.Errorf("failed to unmarshal user's metadata: %w", err)
	}

	return Page{
		ID:          n.ID,
		Title:       n.Title,
		URL:         meta.URL,
		BodyHTML:    body,
		CreatedAt:   n.CreatedTime,
		PublishedAt: meta.PublishedAt,
	}, nil
}
