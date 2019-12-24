package main

// Code is heavily inspired from https://github.com/cloud-native-nordics/meetups/generator

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

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
	fmt.Println(cfg.MeetupGroups)
	return nil
}

func load(speakersFile, companiesFile, meetupsDir string) (*Config, error) {
	speakers := []Speaker{}
	speakersData, err := ioutil.ReadFile(speakersFile)
	if err != nil {
		return nil, err
	}
	if err := unmarshal(speakersData, &speakers); err != nil {
		return nil, err
	}

	companies := []Company{}
	companiesData, err := ioutil.ReadFile(companiesFile)
	if err != nil {
		return nil, err
	}
	if err := unmarshal(companiesData, &companies); err != nil {
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
	if err != nil {
		return nil, err
	}

	return &Config{
		Speakers:     speakers,
		Companies:    companies,
		MeetupGroups: meetupGroups,
	}, nil
}
