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

	contents := string(bytes)

	// The following template is allowed:
	//
	//   key:val&key2:val2
	//
	// The delimiter is the `&` character for multiple template updates, and
	// the delimiter is the `:` character for key/val for each template update.
	if template != "" {
		kvs := strings.Split(template, "&")
		for _, k := range kvs {
			tmpl := strings.Split(k, ":")
			if len(tmpl) != 2 {
				return errors.New("A template must be in the form of key:value.")
			}

			contents = strings.Replace(contents, fmt.Sprintf("{{%s}}", tmpl[0]), tmpl[1], -1)
		}
	}

	file, err := hcl.Parse(contents)
	if err != nil {
		return err
	}

	return hcl.DecodeObject(&config, file)
}
