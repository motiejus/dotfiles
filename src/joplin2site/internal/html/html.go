package html

import (
	"fmt"

	"github.com/motiejus/dotfiles/joplin2site/internal/note"
)

type Files map[string][]byte

// Render accepts a TLD and returns URL -> []bytes mapping.
func Render(dir, tld string) (Files, error) {
	notes1, err := note.ListNotes(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to get notes: %w", err)
	}
	_, err := notes1.Shake(tld)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
