package main

// Code is heavily inspired from https://github.com/cloud-native-nordics/meetups/generator

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"github.com/spf13/pflag"
	"sigs.k8s.io/yaml"
)

var speakersFile = pflag.String("speakers-file", "speakers.yaml", "File that contains people who have spoken at meetup groups")
var companiesFile = pflag.String("companies-file", "companies.yaml", "File that contains companies who have sponsored and spoken at meetup groups")
var rootDir = pflag.String("meetups-dir", ".", "Directory that has all meetup groups as subfolders, each subfolder contains a meetup.yaml file")
var unmarshal = yaml.UnmarshalStrict

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func run() error {
	pflag.Parse()
	cfg, err := load(*speakersFile, *companiesFile, *rootDir)
	if err != nil {
		return err
	}

	out, err := exec(cfg)
	if err != nil {
		return err
	}
	return apply(out, *rootDir)
}

func load(speakersFile, companiesFile, meetupsDir string) (*Config, error) {

	companies := []Company{}
	companiesData, err := ioutil.ReadFile(companiesFile)
	if err != nil {
		return nil, err
	}
	if err := unmarshal(companiesData, &companies); err != nil {
		return nil, err
	}

	speakers := []Speaker{}
	speakersData, err := ioutil.ReadFile(speakersFile)
	if err != nil {
		return nil, err
	}
	if err := unmarshal(speakersData, &speakers); err != nil {
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

		meetupsFile := filepath.Join(path, "meetup.yaml")
		if _, err := os.Stat(meetupsFile); os.IsNotExist(err) {
			return nil
		} else if err != nil {
			return err
		}

		mg := MeetupGroup{}
		mg.FilePath = path
		mgData, err := ioutil.ReadFile(meetupsFile)
		if err != nil {
			return err
		}
		if err := unmarshal(mgData, &mg); err != nil {
			return err
		}
		meetupGroups = append(meetupGroups, mg)

		return nil
	})

	fmt.Println(meetupGroups[0].Meetups["20190429"].Presentations[0].Speakers[0])
	if err != nil {
		return nil, err
	}

	return &Config{
		Speakers:     speakers,
		Companies:    companies,
		MeetupGroups: meetupGroups,
	}, nil
}

func exec(cfg *Config) (map[string][]byte, error) {
	result := map[string][]byte{}
	for _, mg := range cfg.MeetupGroups {
		err := mg.SetMeetupList()
		if err != nil {
			return nil, err
		}

		b, err := tmpl(readmeTmpl, mg)
		if err != nil {
			return nil, err
		}

		path := filepath.Join(mg.FilePath, "README.md")
		result[path] = b
	}

	return result, nil
}

func tmpl(t *template.Template, obj interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := t.Execute(&buf, obj); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
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

func writeFile(path string, b []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	return ioutil.WriteFile(path, b, 0644)
}
