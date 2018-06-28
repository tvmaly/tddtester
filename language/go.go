package language

import (
	//"bufio"
	//"github.com/fatih/color"
	"io"
	"strings"
)

type GoColorizer struct{}

func (c *GoColorizer) Colorize(r io.Reader, w io.Writer) {
	return
}

func (c *GoColorizer) IsFailure(line string) bool {

	if strings.Index(line, "FAIL") != -1 {
		return true
	}

	return false

}

func (c *GoColorizer) IsSuccess(line string) bool {

	if strings.Index(line, "PASS") != -1 {
		return true
	}

	return false

}
