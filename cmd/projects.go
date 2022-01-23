package cmd

import "github.com/urfave/cli/v2"

func NewProjectsCmd() *cli.Command {
	return &cli.Command{
		Name:    "projects",
		Aliases: []string{"p"},
		Usage:   "operate on projects",
	}
}
