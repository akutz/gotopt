package gotopt

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGotOptParserParse(t *testing.T) {
	p := NewParser()
	p.Opt('n', "name", NoArgument, "", "")
	p.Opt('t', "time", RequiredArgument, "", "")

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
}

func TestGotOptParserParseAll(t *testing.T) {
	p := NewParser()
	p.Opt('n', "name", NoArgument, "", "")
	p.Opt('t', "time", RequiredArgument, "", "")

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
}

func TestPrintUsage(t *testing.T) {
	p := NewParser()
	p.Opt('n', "name", NoArgument, "", "A flag indicating the name is a trailing arg")
	p.Opt('t', "time", RequiredArgument, "epoch", "The epoch")
	p.Opt('x', "xist", OptionalArgument, "val", "A value with the answer to our existence")
	p.Opt(0, "play", NoArgument, "", "Whether or not it's time to play")
	p.Opt(0, "fast", OptionalArgument, "mph", "How fast to go")
	p.Opt(0, "slow", RequiredArgument, "mph", "How slow to go")
	p.Opt('h', "", RequiredArgument, "00-23", "The current hour")
	p.Opt('o', "", OptionalArgument, "asc|desc", "The current order")
	p.Opt('z', "", NoArgument, "", "Drop the zero, and get with the hero")

	exp := `    -n, --name       A flag indicating the name is a trailing arg
    -t, --time epoch The epoch
    -x, --xist [val] A value with the answer to our existence
        --play       Whether or not it's time to play
        --fast [mph] How fast to go
        --slow mph   How slow to go
    -h 00-23         The current hour
    -o [asc|desc]    The current order
    -z               Drop the zero, and get with the hero
`
	assert.NoError(t, p.PrintUsage(os.Stdout))
	assert.Equal(t, exp, p.Usage())

	exp = `-n, --name       A flag indicating the name is a trailing arg
-t, --time epoch The epoch
-x, --xist [val] A value with the answer to our existence
    --play       Whether or not it's time to play
    --fast [mph] How fast to go
    --slow mph   How slow to go
-h 00-23         The current hour
-o [asc|desc]    The current order
-z               Drop the zero, and get with the hero
`

	b := &bytes.Buffer{}
	assert.NoError(t, p.PrintAndIndentUsage(b, 0))
	assert.Equal(t, exp, b.String())
}

func TestParserOk(t *testing.T) {
	assertParseOk(t, testParse(t, "tipok01", "--time=37", "-n", "effie"))
	assertParseOk(t, testParse(t, "tipok02", "-n", "--time=37", "effie"))
	assertParseOk(t, testParse(t, "tipok03", "--time=37", "effie", "-n"))
	assertParseOk(t, testParse(t, "tipok04", "effie", "-n", "--time=37"))

	assertParseOk(t, testParse(t, "tipok05", "--time", "37", "-n", "effie"))
	assertParseOk(t, testParse(t, "tipok06", "-n", "--time", "37", "effie"))
	assertParseOk(t, testParse(t, "tipok07", "--time", "37", "effie", "-n"))
	assertParseOk(t, testParse(t, "tipok08", "effie", "-n", "--time", "37"))

	assertParseOk(t, testParse(t, "tipok09", "--ti", "37", "-n", "effie"))
	assertParseOk(t, testParse(t, "tipok10", "-n", "--tim", "37", "effie"))
	assertParseOk(t, testParse(t, "tipok11", "--t", "37", "effie", "-n"))

	assertParseOk(t, testParse(t, "tipok12", "--time", "37", "-n", "effie"))
	assertParseOk(t, testParse(t, "tipok13", "--na", "--time", "37", "effie"))
	assertParseOk(t, testParse(t, "tipok14", "--time", "37", "effie", "--nam"))
	assertParseOk(t, testParse(t, "tipok15", "effie", "--name", "--time", "37"))
}

