package gotopt

import (
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strings"
)

// Parser can be used to parse multiple argument slices.
type Parser interface {
	// Parse immediately returns a channel on which ParserState instances are
	// receives as they are parsed from the supplied arguments.
	Parse(argv []string) (<-chan ParserState, error)

	// ParseAll parses all arguments and then returns the final ParserState.
	ParseAll(argv []string) (ParserState, error)

	// Opt registers an option with the parser.
	Opt(opt int, longName string, optType OptionTypes, argText, usage string)

	// Usage returns the usage text.
	Usage() string

	// PrintUsage writes the usage text to the provided stream.
	PrintUsage(w io.Writer) error

	// PrintAndIndentUsage writes the usage text to the provided stream with
	// each line indented by 'indent' number of white space characters.
	PrintAndIndentUsage(w io.Writer, indent int) error
}

// ParserState is the current state of the parser.
//
// A ParserState is a node in a double-linked list, with each node having a
// pointer to the first and last elements in the list as well as the previous
// and next ParseState instances as well.
//
// If a ParserState object is received on the channel returned from the Parse
// operation, the ParserState's references to the first, last, previous, and
// next nodes in the list may not yet be established. The list to which a
// ParserState belongs cannot be safely traversed until the entire Parse
// operation completes. Any attempts to do so should be considered unreliable
// and unsupported.
//
// The ParseAll operation does not return until all ParserStates are discovered,
// thus it's safe to immediately begin traversing the list of ParserStates using
// the navigation methods as soon as the ParseAll operation completes.
type ParserState interface {
	// Value returns the result of the interation of the GetOpt loop that this
	// ParserState represents. The value can be an Option, an error, or if
	// there are non-option arguments remaining during the final iteration of
	// the GetOpt loop, an array of strings ([]string).
	Value() interface{}

	// Index returns the index of the ParserState with respect to the total
	// number of ParserState instances created as a result of a Parse or
	// ParseAll operation. If five instances are created, the first ParserState
	// would have an index of 0, the second, 1, the third, 2, the fourth 3, and
	// the fifth, 4.
	Index() int

	// First returns the first ParserState created as a result of a Parse or
	// ParseAll operation.
	First() ParserState

	// Prev returns a flag indicating whether or not there is a previous node
	// and a reference to that node.
	Prev() (ParserState, bool)

	// Next returns a flag indicating whether or not there is a next node
	// and a reference to that node.
	Next() (ParserState, bool)

	// Last returns the last ParserState created as a result of a Parse or
	// ParseAll operation.
	Last() ParserState

	// LookupOpt looks up all the options that match the given option character.
	LookupOpt(opt int) []ParserState

	// LookupOptLong looks up all the options that match the given option name.
	LookupOptLong(opt string) []ParserState
}

// Option is the representation of an option as sent to clients receiving the
// results of a Parse or ParseAll operation.
//
// The value returned by the Index() function is not necessarily the same value
// as as the values returned by the corresponding ParserState index. For
// example:
//
//     -n --time 37 -n
//
// In the above argument list there are three options:
//
//   -n
//   --time
//   -n
//
// Those options will generate three ParserState instances with indices of
// 0, 1, 2.
//
// Those same options will also generate three Option instances, one
// for each ParserState. Those Option instances will also have Index values
// set, but those will be 0, 0, 1.
type Option interface {
	// Opt returns the option character if one was provided; otherwise this
	// function returns zero.
	Opt() int

	// LongName returns the option's long name if one was provided; otherwise
	// this function returns an empty string.
	LongName() string

	// Type returns a value indicating whether or not the option required an
	// argument, took an optional argument, or was simply a flag.
	Type() OptionTypes

	// Value returns the option's argument if one was parsed.
	Value() string

	// Index returns the index of the Option with respect to the total
	// number of Option instances created with the same Opt or LongName
	// values as a result of a Parse or ParseAll operation. If five instances
	// are created, the first ParserState would have an index of 0, the second,
	// 1, the third, 2, the fourth 3, and the fifth, 4.
	Index() int
}

