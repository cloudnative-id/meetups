package main

import (
	"fmt"
	"sort"
	"time"
)

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
	FilePath   string
	MeetupID   string            `json:"meetupID"`
	MeetupName string            `json:"meetupName"`
	Organizers []string          `json:"organizers"`
	Meetups    map[string]Meetup `json:"meetups"`
	MeetupList MeetupList        `json:"-"`
}

func (mg *MeetupGroup) SetMeetupList() error {
	mlist := []Meetup{}
	for d, m := range mg.Meetups {
		t, err := time.Parse("20060102", d)
		if err != nil {
			return err
		}
		m.Date = Time{t}
		mlist = append(mlist, m)
	}
	mg.MeetupList = MeetupList(mlist)
	sort.Sort(mg.MeetupList)
	return nil
}

// MeetupList is a slice of meetups implementing sort.Interface
type MeetupList []Meetup

var _ sort.Interface = MeetupList{}

func (ml MeetupList) Len() int {
	return len(ml)
}

func (ml MeetupList) Less(i, j int) bool {
	return ml[i].Date.Time.After(ml[j].Date.Time)
}

func (ml MeetupList) Swap(i, j int) {
	ml[i], ml[j] = ml[j], ml[i]
}

type Meetup struct {
	Date          Time            `json:"date,omitempty"`
	Title         string          `json:"title"`
	Recording     string          `json:"recording"`
	Sponsors      []MeetupSponsor `json:"sponsors"`
	Presentations []Presentation  `json:"presentations"`
}

func (m *Meetup) DateTime() string {
	d := m.Date.UTC()
	year, month, day := d.Date()
	weekday := d.Weekday()
	return fmt.Sprintf("%s, %d %s %d", weekday, day, month, year)
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
