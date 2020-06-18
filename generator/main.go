package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"text/template"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"sigs.k8s.io/yaml"
)

var speakersFile = pflag.String("speakers-file", "speakers.yaml", "Point to the speakers.yaml file")
var companiesFile = pflag.String("companies-file", "companies.yaml", "Point to the companies.yaml file")
var rootDir = pflag.String("meetups-dir", ".", "Point to the directory that has all meetup groups as subfolders, each with a meetup.yaml file")
var dryRun = pflag.Bool("dry-run", true, "Whether to actually apply the changes or not")
var validateFlag = pflag.Bool("validate", false, "Whether to validate the current state of the repo content with the spec")
var unmarshal = yaml.UnmarshalStrict

// this maps the locations returned from meetup.com to what we want to use here.
// TODO: Maybe skip this and just use "Århus" directly in our
var cityNameExceptions = map[string]string{
	"Århus": "Aarhus",
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func run() error {
	pflag.Parse()
	cfg, err := load(*companiesFile, *speakersFile, *rootDir)
	if err != nil {
		return err
	}
	if err := update(cfg); err != nil {
		return err
	}
	out, err := exec(cfg)
	if err != nil {
		return err
	}
	if *validateFlag {
		return validate(out, *rootDir)
	}
	return apply(out, *rootDir)
}

func load(companiesPath, speakersPath, meetupsDir string) (*Config, error) {
	companies := []Company{}
	companiesContent, err := ioutil.ReadFile(companiesPath)
	if err != nil {
		return nil, err
	}
	if err := unmarshal(companiesContent, &companies); err != nil {
		return nil, err
	}
	speakers := []Speaker{}
	speakersContent, err := ioutil.ReadFile(speakersPath)
	if err != nil {
		return nil, err
	}
	if err := unmarshal(speakersContent, &speakers); err != nil {
		return nil, err
	}
	meetupGroups := []MeetupGroup{}

	err = filepath.Walk(meetupsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return nil
		}
		// Consider only subdirectories of the root path
		if filepath.Dir(path) != "." {
			return nil
		}
		meetupsFile := filepath.Join(path, "meetup.yaml")
		if _, err := os.Stat(meetupsFile); os.IsNotExist(err) {
			return nil
		} else if err != nil {
			return err
		}
		mg := MeetupGroup{}
		mgContent, err := ioutil.ReadFile(meetupsFile)
		if err != nil {
			return err
		}
		if err := unmarshal(mgContent, &mg); err != nil {
			return err
		}
		meetupGroups = append(meetupGroups, mg)
		return nil
	})
	if err != nil {
		return nil, err
	}
	var wg sync.WaitGroup
	wg.Add(len(meetupGroups))
	// Run the fetching from the meetup API in parallel for all meetup groups to speed things up
	for i := range meetupGroups {
		go func(mg *MeetupGroup) {
			defer wg.Done()
			mg.AutogenMeetupGroup, err = GetMeetupInfoFromAPI(*mg)
			if err != nil {
				log.Fatal(err)
			}
			mg.ApplyGeneratedData()
		}(&meetupGroups[i])
	}
	wg.Wait()

	return &Config{
		Speakers:     speakers,
		Companies:    companies,
		MeetupGroups: meetupGroups,
	}, nil
}

func apply(files map[string][]byte, rootDir string) error {
	for path, fileContent := range files {
		fullPath := filepath.Join(rootDir, path)
		if err := writeFile(fullPath, fileContent); err != nil {
			return err
		}
	}
	return nil
}

func validate(files map[string][]byte, rootDir string) error {
	for path, fileContent := range files {
		fullPath := filepath.Join(rootDir, path)
		actual, err := ioutil.ReadFile(fullPath)
		if err != nil {
			return err
		}
		if !bytes.Equal(actual, fileContent) {
			return fmt.Errorf("%s differs from expected state. expected: \"%s\", actual: \"%s\"", fullPath, fileContent, actual)
		}
	}
	log.Info("Validation succeeded!")
	return nil
}

func tmpl(t *template.Template, obj interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := t.Execute(&buf, obj); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func exec(cfg *Config) (map[string][]byte, error) {
	result := map[string][]byte{}
	shouldMarshalAutoMeetup = false
	for _, mg := range cfg.MeetupGroups {
		mg.SetMeetupList()
		b, err := tmpl(readmeTmpl, mg)
		if err != nil {
			return nil, err
		}
		city := mg.CityLowercase()
		path := filepath.Join(city, "README.md")
		result[path] = b

		path = filepath.Join(city, "meetup.yaml")
		mg.AutogenMeetupGroup = nil
		meetupYAML, err := yaml.Marshal(mg)
		if err != nil {
			return nil, err
		}
		result[path] = meetupYAML
	}
	companiesYAML, err := yaml.Marshal(cfg.Companies)
	if err != nil {
		return nil, err
	}
	result["companies.yaml"] = companiesYAML
	speakersYAML, err := yaml.Marshal(cfg.Speakers)
	if err != nil {
		return nil, err
	}
	result["speakers.yaml"] = speakersYAML
	readmeBytes, err := tmpl(toplevelTmpl, cfg)
	if err != nil {
		return nil, err
	}
	result["README.md"] = readmeBytes
	shouldMarshalAutoMeetup = true
	configJSON, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return nil, err
	}
	result["config.json"] = configJSON
	stats, err := aggregateStats(cfg)
	if err != nil {
		return nil, err
	}
	statsJSON, err := json.MarshalIndent(stats, "", "  ")
	if err != nil {
		return nil, err
	}
	result["stats.json"] = statsJSON
	return result, nil
}

func update(cfg *Config) error {
	for i := range cfg.MeetupGroups {
		mg := &cfg.MeetupGroups[i]

		calcSponsorTiers(mg)

		for j, m := range mg.Meetups {
			if err := setPresentationTimestamps(&m); err != nil {
				return err
			}
			mg.Meetups[j] = m
		}
	}
	return nil
}

func calcSponsorTiers(mg *MeetupGroup) {
	mg.SponsorTiers = map[CompanyID]SponsorTier{}
	for _, c := range mg.EcosystemMembers {
		if c.Company != nil {
			mg.SponsorTiers[c.ID] = SponsorTierEcosystemMember
		}
	}
	for _, m := range mg.Meetups {
		for _, p := range m.Presentations {
			for _, s := range p.Speakers {
				if s.Company.Company != nil {
					mg.SponsorTiers[s.Company.ID] = SponsorTierSpeakerProvider
				}
			}
		}
	}
	for _, o := range mg.Organizers {
		if o.Company.Company != nil {
			mg.SponsorTiers[o.Company.ID] = SponsorTierMeetup
		}
	}
	for _, m := range mg.Meetups {
		for _, s := range m.Sponsors {
			if s.Company.Company != nil {
				if s.Role == SponsorRoleLongterm {
					mg.SponsorTiers[s.Company.ID] = SponsorTierLongterm
				} else {
					mg.SponsorTiers[s.Company.ID] = SponsorTierMeetup
				}
			}
		}
	}
}

func writeFile(path string, b []byte) error {
	if *dryRun {
		fmt.Printf("Would write file %q with contents \"%s\"\n", path, string(b))
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	return ioutil.WriteFile(path, b, 0644)
}
