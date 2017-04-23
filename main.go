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

	// Truncate all tables specified in the configuration file
	for _, table := range config.Truncate {
		fmt.Println("Truncating:", table)
		if err = db.Exec(fmt.Sprintf("TRUNCATE %s", table)).Error; err != nil {
			fmt.Println(err)
		}
	}

	// Run all updates
	for table, rows := range config.Table {
		for _, row := range rows {

			var (
				where string
				and   string
			)

			for k, v := range row.Where {
				if where != "" && and == "" {
					and = "AND"
				}

				where = fmt.Sprintf(
					"%s %s `%s` = '%s'",
					where,
					and,
					k,
					v,
				)
			}

			fmt.Println("Updating row on", table, "where:", where)
			db.Table(table).Where(where).Updates(row.Update)
		}
	}

	return nil
}
