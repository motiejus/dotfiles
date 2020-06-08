package note

import (
	"fmt"
	"io/ioutil"
	"strings"

	"time"

	"gopkg.in/yaml.v2"
)

// Note is how Joplin understands the note.
type Note struct {
	ID                   string    `yaml:"id"`
	ParentID             string    `yaml:"parent_id"`
	Title                string    `yaml:"-"`
	Body                 string    `yaml:"-"`
	CreatedTime          time.Time `yaml:"created_time"`
	UpdatedTime          time.Time `yaml:"updated_time"`
	IsConflict           int       `yaml:"is_conflict"`
	Latitude             float64   `yaml:"latitude"`
	Longitude            float64   `yaml:"longitude"`
	Altitude             float64   `yaml:"altitude"`
	Author               string    `yaml:"author"`
	SourceUrl            string    `yaml:"source_url"`
	IsTODO               int       `yaml:"is_todo"`
	TODODue              int       `yaml:"todo_due"`
	TODOCompleted        int       `yaml:"todo_completed"`
	Source               string    `yaml:"source"`
	SourceApplication    string    `yaml:"source_application"`
	ApplicationData      string    `yaml:"application_data"`
	Order                int       `yaml:"order"`
	UserCreatedTime      time.Time `yaml:"user_created_time"`
	UserUpdatedTime      time.Time `yaml:"user_updated_time"`
	EncryptionCipherText string    `yaml:"encryption_cipher_text"`
	EncryptionApplied    int       `yaml:"encryption_applied"`
	MarkupLanguage       int       `yaml:"markup_language"`
	IsShared             int       `yaml:"is_shared"`
	BodyHTML             string    `yaml:"body_html"`
	// BaseURL is if `body_html` is provided and contains relative URLs, provide the `base_url` parameter too so that all the URLs can be converted to absolute ones. The base URL is basically where the HTML was fetched from, minus the query (everything after the '?'). For example if the original page was `https://stackoverflow.com/search?q=%5Bjava%5D+test`, the base URL is `https://stackoverflow.com/search`.
	BaseUrl string `yaml:"base_url"`
	// ImageDataUrl contains an image to attach to the note, in [Data URL](https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/Data_URIs) format.
	ImageDataUrl string `yaml:"image_data_url"`
	// CropRect is, if an image is provided, you can also specify an optional rectangle that will be used to crop the image. In format `{ x: x, y: y, width: width, height: height }`
	CropRect string   `yaml:"crop_rect"`
	Type     ItemType `yaml:"type_"`
}

// ItemType is a type of the note. Copied from readme/api.md
type ItemType int

const (
	ItemTypeNote               ItemType = 1
	ItemTypeFolder                      = 2
	ItemTypeSetting                     = 3
	ItemTypeResource                    = 4
	ItemTypeTag                         = 5
	ItemTypeNoteTag                     = 6
	ItemTypeSearch                      = 7
	ItemTypeAlarm                       = 8
	ItemTypeMasterKey                   = 9
	ItemTypeItemChange                  = 10
	ItemTypeNoteResource                = 11
	ItemTypeResourceLocalState          = 12
	ItemTypeRevision                    = 13
	ItemTypeMigration                   = 14
	ItemTypeSmartFilter                 = 15
)

func Parse(in string) (Note, error) {
	var title, body, params string
	titleIdx := strings.Index(in, "\n\n")
	paramsIdx := strings.LastIndex(in, "\n\n")
	title, body, params = in[0:titleIdx], in[min(titleIdx+2, paramsIdx):paramsIdx], in[paramsIdx:]

	var note Note
	if err := yaml.Unmarshal([]byte(params), &note); err != nil {
		return Note{}, err
	}

	note.Title = title
	note.Body = body

	return note, nil
}

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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