func TestParserOkW(t *testing.T) {
	assertParseOk(t, testParse(t, "tipokw01", "-W", "time=37", "-n", "effie"))
	assertParseOk(t, testParse(t, "tipokw02", "-n", "-W", "time=37", "effie"))
	assertParseOk(t, testParse(t, "tipokw03", "-W", "time=37", "effie", "-n"))
	assertParseOk(t, testParse(t, "tipokw04", "effie", "-n", "-W", "time=37"))

	assertParseOk(t, testParse(t, "tipokw05", "-W", "time", "37", "-n", "effie"))
	assertParseOk(t, testParse(t, "tipokw06", "-n", "-W", "time", "37", "effie"))
	assertParseOk(t, testParse(t, "tipokw07", "-W", "time", "37", "effie", "-n"))
	assertParseOk(t, testParse(t, "tipokw08", "effie", "-n", "-W", "time", "37"))

	assertParseOk(t, testParse(t, "tipokw09", "-W", "ti", "37", "-n", "effie"))
	assertParseOk(t, testParse(t, "tipokw10", "-n", "-W", "tim", "37", "effie"))
	assertParseOk(t, testParse(t, "tipokw11", "-W", "t", "37", "effie", "-n"))

	assertParseOk(t, testParse(t, "tipokw12", "-W", "time", "37", "-W", "n", "effie"))
	assertParseOk(t, testParse(t, "tipokw13", "-W", "na", "-W", "time", "37", "effie"))
	assertParseOk(t, testParse(t, "tipokw14", "-W", "time", "37", "effie", "-W", "nam"))
	assertParseOk(t, testParse(t, "tipokw15", "effie", "-W", "name", "-W", "time", "37"))
}

func TestParserMissingName(t *testing.T) {
	assertParseNoName(t, testParse(t, "tipnn01", "--time=37", "-n"))
	assertParseNoName(t, testParse(t, "tipnn02", "-n", "--time=37"))
	assertParseNoName(t, testParse(t, "tipnn03", "--time=37", "-n"))
	assertParseNoName(t, testParse(t, "tipnn04", "-n", "--time=37"))

	assertParseNoName(t, testParse(t, "tipnn05", "--time", "37", "-n"))
	assertParseNoName(t, testParse(t, "tipnn06", "-n", "--time", "37"))
	assertParseNoName(t, testParse(t, "tipnn07", "--time", "37", "-n"))
	assertParseNoName(t, testParse(t, "tipnn08", "-n", "--time", "37"))

	assertParseNoName(t, testParse(t, "tipnn09", "--ti", "37", "-n"))
	assertParseNoName(t, testParse(t, "tipnn10", "-n", "--tim", "37"))
	assertParseNoName(t, testParse(t, "tipnn11", "--t", "37", "-n"))
}

func TestParserMissingNameW(t *testing.T) {
	assertParseNoName(t, testParse(t, "tipnnw01", "-W", "time=37", "-n"))
	assertParseNoName(t, testParse(t, "tipnnw02", "-n", "-W", "time=37"))
	assertParseNoName(t, testParse(t, "tipnnw03", "-W", "time=37", "-W", "n"))
	assertParseNoName(t, testParse(t, "tipnnw04", "-W", "na", "-W", "time=37"))

	assertParseNoName(t, testParse(t, "tipnnw05", "-W", "time", "37", "-n"))
	assertParseNoName(t, testParse(t, "tipnnw06", "-n", "-W", "time", "37"))
	assertParseNoName(t, testParse(t, "tipnnw07", "-W", "time", "37", "-W", "n"))
	assertParseNoName(t, testParse(t, "tipnnw08", "-W", "na", "-W", "time", "37"))

	assertParseNoName(t, testParse(t, "tipnnw09", "-W", "ti", "37", "-n"))
	assertParseNoName(t, testParse(t, "tipnnw10", "-W", "nam", "-W", "tim", "37"))
	assertParseNoName(t, testParse(t, "tipnnw11", "-W", "t", "37", "-W", "name"))
}