// parsedOpt is an option that's been parsed.
type parsedOpt struct {
	optDef
	value string
	index int
}

func (o *parsedOpt) Opt() int {
	return o.opt
}
func (o *parsedOpt) LongName() string {
	return o.longName
}
func (o *parsedOpt) Type() OptionTypes {
	return o.optType
}
func (o *parsedOpt) Value() string {
	return o.value
}
func (o *parsedOpt) Index() int {
	return o.index
}
func (o *parsedOpt) String() string {
	b := &bytes.Buffer{}
	b.WriteString("&{")
	fmt.Fprintf(b, "Opt:%[1]d|%[1]c", o.opt)
	fmt.Fprintf(b, " LongName:%s", o.longName)
	fmt.Fprintf(b, " Index:%d", o.index)
	fmt.Fprintf(b, " Type:%+v", o.optType)
	fmt.Fprintf(b, " Value:%s", o.value)
	b.WriteString("}")
	return b.String()
}

// the backing struct for the ParserState interface
type parserState struct {
	value interface{}
	index int
	first *parserState
	prev  *parserState
	next  *parserState
	last  *parserState
}

func (p *parserState) Value() interface{} {
	return p.value
}
func (p *parserState) Index() int {
	return p.index
}
func (p *parserState) First() ParserState {
	return p.first
}
func (p *parserState) Prev() (ParserState, bool) {
	if p.prev == nil {
		return nil, false
	}
	return p.prev, true
}
func (p *parserState) Next() (ParserState, bool) {
	if p.next == nil {
		return nil, false
	}
	return p.next, true
}
func (p *parserState) Last() ParserState {
	return p.last
}

func (p *parserState) LookupOpt(opt int) []ParserState {
	v := []ParserState{}
	c := p.first
	for {
		if c == nil {
			debugln("c == nil")
			break
		}
		switch i := c.value.(type) {
		case Option:
			debugf("opt=%c found", opt)
			if i.Opt() == opt {
				v = append(v, c)
			}
		}
		c = c.next
	}
	return v
}

func (p *parserState) LookupOptLong(opt string) []ParserState {
	v := []ParserState{}
	c := p.first
	for {
		if c == nil {
			break
		}
		switch i := c.value.(type) {
		case Option:
			if i.LongName() == opt {
				v = append(v, c)
			}
		}
		c = p.next
	}
	return v
}

// parser is the backing struct for the Parser interface.
type parser struct {
	parsed      bool
	opts        map[*optDef]*optDef
	optsOrdered []*optDef
	shortOpts   map[int]*optDef
	longOpts    map[string]*optDef
	maxUsageLen int
}

// NewParser returns a new parser.
func NewParser() Parser {
	return &parser{
		opts:        map[*optDef]*optDef{},
		shortOpts:   map[int]*optDef{},
		longOpts:    map[string]*optDef{},
		optsOrdered: []*optDef{},
	}
}

// Parse parses the supplied arguments.
func (p *parser) Parse(argv []string) (<-chan ParserState, error) {
	if len(argv) == 0 {
		return nil, ErrEmptyArgList
	}
	c := make(chan ParserState)
	go func() {
		p.parse(argv, c)
		close(c)
	}()
	return c, nil
}

// ParseAll parses the supplied arguments and returns the final ParserState.
func (p *parser) ParseAll(argv []string) (ParserState, error) {
	c, err := p.Parse(argv)
	if err != nil {
		return nil, err
	}
	var ps ParserState
	for ps = range c {
		// do nothing
	}
	return ps, nil
}

