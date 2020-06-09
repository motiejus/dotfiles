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
	ErrMetaStart   = errors.New("missing metadata opening tag")
	ErrMetaEnd     = errors.New("missing metadata closing tag")
	ErrUnpublished = errors.New("not yet published")
	ErrBadType     = errors.New("invalid note type")
)

const (
	_metaPrefix = "<!--\n"
	_metaSuffix = "\n-->\n"
)

// FromNote converts a Joplin Note to a Page
func FromNote(note inote.Note) (Page, error) {
	if note.Type != inote.ItemTypeNote {
		return Page{}, fmt.Errorf("got %d, expected %d: %w",
			note.Type, inote.ItemTypeNote, ErrBadType)
	}

	if !strings.HasPrefix(note.Body, _metaPrefix) {
		return Page{}, ErrMetaStart
	}
	endIdx := strings.Index(note.Body, _metaSuffix)
	if endIdx == -1 {
		return Page{}, ErrMetaEnd
	}
	metaS := note.Body[len(_metaPrefix):endIdx]
	body := note.Body[endIdx+len(_metaSuffix):]

	var meta userMeta
	if err := yaml.UnmarshalStrict([]byte(metaS), &meta); err != nil {
		return Page{}, fmt.Errorf("bad user's metadata: %w", err)
	}

	if page.PublishedAt.After(time.Now()) {
		return Page{}, ErrUnpublished
	}

	return Page{
		ID:          note.ID,
		Title:       note.Title,
		URL:         meta.URL,
		Body:        body,
		CreatedAt:   note.CreatedTime,
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
		page, err := FromNote(note)
		if err != nil && (errors.Is(err, ErrBadType) || errors.Is(err, ErrUnpublished)) {
			continue
		} else if err != nil {
			return nil, fmt.Errorf("failed to convert note to page: %w", err)
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
		page, err := ipage.FromNote(note)
		if err != nil && (errors.Is(err, ErrBadType) || errors.Is(err, ErrUnpublished)) {
			continue
		} else if err != nil {
			return nil, fmt.Errorf("failed to convert note to page: %w", err)
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
func (pages Pages) Sequenced(notes inote.Notes, tree inote.NoteTree) map[string]PrevNext {
	// folders maps folder ID -> []noteID
	var folders map[string][]string
	for _, page := range pages {
		parentID := notes[page.ID].ParentID
		if _, ok := folders[parentID]; !ok {
			folders[parentID] = []string{}
		}
		folders[parentID] = append(folders[parentID], page.ID)
	}
}