func TestParserNoTime(t *testing.T) {
	a1 := func(t *testing.T, r *parseTestResult) {
		assert.False(t, r.tfnd)
		assert.IsType(t, &ErrRequiredArg{}, r.err)
		err := r.err.(*ErrRequiredArg)
		assert.EqualValues(t, 't', err.Opt)
		assert.NotEqual(t, "37", r.nsecs)
		assert.False(t, r.nfnd)
		assert.Equal(t, "", r.name)
	}
	a1(t, testParse(t, "tipnt01", "--t"))
	a1(t, testParse(t, "tipnt02", "--ti"))
	a1(t, testParse(t, "tipnt03", "--tim"))
	a1(t, testParse(t, "tipnt04", "--time"))

	a2 := func(t *testing.T, r *parseTestResult) {
		assert.False(t, r.tfnd)
		assert.IsType(t, &ErrRequiredArg{}, r.err)
		err := r.err.(*ErrRequiredArg)
		assert.EqualValues(t, 't', err.Opt)
		assert.NotEqual(t, "37", r.nsecs)
		assert.True(t, r.nfnd)
		assert.Equal(t, "", r.name)
	}
	a2(t, testParse(t, "tipnt05", "-n", "--t"))
	a2(t, testParse(t, "tipnt06", "-n", "--ti"))
	a2(t, testParse(t, "tipnt07", "-n", "--tim"))
	a2(t, testParse(t, "tipnt08", "-n", "--time"))
}

func TestParserUnknownOpt(t *testing.T) {
	a1 := func(t *testing.T, r *parseTestResult, u string) {
		assert.True(t, r.tfnd)
		assert.Equal(t, "37", r.nsecs)
		assert.IsType(t, &ErrUnknownOpt{}, r.err)
		err := r.err.(*ErrUnknownOpt)
		if err.LongName == "" {
			assert.Equal(t, u, fmt.Sprintf("%c", err.Opt))
		} else {
			assert.Equal(t, u, err.LongName)
		}
		assert.False(t, r.nfnd)
		assert.Equal(t, "", r.name)
	}
	a1(t, testParse(t, "tipunkn01a", "--t=37", "-f", "-n", "effie"), "f")
	a1(t, testParse(t, "tipunkn02a", "--ti=37", "-fu", "-n", "effie"), "f")
	a1(t, testParse(t, "tipunkn03a", "--tim=37", "-fub", "-n", "effie"), "f")
	a1(t, testParse(t, "tipunkn04a", "--time=37", "-fubar", "-n", "effie"), "f")

	a1(t, testParse(t, "tipunkn01b", "--t=37", "-W", "f", "-n", "effie"), "f")
	a1(t, testParse(t, "tipunkn02b", "--ti=37", "-W", "fu", "-n", "effie"), "fu")
	a1(t, testParse(t, "tipunkn03b", "--tim=37", "-W", "fub", "-n", "effie"), "fub")
	a1(t, testParse(t, "tipunkn04b", "--time=37", "-W", "fubar", "-n", "effie"), "fubar")

	a2 := func(t *testing.T, r *parseTestResult, u string) {
		assert.True(t, r.tfnd)
		assert.Equal(t, "37", r.nsecs)
		assert.True(t, r.nfnd)
		assert.Equal(t, "", r.name)
		assert.IsType(t, &ErrUnknownOpt{}, r.err)
		err := r.err.(*ErrUnknownOpt)
		if err.LongName == "" {
			assert.Equal(t, u, fmt.Sprintf("%c", err.Opt))
		} else {
			assert.Equal(t, u, err.LongName)
		}
	}
	a2(t, testParse(t, "tipunkn05", "--t=37", "-n", "effie", "-f"), "f")
	a2(t, testParse(t, "tipunkn06", "--ti=37", "-n", "effie", "-W", "f"), "f")
	a2(t, testParse(t, "tipunkn07", "--tim=37", "-n", "effie", "-fu"), "f")
	a2(t, testParse(t, "tipunkn08", "--time=37", "-n", "effie", "-W", "fubar"), "fubar")

	a3 := func(t *testing.T, r *parseTestResult) {
		assert.True(t, r.tfnd)
		assert.Equal(t, "37", r.nsecs)
		assert.IsType(t, &ErrUnknownOpt{}, r.err)
		err := r.err.(*ErrUnknownOpt)
		assert.Equal(t, 0, err.Opt)
		assert.Equal(t, "hello", err.LongName)
		assert.False(t, r.nfnd)
		assert.Equal(t, "", r.name)
	}
	a3(t, testParse(t, "tipunkn09", "--t=37", "--hello", "-f", "-n", "effie"))
	a3(t, testParse(t, "tipunkn10", "--ti=37", "--hello", "-f", "-n", "effie"))
	a3(t, testParse(t, "tipunkn11", "--tim=37", "--hello", "-f", "-n", "effie"))
	a3(t, testParse(t, "tipunkn12", "--time=37", "--hello", "-f", "-n", "effie"))
}

