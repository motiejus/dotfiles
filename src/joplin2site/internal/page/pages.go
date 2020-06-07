package page

import (
	"fmt"
	"sort"
	"time"

	"github.com/motiejus/dotfiles/joplin2site/internal/note"
)

type Pages []Page

// SubPages returns immediate Pages of a tree
func SubPages(notes note.Notes, title string) (Pages, error) {
	// Find the parent note that will be the parent of the sub-notebook.
	parentID := notes.GetFolderID(title)
	if parentID == "" {
		return nil, fmt.Errorf("sub-page %q not found", title)
	}

	var pages []Page
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
