package html

import "text/template"

const _indexTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <title>{{ .Title }}</title>
  <meta name="viewport" content="width=device-width,initial-scale=1">
  <style>
  .pd{width:4em;flex-shrink:0;padding-bottom:.9em}
  .par div{display:flex}
  </style>
</head>
<body>
  <div class="par">
  {{ range _, $page = .Pages }}
    <div>
      <div class="pd">{{ $page.PublishedAt }}</div>
      <a href="{{ $page.URL }}">{{ $page.Title | html }}</a> 
    </div>
  {{ end }}
  </div>
</body>
</html>`

var index = template.Must(template.New("index").Parse(_indexTemplate))
