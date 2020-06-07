package page

import "html/template"

const _indexForTemplate = `
  <div class="par">
  {{ range . }}
    <div>
      <div class="pd">{{ .PublishedAt }}</div>
      <a href="{{ .URL }}">{{ .Title | html }}</a> 
    </div>
  {{ end }}
  </div>
`

var _indexFor = template.Must(template.New("indexFor").Parse(_indexForTemplate))
