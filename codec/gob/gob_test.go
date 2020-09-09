package gob

import (
	"testing"

	"github.com/zgwit/storm/codec/internal"
)

func TestGob(t *testing.T) {
	internal.RoundtripTester(t, Codec)
}
