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

- Date: {{ .DateTime }}
{{ range .Sponsors }}- {{ .Role }} sponsor: {{ .Company }}{{ end }}

#### Agenda
{{ range .Presentations }}
- **{{ .Title }}**: {{ if .Slides }} [[Slides]]({{ .Slides }}){{ end }}{{ if .Recording }} [[Recording]]({{ .Recording }}){{ end }}{{ range .Speakers }}
	- {{ . }}{{ end }}{{ end }}
{{ end }}`
)
