package gotopt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOptLongOk(t *testing.T) {
	assertLongOk(t, testGetOptLongGlobal(t, "tgok01", "--time=37", "-n", "effie"))
	assertLongOk(t, testGetOptLongGlobal(t, "tgok02", "-n", "--time=37", "effie"))
	assertLongOk(t, testGetOptLongGlobal(t, "tgok03", "--time=37", "effie", "-n"))
	assertLongOk(t, testGetOptLongGlobal(t, "tgok04", "effie", "-n", "--time=37"))

	assertLongOk(t, testGetOptLongGlobal(t, "tgok05", "--time", "37", "-n", "effie"))
	assertLongOk(t, testGetOptLongGlobal(t, "tgok06", "-n", "--time", "37", "effie"))
	assertLongOk(t, testGetOptLongGlobal(t, "tgok07", "--time", "37", "effie", "-n"))
	assertLongOk(t, testGetOptLongGlobal(t, "tgok08", "effie", "-n", "--time", "37"))

	assertLongOk(t, testGetOptLongGlobal(t, "tgok09", "--ti", "37", "-n", "effie"))
	assertLongOk(t, testGetOptLongGlobal(t, "tgok10", "-n", "--tim", "37", "effie"))
	assertLongOk(t, testGetOptLongGlobal(t, "tgok11", "--t", "37", "effie", "-n"))

	assertLongOk(t, testGetOptLongGlobal(t, "tgok12", "--time", "37", "-n", "effie"))
	assertLongOk(t, testGetOptLongGlobal(t, "tgok13", "--na", "--time", "37", "effie"))
	assertLongOk(t, testGetOptLongGlobal(t, "tgok14", "--time", "37", "effie", "--nam"))
	assertLongOk(t, testGetOptLongGlobal(t, "tgok15", "effie", "--name", "--time", "37"))

	assertLongOk(t, testGetOptLongInstance(t, "tiok01", "--time=37", "-n", "effie"))
	assertLongOk(t, testGetOptLongInstance(t, "tiok02", "-n", "--time=37", "effie"))
	assertLongOk(t, testGetOptLongInstance(t, "tiok03", "--time=37", "effie", "-n"))
	assertLongOk(t, testGetOptLongInstance(t, "tiok04", "effie", "-n", "--time=37"))

	assertLongOk(t, testGetOptLongInstance(t, "tiok05", "--time", "37", "-n", "effie"))
	assertLongOk(t, testGetOptLongInstance(t, "tiok06", "-n", "--time", "37", "effie"))
	assertLongOk(t, testGetOptLongInstance(t, "tiok07", "--time", "37", "effie", "-n"))
	assertLongOk(t, testGetOptLongInstance(t, "tiok08", "effie", "-n", "--time", "37"))

	assertLongOk(t, testGetOptLongInstance(t, "tiok09", "--ti", "37", "-n", "effie"))
	assertLongOk(t, testGetOptLongInstance(t, "tiok10", "-n", "--tim", "37", "effie"))
	assertLongOk(t, testGetOptLongInstance(t, "tiok11", "--t", "37", "effie", "-n"))

	assertLongOk(t, testGetOptLongInstance(t, "tiok12", "--time", "37", "-n", "effie"))
	assertLongOk(t, testGetOptLongInstance(t, "tiok13", "--na", "--time", "37", "effie"))
	assertLongOk(t, testGetOptLongInstance(t, "tiok14", "--time", "37", "effie", "--nam"))
	assertLongOk(t, testGetOptLongInstance(t, "tiok15", "effie", "--name", "--time", "37"))
}

