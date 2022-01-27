package jobs

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/ROUKIEN/go-rundeck/gorundeck"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

func NewJobsRunCmd() *cli.Command {
	return &cli.Command{
		Name:  "run",
		Usage: "run a command by its ID",
		Action: func(c *cli.Context) error {
			jobUuid := c.Args().Get(0)
			if jobUuid == "" {
				return errors.New("job uuid must be provided")
			}

			jobs, err := filesToJobs(c.String("workdir"))
			if err != nil {
				log.Fatal(err)
				return err
			}

			job := findJobByID(jobUuid, jobs)
			if job == nil {
				return errors.New("no job found")
			}

			fmt.Printf("Will execute the following commands:\n")
			for i, cmd := range *job.Sequence.Commands {
				fmt.Printf("\t %d. %s\n", i, cmd.ToString())
			}

			return nil
		},
	}
}

func findJobByID(ID string, jobs []*gorundeck.Job) *gorundeck.Job {
	for _, job := range jobs {
		if job.ID == ID {
			return job
		}
	}

	return nil
}

func filesToJobs(path string) ([]*gorundeck.Job, error) {
	files, err := ioutil.ReadDir(path + "/jobs/")
	if err != nil {
		return nil, err
	}

	jobs := make([]*gorundeck.Job, 0)
	for _, file := range files {
		filename := filepath.Base(file.Name())
		yamlFile, err := ioutil.ReadFile(path + "/jobs/" + filename)

		if err != nil {
			return nil, err
		}

		var job gorundeck.Job

		err = yaml.Unmarshal(yamlFile, &job)
		if err != nil {
			return nil, err
		}

		jobs = append(jobs, &job)
	}

	return jobs, nil
}
