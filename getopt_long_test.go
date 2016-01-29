package gotopt

import (
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

	assertLongOk(t, testGetOptLongInstance(t, "tgok09", "--ti", "37", "-n", "effie"))
	assertLongOk(t, testGetOptLongInstance(t, "tgok10", "-n", "--tim", "37", "effie"))
	assertLongOk(t, testGetOptLongInstance(t, "tgok11", "--t", "37", "effie", "-n"))
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

func TestGetOptLongNoTime(t *testing.T) {
	a1 := func(t *testing.T, r *getOptLongTestResult) {
		assert.False(t, r.tfnd)
		assert.IsType(t, &ErrRequiredArg{}, r.err)
		err := r.err.(*ErrRequiredArg)
		assert.EqualValues(t, 't', err.Opt)
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
		assert.EqualValues(t, 't', err.Opt)
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
	a1 := func(t *testing.T, r *getOptLongTestResult) {
		assert.True(t, r.tfnd)
		assert.Equal(t, "37", r.nsecs)
		assert.IsType(t, &ErrUnknownOpt{}, r.err)
		err := r.err.(*ErrUnknownOpt)
		assert.EqualValues(t, 'f', err.Opt)
		assert.False(t, r.nfnd)
		assert.Equal(t, "", r.name)
	}
	a1(t, testGetOptLongGlobal(t, "tgunkn01", "--t=37", "-f", "-n", "effie"))
	a1(t, testGetOptLongGlobal(t, "tgunkn02", "--ti=37", "-f", "-n", "effie"))
	a1(t, testGetOptLongGlobal(t, "tgunkn03", "--tim=37", "-f", "-n", "effie"))
	a1(t, testGetOptLongGlobal(t, "tgunkn04", "--time=37", "-f", "-n", "effie"))

	a1(t, testGetOptLongInstance(t, "tiunkn01", "--t=37", "-f", "-n", "effie"))
	a1(t, testGetOptLongInstance(t, "tiunkn02", "--ti=37", "-f", "-n", "effie"))
	a1(t, testGetOptLongInstance(t, "tiunkn03", "--tim=37", "-f", "-n", "effie"))
	a1(t, testGetOptLongInstance(t, "tiunkn04", "--time=37", "-f", "-n", "effie"))

	a2 := func(t *testing.T, r *getOptLongTestResult) {
		assert.True(t, r.tfnd)
		assert.Equal(t, "37", r.nsecs)
		assert.True(t, r.nfnd)
		assert.Equal(t, "", r.name)
		assert.IsType(t, &ErrUnknownOpt{}, r.err)
		err := r.err.(*ErrUnknownOpt)
		assert.EqualValues(t, 'f', err.Opt)
	}
	a2(t, testGetOptLongGlobal(t, "tgunkn05", "--t=37", "-n", "effie", "-f"))
	a2(t, testGetOptLongGlobal(t, "tgunkn06", "--ti=37", "-n", "effie", "-f"))
	a2(t, testGetOptLongGlobal(t, "tgunkn07", "--tim=37", "-n", "effie", "-f"))
	a2(t, testGetOptLongGlobal(t, "tgunkn08", "--time=37", "-n", "effie", "-f"))

	a2(t, testGetOptLongInstance(t, "tiunkn05", "--t=37", "-n", "effie", "-f"))
	a2(t, testGetOptLongInstance(t, "tiunkn06", "--ti=37", "-n", "effie", "-f"))
	a2(t, testGetOptLongInstance(t, "tiunkn07", "--tim=37", "-n", "effie", "-f"))
	a2(t, testGetOptLongInstance(t, "tiunkn08", "--time=37", "-n", "effie", "-f"))

	a3 := func(t *testing.T, r *getOptLongTestResult) {
		assert.True(t, r.tfnd)
		assert.Equal(t, "37", r.nsecs)
		assert.IsType(t, &ErrUnknownOpt{}, r.err)
		err := r.err.(*ErrUnknownOpt)
		assert.EqualValues(t, 0, err.Opt)
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
		opt := GetOptLong(argv, ":nt:", longOpts, &longInd)

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
		default: // ?
			r.optArg = OptArg
			r.err = &ErrUnknownOpt{OptOpt}
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
		opt := p.GetOptLong(argv, ":nt:", longOpts, &longInd)

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
		default: // ?
			r.optArg = p.OptArg
			r.err = &ErrUnknownOpt{p.OptOpt}
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
