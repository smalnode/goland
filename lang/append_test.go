package lang

import "testing"

func TestAppend(t *testing.T) {
	a := []int{0, 1, 2, 3, 4}
	b := []int{5, 6}
	ap := &a
	a = append(a, b...)
	ap2 := &a
	t.Logf("adr before append %p ", ap)
	t.Logf("adr after append %p ", ap2)
	if ap2 == ap {
		t.Log("append doesn't change ptr ")
	} else {
		t.Log("append change ptr")
	}

	for i := 0; i < 100000; i++ {
		a = append(a, i)
		ap2 = &a
		if ap2 != ap {
			t.Log("append change ptr when len(a) = ", len(a))
		}
	}
	if ap2 == ap {
		t.Log("append doesn't change ptr whatever len() under ", len(a))
	}

}
