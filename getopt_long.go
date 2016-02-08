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

// GetOptLong behaves identically to GetOpt except long options beginning with
// two dashes '--' are also accepted.
func GetOptLong(
	argv []string, optString string,
	longOpts []*LongOption, longInd *int) int {

	return getOptParser.GetOptLong(argv, optString, longOpts, longInd)
}

// GetOptLongOnly behavs identically to GetOptLong except that only long
// options are accepted.
func GetOptLongOnly(
	argv []string, longOpts []*LongOption, longInd *int) int {

	return getOptParser.GetOptLongOnly(argv, longOpts, longInd)
}

// GetOptLong behaves identically to the global GetOptLong function except all
// reads and writes from and to the fields OptArg, OptInd, OptErr, OptOpt are
// instance operations.
func (p *GetOptParser) GetOptLong(
	argv []string, optString string,
	longOpts []*LongOption, longInd *int) int {

	return p.getOptInternal(argv, optString, longOpts, longInd, false, false)
}

// GetOptLongOnly behaves identically to the global GetOptLongOnly function
// except all reads and writes from and to the fields OptArg, OptInd, OptErr,
// OptOpt are instance operations.
func (p *GetOptParser) GetOptLongOnly(
	argv []string, longOpts []*LongOption, longInd *int) int {

	return p.getOptInternal(argv, "", longOpts, longInd, true, false)
}