func TestGetOptLongW(t *testing.T) {
	assertLongOk(t, testGetOptLongGlobal(t, "tgokw01", "-W", "time=37", "-n", "effie"))
	assertLongOk(t, testGetOptLongGlobal(t, "tgokw02", "-n", "-W", "time=37", "effie"))
	assertLongOk(t, testGetOptLongGlobal(t, "tgokw03", "-W", "time=37", "effie", "-n"))
	assertLongOk(t, testGetOptLongGlobal(t, "tgokw04", "effie", "-n", "-W", "time=37"))

	assertLongOk(t, testGetOptLongGlobal(t, "tgokw05", "-W", "time", "37", "-n", "effie"))
	assertLongOk(t, testGetOptLongGlobal(t, "tgokw06", "-n", "-W", "time", "37", "effie"))
	assertLongOk(t, testGetOptLongGlobal(t, "tgokw07", "-W", "time", "37", "effie", "-n"))
	assertLongOk(t, testGetOptLongGlobal(t, "tgokw08", "effie", "-n", "-W", "time", "37"))

	assertLongOk(t, testGetOptLongGlobal(t, "tgokw09", "-W", "ti", "37", "-n", "effie"))
	assertLongOk(t, testGetOptLongGlobal(t, "tgokw10", "-n", "-W", "tim", "37", "effie"))
	assertLongOk(t, testGetOptLongGlobal(t, "tgokw11", "-W", "t", "37", "effie", "-n"))

	assertLongOk(t, testGetOptLongGlobal(t, "tgokw12", "-W", "time", "37", "-W", "n", "effie"))
	assertLongOk(t, testGetOptLongGlobal(t, "tgokw13", "-W", "na", "-W", "time", "37", "effie"))
	assertLongOk(t, testGetOptLongGlobal(t, "tgokw14", "-W", "time", "37", "effie", "-W", "nam"))
	assertLongOk(t, testGetOptLongGlobal(t, "tgokw15", "effie", "-W", "name", "-W", "time", "37"))

	assertLongOk(t, testGetOptLongInstance(t, "tiokw01", "-W", "time=37", "-n", "effie"))
	assertLongOk(t, testGetOptLongInstance(t, "tiokw02", "-n", "-W", "time=37", "effie"))
	assertLongOk(t, testGetOptLongInstance(t, "tiokw03", "-W", "time=37", "effie", "-n"))
	assertLongOk(t, testGetOptLongInstance(t, "tiokw04", "effie", "-n", "-W", "time=37"))

	assertLongOk(t, testGetOptLongInstance(t, "tiokw05", "-W", "time", "37", "-n", "effie"))
	assertLongOk(t, testGetOptLongInstance(t, "tiokw06", "-n", "-W", "time", "37", "effie"))
	assertLongOk(t, testGetOptLongInstance(t, "tiokw07", "-W", "time", "37", "effie", "-n"))
	assertLongOk(t, testGetOptLongInstance(t, "tiokw08", "effie", "-n", "-W", "time", "37"))

	assertLongOk(t, testGetOptLongInstance(t, "tiokw09", "-W", "ti", "37", "-n", "effie"))
	assertLongOk(t, testGetOptLongInstance(t, "tiokw10", "-n", "-W", "tim", "37", "effie"))
	assertLongOk(t, testGetOptLongInstance(t, "tiokw11", "-W", "t", "37", "effie", "-n"))

	assertLongOk(t, testGetOptLongInstance(t, "tiokw12", "-W", "time", "37", "-W", "n", "effie"))
	assertLongOk(t, testGetOptLongInstance(t, "tiokw13", "-W", "na", "-W", "time", "37", "effie"))
	assertLongOk(t, testGetOptLongInstance(t, "tiokw14", "-W", "time", "37", "effie", "-W", "nam"))
	assertLongOk(t, testGetOptLongInstance(t, "tiokw15", "effie", "-W", "name", "-W", "time", "37"))
}

