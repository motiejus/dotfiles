package page

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/motiejus/dotfiles/joplin2site/internal/note"
)

type Tree struct {
	Notes []note.Note
}

// Buildtree builds a tree from the directory
func BuildTree(dir string) (Tree, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return Tree{}, err
	}
	notes := make([]note.Note, 0, len(files))
	for _, file := range files {
		body, err := ioutil.ReadFile(file.Name())
		if err != nil {
			return Tree{}, fmt.Errorf("failed to read %q: %w", file.Name(), err)
		}
		subnote, err := note.Parse(string(body))
		if err != nil {
			return Tree{}, fmt.Errorf("failed to parse %q: %w", file.Name(), err)
		}

		notes = append(notes, subnote)
	}

	return Tree{
		Notes: notes,
	}, nil
}

// SubPages returns immediate Pages of a tree
func (t *Tree) SubPages(name string) ([]Page, error) {
	// Find the parent note that will be the parent of the sub-notebook.
	var parentID string
	for _, inote := range t.Notes {
		if inote.Type != note.ItemTypeFolder {
			continue
		}
		if inote.Title == name {
			parentID = inote.ID
		}
	}

	if parentID == "" {
		return nil, fmt.Errorf("sub-page %q not found", name)
	}

	var retPages []Page
	for _, inote := range t.Notes {
		if inote.ParentID != parentID {
			continue
		}
		if inote.Type != note.ItemTypeNote {
			continue
		}

		page, err := FromNote(inote)
		if err != nil {
			return nil, fmt.Errorf("failed to convert note to page: %w", err)
		}
		if page.PublishedAt.After(time.Now()) {
			continue
		}

		retPages = append(retPages, page)
	}
	return retPages, nil
}
