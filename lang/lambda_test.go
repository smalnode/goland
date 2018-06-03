package lang

import "testing"

func f(n int) func(int) int {
	return func(i int) int {
		n += i
		return n
	}
}

func TestFn(t *testing.T) {
	ff := f(10)
	t.Log(ff(1))
	t.Log(ff(2))
}

func newFib() func() (int, int) {
	p, q, i := 0, 1, 0
	return func() (int, int) {
		p, q = q, p+q
		i++
		return i, p
	}
}

func fibn(n int) int {
	p, q := 0, 1
	for i := 0; i < n; i++ {
		p, q = q, p+q
	}
	return p
}

func TestFib(t *testing.T) {
	fib := newFib()
	for i := 0; i < 10; i++ {
		k, v := fib()
		if v != fibn(k) {
			t.Errorf("mismatch at %d, expected %d, got %d ", k, fibn(k), v)
		}
	}
}
