package msgpack

import (
	"testing"

	"github.com/zgwit/storm/v3/codec/internal"
)

func TestMsgpack(t *testing.T) {
	internal.RoundtripTester(t, Codec)
}