func (p *parser) parse(argv []string, c chan<- ParserState) {

	b := &bytes.Buffer{}
	longOpts := []*LongOption{}
	b.WriteString(":W;")

	for _, o := range p.opts {
		debugf("opt.Opt=%[1]d|%[1]c, opt.LongName=%s", o.opt, o.longName)

		if o.opt > 0 {
			b.WriteByte(byte(o.opt))
			if o.optType == RequiredArgument || o.optType == OptionalArgument {
				b.WriteByte(':')
				if o.optType == OptionalArgument {
					b.WriteByte(':')
				}
			}
		}
		if o.longName != "" {
			lo := &LongOption{Name: o.longName, Type: o.optType}
			if o.opt > 0 {
				lo.Val = o.opt
				lo.Flag = nil
			}
			longOpts = append(longOpts, lo)
		}
	}

	longInd := 0
	optString := b.String()
	gop := NewGetOptParser()
	var pf func() int
	if len(longOpts) > 0 {
		pf = func() int {
			return gop.GetOptLong(argv, optString, longOpts, &longInd)
		}
	} else {
		pf = func() int {
			return gop.GetOpt(argv, optString)
		}
	}

	var (
		psPrev *parserState
		psInd  int
	)

	optIndices := map[*optDef]int{}

	for {
		opt := pf()
		if opt == -1 {
			break
		}

		debugf(
			"opt=%[1]d|%[1]c, OptOpt=%[2]d|%[2]c, OptArg=%s",
			opt, gop.OptOpt, gop.OptArg)

		var psCurr *parserState

		switch opt {
		case 0:
			if longInd > -1 && longInd < len(longOpts) {
				if o, ok := p.longOpts[longOpts[longInd].Name]; ok {
					optIdx, optIdxOk := optIndices[o]
					if optIdxOk {
						optIdx++
					} else {
						optIdx = 0
					}
					optIndices[o] = optIdx
					psCurr = &parserState{
						value: &parsedOpt{
							optDef: optDef{
								opt:      o.opt,
								longName: o.longName,
							},
							value: gop.OptArg,
							index: optIdx,
						},
					}
				}
			}
		case ':':
			psCurr = &parserState{
				value: &ErrRequiredArg{gop.OptOpt},
			}
		case '?':
			psCurr = &parserState{
				value: &ErrUnknownOpt{gop.OptOpt, gop.OptArg},
			}
		case 'W':
			psCurr = &parserState{
				value: &ErrUnknownOpt{gop.OptOpt, gop.OptArg},
			}
		default:
			if o, ok := p.shortOpts[opt]; ok {
				optIdx, optIdxOk := optIndices[o]
				if optIdxOk {
					optIdx++
				} else {
					optIdx = 0
				}
				optIndices[o] = optIdx
				psCurr = &parserState{
					value: &parsedOpt{
						optDef: optDef{
							opt:      o.opt,
							longName: o.longName,
						},
						value: gop.OptArg,
						index: optIdx,
					},
				}
			} else {
				psCurr = &parserState{
					value: &ErrUnknownOpt{gop.OptOpt, gop.OptArg},
				}
			}
		}

		if psCurr != nil {
			psCurr.index = psInd
			psCurr.next = nil
			psInd++
			if psPrev == nil {
				psCurr.first = psCurr
			} else {
				psCurr.prev = psPrev
				psPrev.next = psCurr
				psCurr.first = psPrev.first
			}
			psPrev = psCurr
			c <- psCurr
		}
	}

	if gop.OptInd < len(argv) {
		psCurr := &parserState{
			index: psInd,
			value: argv[gop.OptInd:],
		}
		if psPrev == nil {
			psCurr.first = psCurr
		} else {
			psCurr.prev = psPrev
			psPrev.next = psCurr
			psCurr.first = psPrev.first
		}
		psPrev = psCurr
		c <- psCurr
	}

	if psPrev != nil {
		psCurr := psPrev.first
		for {
			psCurr.last = psPrev
			psCurr = psCurr.next
			if psCurr == nil {
				break
			}
		}
	}
}

