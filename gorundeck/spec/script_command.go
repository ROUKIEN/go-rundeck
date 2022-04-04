package spec

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

type ScriptCommand struct {
	Script string `yaml:"script"`
	Args   string `yaml:"args"`
}

func (sc *ScriptCommand) Execute() (io.Reader, error) {
	// 1.a create a tmp file with the script content (and chmod +x)
	file, err := ioutil.TempFile("", "go-rundeck*.sh")
	if err != nil {
		return nil, err
	}
	defer os.Remove(file.Name())
	// 1.b write script content in tmp file
	writer := bufio.NewWriter(file)

	for _, line := range strings.Split(sc.Script, "\n") {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return nil, err
		}
	}
	writer.Flush()
	if err := file.Close(); err != nil {
		return nil, err
	}

	err = os.Chmod(file.Name(), 0500)
	if err != nil {
		log.Fatal(err)
	}

	// 2. execute the script with args
	cmd := exec.Command(file.Name(), sc.Args)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	cmd.Stderr = cmd.Stdout

	done := make(chan bool)

	scanner := bufio.NewScanner(stdout)

	go func() {
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Printf("%s\n", line)
		}

		done <- true
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

func (ec *ScriptCommand) ScriptInterpreter() string {
	return "bash"
}

func (s *ScriptCommand) ToString() string {
	return strings.Replace(s.Script, "\n", "\\n", -1)
}
