package gotopt

import (
	"errors"
	"fmt"
)

// ErrRequiredArg is the error for when a required argument is missing.
type ErrRequiredArg struct {
	Opt int
}

func (e *ErrRequiredArg) Error() string {
	return fmt.Sprintf("arg required for opt '%c'", e.Opt)
}

// ErrUnknownOpt is the error for when an unknown option is encountered.
type ErrUnknownOpt struct {
	Opt      int
	LongName string
}

func (e *ErrUnknownOpt) Error() string {
	if e.LongName == "" {
		return fmt.Sprintf("unknown option '-%c'", e.Opt)
	}
	return fmt.Sprintf("unknown option '--%s'", e.LongName)
}

var (
	// ErrEmptyArgList is returned by Parser.Parse and Parser.ParseAll when
	// there is an empty argument list.
	ErrEmptyArgList = errors.New("empty arg list")
)
