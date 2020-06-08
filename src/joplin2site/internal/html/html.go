package html

import (
	"fmt"

	inote "github.com/motiejus/dotfiles/joplin2site/internal/note"
	ipage "github.com/motiejus/dotfiles/joplin2site/internal/page"
)

type (
	// Webpage is the resulting URL -> []byte map.
	Webpage map[string][]byte

	// pageContext contains everything that's necessary to render a page.
	pageContext struct {
		ipage.Page

		Prev *ipage.Page
		Next *ipage.Page
	}
)

// Render accepts a TLD and returns URL -> []bytes mapping.
func Render(dir, tld string) (Webpage, error) {
	notes, err := inote.ListNotes(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to get notes: %w", err)
	}
	tree := inote.BuildTree(notes)
	if err := notes.Shake(tree, tld); err != nil {
		return nil, err
	}

	// Convert all publish-able notes to pages.
	pages, err := ipage.ToPublishablePages(notes)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
