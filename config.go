package main

import (
	"errors"
	"github.com/hashicorp/hcl"
	"gopkg.in/urfave/cli.v2"
	"io/ioutil"
)

var (
	path   string
	config *Config
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

	file, err := hcl.ParseBytes(bytes)
	if err != nil {
		return err
	}

	return hcl.DecodeObject(&config, file)
}