// optDef is the definition of an option as recorded when registering options.
type optDef struct {
	opt      int
	longName string
	optType  OptionTypes
	argText  string
	desc     string
	usageLen int
}

var optionalArgRx = regexp.MustCompile(`^[\[].+[\]]$`)

// Opt registers an option with the parser.
func (p *parser) Opt(
	opt int,
	longName string,
	optType OptionTypes,
	argText, usage string) {

	if opt <= 0 && longName == "" {
		panic("opt and longName invalid")
	}

	if argText == "" && optType != NoArgument {
		argText = "arg"
	}

	if optType == OptionalArgument && !optionalArgRx.MatchString(argText) {
		argText = fmt.Sprintf("[%s]", argText)
	}

	o := &optDef{
		opt:      opt,
		longName: longName,
		optType:  optType,
		desc:     usage,
		argText:  argText,
	}

	lln := len(o.longName)

	if o.opt > 0 {
		p.shortOpts[o.opt] = o
	}

	if lln > 0 {
		p.longOpts[o.longName] = o
	}

	if o.opt > 0 {
		o.usageLen += 2
		if lln > 0 {
			o.usageLen += 2
		} else if o.optType != NoArgument {
			o.usageLen++
		}
	}

	if lln > 0 {
		o.usageLen += lln + 2
		if o.optType != NoArgument {
			o.usageLen++
		}
	}

	if o.optType != NoArgument {
		o.usageLen += len(argText)
	}

	if o.usageLen > p.maxUsageLen {
		p.maxUsageLen = o.usageLen
	}
	p.opts[o] = o
	p.optsOrdered = append(p.optsOrdered, o)
}

func (p *parser) Usage() string {
	b := &bytes.Buffer{}
	p.PrintUsage(b)
	return b.String()
}

func (p *parser) PrintUsage(w io.Writer) error {
	return p.PrintAndIndentUsage(w, 4)
}

func (p *parser) PrintAndIndentUsage(w io.Writer, indent int) error {

	// "INDENT[-OPT][, [--LONGOPT][ ARG]][VARSPACE][DESCRIP]"
	// :ntx::
	//     -n, --name       The name description.
	//     -t, --time arg   The time description.
	//     -x, --xist [arg] The xist description.
	//         --pulp       The pulp description.

	hasOpt := len(p.shortOpts) > 0
	indentStr := []byte(strings.Repeat(" ", indent))
	maxUsage := p.maxUsageLen + indent + 1

	for _, o := range p.optsOrdered {

		var (
			n, nn int
			err   error
		)

		// write the indent
		if n, err = w.Write(indentStr); err != nil {
			return err
		}
		nn += n

		if o.opt > 0 {
			if n, err = fmt.Fprintf(w, "-%c", o.opt); err != nil {
				return err
			}
			nn += n
		} else if hasOpt {
			if n, err = fmt.Fprintf(w, "  "); err != nil {
				return err
			}
			nn += n
		}

		if o.opt > 0 {
			if o.longName != "" {
				if n, err = fmt.Fprintf(w, ", "); err != nil {
					return err
				}
				nn += n
			}
		} else if hasOpt {
			if n, err = fmt.Fprintf(w, "  "); err != nil {
				return err
			}
			nn += n
		}

		if o.longName != "" {
			if n, err = fmt.Fprintf(w, "--%s", o.longName); err != nil {
				return err
			}
			nn += n
		}

		if o.optType != NoArgument {
			if n, err = fmt.Fprintf(w, " "); err != nil {
				return err
			}
			nn += n
			if n, err = fmt.Fprintf(w, o.argText); err != nil {
				return err
			}
			nn += n
		}

		ws := strings.Repeat(" ", maxUsage-nn)
		if _, err = fmt.Fprintf(w, ws); err != nil {
			return err
		}

		if _, err = fmt.Fprintf(w, o.desc); err != nil {
			return err
		}

		if _, err = fmt.Fprintln(w); err != nil {
			return err
		}
	}

	return nil
}
