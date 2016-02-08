/*
Package gotopt is a pure-go port of the GNU C getopt library
https://goo.gl/cQI7VN. This package also includes functionality built on top
of the getopt port, such as:

* Parser, a channel-based parser for receiving options as they are parsed along
with a navigable, doubly-linked list of the parser state at the time which the
option was parsed.

* The package "github.com/akutz/gotopt/flag", a drop-in replacement for the
golang stdlib "flag" package (https://goo.gl/gKusxh) as well as its de facto
replacement package, "pflag" (https://goo.gl/9wvL9j).

It is also possible to entirely ignore all of the additional features and
use purely the same logic one would with the original getopt functions. The
following example parses the command line using the traditional, getopt logic:

    argv := []string{"ProgramName", "-nt37", "effie"}
    for {
        opt := GetOpt(argv, ":nt:")
        if opt == -1 {
            break
        }
        switch opt {
        case 'n':
            fmt.Println("-n detected")
        case 't':
            fmt.Printf("-t detected, arg=%s\n", OptArg)
        }
    }
    if OptInd < len(argv) {
        fmt.Printf("name is %s\n", argv[OptInd])
    }

The following text will be printed to stdout:

  -n detected
  -t detected, arg=37
  name is effie

The above example can be executed locally by executing the following command
in a shell from inside a cloned version of this repository:

  $ go test -run TestGetOptLoop -v

The example uses standard getopt loop to process arguments using an option
string until the return value is -1.  The OptInd, OptArg, OptOpt, and OptErr
fields exist just as they do in the original getopt library.

It's also possible to use the NewGetOptParser to create an instance of the
GetOptParser type in order to parse multiple argument slices in multiple
goroutines instead of the thread-unsafe, package-level GetOpt, GetOptLong,
and GetOptLongOnly functions. The GetOptParser type exposed all of the same
functions as well as instance-specific versions of the OptInd, OptArg, OptOpt,
and OptErr fields.

For more information on the way the getopt loop behaves, the option string
format, or anything else related to the getopt logic, please see
http://goo.gl/uXjKx9.

This package introduces new functionality build on top of getopt, starting with
the type Parser. Using the NewParser function, a Parser is created that can be
used to parse multiple argument slices across multiple goroutines. However,
unlike the GetOpt loop, Parser returns a channel from the Parse function that
receives options as they are parsed. Specifically, the channel receives
ParserState instances, which may indicate an Option has been parsed, an
error has occurred, or if there are no more options remaining, and there are
any non-option arguments, an array of strings:

    p := NewParser()
    p.Opt('n', "name", NoArgument)
    p.Opt('t', "time", RequiredArgument)

    c, _ := p.Parse([]string{"ProgramName", "-nt37", "effie"})

    for ps := range c {
        switch tv := ps.Value().(type) {
        case Option:
            fmt.Printf("-%c detected", tv.Opt())
            if tv.Opt() == 't' {
                fmt.Printf(", arg=%s", tv.Value())
            }
            fmt.Println()
        case []string:
            fmt.Printf("name is %s\n", tv[0])
        }
    }

The following text will be printed to stdout:

  -n detected
  -t detected, arg=37
  name is effie

The above example can be executed locally by executing the following command
in a shell from inside a cloned version of this repository:

  $ go test -run ^TestGotOptParserParse$ -v

The Parser interface also defines the function ParseAll. Using Parse internally,
ParseAll doesn't return until all of the options and their arguments have been
parsed. ParseAll returns the last, received ParserState, which can then be used
to traverse the list of parsed data:

    p := NewParser()
    p.Opt('n', "name", NoArgument)
    p.Opt('t', "time", RequiredArgument)

    ps, _ := p.ParseAll([]string{"ProgramName", "-nt37", "effie"})
    ps = ps.First()

    for {
        switch tv := ps.Value().(type) {
        case Option:
            fmt.Printf("-%c detected", tv.Opt())
            if tv.Opt() == 't' {
                fmt.Printf(", arg=%s", tv.Value())
            }
            fmt.Println()
        case []string:
            fmt.Printf("name is %s\n", tv[0])
        }
        var ok bool
        if ps, ok = ps.Next(); !ok {
            break
        }
    }

The following text will be printed to stdout:

  -n detected
  -t detected, arg=37
  name is effie

The above example can be executed locally by executing the following command
in a shell from inside a cloned version of this repository:

  $ go test -run ^TestGotOptParserParseAll$ -v
*/
package gotopt
