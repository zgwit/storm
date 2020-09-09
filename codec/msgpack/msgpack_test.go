package msgpack

import (
	"testing"

	"github.com/zgwit/storm/codec/internal"
)

func TestMsgpack(t *testing.T) {
	internal.RoundtripTester(t, Codec)
}
