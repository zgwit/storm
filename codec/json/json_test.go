package json

import (
	"testing"

	"github.com/zgwit/storm/codec/internal"
)

func TestJSON(t *testing.T) {
	internal.RoundtripTester(t, Codec)
}
