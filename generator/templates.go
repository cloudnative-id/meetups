package main

import "text/template"

var (
	readmeTmpl = template.Must(template.New("").Parse(readmeTmplStr))
)

const (
	readmeTmplStr = `# Meetups organized by {{ .MeetupName }}

## Organizers
{{ range .Organizers }}- {{ . }}
{{ end }}`
)
