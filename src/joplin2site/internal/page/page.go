package page

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	inote "github.com/motiejus/dotfiles/joplin2site/internal/note"
	"gopkg.in/yaml.v2"
)

type (
	// Page is a stand-alone converted note.
	Page struct {
		ID          string
		Title       string
		URL         string
		Body        string
		CreatedAt   time.Time
		PublishedAt time.Time
	}

	// Pages is a slice of pages ordered by publish date.
	Pages []Page
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
func FromNote(n inote.Note) (Page, error) {
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
func SubPages(notes inote.Notes, title string) (Pages, error) {
	// Find the parent note that will be the parent of the sub-notebook.
	parentID := notes.GetFolderID(title)
	if parentID == "" {
		return nil, fmt.Errorf("sub-page %q not found", title)
	}

	var pages Pages
	for _, note := range notes {
		if note.ParentID != parentID {
			continue
		}
		if note.Type != inote.ItemTypeNote {
			continue
		}
		page, err := FromNote(note)
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

// ToPublishablePages returns notes that can be published.
func ToPublishablePages(notes inote.Notes) (Pages, error) {
	pages := make([]Page, len(notes))
	for id, note := range notes {
		if note.Type != inote.ItemTypeNote {
			continue
		}
		page, err := ipage.FromNote(note)
		if err != nil {
			return nil, fmt.Errorf("failed to convert note to page: %w", err)
		}
		if page.PublishedAt.After(time.Now()) {
			continue
		}

		pages = append(pages, page)
	}
	return pages, nil
}

type PrevNext struct {
	Prev *page.Page
	Next *page.Page
}

// Sequenced returns sequenced pages within the same folder.
func (pages Pages) Sequenced(tree inote.NoteTree) map[string]PrevNext {
	// folders maps folder ID -> set(note id)
	var folders map[string]map[string]struct{}
	for _, page := range pages {
		parent := tree[
	}
}
