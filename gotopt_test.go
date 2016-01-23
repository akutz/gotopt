package gotopt

import (
	"testing"
)

func TestExchange(t *testing.T) {
	g := NewGotOpt("-tzvf", "archive.tgz")
	g.exchange(g.argv)
}
