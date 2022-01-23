package gorundeck

type Job struct {
	Name        string  `yaml:"name"`
	Description *string `yaml:"description"`
	LogLevel    string  `yaml:"loglevel"`
	Sequence
}

type Sequence struct {
	KeepGoing bool       `yaml:"keepgoing"`
	Strategy  string     `yaml:"strategy"`
	Commands  *[]Command `yaml:"commands"`
}

type Command interface{}

type ExecCommand struct {
	Exec string `yaml:"exec"`
}

type ScriptCommand struct {
	Script string `yaml:"script"`
	Args   string `yaml:"args"`
}

type ScriptFileCommand struct {
	ScriptFile string `yaml:"scriptfile"`
	Args       string `yaml:"args"`
}

type ScriptUrlCommand struct {
	ScriptUrl string `yaml:"scripturl"`
	Args      string `yaml:"args"`
}
