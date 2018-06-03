package main

import (
	"testing"
	"unsafe"
)

func TestPointer1(t *testing.T) {
	a := 2
	c := (*string)(unsafe.Pointer(&a))
	*c = "44"
	t.Log(*c)
	// print(*c) panic
	// t.Log(*c) "44"
	// t.Log(*c) ""
}

func TestPointer2(t *testing.T) {
	a := "654"
	c := (*string)(unsafe.Pointer(&a))
	*c = "44"
	t.Log(*c)
}

func TestPointer3(t *testing.T) {
	a := 3
	c := *(*string)(unsafe.Pointer(&a))
	c = "445"
	t.Log(c)
}
