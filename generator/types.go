package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/rs/zerolog/log"
)

type SpeakerID string
type CompanyID string
type MeetupGroupID string

var (
	globalCompanyMap = map[CompanyID]*Company{}
	globalSpeakerMap = map[SpeakerID]*Speaker{}
)

type SponsorRole string

var (
	SponsorRoleVenue SponsorRole = "Venue"
	SponsorRoleFood  SponsorRole = "Food"
	SponsorRoleOther SponsorRole = "Other"
)

type Speaker struct {
	speakerInternal
}

type speakerInternal struct {
	ID       SpeakerID  `json:"id"`
	Name     string     `json:"name"`
	Title    string     `json:"title,omitempty"`
	Company  CompanyRef `json:"company"`
	Github   string     `json:"github"`
	Twitter  string     `json:"twitter,omitempty"`
	Linkedin string     `json:"linkedin,omitempty"`
}

func (s *Speaker) UnmarshalJSON(b []byte) error {
	stest := speakerInternal{}
	if err := json.Unmarshal(b, &stest); err != nil {
		return fmt.Errorf("couldn't marshal speaker %q: %v", string(b), err)
	}
	s.speakerInternal = stest
	if _, ok := globalSpeakerMap[s.ID]; ok {
		log.Fatal().Msgf("Duplicate speaker found: %q", s.ID)
	}
	if s.Company.Company == nil {
		log.Warn().Msgf("Speaker %q doesn't have a company", s.ID)
	}
	globalSpeakerMap[s.ID] = s
	return nil
}

func (s Speaker) String() string {
	str := s.Name
	if len(s.Github) != 0 {
		str += fmt.Sprintf(" [@%s](https://github.com/%s)", s.Github, s.Github)
	}
	if len(s.Title) != 0 {
		str += fmt.Sprintf(", %s", s.Title)
	}
	if s.Company.Company != nil {
		str += fmt.Sprintf(", [%s](%s)", s.Company.Name, s.Company.WebsiteURL)
	}
	return str
}

type SpeakerRef struct {
	*Speaker `json:"-"`
}

func (s SpeakerRef) MarshalJSON() ([]byte, error) {
	if s.Speaker == nil {
		return []byte(`""`), nil
	}
	return []byte(`"` + s.ID + `"`), nil
}

func (s *SpeakerRef) UnmarshalJSON(b []byte) error {
	if string(b) == "null" || string(b) == `""` {
		*s = SpeakerRef{}
		return nil
	}
	sid := SpeakerID("")
	if err := json.Unmarshal(b, &sid); err != nil {
		return fmt.Errorf("couldn't marshal speaker %q: %v", string(b), err)
	}
	speaker, ok := globalSpeakerMap[sid]
	if !ok {
		log.Fatal().Msgf("Speaker reference not found %q: %q", sid, string(b))
	}
	*s = SpeakerRef{speaker}
	return nil
}

type Company struct {
	companyInternal
}

func (c Company) String() string {
	var str string
	if len(c.Name) != 0 {
		str += fmt.Sprintf(" [%s](%s)", c.Name, c.WebsiteURL)
	}
	return str
}

type companyInternal struct {
	ID         CompanyID `json:"id"`
	Name       string    `json:"name"`
	WebsiteURL string    `json:"websiteURL"`
	LogoURL    string    `json:"logoURL"`
}

func (c *Company) UnmarshalJSON(b []byte) error {
	ctest := companyInternal{}
	if err := json.Unmarshal(b, &ctest); err != nil {
		return fmt.Errorf("couldn't marshal company %q: %v", string(b), err)
	}
	c.companyInternal = ctest
	if _, ok := globalCompanyMap[c.ID]; ok {
		log.Fatal().Msgf("Duplicate company found: %q", c.ID)
	}
	globalCompanyMap[c.ID] = c
	return nil
}

type CompanyRef struct {
	*Company `json:"-"`
}

func (c CompanyRef) MarshalJSON() ([]byte, error) {
	if c.Company == nil {
		return []byte(`""`), nil
	}
	return []byte(`"` + c.ID + `"`), nil
}

func (c *CompanyRef) UnmarshalJSON(b []byte) error {
	if string(b) == "null" || string(b) == `""` {
		*c = CompanyRef{}
		return nil
	}
	cid := CompanyID("")
	if err := json.Unmarshal(b, &cid); err != nil {
		return fmt.Errorf("couldn't marshal company %q: %v", string(b), err)
	}

	company, ok := globalCompanyMap[cid]
	if !ok {
		log.Fatal().Msgf("Company reference not found %q: %q", cid, string(b))
	}
	*c = CompanyRef{company}
	return nil
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
	Company CompanyRef  `json:"company"`
}

type Presentation struct {
	Duration  string       `json:"duration"`
	Title     string       `json:"title"`
	Slides    string       `json:"slides"`
	Recording string       `json:"recording,omitempty"`
	Speakers  []SpeakerRef `json:"speakers"`
}

type Config struct {
	Companies    []Company     `json:"companies"`
	Speakers     []Speaker     `json:"speakers"`
	MeetupGroups []MeetupGroup `json:"meetupGroups"`
}
