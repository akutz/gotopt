package gotopt

type getOptData struct {
	optInd int
	optErr bool
	optOpt int
	optArg string

	initialized    bool
	nextChar       *int
	ordering       OrderTypes
	posixlyCorrect bool
	firstNonOpt    int
	lastNonOpt     int
}

// GetOptParser can be used to parse multiple argument slices.
type GetOptParser struct {
	OptInd int
	OptErr bool
	OptOpt int
	OptArg string

	data *getOptData
}

// OrderTypes are RequireOrder, Permute, and ReturnInOrder
type OrderTypes int

const (
	// RequireOrder means don't recognize them as options;
	// stop option processing when the first non-option is seen.
	//
	// This mode of operation is selected by either setting the environment
	// variable POSIXLY_CORRECT, or using `+' as the first character
	// of the list of option characters.
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
	Val  int
}

type longOptList struct {
	p    *LongOption
	next *longOptList
}
