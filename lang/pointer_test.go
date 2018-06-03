package lang

import (
	"testing"
)

func TestPointer(t *testing.T) {
	var p *s
	p = &s{"1", 1}
	// *p = s{"1", 1} is not allowed
	p1 := *p
	p = &s{"1", 1}
	p2 := *p
	if p1 == p2 {
		t.Log("struct is equal if their fields are all equal ")
	} else {
		t.Log("struct isn't equal even if their fields are all equal ")
	}
}
