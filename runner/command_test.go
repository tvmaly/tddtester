package runner

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCommand(t *testing.T) {

	cmd := NewCommand("go test")

	assert.NotNil(t, cmd)

}
