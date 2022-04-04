package spec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	cmd := ExecCommand{
		Exec: `echo "Hello world"`,
	}

	_, err := cmd.Execute()
	assert.Nil(t, err)
}
