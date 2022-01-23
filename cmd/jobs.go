package cmd

import (
	"fmt"

	"github.com/ROUKIEN/go-rundeck/cmd/jobs"
	"github.com/urfave/cli/v2"
)

func NewJobsCmd() *cli.Command {
	return &cli.Command{
		Name:    "jobs",
		Aliases: []string{"j"},
		Usage:   "operate on registered jobs",
		Subcommands: []*cli.Command{
			jobs.NewJobsListCmd(),
			{
				Name:  "remove",
				Usage: "remove an existing template",
				Action: func(c *cli.Context) error {
					fmt.Println("removed task template: ", c.Args().First())
					return nil
				},
			},
		},
	}
}
