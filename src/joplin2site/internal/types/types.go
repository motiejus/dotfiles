package types

import "time"

// Note is a joplin entry
type Note struct {
	Body string

	Id          string    `yaml:"id"`
	ParentId    string    `yaml:"parent_id"`
	CreatedTime time.Time `yaml:"created_time"`
	Type        int       `yaml:"type"`
}

// NoteType is a type of the note. Copied from readme/api.md
type NoteType int

const (
	Note               NoteType = 1
	Folder             NoteType = 2
	Setting            NoteType = 3
	Resource           NoteType = 4
	Tag                NoteType = 5
	NoteTag            NoteType = 6
	Search             NoteType = 7
	Alarm              NoteType = 8
	MasterKey          NoteType = 9
	ItemChange         NoteType = 10
	NoteResource       NoteType = 11
	ResourceLocalState NoteType = 12
	Revision           NoteType = 13
	Migration          NoteType = 14
	SmartFilter        NoteType = 15
)
