package note

import "time"

// JoplinNote is how Joplin understands the note.
type JoplinNote struct {
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
	IsTodo               int       `yaml:"is_todo"`
	TodoDue              int       `yaml:"todo_due"`
	TodoCompleted        int       `yaml:"todo_completed"`
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
	BodyHtml             string    `yaml:"body_html"`
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
