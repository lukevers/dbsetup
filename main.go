package main

import (
	"fmt"
	"github.com/mgutz/ansi"
	"gopkg.in/urfave/cli.v2"
	"os"
)

const (
	Version = "0.2.0"
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
			&cli.StringFlag{
				Name:        "template",
				Aliases:     []string{"t"},
				Usage:       "Set a template variable in the configuration file.",
				Destination: &template,
			},
		},
	}).Run(os.Args)
}

func Run(ctx *cli.Context) error {
	// Open and (eventually) close the database connection
	err := config.Connection.Connect()
	defer config.Connection.Close()

	if err != nil {
		return err
	}

	// Turn off FK checks
	if err = db.Exec("SET FOREIGN_KEY_CHECKS = 0;").Error; err != nil {
		fmt.Println(ansi.Color(err.Error(), "red"))
		os.Exit(1)
	}

	// Truncate all tables specified in the configuration file
	for _, table := range config.Truncate {
		fmt.Println(ansi.Color("Truncating:", "green"), table)
		if err = db.Exec(fmt.Sprintf("TRUNCATE %s", table)).Error; err != nil {
			fmt.Println(ansi.Color(err.Error(), "red"))
			os.Exit(1)
		}
	}

	// Turn FK checks back on
	if err = db.Exec("SET FOREIGN_KEY_CHECKS = 1;").Error; err != nil {
		fmt.Println(ansi.Color(err.Error(), "red"))
		os.Exit(1)
	}

	// Run all updates
	for table, rows := range config.Table {
		for _, row := range rows {
			fmt.Println(ansi.Color("Updating:", "green"), "where", row.String(row.Where))
			if err = db.Table(table).Where(row.Sanitize(row.Where)).Updates(row.Sanitize(row.Update)).Error; err != nil {
				fmt.Println(ansi.Color(err.Error(), "red"))
				os.Exit(1)
			}
		}
	}

	return nil
}
