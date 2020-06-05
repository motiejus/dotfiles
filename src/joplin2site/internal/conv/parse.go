package conv

import (
	"strings"

	"github.com/motiejus/dotfiles/joplin2site/internal/note"
	"gopkg.in/yaml.v2"
)

func Parse(in string) (note.JoplinNote, error) {
	var title, body, params string
	titleIdx := strings.Index(in, "\n\n")
	paramsIdx := strings.LastIndex(in, "\n\n")
	title, body, params = in[0:titleIdx], in[min(titleIdx+2, paramsIdx):paramsIdx], in[paramsIdx:]

	var jnote note.JoplinNote
	if err := yaml.Unmarshal([]byte(params), &jnote); err != nil {
		return note.JoplinNote{}, err
	}

	jnote.Title = title
	jnote.Body = body

	return jnote, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
