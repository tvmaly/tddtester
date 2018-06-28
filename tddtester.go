package tddtester

import (
	"io"
)

type Colorizer interface {
	Colorize(r io.Reader, w io.Writer)
	IsFailure(line string) bool
	IsSuccess(line string) bool
}
