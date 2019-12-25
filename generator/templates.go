package main

import "text/template"

var (
	readmeTmpl = template.Must(template.New("").Parse(readmeTmplStr))
)

const (
	readmeTmplStr = `# Meetups organized by {{ .MeetupName }}

## Organizers
{{ range .Organizers }}- {{ . }}
{{ end }}

## Meetups
{{ range .MeetupList }}
### {{ .Title }}

- Date: {{ .DateTime }}{{ if .Recording }}
- Recording: {{ .Recording }}{{ end }}
{{ range .Sponsors }}- {{ .Role }} sponsor: {{ .Company }}{{ end }}
{{ end }}`
)
