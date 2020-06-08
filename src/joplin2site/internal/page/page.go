package page

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/motiejus/dotfiles/joplin2site/internal/note"
	"gopkg.in/yaml.v2"
)

type (
	// Page contains everything that's necessary to render a page.
	Page struct {
		ID          string
		Title       string
		URL         string
		Body        string
		CreatedAt   time.Time
		PublishedAt time.Time
	}

	// Pages is a slice of pages.
	Pages []Page

	// PageHierarchy is a map of: folderID -> ordered list of pages.
	PageHierarchy map[string][]*Page
)

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

// FromNote converts a Joplin Note to a Page
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

// SubPages returns immediate Pages of a tree
func SubPages(notes note.Notes, title string) (Pages, error) {
	// Find the parent note that will be the parent of the sub-notebook.
	parentID := notes.GetFolderID(title)
	if parentID == "" {
		return nil, fmt.Errorf("sub-page %q not found", title)
	}

	var pages Pages
	for _, inote := range notes {
		if inote.ParentID != parentID {
			continue
		}
		if inote.Type != note.ItemTypeNote {
			continue
		}
		page, err := FromNote(*inote)
		if err != nil {
			return nil, fmt.Errorf("failed to convert note to page: %w", err)
		}
		if page.PublishedAt.After(time.Now()) {
			continue
		}
		pages = append(pages, page)
	}
	sort.Slice(pages, func(i, j int) bool {
		return pages[i].PublishedAt.Before(pages[j].PublishedAt)
	})
	return pages, nil
}
