package spec

import (
	"fmt"
	"io"
)

type ScriptUrlCommand struct {
	ScriptUrl string `yaml:"scripturl"`
	Args      string `yaml:"args"`
}

func (ec *ScriptUrlCommand) Execute() (io.Reader, error) {
	fmt.Printf("[DEBUG] ScriptUrlCommand Execute() is not implemented yet.\n")
	return nil, nil
}

func (ec *ScriptUrlCommand) ScriptInterpreter() string {
	return "bash"
}

func (s *ScriptUrlCommand) ToString() string {
	return s.ScriptUrl
}
