package gotopt

import (
	"io"
)

// GetOptParser is used to parse options and their possible arguments. The
// GetOpt function can be used to return
type GetOptParser interface {

	// GetOpt returns a channel that returns all the parsed options and
	// their possible arguments.
	GetOpt() chan<- *OptionInfo

	// SetOutStream sets the stream to which the parser should emit standard
	// output.
	SetOutStream(w io.Writer)

	// SetErrStream sets the stream to which the parser should emit errors.
	SetErrStream(w io.Writer)

	// ShowErrors sets a flag indicating whether or not to print error
	// messages for unrecognized options.
	ShowErrors(b bool)
}

// OrderTypes are RequireOrder, Permute, and ReturnInOrder
type OrderTypes int

const (
	// RequireOrder means don't recognize them as options;
	// stop option processing when the first non-option is seen.
	//
	// This mode of operation is selected by either setting the environment
	// variable POSIXLY_CORRECT, or using `+' as the first character
	// of the list of option characters. */
	RequireOrder OrderTypes = iota

	// Permute is the default. We permute the contents of argv as we scan,
	// so that eventually all the non-options are at the end. This allows
	// options to be given in any order, even with programs that were not
	// written to expect this.
	Permute

	// ReturnInOrder is an option available to programs that were written
	// to expect options and other argv-elements in any order and that care
	// about the ordering of the two. We describe each non-option argv-element
	// as if it were the argument of an option with character code 1.
	// Using `-' as the first character of the list of option characters
	// selects this mode of operation.
	ReturnInOrder
)

// OptionTypes are NoArgument, RequiredArgument, OptionalArgument
type OptionTypes int

const (
	// NoArgument is for options that do not take an argument
	NoArgument OptionTypes = iota

	// RequiredArgument is for options that require an argument
	RequiredArgument

	// OptionalArgument is for options that take an optional argument
	OptionalArgument
)

// OptionInfo contains information about a parsed option and it's possible
// argument.
type OptionInfo struct {
	Opt          rune
	LongOpt      string
	Arg          string
	Type         OptionTypes
	Unrecognized bool
}

// LongOption describes the long-named options requested by the application.
// The longOpts argument to GetOptLong or GetOptLongOnly is a slice of
// LongOption structs.
//
// If the field 'flag' is not nil, it points to a variable that is set
// to the value given in the field 'val' when the option is found, but
// left unchanged if the option is not found.
//
// To have a long-named option do something other than set an `int' to
// a compiled-in constant, such as set a value from 'optarg', set the
// option's 'flag' field to zero and its 'val' field to a nonzero
// value (the equivalent single-letter option character, if there is
// one). For long options that have a zero 'flag' field, 'getopt'
// returns the contents of the 'val' field.
type LongOption struct {
	Name string
	Type OptionTypes
	Flag *int
	Val  rune
}

type longOptList struct {
	p    *LongOption
	next *longOptList
}

// getOptState contains state information about the options.
type getOptState struct {
	argv           []string
	optString      string
	longOpts       []*LongOption
	longOnly       bool
	outStream      io.Writer
	errStream      io.Writer
	optErr         bool
	posixlyCorrect bool
	ordering       OrderTypes
}

type getOptParser struct {
	argv      []string
	optString string
	longOpts  []*LongOption
	longOnly  bool
	outStream io.Writer
	errStream io.Writer
	optErr    bool
	optIndex  int
	optChan   chan<- *OptionInfo

	// if the POSIXLY_CORRECT environment variable is set
	posixlyCorrect bool

	// ordering describes how to deal with options that follow non-option
	// argv-elements.
	//
	// If the caller did not specify anything, the default is RequireOrder if
	// the environment variable POSIXLY_CORRECT is defined, Permute otherwise.
	//
	// The special argument `--' forces an end of option-scanning regardless
	// of the value of `ordering'. In the case of ReturnInOrder, only
	// `--' can cause `getopt' to return EOF with `optInd' != ARGC.
	ordering OrderTypes

	// nextChar is the next char to be scanned in the option-element
	// in which the last option character we returned was found.
	// This allows us to pick up the scan where we left off.
	//
	// If this is -1 it means resume the scan by advancing to the next
	// argv-element.
	nextChar int

	// firstNonOpt is the index in argv of the first of the non-options that
	// have been skipped
	firstNonOpt int

	// lastNonOpt is the index in argv of the last of the non-options that
	// have been skipped
	lastNonOpt int
}
