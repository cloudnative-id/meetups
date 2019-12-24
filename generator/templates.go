package main

import "text/template"

var (
	readmeTmpl = template.Must(template.New("").Parse(readmeTmplStr))
)

const (
	readmeTmplStr = `# {{ .MeetupName }} Meetup`
)
