package spec

import (
	"fmt"
	"io"
)

type ScriptFileCommand struct {
	ScriptFile string `yaml:"scriptfile"`
	Args       string `yaml:"args"`
}

func (ec *ScriptFileCommand) Execute() (io.Reader, error) {
	fmt.Printf("[DEBUG] ScriptFileCommand Execute() is not implemented yet.\n")
	return nil, nil
}

func (ec *ScriptFileCommand) ScriptInterpreter() string {
	return "bash"
}

func (s *ScriptFileCommand) ToString() string {
	return s.ScriptFile
}
