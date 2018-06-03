package lang

import (
	"testing"
	"fmt"
)

type bb byte

func (b bb) Square() int64 {
	return int64(b) * int64(b)
}

func TestBB(t *testing.T) {
	var b byte
	b = 128
	t.Log(bb(b).Square())
}

type aa struct{}

func (a aa) Foo() {
	fmt.Println("aa.Foo()")
}

type cc = aa
func TestCC(t *testing.T) {
	cc{}.Foo()
}


type dd aa

func TestDD(t *testing.T) {
	aa(dd{}).Foo()
}