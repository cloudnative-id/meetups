package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

var (
	globalSpeakerMap        = map[SpeakerID]*Speaker{}
	globalCompanyMap        = map[CompanyID]*Company{}
	shouldMarshalAutoMeetup = false
)

type CompanyID string
type SpeakerID string

type StatsFile struct {
	MeetupGroups uint64                 `json:"meetupGroups"`
	AllMeetups   MeetupStats            `json:"allMeetups"`
	PerMeetup    map[string]MeetupStats `json:"perMeetup"`
}

type MeetupStats struct {
	Sponsors      uint64                 `json:"sponsors"`
	SponsorByTier map[SponsorTier]uint64 `json:"sponsorByTier,omitempty"`
	Speakers      uint64                 `json:"speakers"`
	Meetups       uint64                 `json:"meetups"`
	Members       uint64                 `json:"members"`
	TotalRSVPs    uint64                 `json:"totalRSVPs"`
	AverageRSVPs  uint64                 `json:"averageRSVPs"`
	UniqueRSVPs   uint64                 `json:"uniqueRSVPs"`
}

type Config struct {
	Companies    []Company     `json:"companies"`
	Speakers     []Speaker     `json:"speakers"`
	MeetupGroups []MeetupGroup `json:"meetupGroups"`
}

var _ json.Marshaler = &CompanyRef{}
var _ json.Unmarshaler = &CompanyRef{}
var _ json.Unmarshaler = &Company{}
var _ json.Marshaler = &SpeakerRef{}
var _ json.Unmarshaler = &SpeakerRef{}
var _ json.Unmarshaler = &Speaker{}

type SponsorRole string

var (
	SponsorRoleVenue    SponsorRole = "Venue"
	SponsorRoleLongterm SponsorRole = "Longterm"
	SponsorRoleCloud    SponsorRole = "Cloud"
	SponsorRoleFood     SponsorRole = "Food"
	SponsorRoleOther    SponsorRole = "Other"

	ValidSponsorRoles = map[SponsorRole]struct{}{
		SponsorRoleVenue:    {},
		SponsorRoleLongterm: {},
		SponsorRoleCloud:    {},
		SponsorRoleFood:     {},
		SponsorRoleOther:    {},
	}
)

func (c *SponsorRole) UnmarshalJSON(b []byte) error {
	str := ""
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	if _, ok := ValidSponsorRoles[SponsorRole(str)]; !ok {
		return fmt.Errorf("not a valid sponsor role: %q", str)
	}
	*c = SponsorRole(str)
	return nil
}

type SponsorTier string

var (
	SponsorTierLongterm        SponsorTier = "Longterm"
	SponsorTierMeetup          SponsorTier = "Meetup"
	SponsorTierSpeakerProvider SponsorTier = "SpeakerProvider"
	SponsorTierEcosystemMember SponsorTier = "EcosystemMember"
)

type Company struct {
	companyInternal
}

type companyInternal struct {
	ID         CompanyID `json:"id"`
	Name       string    `json:"name"`
	WebsiteURL string    `json:"websiteURL"`
	LogoURL    string    `json:"logoURL"`
	WhiteLogo  bool      `json:"whiteLogo,omitempty"`
}

func (c *Company) UnmarshalJSON(b []byte) error {
	ctest := companyInternal{}
	if err := json.Unmarshal(b, &ctest); err != nil {
		return fmt.Errorf("couldn't marshal company %q: %v", string(b), err)
	}
	c.companyInternal = ctest
	if _, ok := globalCompanyMap[c.ID]; ok {
		log.Fatalf("Duplicate company found: %q", c.ID)
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
		log.Fatalf("Company reference not found %q: %q", cid, string(b))
	}
	*c = CompanyRef{company}
	return nil
}

type Speaker struct {
	speakerInternal
}

type speakerInternal struct {
	ID             SpeakerID  `json:"id"`
	Name           string     `json:"name"`
	Title          string     `json:"title,omitempty"`
	Email          string     `json:"email"`
	Company        CompanyRef `json:"company"`
	Github         string     `json:"github"`
	Twitter        string     `json:"twitter,omitempty"`
	SpeakersBureau string     `json:"speakersBureau"`
}

