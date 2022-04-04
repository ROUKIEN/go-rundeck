package jobs

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/ROUKIEN/go-rundeck/gorundeck/spec"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

func NewJobsRunCmd() *cli.Command {
	return &cli.Command{
		Name:   "run",
		Usage:  "run a command by its ID",
		Action: executeJobsRun,
	}
}

func executeJobsRun(c *cli.Context) error {
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
		stdout, err := cmd.Execute()
		if err != nil {
			return err
		}

		fmt.Printf("Command output:\n")
		if stdout != nil {
			// scanner := bufio.NewScanner(stdout)
			// for scanner.Scan() {
			// 	line := scanner.Text()
			// 	fmt.Printf("%s\n", line)
			// }
		} else {
			fmt.Printf("%s\n", "no output available")
		}
	}

	return nil
}

func findJobByID(ID string, jobs []*spec.Job) *spec.Job {
	for _, job := range jobs {
		if job.ID == ID {
			return job
		}
	}

	return nil
}

func filesToJobs(path string) ([]*spec.Job, error) {
	files, err := ioutil.ReadDir(path + "/jobs/")
	if err != nil {
		return nil, err
	}

	jobs := make([]*spec.Job, 0)
	for _, file := range files {
		filename := filepath.Base(file.Name())
		yamlFile, err := ioutil.ReadFile(path + "/jobs/" + filename)

		if err != nil {
			return nil, err
		}

		var job spec.Job

		err = yaml.Unmarshal(yamlFile, &job)
		if err != nil {
			return nil, err
		}

		jobs = append(jobs, &job)
	}

	return jobs, nil
}
