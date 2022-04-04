package spec

import (
	"fmt"
	"io"

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
	ScriptInterpreter() string
	ToString() string
	Execute() (io.Reader, error)
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
