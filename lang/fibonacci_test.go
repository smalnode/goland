package lang

import "testing"

func fibgen() func() (int, int) {
	p, q := 0, 1
	i := 0
	return func() (int, int) {
		p, q = q, p+q
		i++
		return i, p
	}
}

func TestFibGen(t *testing.T) {
	fib := fibgen()
	for i := 0; i < 50; i++ {
		k, v := fib()
		t.Logf("#%02d %15d", k, v)
	}
}