func TestGetOptLongMissingName(t *testing.T) {
	assertLongNoName(t, testGetOptLongGlobal(t, "tgnn01", "--time=37", "-n"))
	assertLongNoName(t, testGetOptLongGlobal(t, "tgnn02", "-n", "--time=37"))
	assertLongNoName(t, testGetOptLongGlobal(t, "tgnn03", "--time=37", "-n"))
	assertLongNoName(t, testGetOptLongGlobal(t, "tgnn04", "-n", "--time=37"))

	assertLongNoName(t, testGetOptLongGlobal(t, "tgnn05", "--time", "37", "-n"))
	assertLongNoName(t, testGetOptLongGlobal(t, "tgnn06", "-n", "--time", "37"))
	assertLongNoName(t, testGetOptLongGlobal(t, "tgnn07", "--time", "37", "-n"))
	assertLongNoName(t, testGetOptLongGlobal(t, "tgnn08", "-n", "--time", "37"))

	assertLongNoName(t, testGetOptLongGlobal(t, "tgok09", "--ti", "37", "-n"))
	assertLongNoName(t, testGetOptLongGlobal(t, "tgok10", "-n", "--tim", "37"))
	assertLongNoName(t, testGetOptLongGlobal(t, "tgok11", "--t", "37", "-n"))

	assertLongNoName(t, testGetOptLongInstance(t, "tinn01", "--time=37", "-n"))
	assertLongNoName(t, testGetOptLongInstance(t, "tinn02", "-n", "--time=37"))
	assertLongNoName(t, testGetOptLongInstance(t, "tinn03", "--time=37", "-n"))
	assertLongNoName(t, testGetOptLongInstance(t, "tinn04", "-n", "--time=37"))

	assertLongNoName(t, testGetOptLongInstance(t, "tinn05", "--time", "37", "-n"))
	assertLongNoName(t, testGetOptLongInstance(t, "tinn06", "-n", "--time", "37"))
	assertLongNoName(t, testGetOptLongInstance(t, "tinn07", "--time", "37", "-n"))
	assertLongNoName(t, testGetOptLongInstance(t, "tinn08", "-n", "--time", "37"))

	assertLongNoName(t, testGetOptLongInstance(t, "tgok09", "--ti", "37", "-n"))
	assertLongNoName(t, testGetOptLongInstance(t, "tgok10", "-n", "--tim", "37"))
	assertLongNoName(t, testGetOptLongInstance(t, "tgok11", "--t", "37", "-n"))
}

func TestGetOptLongMissingNameW(t *testing.T) {
	assertLongNoName(t, testGetOptLongGlobal(t, "tgnnw01", "-W", "time=37", "-n"))
	assertLongNoName(t, testGetOptLongGlobal(t, "tgnnw02", "-n", "-W", "time=37"))
	assertLongNoName(t, testGetOptLongGlobal(t, "tgnnw03", "-W", "time=37", "-W", "n"))
	assertLongNoName(t, testGetOptLongGlobal(t, "tgnnw04", "-W", "na", "-W", "time=37"))

	assertLongNoName(t, testGetOptLongGlobal(t, "tgnnw05", "-W", "time", "37", "-n"))
	assertLongNoName(t, testGetOptLongGlobal(t, "tgnnw06", "-n", "-W", "time", "37"))
	assertLongNoName(t, testGetOptLongGlobal(t, "tgnnw07", "-W", "time", "37", "-W", "n"))
	assertLongNoName(t, testGetOptLongGlobal(t, "tgnnw08", "-W", "na", "-W", "time", "37"))

	assertLongNoName(t, testGetOptLongGlobal(t, "tgokw09", "-W", "ti", "37", "-n"))
	assertLongNoName(t, testGetOptLongGlobal(t, "tgokw10", "-W", "nam", "-W", "tim", "37"))
	assertLongNoName(t, testGetOptLongGlobal(t, "tgokw11", "-W", "t", "37", "-W", "name"))

	assertLongNoName(t, testGetOptLongInstance(t, "tinnw01", "-W", "time=37", "-n"))
	assertLongNoName(t, testGetOptLongInstance(t, "tinnw02", "-n", "-W", "time=37"))
	assertLongNoName(t, testGetOptLongInstance(t, "tinnw03", "-W", "time=37", "-W", "n"))
	assertLongNoName(t, testGetOptLongInstance(t, "tinnw04", "-W", "na", "-W", "time=37"))

	assertLongNoName(t, testGetOptLongInstance(t, "tinnw05", "-W", "time", "37", "-n"))
	assertLongNoName(t, testGetOptLongInstance(t, "tinnw06", "-n", "-W", "time", "37"))
	assertLongNoName(t, testGetOptLongInstance(t, "tinnw07", "-W", "time", "37", "-W", "n"))
	assertLongNoName(t, testGetOptLongInstance(t, "tinnw08", "-W", "na", "-W", "time", "37"))

	assertLongNoName(t, testGetOptLongInstance(t, "tiokw09", "-W", "ti", "37", "-n"))
	assertLongNoName(t, testGetOptLongInstance(t, "tiokw10", "-W", "nam", "-W", "tim", "37"))
	assertLongNoName(t, testGetOptLongInstance(t, "tiokw11", "-W", "t", "37", "-W", "name"))
}

