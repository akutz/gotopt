package gotopt

import (
	"fmt"
)

// ErrRequiredArg is the error for when a required argument is missing.
type ErrRequiredArg struct {
	Opt int
}

func (e *ErrRequiredArg) Error() string {
	return fmt.Sprintf("arg required for opt %c", e.Opt)
}

// ErrUnknownOpt is the error for when an unknown option is encountered.
type ErrUnknownOpt struct {
	Opt int
}

func (e *ErrUnknownOpt) Error() string {
	return fmt.Sprintf("unknown option %c", e.Opt)
}
