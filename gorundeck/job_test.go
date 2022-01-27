package gorundeck

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

const jobDefinition = `
---

id: 47f24c4b-34bb-4416-a9f3-db913cd11044
name: Another job
description: hello
loglevel: INFO
sequence:
  keepgoing: true
  strategy: node-first
  commands:
  - exec: ls
  - script: |-
      #!/bin/bash

      echo "hello world"
    args: -h
  - scriptfile: /usr/bin/foo.sh
    args: -h
  - scripturl: https://getcomposer.org/installer
    args: --quiet
  - exec: uname
`

func TestCustomJobUnmarshaller(t *testing.T) {
	var job Job
	err := yaml.Unmarshal([]byte(jobDefinition), &job)
	assert.Nil(t, err)
	assert.Len(t, *job.Sequence.Commands, 5)
	cmds := *job.Sequence.Commands

	assert.IsType(t, &ExecCommand{}, cmds[0])
	assert.IsType(t, &ScriptCommand{}, cmds[1])
	assert.IsType(t, &ScriptFileCommand{}, cmds[2])
	assert.IsType(t, &ScriptUrlCommand{}, cmds[3])
	assert.IsType(t, &ExecCommand{}, cmds[4])
}