func TestGetOptLongNoTime(t *testing.T) {
	a1 := func(t *testing.T, r *getOptLongTestResult) {
		assert.False(t, r.tfnd)
		assert.IsType(t, &ErrRequiredArg{}, r.err)
		err := r.err.(*ErrRequiredArg)
		assert.EqualValues(t, 't', err.OptOpt)
		assert.NotEqual(t, "37", r.nsecs)
		assert.False(t, r.nfnd)
		assert.Equal(t, "", r.name)
	}
	a1(t, testGetOptLongGlobal(t, "tgnt01", "--t"))
	a1(t, testGetOptLongGlobal(t, "tgnt02", "--ti"))
	a1(t, testGetOptLongGlobal(t, "tgnt03", "--tim"))
	a1(t, testGetOptLongGlobal(t, "tgnt04", "--time"))

	a1(t, testGetOptLongInstance(t, "tint01", "--t"))
	a1(t, testGetOptLongInstance(t, "tint02", "--ti"))
	a1(t, testGetOptLongInstance(t, "tint03", "--tim"))
	a1(t, testGetOptLongInstance(t, "tint04", "--time"))

	a2 := func(t *testing.T, r *getOptLongTestResult) {
		assert.False(t, r.tfnd)
		assert.IsType(t, &ErrRequiredArg{}, r.err)
		err := r.err.(*ErrRequiredArg)
		assert.EqualValues(t, 't', err.OptOpt)
		assert.NotEqual(t, "37", r.nsecs)
		assert.True(t, r.nfnd)
		assert.Equal(t, "", r.name)
	}
	a2(t, testGetOptLongGlobal(t, "tgnt05", "-n", "--t"))
	a2(t, testGetOptLongGlobal(t, "tgnt06", "-n", "--ti"))
	a2(t, testGetOptLongGlobal(t, "tgnt07", "-n", "--tim"))
	a2(t, testGetOptLongGlobal(t, "tgnt08", "-n", "--time"))

	a2(t, testGetOptLongInstance(t, "tint05", "-n", "--t"))
	a2(t, testGetOptLongInstance(t, "tint06", "-n", "--ti"))
	a2(t, testGetOptLongInstance(t, "tint07", "-n", "--tim"))
	a2(t, testGetOptLongInstance(t, "tint08", "-n", "--time"))
}

