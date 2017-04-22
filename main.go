package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v2"
	"os"
)

const (
	Version = "0.0.0"
)

func main() {
	(&cli.App{
		Name:    "dbsetup",
		Usage:   "A CLI for setting up databases",
		Version: Version,
		Action:  Run,
		Before:  LoadConfig,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Usage:       "Path to HCL configuration file.",
				Destination: &path,
			},
		},
	}).Run(os.Args)
}

func Run(ctx *cli.Context) error {

	fmt.Println(*config)

	return nil
}
