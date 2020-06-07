package html

import (
	"fmt"

	"github.com/motiejus/dotfiles/joplin2site/internal/page"
)

type Files map[string][]byte

// Render accepts a TLD and returns URL -> []bytes mapping.
func Render(dir, tld string) (Files, error) {
	noteTree, err := page.BuildTree(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to build template tree: %w", err)
	}
	pages, err := noteTree.SubPages(tld)
	if err != nil {
		return nil, err
	}

	// find index page
	var indexPage page.Page
	for _, ipage := range pages {
		if ipage.URL == "/" {
			indexPage = ipage
		}
	}

	p, err := indexPage.Render(&noteTree)
	if err != nil {
		fmt.Errorf("failed to render index page: %w", err)
	}

	return Files{"/": p}, nil
}
