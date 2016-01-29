package gotopt

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

// GetOptLong TODO
func GetOptLong(
	argv []string, optString string,
	longOpts []*LongOption, longInd *int) int {

	return parser.GetOptLong(argv, optString, longOpts, longInd)
}

// GetOptLongOnly TODO
func GetOptLongOnly(
	argv []string, optString string,
	longOpts []*LongOption, longInd *int) int {

	return parser.GetOptLongOnly(argv, optString, longOpts, longInd)
}

// GetOptLong TODO
func (p *GetOptParser) GetOptLong(
	argv []string, optString string,
	longOpts []*LongOption, longInd *int) int {

	return p.getOptInternal(argv, optString, longOpts, longInd, false, false)
}

// GetOptLongOnly TODO
func (p *GetOptParser) GetOptLongOnly(
	argv []string, optString string,
	longOpts []*LongOption, longInd *int) int {

	return p.getOptInternal(argv, optString, longOpts, longInd, true, false)
}
