package spec

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
)

type ExecCommand struct {
	Exec string `yaml:"exec"`
}

func (ec *ExecCommand) Execute() (io.Reader, error) {
	cmd := exec.Command(ec.ScriptInterpreter(), "-c", ec.Exec)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	done := make(chan struct{})

	scanner := bufio.NewScanner(stdout)

	go func() {
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Printf("%s\n", line)
		}

		done <- struct{}{}
	}()

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	<-done

	if err := cmd.Wait(); err != nil {
		return nil, err
	}

	return stdout, nil
}

func (ec *ExecCommand) ScriptInterpreter() string {
	return "bash"
}

func (s *ExecCommand) ToString() string {
	return s.Exec
}
