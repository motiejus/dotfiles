package html

const _pageTemplate = `
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
{{ block "body" }}
</body>
</html>
`
