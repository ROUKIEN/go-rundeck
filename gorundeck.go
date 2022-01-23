package main

import (
	"log"
	"os"
	"sort"

	"github.com/ROUKIEN/go-rundeck/cmd"
	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "workdir",
				Aliases: []string{"w"},
				Value:   ".",
				Usage:   "Load configuration from `FILE`",
			},
		},
		Commands: []*cli.Command{
			cmd.NewJobsCmd(),
			cmd.NewProjectsCmd(),
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
