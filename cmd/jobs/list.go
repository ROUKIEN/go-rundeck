package jobs

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/ROUKIEN/go-rundeck/gorundeck"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
)

func NewJobsListCmd() *cli.Command {
	return &cli.Command{
		Name:  "list",
		Usage: "list jobs",
		Action: func(c *cli.Context) error {
			files, err := ioutil.ReadDir(c.String("workdir") + "/jobs")
			if err != nil {
				log.Fatal(err)
				return err
			}

			fmt.Printf("%d jobs: \n", len(files))
			for _, file := range files {
				filename := filepath.Base(file.Name())
				yamlFile, err := ioutil.ReadFile(c.String("workdir") + "/jobs/" + filename)

				if err != nil {
					return err
				}

				var job gorundeck.Job

				err = yaml.Unmarshal(yamlFile, &job)
				if err != nil {
					return err
				}

				fmt.Printf(" - %s\n", job.Name)
			}

			return nil
		},
	}
}
