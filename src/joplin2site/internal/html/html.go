package html

import (
	"fmt"

	"github.com/motiejus/dotfiles/joplin2site/internal/note"
	"github.com/motiejus/dotfiles/joplin2site/internal/page"
)

type Files map[string][]byte

// Decorate accepts a TLD and returns files
func Decorate(dir, tld string) (Files, error) {
	noteTree, err := page.BuildTree(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to build template tree: %w", err)
	}
	notes := noteTree.SubNotes(tld)

	// find index page
	var indexNote note.Note
	for _, note := range notes {
		if note.URL == "/" {
			indexNote = note
		}
	}

	indexPage, err := page.FromNote(indexNote)
	if err != nil {
		return err
	}

	return nil
}