func TestParseOptional(t *testing.T) {

	a1 := func(t *testing.T, r *parseTestResult) {
		assert.True(t, r.tfnd)
		assert.Equal(t, "37", r.nsecs)
		assert.True(t, r.nfnd)
		assert.Equal(t, "effie", r.name)
		assert.True(t, r.xfnd)
		assert.Equal(t, "play", r.xist)
	}
	a1(t, testParse(t, "tipopt01", "--time=37", "-n", "effie", "-xplay"))
	a1(t, testParse(t, "tipopt02", "-n", "--time=37", "effie", "--xist=play"))
	a1(t, testParse(t, "tipopt03", "--time=37", "effie", "-nxplay"))

	// NOTE: This will fail due to getopt not parsing optional parameters for
	// longOpts unless using an explicit equals sign. This StackOverflow
	// post explains it - http://stackoverflow.com/questions/1052746
	//
	// a1(t, testParse(t, "tipopt04", "effie", "-n", "--time=37", "--xist", "play"))
}

func TestParseLongOnly(t *testing.T) {

	a1 := func(t *testing.T, r *parseTestResult) {
		assert.True(t, r.tfnd)
		assert.Equal(t, "37", r.nsecs)
		assert.True(t, r.nfnd)
		assert.Equal(t, "effie", r.name)
		assert.True(t, r.xfnd)
		assert.Equal(t, "play", r.xist)
		assert.True(t, r.pfnd)
	}
	a1(t, testParse(t, "tplo01", "--time=37", "-n", "effie", "-xplay", "--pulp"))
	a1(t, testParse(t, "tplo02", "-n", "--time=37", "effie", "--xist=play", "--p"))
	a1(t, testParse(t, "tplo03", "--time=37", "effie", "-nxplay", "--pu"))
}

func TestParserStateCount(t *testing.T) {
	a1 := func(t *testing.T, ps ParserState, nct int) {
		fps := ps.LookupOpt('n')
		assert.Equal(t, nct, len(fps))
		for _, i := range fps {
			t.Logf("fpts=%+v", i)
		}
	}
	a1(t, testParseAll(t, "tpct01", "--time=37", "-nn", "effie"), 2)
	a1(t, testParseAll(t, "tpct02", "-n", "-nnn", "--time=37", "effie"), 4)
	a1(t, testParseAll(t, "tpct03", "--name", "--time=37", "effie", "-nnn", "--name"), 5)
	a1(t, testParseAll(t, "tpct04", "effie", "-nnn", "-nnnt", "37"), 6)
}

type testParserCountAndIndexData struct {
	Argv    []string
	Count   int
	Indices []int
}

func TestParserCountAndIndex(t *testing.T) {

	payloads := []*testParserCountAndIndexData{
		&testParserCountAndIndexData{
			Argv:    []string{"--time=37", "-nn", "effie"},
			Count:   2,
			Indices: []int{1, 2},
		},
		&testParserCountAndIndexData{
			Argv:    []string{"-n", "-nnn", "--time=37", "effie"},
			Count:   4,
			Indices: []int{0, 1, 2, 3},
		},
		&testParserCountAndIndexData{
			Argv:    []string{"--name", "--time=37", "effie", "-nnn", "--name"},
			Count:   5,
			Indices: []int{0, 2, 3, 4, 5},
		},
		&testParserCountAndIndexData{
			Argv:    []string{"effie", "-nnn", "-nnnt", "37"},
			Count:   6,
			Indices: []int{0, 1, 2, 3, 4, 5},
		},
	}

	testOpt := (*Option)(nil)
	pp := 0

	for x, td := range payloads {
		assert.Equal(t, td.Count, len(td.Indices), "td.Count == len(td.Indices)")
		argv := []string{fmt.Sprintf("tpctidx0%d", x)}
		argv = append(argv, td.Argv...)
		ps := testParseAll(t, argv...)
		fps := ps.LookupOpt('n')
		assert.Equal(t, td.Count, len(fps), "Count")
		for y, fpsi := range fps {
			assert.Implements(t, testOpt, fpsi.Value())
			o := fpsi.Value().(Option)
			assert.Equal(t, td.Indices[y], fpsi.Index(), "ParserState.Index")
			assert.Equal(t, y, o.Index(), "Option.Index")
		}
		pp++
	}

	assert.Equal(t, 4, pp)
}

