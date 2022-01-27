package gorundeck

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

type Job struct {
	Name        string  `yaml:"name"`
	ID          string  `yaml:"id"`
	Description *string `yaml:"description"`
	LogLevel    string  `yaml:"loglevel"`
	Sequence
}

type Sequence struct {
	KeepGoing bool      `yaml:"keepgoing"`
	Strategy  string    `yaml:"strategy"`
	Commands  *Commands `yaml:"commands"`
}

type Command interface {
	ToString() string
}

type Commands []Command

func (c *Commands) UnmarshalYAML(value *yaml.Node) error {
	type tmpCmd struct {
		Exec       *string `yaml:"exec"`
		Script     *string `yaml:"script"`
		ScriptFile *string `yaml:"scriptfile"`
		ScriptURL  *string `yaml:"scripturl"`
	}

	for _, content := range value.Content {
		var cmd tmpCmd
		err := content.Decode(&cmd)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
		}

		var ec ExecCommand
		if cmd.Exec != nil {
			content.Decode(&ec)
			*c = append(*c, &ec)
			continue
		}

		var sc ScriptCommand
		if cmd.Script != nil {
			content.Decode(&sc)
			*c = append(*c, &sc)
			continue
		}

		var sfc ScriptFileCommand
		if cmd.ScriptFile != nil {
			content.Decode(&sfc)
			*c = append(*c, &sfc)
			continue
		}

		var suc ScriptUrlCommand
		if cmd.ScriptURL != nil {
			content.Decode(&suc)
			*c = append(*c, &suc)
			continue
		}
	}

	return nil
}

type ExecCommand struct {
	Exec string `yaml:"exec"`
}

func (s *ExecCommand) ToString() string {
	return s.Exec
}

type ScriptCommand struct {
	Script string `yaml:"script"`
	Args   string `yaml:"args"`
}

func (s *ScriptCommand) ToString() string {
	return strings.Replace(s.Script, "\n", "\\n", -1)
}

type ScriptFileCommand struct {
	ScriptFile string `yaml:"scriptfile"`
	Args       string `yaml:"args"`
}

func (s *ScriptFileCommand) ToString() string {
	return s.ScriptFile
}

type ScriptUrlCommand struct {
	ScriptUrl string `yaml:"scripturl"`
	Args      string `yaml:"args"`
}

func (s *ScriptUrlCommand) ToString() string {
	return s.ScriptUrl
}
