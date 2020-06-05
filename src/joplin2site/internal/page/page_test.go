package page

import (
	"testing"
	"time"

	"github.com/motiejus/dotfiles/joplin2site/internal/note"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFromNote(t *testing.T) {
	tests := []struct {
		name     string
		note     note.Note
		wantPage Page
		wantErr  string
	}{
		{
			name: "ok minimal",
			note: note.Note{
				Body: `<!--
url: /about/
published_at: 2020-06-05T12:42:00Z
-->
This is body`,
			},
			wantPage: Page{
				URL:         "/about/",
				BodyHTML:    "This is body",
				PublishedAt: time.Date(2020, 6, 5, 12, 42, 00, 0, time.UTC),
			},
		},

		{
			name:    "missing metadata opening tag",
			note:    note.Note{},
			wantErr: `missing metadata opening tag`,
		},
		{
			name: "missing metadata closing tag",
			note: note.Note{
				Body: "<!--\n foo bar noend\n\n",
			},
			wantErr: `missing metadata closing tag`,
		},
		{
			name: "bad user's metadata",
			note: note.Note{
				Body: `<!--
foo: /non-existing-key/
-->
`,
			},
			wantErr: "failed to unmarshal user's metadata: yaml: unmarshal errors:\n  line 1: field foo not found in type page.userMeta",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			page, err := FromNote(tt.note)
			if tt.wantErr != "" {
				assert.EqualError(t, err, tt.wantErr)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.wantPage, page)
		})
	}

}
