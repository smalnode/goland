package lang_test

import (
	"testing"
)

func TestFormt(t *testing.T) {
	var p *[10]int
	t.Logf("%v\n", p)
	t.Logf("%d\n", len(p))
	t.Logf("p ==nil ? %t", p == nil)

	for i := range p {
		t.Log(i)
	}
}
