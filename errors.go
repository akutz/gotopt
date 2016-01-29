package gotopt

import (
	"fmt"
)

// ErrRequiredArg is the error for when a required argument is missing.
type ErrRequiredArg struct {
	OptOpt int
}

func (e *ErrRequiredArg) Error() string {
	return fmt.Sprintf("arg required for opt '%c'", e.OptOpt)
}

// ErrUnknownOpt is the error for when an unknown option is encountered.
type ErrUnknownOpt struct {
	OptOpt int
	OptArg string
}

func (e *ErrUnknownOpt) Error() string {
	if OptArg == "" {
		return fmt.Sprintf("unknown option '%c'", e.OptOpt)
	}
	return fmt.Sprintf("unknown option '%s'", e.OptArg)
}