func (s *Speaker) UnmarshalJSON(b []byte) error {
	stest := speakerInternal{}
	if err := json.Unmarshal(b, &stest); err != nil {
		return fmt.Errorf("couldn't marshal speaker %q: %v", string(b), err)
	}
	s.speakerInternal = stest
	if _, ok := globalSpeakerMap[s.ID]; ok {
		log.Fatalf("Duplicate speaker found: %q", s.ID)
	}
	if s.Company.Company == nil {
		log.Warnf("Speaker %q doesn't have a company", s.ID)
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
	if len(s.SpeakersBureau) != 0 {
		str += fmt.Sprintf(", [Contact](https://www.cncf.io/speaker/%s)", s.SpeakersBureau)
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
		log.Fatalf("Speaker reference not found %q: %q", sid, string(b))
	}
	*s = SpeakerRef{speaker}
	return nil
}

type AutogenMeetupGroup struct {
	Photo        string                    `json:"photo,omitempty"`
	Name         string                    `json:"name"`
	City         string                    `json:"city"`
	Country      string                    `json:"country"`
	Description  string                    `json:"description"`
	SponsorTiers map[CompanyID]SponsorTier `json:"sponsorTiers"`
	AutoMeetups  map[string]AutogenMeetup  `json:"-"`

	members uint64
}

type MeetupGroup struct {
	*AutogenMeetupGroup `json:",inline,omitempty"`

	MeetupID          string            `json:"meetupID"`
	Organizers        []SpeakerRef      `json:"organizers"`
	IgnoreMeetupDates []string          `json:"ignoreMeetupDates,omitempty"`
	CFP               string            `json:"cfpLink"`
	Latitude          float64           `json:"latitude"`
	Longitude         float64           `json:"longitude"`
	EcosystemMembers  []CompanyRef      `json:"ecosystemMembers"`
	Meetups           map[string]Meetup `json:"meetups"`
	MeetupList        MeetupList        `json:"-"`
}

func (mg *MeetupGroup) ApplyGeneratedData() {
	for key := range mg.AutoMeetups {
		m, ok := mg.Meetups[key]
		if !ok {
			found := false
			for _, date := range mg.IgnoreMeetupDates {
				if key == date {
					found = true
				}
			}
			if !found {
				log.Warnf("Didn't find information about meetup at %s on date %q\n", mg.Name, key)
			}
			continue
		}
		autoMeetup := mg.AutoMeetups[key]
		m.AutogenMeetup = &autoMeetup
		mg.Meetups[key] = m
	}
}

// CityLowercase gets the lowercase variant of the city
func (mg *MeetupGroup) CityLowercase() string {
	return strings.ToLower(mg.City)
}

func (mg *MeetupGroup) SetMeetupList() {
	marr := []Meetup{}
	for _, m := range mg.Meetups {
		marr = append(marr, m)
	}
	mg.MeetupList = MeetupList(marr)
	sort.Sort(mg.MeetupList)
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

type AutogenMeetup struct {
	ID        uint64   `json:"id"`
	Photo     string   `json:"photo,omitempty"`
	Name      string   `json:"name"`
	Date      Time     `json:"date,omitempty"`
	Duration  Duration `json:"duration,omitempty"`
	Attendees uint64   `json:"attendees,omitempty"`
	Address   string   `json:"address"`

	// rsvps map the user ID to how many rsvp's they used at this event (themselves + guests)
	rsvps map[uint64]uint64
}

type HumanMeetup struct {
	Recording     string          `json:"recording"`
	Sponsors      []MeetupSponsor `json:"sponsors"`
	Presentations []Presentation  `json:"presentations"`
}

type Meetup struct {
	*AutogenMeetup `json:",inline,omitempty"`
	HumanMeetup    `json:",inline"`
}

type fullMeetup struct {
	*AutogenMeetup `json:",inline,omitempty"`
	HumanMeetup    `json:",inline"`
}

func (m Meetup) MarshalJSON() ([]byte, error) {
	if shouldMarshalAutoMeetup {
		return json.Marshal(fullMeetup{
			AutogenMeetup: m.AutogenMeetup,
			HumanMeetup:   m.HumanMeetup,
		})
	}
	return json.Marshal(m.HumanMeetup)
}

type MeetupSponsor struct {
	Role    SponsorRole `json:"role"`
	Company CompanyRef  `json:"company"`
}

func (m *Meetup) DateTime() string {
	d := m.Date.UTC()
	year, month, day := d.Date()
	hour, min, _ := d.Clock()
	hour2, min2, _ := d.Add(m.Duration.Duration).Clock()
	return fmt.Sprintf("%d %s, %d at %d:%02d - %d:%02d", day, month, year, hour, min, hour2, min2)
}

type Presentation struct {
	Duration  Duration     `json:"duration"`
	Delay     *Duration    `json:"delay,omitempty"`
	Title     string       `json:"title"`
	Slides    string       `json:"slides"`
	Recording string       `json:"recording,omitempty"`
	Speakers  []SpeakerRef `json:"speakers"`

	start time.Time
	end   time.Time
}

func (p *Presentation) StartTime() string {
	return fmt.Sprintf("%d:%02d", p.start.UTC().Hour(), p.start.UTC().Minute())
}

func (p *Presentation) EndTime() string {
	return fmt.Sprintf("%d:%02d", p.end.UTC().Hour(), p.end.UTC().Minute())
}
