package main

import "text/template"

var (
	readmeTmpl   = template.Must(template.New("").Parse(readmeTmplStr))
	toplevelTmpl = template.Must(template.New("").Parse(toplevelTmplStr))
)

const (
	readmeTmplStr = `# Meetups organized in {{ .City }}

<img width="50%" align="right" alt="Meetup Group Logo" src="{{ .Photo }}">

## Description

{{ .Description }}

{{if .CFP}}## Submit a talk

If you're interested in speaking in this meetup, fill out this form: {{.CFP}}
{{end}}
## Organizers

{{ range .Organizers }}- {{ . }}
{{end}}{{ range .MeetupList }}
### {{ .Name }}

- Date: {{ .DateTime }}
- Meetup link: https://www.meetup.com/{{ $.MeetupID }}/events/{{ .ID }}{{ if .Recording }}
- Recording: {{ .Recording }}{{end}}{{ if .Attendees }}
- Attendees (according to meetup.com): {{ .Attendees }}{{end}}
{{ range .Sponsors }}{{ if .Company }}- {{ .Role }} sponsor: [{{ .Company.Name }}]({{ .Company.WebsiteURL }}){{end}}
{{end}}
#### Agenda

{{ range .Presentations }}- {{ .StartTime }} - {{ .EndTime }}: {{ .Title }} {{ range .Speakers }}
  - {{ . }}{{end}}{{ if .Slides }}
  - Slides: {{ .Slides }}{{end}}{{ if .Recording }}
  - Recording: {{ .Recording }}{{end}}
{{end}}{{end}}`

	toplevelTmplStr = `# Cloud Native Nordics Meetups

Repository to gather all meetup information and slides from Cloud Native Nordic meetups:

{{ range .MeetupGroups }}* [{{ .City }}]({{ .CityLowercase }}/README.md){{ range .Organizers }}
  * {{ . }}{{end}}
{{end}}
## Join our Community!

### Slack

To facilitate and help each other in between meetups and different geographical locations, we have set up a joined Slack Community.

In order to sign-up, go to [www.cloudnativenordics.com](https://www.cloudnativenordics.com) and enter your e-mail. Shortly hereafter you will receive an email with instructions to join the community.

### Mailing List

In order to share documents and calendar invites across our community, we have set up a Mailing List using Google Groups.

Please join our group at [#cloud-native-nordics](https://groups.google.com/forum/#!forum/cloud-native-nordics)!

### Speaking Opportunities

If you'd like to speak at a meetup, please join our [#cloud-native-nordics-speakers](https://groups.google.com/forum/#!forum/cloud-native-nordics-speakers) Mailing List. In this low-traffic group you can get information about speaking opportunities
across all of the Nordic countries!

### Monthly Calls

We're organizing public monthly community calls where everybody is invited to join.
The calls are recorded and will be uploaded to YouTube afterwards.

The **[meeting agenda](https://docs.google.com/document/d/1JxAZcNrGrK89-ErVOKku7Ik76ccSXPa66S8zdbVDr2g/edit#heading=h.cdvsk7jju5f9)** 
is publicly available, please join the [#cloud-native-nordics](https://groups.google.com/forum/#!forum/cloud-native-nordics) mailing
list to get write-access.
`
)