func TestParserDiffStates(t *testing.T) {
	ps := testParseAll(
		t, "tpdifs01", "--time=37", "-nn", "-xplay", "-t", "47", "effie")

	testOpt := (*Option)(nil)
	testStrArr := []string{}

	psArgs := ps.Last()
	assert.NotNil(t, psArgs)
	assert.NotNil(t, psArgs.Value())
	assert.IsType(t, testStrArr, psArgs.Value())
	assert.Len(t, psArgs.Value(), 1)
	assert.Equal(t, "effie", psArgs.Value().([]string)[0])

	psT0 := ps.First()
	assert.NotNil(t, psT0)
	assert.NotNil(t, psT0.Value())
	assert.Implements(t, testOpt, psT0.Value())
	o := psT0.Value().(Option)
	assert.EqualValues(t, 't', o.Opt())
	assert.Equal(t, "37", o.Value())

	psN0, _ := psT0.Next()
	assert.NotNil(t, psN0)
	assert.NotNil(t, psN0.Value())
	assert.Implements(t, testOpt, psN0.Value())
	o = psN0.Value().(Option)
	assert.EqualValues(t, 'n', o.Opt())

	psN1, _ := psN0.Next()
	assert.NotNil(t, psN1)
	assert.NotNil(t, psN1.Value())
	assert.Implements(t, testOpt, psN1.Value())
	o = psN1.Value().(Option)
	assert.EqualValues(t, 'n', o.Opt())

	psX0, _ := psN1.Next()
	assert.NotNil(t, psX0)
	assert.NotNil(t, psX0.Value())
	assert.Implements(t, testOpt, psX0.Value())
	o = psX0.Value().(Option)
	assert.EqualValues(t, 'x', o.Opt())
	assert.Equal(t, "play", o.Value())

	psT1, _ := psX0.Next()
	assert.NotNil(t, psT1)
	assert.NotNil(t, psT1.Value())
	assert.Implements(t, testOpt, psT1.Value())
	o = psT1.Value().(Option)
	assert.EqualValues(t, 't', o.Opt())
	assert.Equal(t, "47", o.Value())
}

func newTestParser() Parser {
	p := NewParser()
	p.Opt('n', "name", NoArgument, "", "")
	p.Opt('t', "time", RequiredArgument, "", "")
	p.Opt('x', "xist", OptionalArgument, "", "")
	p.Opt(0, "pulp", NoArgument, "", "")
	return p
}

func testParseAll(t *testing.T, argv ...string) ParserState {
	t.Logf("argv=%v\n", argv)
	ps, _ := newTestParser().ParseAll(argv)
	return ps
}

func testParse(t *testing.T, argv ...string) *parseTestResult {

	t.Logf("argv=%v\n", argv)

	r := &parseTestResult{}
	c, _ := newTestParser().Parse(argv)
	for i := range c {
		t.Logf("i=%[1]T %[1]v\n", i)

		switch ti := i.Value().(type) {
		case Option:
			switch ti.Opt() {
			case 0:
				if ti.LongName() == "pulp" {
					r.pfnd = true
				}
			case 't':
				r.tfnd = true
				r.nsecs = ti.Value()
			case 'n':
				r.nfnd = true
				r.nct++
			case 'x':
				r.xfnd = true
				r.xist = ti.Value()
			}
		case error:
			r.err = ti
			return r
		case []string:
			r.name = ti[0]
		default:
			t.Fatalf("unknown chan val: %v", ti)
		}
	}
	return r
}

func assertParseOk(t *testing.T, r *parseTestResult) {
	assert.True(t, r.tfnd)
	assert.True(t, r.nfnd)
	assert.Equal(t, "37", r.nsecs)
	assert.Equal(t, "effie", r.name)
}

func assertParseNoName(t *testing.T, r *parseTestResult) {
	assert.True(t, r.tfnd)
	assert.True(t, r.nfnd)
	assert.Equal(t, "37", r.nsecs)
	assert.Equal(t, "", r.name)
}

type parseTestResult struct {
	argv  []string
	argc  int
	nfnd  bool
	nsecs string
	tfnd  bool
	name  string
	xfnd  bool
	xist  string
	pfnd  bool
	err   error
	nct   int
}
