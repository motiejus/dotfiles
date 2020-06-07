package html

import (
	"sort"

	"github.com/motiejus/dotfiles/joplin2site/internal/page"
)

type Files map[string][]byte

// Decorate accepts valid pages and prepares output
func Decorate(pages []page.Page) Files {
	sort.Slice(pages, func(i, j int) bool {
		return pages[i].PublishedAt.Before(pages[j].PublishedAt)
	})

	return nil
}