func TestGetOptLongUnknownOpt(t *testing.T) {
	a1 := func(t *testing.T, r *getOptLongTestResult, u string) {
		assert.True(t, r.tfnd)
		assert.Equal(t, "37", r.nsecs)
		assert.IsType(t, &ErrUnknownOpt{}, r.err)
		err := r.err.(*ErrUnknownOpt)
		if err.OptArg == "" {
			assert.EqualValues(t, u, fmt.Sprintf("%c", err.OptOpt))
		} else {
			assert.EqualValues(t, u, err.OptArg)
		}
		assert.False(t, r.nfnd)
		assert.Equal(t, "", r.name)
	}
	a1(t, testGetOptLongGlobal(t, "tgunkn01a", "--t=37", "-f", "-n", "effie"), "f")
	a1(t, testGetOptLongGlobal(t, "tgunkn02a", "--ti=37", "-fu", "-n", "effie"), "f")
	a1(t, testGetOptLongGlobal(t, "tgunkn03a", "--tim=37", "-fub", "-n", "effie"), "f")
	a1(t, testGetOptLongGlobal(t, "tgunkn04a", "--time=37", "-fubar", "-n", "effie"), "f")

	a1(t, testGetOptLongGlobal(t, "tgunkn01b", "--t=37", "-W", "f", "-n", "effie"), "f")
	a1(t, testGetOptLongGlobal(t, "tgunkn02b", "--ti=37", "-W", "fu", "-n", "effie"), "fu")
	a1(t, testGetOptLongGlobal(t, "tgunkn03b", "--tim=37", "-W", "fub", "-n", "effie"), "fub")
	a1(t, testGetOptLongGlobal(t, "tgunkn04b", "--time=37", "-W", "fubar", "-n", "effie"), "fubar")

	a1(t, testGetOptLongInstance(t, "tiunkn01a", "--t=37", "-f", "-n", "effie"), "f")
	a1(t, testGetOptLongInstance(t, "tiunkn02a", "--ti=37", "-fu", "-n", "effie"), "f")
	a1(t, testGetOptLongInstance(t, "tiunkn03a", "--tim=37", "-fub", "-n", "effie"), "f")
	a1(t, testGetOptLongInstance(t, "tiunkn04a", "--time=37", "-fubar", "-n", "effie"), "f")

	a1(t, testGetOptLongInstance(t, "tiunkn01b", "--t=37", "-W", "f", "-n", "effie"), "f")
	a1(t, testGetOptLongInstance(t, "tiunkn02b", "--ti=37", "-W", "fu", "-n", "effie"), "fu")
	a1(t, testGetOptLongInstance(t, "tiunkn03b", "--tim=37", "-W", "fub", "-n", "effie"), "fub")
	a1(t, testGetOptLongInstance(t, "tiunkn04b", "--time=37", "-W", "fubar", "-n", "effie"), "fubar")

	a2 := func(t *testing.T, r *getOptLongTestResult, u string) {
		assert.True(t, r.tfnd)
		assert.Equal(t, "37", r.nsecs)
		assert.True(t, r.nfnd)
		assert.Equal(t, "", r.name)
		assert.IsType(t, &ErrUnknownOpt{}, r.err)
		err := r.err.(*ErrUnknownOpt)
		if err.OptArg == "" {
			assert.EqualValues(t, u, fmt.Sprintf("%c", err.OptOpt))
		} else {
			assert.EqualValues(t, u, err.OptArg)
		}
	}
	a2(t, testGetOptLongGlobal(t, "tgunkn05", "--t=37", "-n", "effie", "-f"), "f")
	a2(t, testGetOptLongGlobal(t, "tgunkn06", "--ti=37", "-n", "effie", "-W", "f"), "f")
	a2(t, testGetOptLongGlobal(t, "tgunkn07", "--tim=37", "-n", "effie", "-fu"), "f")
	a2(t, testGetOptLongGlobal(t, "tgunkn08", "--time=37", "-n", "effie", "-W", "fubar"), "fubar")

	a2(t, testGetOptLongInstance(t, "tiunkn05", "--t=37", "-n", "effie", "-f"), "f")
	a2(t, testGetOptLongInstance(t, "tiunkn06", "--ti=37", "-n", "effie", "-W", "f"), "f")
	a2(t, testGetOptLongInstance(t, "tiunkn07", "--tim=37", "-n", "effie", "-fu"), "f")
	a2(t, testGetOptLongInstance(t, "tiunkn08", "--time=37", "-n", "effie", "-W", "fubar"), "fubar")

	a3 := func(t *testing.T, r *getOptLongTestResult) {
		assert.True(t, r.tfnd)
		assert.Equal(t, "37", r.nsecs)
		assert.IsType(t, &ErrUnknownOpt{}, r.err)
		err := r.err.(*ErrUnknownOpt)
		assert.EqualValues(t, 0, err.OptOpt)
		assert.EqualValues(t, "hello", r.optArg)
		assert.False(t, r.nfnd)
		assert.Equal(t, "", r.name)
	}
	a3(t, testGetOptLongGlobal(t, "tgunkn09", "--t=37", "--hello", "-f", "-n", "effie"))
	a3(t, testGetOptLongGlobal(t, "tgunkn10", "--ti=37", "--hello", "-f", "-n", "effie"))
	a3(t, testGetOptLongGlobal(t, "tgunkn11", "--tim=37", "--hello", "-f", "-n", "effie"))
	a3(t, testGetOptLongGlobal(t, "tgunkn12", "--time=37", "--hello", "-f", "-n", "effie"))

	a3(t, testGetOptLongInstance(t, "tiunkn09", "--t=37", "--hello", "-f", "-n", "effie"))
	a3(t, testGetOptLongInstance(t, "tiunkn10", "--ti=37", "--hello", "-f", "-n", "effie"))
	a3(t, testGetOptLongInstance(t, "tiunkn11", "--tim=37", "--hello", "-f", "-n", "effie"))
	a3(t, testGetOptLongInstance(t, "tiunkn12", "--time=37", "--hello", "-f", "-n", "effie"))
}

