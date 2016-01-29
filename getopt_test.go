package gotopt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOptOk(t *testing.T) {
	assertOk(t, testGetOptGlobal(t, "tgok01", "-t", "37", "-n", "effie"))
	assertOk(t, testGetOptGlobal(t, "tgok02", "-nt", "37", "effie"))
	assertOk(t, testGetOptGlobal(t, "tgok03", "-t", "37", "effie", "-n"))
	assertOk(t, testGetOptGlobal(t, "tgok04", "effie", "-nt", "37"))
	assertOk(t, testGetOptGlobal(t, "tgok05", "effie", "-n", "-t", "37"))

	assertOk(t, testGetOptGlobal(t, "tgok06", "-t37", "-n", "effie"))
	assertOk(t, testGetOptGlobal(t, "tgok07", "-nt37", "effie"))
	assertOk(t, testGetOptGlobal(t, "tgok08", "-t37", "effie", "-n"))
	assertOk(t, testGetOptGlobal(t, "tgok09", "effie", "-nt37"))

	assertOk(t, testGetOptInstance(t, "tiok01", "-t", "37", "-n", "effie"))
	assertOk(t, testGetOptInstance(t, "tiok02", "-nt", "37", "effie"))
	assertOk(t, testGetOptInstance(t, "tiok03", "-t", "37", "effie", "-n"))
	assertOk(t, testGetOptInstance(t, "tiok04", "effie", "-nt", "37"))
	assertOk(t, testGetOptInstance(t, "tiok05", "effie", "-n", "-t", "37"))

	assertOk(t, testGetOptGlobal(t, "tiok06", "-t37", "-n", "effie"))
	assertOk(t, testGetOptGlobal(t, "tiok07", "-nt37", "effie"))
	assertOk(t, testGetOptGlobal(t, "tiok08", "-t37", "effie", "-n"))
	assertOk(t, testGetOptGlobal(t, "tiok09", "effie", "-nt37"))
}

func TestGetOptMissingName(t *testing.T) {
	assertNoName(t, testGetOptGlobal(t, "tgnn1", "-t", "37", "-n"))
	assertNoName(t, testGetOptGlobal(t, "tgnn2", "-nt", "37"))
	assertNoName(t, testGetOptGlobal(t, "tgnn3", "-t", "37", "-n"))
	assertNoName(t, testGetOptGlobal(t, "tgnn4", "-nt", "37"))
	assertNoName(t, testGetOptGlobal(t, "tgnn5", "-n", "-t", "37"))

	assertNoName(t, testGetOptInstance(t, "tinn1", "-t", "37", "-n"))
	assertNoName(t, testGetOptInstance(t, "tinn2", "-nt", "37"))
	assertNoName(t, testGetOptInstance(t, "tinn3", "-t", "37", "-n"))
	assertNoName(t, testGetOptInstance(t, "tinn4", "-nt", "37"))
	assertNoName(t, testGetOptInstance(t, "tinn5", "-n", "-t", "37"))
}

func TestGetOptNoTime(t *testing.T) {
	{
		r := testGetOptGlobal(t, "tgnt1", "-t")
		assert.False(t, r.tfnd)
		assert.IsType(t, &ErrRequiredArg{}, r.err)
		err := r.err.(*ErrRequiredArg)
		assert.EqualValues(t, 't', err.Opt)
		assert.NotEqual(t, "37", r.nsecs)
		assert.False(t, r.nfnd)
		assert.Equal(t, "", r.name)
	}

	{
		r := testGetOptGlobal(t, "tgnt2", "-n", "-t")
		assert.False(t, r.tfnd)
		assert.IsType(t, &ErrRequiredArg{}, r.err)
		err := r.err.(*ErrRequiredArg)
		assert.EqualValues(t, 't', err.Opt)
		assert.NotEqual(t, "37", r.nsecs)
		assert.True(t, r.nfnd)
		assert.Equal(t, "", r.name)
	}

	{
		r := testGetOptInstance(t, "tint1", "-t")
		assert.False(t, r.tfnd)
		assert.IsType(t, &ErrRequiredArg{}, r.err)
		err := r.err.(*ErrRequiredArg)
		assert.EqualValues(t, 't', err.Opt)
		assert.NotEqual(t, "37", r.nsecs)
		assert.False(t, r.nfnd)
		assert.Equal(t, "", r.name)
	}

	{
		r := testGetOptInstance(t, "tint2", "-n", "-t")
		assert.False(t, r.tfnd)
		assert.IsType(t, &ErrRequiredArg{}, r.err)
		err := r.err.(*ErrRequiredArg)
		assert.EqualValues(t, 't', err.Opt)
		assert.NotEqual(t, "37", r.nsecs)
		assert.True(t, r.nfnd)
		assert.Equal(t, "", r.name)
	}
}

