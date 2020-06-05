package note

import (
	"strings"

	"gopkg.in/yaml.v2"
)

func Parse(in string) (JoplinNote, error) {
	var title, body, params string
	titleIdx := strings.Index(in, "\n\n")
	paramsIdx := strings.LastIndex(in, "\n\n")
	title, body, params = in[0:titleIdx], in[min(titleIdx+2, paramsIdx):paramsIdx], in[paramsIdx:]

	var note JoplinNote
	if err := yaml.Unmarshal([]byte(params), &note); err != nil {
		return JoplinNote{}, err
	}

	note.Title = title
	note.Body = body

	return note, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
