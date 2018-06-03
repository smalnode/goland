package lang

import "testing"
import "unsafe"

func TestRangeV(t *testing.T) {
	ss := []s{
		{"D", 0},
	}
	for k, v := range ss {
		t.Logf("%d, %v, %p", k, v, &v)
		v.v++
	}
	for k, v := range ss {
		t.Logf("%d, %v, %p", k, v, &v)
		if v.v != k {
			t.Log("v is a reference ")
		} else {
			t.Log("v is a copy ")
		}
	}
}

func TestRangeNum(t *testing.T) {
	var p *[10]int
	for i := range p {
		if i == 9 {
			t.Log("range reach", i)
		}
	}
}

func TestPtrSize(t *testing.T) {
	var p *[10]int
	t.Log("var p *[10]int, sizeof(p) = ", unsafe.Sizeof(p))
	var f64 float64
	t.Log("sizeof(float64) = ", unsafe.Sizeof(f64))
}
