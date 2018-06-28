package language

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGoColorizeSuccessandFailure(t *testing.T) {

	go_colorizer := &GoColorizer{}

	assert.NotNil(t, go_colorizer)

	assert.True(t, go_colorizer.IsSuccess("PASS"), "colorizer for Go did not identify a passing test line")

	assert.True(t, go_colorizer.IsFailure("FAIL"), "colorizer for Go did not identify a passing test line")
}
