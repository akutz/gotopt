/*
Package gotopt is a port of the GNU C getopt library -
https://sourceware.org/git/?p=glibc.git;a=blob;f=posix/getopt.c;hb=HEAD.

While not a one-to-one port, this project maintains the logic and behavior of
the stdlib getopt library.
*/
package gotopt

import (
	"fmt"
	"os"
	"strings"
)

var (
	// OptInd is
	OptInd = 1

	// OptErr is
	OptErr = true

	// OptOpt is
	OptOpt int = '?'

	// OptArg is
	OptArg string

	parser = NewGetOptParser()
)

// GetOpt TODO
func GetOpt(argv []string, optString string) int {

	return parser.GetOpt(argv, optString)
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

// NewGetOptParser initializes a new instance of the GetOptParser class.
func NewGetOptParser() *GetOptParser {
	return &GetOptParser{
		OptInd: 1,
		OptErr: true,
		OptOpt: '?',
		data:   &getOptData{},
	}
}

// GetOpt TODO
func (p *GetOptParser) GetOpt(argv []string, optString string) int {

	return p.getOptInternal(argv, optString, nil, nil, false, false)
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

func (p *GetOptParser) getOptInternal(argv []string, optString string,
	longOpts []*LongOption, longInd *int,
	longOnly bool, posixlyCorrect bool) int {

	if longInd == nil {
		var tmpLongInd *int
		longInd = tmpLongInd
	}

	// the global parser. init the parser's OptInd and OptErr vars from the
	// global equivalents. then, just before this function returns, update
	// the global OptInd, OptArg, and OptOpt vals with the equivalent values
	// from the parser
	if p == parser {
		p.OptInd = OptInd
		p.OptErr = OptErr
		defer func() {
			OptInd = p.OptInd
			OptArg = p.OptArg
			OptOpt = p.OptOpt
		}()
	}

	p.data.optInd = p.OptInd
	p.data.optErr = p.OptErr

	r := getOptInternalR(
		len(argv), argv, optString,
		longOpts, longInd, longOnly,
		p.data, posixlyCorrect)

	p.OptInd = p.data.optInd
	p.OptArg = p.data.optArg
	p.OptOpt = p.data.optOpt

	return r
}

func getOptInternalR(argc int, argv []string, optString string,
	longOpts []*LongOption, longInd *int,
	longOnly bool, d *getOptData, posixlyCorrect bool) int {

	logf("argc=%d\n", argc)

	printErrors := d.optErr

	if argc < 1 {
		logln("argc < 1")
		return -1
	}

	d.optArg = ""

	if d.optInd == 0 || !d.initialized {
		if d.optInd == 0 {
			// don't scan argv[0], the program name
			d.optInd = 1
		}
		optString = getOptInit(argc, argv, optString, d, posixlyCorrect)
		d.initialized = true
	} else if optString[0] == '-' || optString[0] == '+' {
		if len(optString) > 0 {
			optString = optString[1:]
		}
	}
	if optString[0] == ':' {
		printErrors = false
	}

	nonOptionP := func() bool {
		return argv[d.optInd][0] != '-' || len(argv[d.optInd]) == 1
	}

	if d.nextChar == nil {

		logf("d.optInd=%d\n", d.optInd)
		logln("d.nextChar == nil")

		// advance to the next ARGV-element

		// give FIRST_NONOPT & LAST_NONOPT rational values if OPTIND has been
		// oved back by the user (who may also have changed the arguments).
		if d.lastNonOpt > d.optInd {
			d.lastNonOpt = d.optInd
		}
		if d.firstNonOpt > d.optInd {
			d.firstNonOpt = d.optInd
		}

		logf("d.firstNonOpt=%d\n", d.firstNonOpt)
		logf("d.lastNonOpt=%d\n", d.lastNonOpt)
		logf("d.ordering=%d\n", d.ordering)

		if d.ordering == Permute {

			logln("d.ordering=permute")

			// if we have just processed some options following some
			// non-options, exchange them so that the options come first.
			if d.firstNonOpt != d.lastNonOpt &&
				d.lastNonOpt != d.optInd {

				logln("exchange(argv, d)")
				exchange(argv, d)
			} else if d.lastNonOpt != d.optInd {

				logln("d.firstNonOpt = d.optInd")
				d.firstNonOpt = d.optInd
			}

			// skip any additional non-options and extend the range of
			// non-options previously skipped
			for d.optInd < argc && nonOptionP() {
				d.optInd++
			}

			d.lastNonOpt = d.optInd
			logf("d.optInd=%d\n", d.optInd)
		}

		// the special argv-element `--' means premature end of options.
		// Skip it like a null option, then exchange with previous non-options
		// as if it were an option, then skip everything else like a non-option.
		if d.optInd != argc && argv[d.optInd] == "--" {

			logln(`d.optInd != argc && argv[d.optInd] == "--"`)

			d.optInd++

			if d.firstNonOpt != d.lastNonOpt && d.lastNonOpt != d.optInd {
				exchange(argv, d)
			} else if d.firstNonOpt == d.lastNonOpt {
				d.firstNonOpt = d.optInd
			}
			d.lastNonOpt = argc
			d.optInd = argc
		}

		// if we have done all the argv-elements, stop the scan
		// and back over any non-options that we skipped and permuted.
		if d.optInd == argc {
			// set the next-arg-index to point at the non-options
			// that we previously skipped, so the caller will digest them.
			if d.firstNonOpt != d.lastNonOpt {
				d.optInd = d.firstNonOpt
			}
			logln("d.optInd == argc")
			return -1
		}

		// if we have come to a non-option and did not permute it,
		// either stop the scan or describe it to the caller and pass it by.
		if nonOptionP() {
			if d.ordering == RequireOrder {
				logln("d.ordering == RequireOrder")
				return -1
			}
			d.optArg = argv[d.optInd]
			d.optInd++
			return 1
		}

		// we have found another option-argv-element. skip the initial
		// punctuation.
		nc := 1 + toIntFromBool(len(longOpts) > 0 && argv[d.optInd][1] == '-')
		d.nextChar = &nc
	}

	var c byte
	if d.nextChar != nil {
		c = argv[d.optInd][*d.nextChar]
		*d.nextChar++
	}

	if *d.nextChar >= len(argv[d.optInd]) {
		d.nextChar = nil
	}

	var temp string
	if tempInd := strings.IndexByte(optString, c); tempInd > -1 {
		temp = optString[tempInd:]
	}

	// increment `optind' when we start to process its last character
	if d.nextChar == nil {
		d.optInd++
	}

	if temp == "" || c == ':' || c == ';' {
		if printErrors {
			fmt.Fprintf(
				os.Stderr, "%s: invalid option -- '%c'\n", argv[0], c)
		}
		d.optOpt = int(c)
		return '?'
	}

	lenTemp := len(temp)

	// convenience. Treat POSIX -W foo same as long option --foo
	if lenTemp > 1 && temp[0] == 'W' && temp[1] == ';' {
		if longOpts == nil {
			d.nextChar = nil
			// let the application handle it
			return 'W'
		}
	}

	if lenTemp > 0 && temp[1] == ':' {
		if lenTemp > 2 && temp[2] == ':' {
			// this is an option that accepts an argument optionally
			if d.nextChar != nil {
				d.optArg = argv[d.optInd][*d.nextChar:]
				d.optInd++
			} else {
				d.optArg = ""
				d.nextChar = nil
			}
		} else {
			logln("requires arg")
			// this is an option that requires an argument
			if d.nextChar != nil {
				d.optArg = argv[d.optInd][*d.nextChar:]

				// if we end this ARGV-element by taking the rest as an arg,
				// we must advance to the next element now
				d.optInd++
			} else if d.optInd == argc {
				if printErrors {
					fmt.Fprintf(
						os.Stderr,
						"%s: option requires an argument -- '%c'\n",
						argv[0], c)
				}
				d.optOpt = int(c)

				logf("optString=%s\n", optString)

				if optString[0] == ':' {
					c = ':'
				} else {
					c = '?'
				}
			} else {
				// we already incremented 'optind' once;
				// increment it again when taking next ARGV-elt as argument.
				d.optArg = argv[d.optInd]
				d.optInd++
				logf("opt=%c arg=%s\n", c, d.optArg)
			}

			d.nextChar = nil
		}
	}

	return int(c)
}

func getOptInit(
	argc int,
	argv []string,
	optString string,
	d *getOptData,
	posixlyCorrect bool) string {

	// start processing options with ARGV-element 1 (since ARGV-element 0
	// is the program name); the sequence of previously skipped
	// non-option ARGV-elements is empty.
	d.firstNonOpt = d.optInd
	d.lastNonOpt = d.optInd

	d.nextChar = nil

	d.posixlyCorrect = posixlyCorrect || envVarExists("POSIXLY_CORRECT")

	// determine how to handle the ordering of options and nonoptions.
	if optString[0] == '-' {
		d.ordering = ReturnInOrder
		if len(optString) > 0 {
			optString = optString[1:]
		}
	} else if optString[0] == '+' {
		d.ordering = RequireOrder
		if len(optString) > 0 {
			optString = optString[1:]
		}
	} else if d.posixlyCorrect {
		d.ordering = RequireOrder
	} else {
		d.ordering = Permute
	}

	logf("optString=%s\n", optString)
	return optString
}

// exchange exchanges two adjacent subsequences of argv.
//
// One subsequence is elements [firstNonOpt,lastNonOpt],
// which contains all the non-options that have been skipped so far.
//
// The other is elements [lastNonOpt,optInd], which contains all
// the options processed since those non-options were skipped.
//
// firstNonOpt and lastNonOpt are relocated so that they describe
// the new indices of the non-options in argv after they are moved.
func exchange(argv []string, d *getOptData) {

	bottom := d.firstNonOpt
	middle := d.lastNonOpt
	top := d.optInd
	var tem string

	// exchange the shorter segment with the far end of the longer segment.
	// that puts the shorter segment into the right place. it leaves the longer
	// segment in the right place overall, but it consists of two parts that
	// need to be swapped next.
	for top > middle && middle > bottom {

		if top-middle > middle-bottom {

			logln("bottom segment is the short one")

			// bottom segment is the short one.
			len := middle - bottom

			// swap it with the top part of the top segment.
			for i := 0; i < len; i++ {
				tem = argv[bottom+i]
				argv[bottom+i] = argv[top-(middle-bottom)+i]
				argv[top-(middle-bottom)+i] = tem
			}

			// exclude the moved bottom segment from further swappind.
			top -= len

		} else {

			logln("top segment is the short one")

			// top segment is the short one.
			len := top - middle

			// swap it with the bottom part of the bottom segment.
			for i := 0; i < len; i++ {
				tem = argv[bottom+i]
				argv[bottom+i] = argv[middle+i]
				argv[middle+i] = tem
			}
			// exclude the moved top segment from further swappind.
			bottom += len
		}
	}

	// update records for the slots the non-options now occupy.
	d.firstNonOpt += (d.optInd - d.lastNonOpt)
	d.lastNonOpt = d.optInd
}
