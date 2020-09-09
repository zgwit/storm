package json

import (
	"testing"

	"github.com/zgwit/storm/v3/codec/internal"
)

func TestJSON(t *testing.T) {
	internal.RoundtripTester(t, Codec)
}
