package cmd

import (
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
			jobs.NewJobsRunCmd(),
		},
	}
}