func TestGetOptUnknownOpt(t *testing.T) {
	{
		r := testGetOptGlobal(t, "tgunkn1", "-t", "37", "-f", "-n", "effie")
		assert.True(t, r.tfnd)
		assert.Equal(t, "37", r.nsecs)
		assert.IsType(t, &ErrUnknownOpt{}, r.err)
		err := r.err.(*ErrUnknownOpt)
		assert.EqualValues(t, 'f', err.Opt)
		assert.False(t, r.nfnd)
		assert.Equal(t, "", r.name)
	}

	{
		r := testGetOptGlobal(t, "tgunkn2", "-t", "37", "-n", "effie", "-f")
		assert.True(t, r.tfnd)
		assert.Equal(t, "37", r.nsecs)
		assert.True(t, r.nfnd)
		assert.Equal(t, "", r.name)
		assert.IsType(t, &ErrUnknownOpt{}, r.err)
		err := r.err.(*ErrUnknownOpt)
		assert.EqualValues(t, 'f', err.Opt)
	}

	{
		r := testGetOptInstance(t, "tiunkn1", "-t", "37", "-f", "-n", "effie")
		assert.True(t, r.tfnd)
		assert.Equal(t, "37", r.nsecs)
		assert.IsType(t, &ErrUnknownOpt{}, r.err)
		err := r.err.(*ErrUnknownOpt)
		assert.EqualValues(t, 'f', err.Opt)
		assert.False(t, r.nfnd)
		assert.Equal(t, "", r.name)
	}

	{
		r := testGetOptInstance(t, "tiunkn2", "-t", "37", "-n", "effie", "-f")
		assert.True(t, r.tfnd)
		assert.Equal(t, "37", r.nsecs)
		assert.True(t, r.nfnd)
		assert.Equal(t, "", r.name)
		assert.IsType(t, &ErrUnknownOpt{}, r.err)
		err := r.err.(*ErrUnknownOpt)
		assert.EqualValues(t, 'f', err.Opt)
	}
}

func assertOk(t *testing.T, r *getOptTestResult) {
	assert.True(t, r.tfnd)
	assert.True(t, r.nfnd)
	assert.Equal(t, "37", r.nsecs)
	assert.Equal(t, "effie", r.name)
}

func assertNoName(t *testing.T, r *getOptTestResult) {
	assert.True(t, r.tfnd)
	assert.True(t, r.nfnd)
	assert.Equal(t, "37", r.nsecs)
	assert.Equal(t, "", r.name)
}

func testGetOptInstance(t *testing.T, argv ...string) *getOptTestResult {

	p := NewGetOptParser()

	r := &getOptTestResult{
		argv: argv,
		argc: len(argv),
	}

	t.Logf("argv=%v\n", argv)

	for {
		opt := p.GetOpt(argv, ":nt:")

		if opt == -1 {
			break
		}

		switch opt {
		case 'n':
			r.nfnd = true
		case 't':
			r.tfnd = true
			r.nsecs = p.OptArg
		case ':':
			r.err = &ErrRequiredArg{p.OptOpt}
			return r
		default: // ?
			r.err = &ErrUnknownOpt{p.OptOpt}
			return r
		}
	}

	if p.OptInd < r.argc {
		r.name = argv[p.OptInd]
	}

	return r
}

func testGetOptGlobal(t *testing.T, argv ...string) *getOptTestResult {

	defer func() {
		OptInd = 1
	}()

	r := &getOptTestResult{
		argv: argv,
		argc: len(argv),
	}

	t.Logf("argv=%v\n", argv)

	for {
		opt := GetOpt(argv, ":nt:")

		if opt == -1 {
			break
		}

		switch opt {
		case 'n':
			r.nfnd = true
		case 't':
			r.tfnd = true
			r.nsecs = OptArg
		case ':':
			r.err = &ErrRequiredArg{OptOpt}
			return r
		default: // ?
			r.err = &ErrUnknownOpt{OptOpt}
			return r
		}
	}

	if OptInd < r.argc {
		r.name = argv[OptInd]
	}

	r.optInd = OptInd

	return r
}

type getOptTestResult struct {
	argv   []string
	argc   int
	nfnd   bool
	nsecs  string
	tfnd   bool
	name   string
	err    error
	optInd int
}
