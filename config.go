package main

import (
	"errors"
	"fmt"
	"github.com/hashicorp/hcl"
	"gopkg.in/urfave/cli.v2"
	"io/ioutil"
	"strings"
)

var (
	path     string
	template string
	config   *Config
)

type Config struct {
	Connection Connection
	Truncate   []string
	Table      map[string][]Row
}

func LoadConfig(ctx *cli.Context) error {
	if path == "" {
		return errors.New("Please provide a configuration file path.")
	}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// Convert the file to a string
	contents := string(bytes)

	if template != "" {
		tmpl := strings.Split(template, ":")
		if len(tmpl) != 2 {
			return errors.New("A template must be in the form of key:value.")
		}

		contents = strings.Replace(contents, fmt.Sprintf("{{%s}}", tmpl[0]), tmpl[1], -1)
	}

	file, err := hcl.Parse(contents)
	if err != nil {
		return err
	}

	return hcl.DecodeObject(&config, file)
}
