/*
Package gotopt is a port of the GNU C getopt library -
https://sourceware.org/git/?p=glibc.git;a=blob;f=posix/getopt.c;hb=HEAD.

While not a one-to-one port, this project maintains the logic and behavior of
the stdlib getopt library.
*/
package gotopt

import (
	"log"
	"os"
)

var (
	debug = os.Getenv("GOTOPT_DEBUG") == "true"
)

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

// GotOpt contains state information about the options.
type GotOpt struct {

	// argv are the arguments to parse
	argv []string

	// OptArg is for communication from `getopt' to the caller.
	// When `getopt' finds an option that takes an argument,
	// the argument value is returned here.
	// Also, when `ordering' is ReturnInOrder,
	// each non-option argv-element is returned here.
	OptArg string

	// OptInd is the index in argv of the next element to be scanned.
	// This is used for communication to and from the caller
	// and for communication between successive calls to `getopt'.
	//
	// On entry to `getopt', zero means this is the first call; initialize.
	//
	// When `getopt' returns EOF, this is the index of the first of the
	// non-option elements that the caller should itself scan.
	//
	// Otherwise, `optInd' communicates from one call to the next
	// how much of argv has been scanned so far.
	//
	// XXX 1003.2 says this must be 1 before any call.
	OptInd int

	// OptErr is for callers store zero here to inhibit the error message
	// for unrecognized options.
	OptErr int

	// OptOpt is set to an option character which was unrecognized.
	// This must be initialized on some systems to avoid linking in the
	// system's own getopt implementation.
	OptOpt rune

	// initialized is if the private members have been initialized
	initialized bool

	// nextchar is the next char to be scanned in the option-element
	// in which the last option character we returned was found.
	// This allows us to pick up the scan where we left off.
	//
	// If this is zero, or a null string, it means resume the scan
	// by advancing to the next argv-element.
	nextChar string

	// posixlyCorrect is the value of the POSIXLY_CORRECT environment variable.
	posixlyCorrect string

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

	// firstNonOpt is the index in argv of the first of the non-options that
	// have been skipped
	firstNonOpt int

	// lastNonOpt is the index in argv of the last of the non-options that
	// have been skipped
	lastNonOpt int
}

// NewGotOpt initializes a new GotOpt instance.
func NewGotOpt(argv ...string) *GotOpt {

	g := &GotOpt{
		OptInd: 1,
		OptOpt: '?',

		argv:           argv,
		firstNonOpt:    1,
		lastNonOpt:     1,
		nextChar:       "",
		posixlyCorrect: os.Getenv("POSIXLY_CORRECT"),
	}

	if len(g.argv) == 0 {
		return g
	}

	// determine how to handle the ordering of options and nonoptions.
	if g.argv[0][0] == '-' {
		g.ordering = ReturnInOrder
		g.argv[0] = g.argv[0][1:]
	} else if g.argv[0][0] == '+' {
		g.ordering = RequireOrder
		g.argv[0] = g.argv[0][1:]
	} else if g.posixlyCorrect != "" {
		g.ordering = RequireOrder
	} else {
		g.ordering = Permute
	}
	return g
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
func (g *GotOpt) exchange(argv []string) {

	bottom := g.firstNonOpt
	middle := g.lastNonOpt
	top := g.OptInd
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

			// exclude the moved bottom segment from further swapping.
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
			// exclude the moved top segment from further swapping.
			bottom += len
		}
	}

	// update records for the slots the non-options now occupy.
	g.firstNonOpt += (g.OptInd - g.lastNonOpt)
	g.lastNonOpt = g.OptInd
}

func logln(s ...interface{}) {
	if !debug {
		return
	}
	log.Println(s)
}
