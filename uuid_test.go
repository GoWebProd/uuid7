package uuid7

import (
	"testing"
	"unsafe"
)

func TestCapacity(t *testing.T) {
	uuid := "31648d3c-0933-4845-bdf8-f61b3ae4418b"
	s := *(*[36]byte)(unsafe.Pointer(unsafe.StringData(uuid)))

	if len(s) != 36 {
		t.Fatal("bad uuid bytes len")
	}

	if cap(s) != 36 {
		t.Fatal("bad uuid bytes cap")
	}
}
