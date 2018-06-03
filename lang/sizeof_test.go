package lang

import (
	"testing"
	"unsafe"
)

func TestStructSize(t *testing.T) {
	var v struct{}
	t.Log("empty struct size: ", unsafe.Sizeof(v))

	var w struct{ a uintptr }
	t.Log("uintptr struct size: ", unsafe.Sizeof(w))
}
