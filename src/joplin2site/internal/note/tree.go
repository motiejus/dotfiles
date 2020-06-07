package note

import (
	"fmt"
	"io/ioutil"
)

type Notes map[string]*Note

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

		notes[note.ID] = &note
	}

	return notes, nil
}

// Shake returns a sub-set of notes which eventually parent to "tld"
func (notes Notes) Shake(tld string) (Notes, error) {
	topID := notes.GetFolderID(tld)
	if topID == "" {
		return nil, fmt.Errorf("tld %q not found", tld)
	}

	children := make(children)
	buildTree(notes).flatten(topID, children)

	ret := make(Notes, len(children))
	for noteID := range children {
		ret[noteID] = notes[noteID]
	}

	return ret, nil
}

// GetFolderID returns a folder ID for a particular title
func (notes Notes) GetFolderID(title string) string {
	var parentID string
	for _, note := range notes {
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
func (t tree) flatten(id string, acc children) {
	acc[id] = struct{}{}
	for child := range t {
		t.flatten(child, acc)
	}
}
