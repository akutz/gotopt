package gotopt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLongOptSize(t *testing.T) {
	nameEnd, nameLen := parseLongOptSize("hello=world", 0)
	assert.EqualValues(t, 5, nameEnd)
	assert.EqualValues(t, 5, nameLen)

	nameEnd, nameLen = parseLongOptSize("hello", 0)
	assert.EqualValues(t, -1, nameEnd)
	assert.EqualValues(t, 5, nameLen)

	nameEnd, nameLen = parseLongOptSize("hello", 3)
	assert.EqualValues(t, -1, nameEnd)
	assert.EqualValues(t, 2, nameLen)

	nameEnd, nameLen = parseLongOptSize("hello=world", 3)
	assert.EqualValues(t, 2, nameEnd)
	assert.EqualValues(t, 2, nameLen)
}
