package main

type SpeakerID string
type CompanyID string
type MeetupGroupID string

type SponsorRole string

var (
	SponsorRoleVenue SponsorRole = "Venue"
	SponsorRoleFood  SponsorRole = "Food"
	SponsorRoleOther SponsorRole = "Other"
)

type Speaker struct {
	ID       SpeakerID `json:"id"`
	Name     string    `json:"name"`
	Title    string    `json:"title"`
	Company  string    `json:"company"`
	Github   string    `json:"github"`
	Twitter  string    `json:"twitter"`
	Linkedin string    `json:"linkedin"`
}

type Company struct {
	ID         CompanyID `json:"id"`
	Name       string    `json:"name"`
	WebsiteURL string    `json:"websiteURL"`
	LogoURL    string    `json:"logoURL"`
}

type MeetupGroup struct {
	MeetupID   string            `json:"meetupID"`
	Organizers []string          `json:"organizers"`
	Meetups    map[string]Meetup `json:"meetups"`
}

type Meetup struct {
	Recording     string          `json:"recording"`
	Sponsors      []MeetupSponsor `json:"sponsors"`
	Presentations []Presentation  `json:"presentations"`
}

type MeetupSponsor struct {
	Role    SponsorRole `json:"role"`
	Company string      `json:"company"`
}

type Presentation struct {
	Duration  string   `json:"duration"`
	Title     string   `json:"title"`
	Slides    string   `json:"slides"`
	Recording string   `json:"recording,omitempty"`
	Speakers  []string `json:"speakers"`
}

type Config struct {
	Companies    []Company     `json:"companies"`
	Speakers     []Speaker     `json:"speakers"`
	MeetupGroups []MeetupGroup `json:"meetupGroups"`
}
