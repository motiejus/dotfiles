package conv

import (
	"testing"
	"time"

	"github.com/motiejus/dotfiles/joplin2site/internal/note"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name     string
		note     string
		wantErr  string
		wantNote note.JoplinNote
	}{
		{
			name: "ok, has body",
			note: `Meta

Fonts: https://news.ycombinator.com/item?id=23381513

id: 4c7dd536ce1641afa4df349d87d9d29f
parent_id: 9e651b478a5a43c196c31719300fee6e
type_: 1
`,
			wantNote: note.JoplinNote{
				ID:       "4c7dd536ce1641afa4df349d87d9d29f",
				Title:    "Meta",
				Body:     "Fonts: https://news.ycombinator.com/item?id=23381513",
				ParentID: "9e651b478a5a43c196c31719300fee6e",
				Type:     note.ItemTypeNote,
			},
		},
		{
			name: "ok folder",
			note: `blog

id: 9e651b478a5a43c196c31719300fee6e
updated_time: 2020-06-04T16:07:19.930Z
encryption_applied: 0
type_: 2
`,
			wantNote: note.JoplinNote{
				ID:          "9e651b478a5a43c196c31719300fee6e",
				Title:       "blog",
				Body:        "",
				UpdatedTime: time.Date(2020, 6, 4, 16, 07, 19, 930000000, time.UTC),
				Type:        note.ItemTypeFolder,
			},
		},
		{
			name: "bad yaml",
			note: `blog

bad yaml`,
			wantErr: "yaml: unmarshal errors:\n  line 3: cannot unmarshal !!str `bad yaml` into note.JoplinNote",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			note, err := Parse(tt.note)
			if tt.wantErr != "" {
				assert.EqualError(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.wantNote, note)
		})
	}
}
