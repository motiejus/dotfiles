package note

import (
	"fmt"
	"io/ioutil"
	"time"
)

type Notes map[string]Note

func ListNotes(dir string) (Notes, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	notes := make(Notes, len(files))
	for _, file := range files {
		body, err := ioutil.ReadFile(file.Name())
		if err != nil {
			return nil, fmt.Errorf("failed to read %q: %w", file.Name(), err)
		}
		note, err := Parse(string(body))
		if err != nil {
			return nil, fmt.Errorf("failed to parse %q: %w", file.Name(), err)
		}

		notes[note.ID] = note
	}

	return notes, nil
}

// SubPages returns immediate Pages of a tree
func (n Notes) SubPages(name string) (Notes, error) {
	// Find the parent note that will be the parent of the sub-notebook.
	var parentID string
	for _, note := range n {
		if note.Type != ItemTypeFolder {
			continue
		}
		if note.Title == name {
			parentID = note.ID
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

// Shake returns a sub-set of notes which eventually parent to "tld"
func (n Notes) Shake(tld string) (Notes, error) {
	tree := buildTree(n)
	topID := n.getFolderID(tld)
	if topID == "" {
		return nil, fmt.Errorf("tld %q not found", tld)
	}

	children := make(children)
	tree.flatten(topID, children)

	ret := make(map[string]note.Note, len(children))
	for noteID := range children {
		ret[noteID] = n[noteID]
	}

	return ret, nil
}

// getFolderID returns a folder ID for a particular title
func (n Notes) getFolderID(title string) string {
	var parentID string
	for _, note := range n {
		if note.Type != ItemTypeFolder {
			continue
		}
		if note.Title == title {
			parentID = note.ID
		}
	}

	return parentID
}

type tree map[string]children
type children map[string]struct{}

func buildTree(notes Notes) tree {
	ret := make(tree)
	for _, note := range notes {
		if _, ok := ret[note.ParentID]; !ok {
			ret[note.ParentID] = make(children)
		}
		ret[note.ParentID][note.ID] = struct{}{}
	}
	return ret
}

// flatten returns all children of a tree
func (t tree) flatten(id string, acc tree) {
	acc[id] = struct{}{}
	for child := range t {
		flatten(child, acc)
	}
}
