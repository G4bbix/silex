package main

import (
	"errors"
	"flag"
	"strings"
)

type Config struct {
	regex         string
	casts         []string
	inputFilePath string
	inputStdin 		bool
}

func ParseArgs(args []string) (conf *Config, err error) {
	var config Config
	var castString string

	flags := flag.NewFlagSet("silex", flag.ContinueOnError)
	flags.StringVar(&config.regex, "m", "", "Regex to match")
	flags.StringVar(&config.inputFilePath, "f", "", "File to read data from, if unspecified read from stdin")
	flags.StringVar(&castString, "c", "", "Cast backrefs to types: int, float. e.g. 1 as int; 2 as float")
	error := flags.Parse(args)

	if error != nil {
		return &config, error
	}

	if config.regex == "" {
		return &config, errors.New("No match defined")
	}

	if castString != "" {
		config.casts = strings.Split(castString, ";")
	}

	return &config, nil
}