func assertLongOk(t *testing.T, r *getOptLongTestResult) {
	assert.True(t, r.tfnd)
	assert.True(t, r.nfnd)
	assert.Equal(t, "37", r.nsecs)
	assert.Equal(t, "effie", r.name)
}

func assertLongNoName(t *testing.T, r *getOptLongTestResult) {
	assert.True(t, r.tfnd)
	assert.True(t, r.nfnd)
	assert.Equal(t, "37", r.nsecs)
	assert.Equal(t, "", r.name)
}

func testGetOptLongGlobal(t *testing.T, argv ...string) *getOptLongTestResult {

	defer func() {
		OptInd = 1
	}()

	r := &getOptLongTestResult{
		argv: argv,
		argc: len(argv),
	}

	t.Logf("argv=%v\n", argv)

	zf := 0
	longInd := 0

	longOpts := []*LongOption{
		&LongOption{
			Name: "name",
			Type: NoArgument,
			Flag: &zf,
			Val:  'n',
		},
		&LongOption{
			Name: "time",
			Type: RequiredArgument,
			Flag: &zf,
			Val:  't',
		},
	}

	for {
		opt := GetOptLong(argv, ":W;nt:", longOpts, &longInd)

		if opt == -1 {
			break
		}

		switch opt {
		case 0:
			if longOpts[longInd].Name == "time" {
				r.tfnd = true
				r.nsecs = OptArg
			} else if longOpts[longInd].Name == "name" {
				r.nfnd = true
			}
		case 'n':
			r.nfnd = true
		case 't':
			r.tfnd = true
			r.nsecs = OptArg
		case ':':
			r.err = &ErrRequiredArg{OptOpt}
			return r
		case 'W':
			r.err = &ErrUnknownOpt{OptOpt, OptArg}
			return r
		default: // ?
			r.optArg = OptArg
			r.err = &ErrUnknownOpt{OptOpt, ""}
			return r
		}
	}

	if r.nfnd && OptInd < r.argc {
		r.name = argv[OptInd]
	}

	r.optInd = OptInd

	return r
}

func testGetOptLongInstance(t *testing.T, argv ...string) *getOptLongTestResult {

	p := NewGetOptParser()

	r := &getOptLongTestResult{
		argv: argv,
		argc: len(argv),
	}

	t.Logf("argv=%v\n", argv)

	zf := 0
	longInd := 0

	longOpts := []*LongOption{
		&LongOption{
			Name: "name",
			Type: NoArgument,
			Flag: &zf,
			Val:  'n',
		},
		&LongOption{
			Name: "time",
			Type: RequiredArgument,
			Flag: &zf,
			Val:  't',
		},
	}

	for {
		opt := p.GetOptLong(argv, ":W;nt:", longOpts, &longInd)

		if opt == -1 {
			break
		}

		switch opt {
		case 0:
			if longOpts[longInd].Name == "time" {
				r.tfnd = true
				r.nsecs = p.OptArg
			} else if longOpts[longInd].Name == "name" {
				r.nfnd = true
			}
		case 'n':
			r.nfnd = true
		case 't':
			r.tfnd = true
			r.nsecs = p.OptArg
		case ':':
			r.err = &ErrRequiredArg{p.OptOpt}
			return r
		case 'W':
			r.err = &ErrUnknownOpt{p.OptOpt, p.OptArg}
			return r
		default: // ?
			r.optArg = p.OptArg
			r.err = &ErrUnknownOpt{p.OptOpt, ""}
			return r
		}
	}

	if r.nfnd && p.OptInd < r.argc {
		r.name = argv[p.OptInd]
	}

	return r
}

type getOptLongTestResult struct {
	argv    []string
	argc    int
	nfnd    bool
	nsecs   string
	tfnd    bool
	name    string
	err     error
	optInd  int
	longOpt *LongOption
	optArg  string
}
