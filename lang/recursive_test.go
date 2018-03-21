package lang

import "testing"

func sum(ceil int) int {
	if ceil > 0 {
		return ceil + sum(ceil-1)
	}
	return 0
}

var memo = make(map[int]int)

func fn(n int) int {
	res, ok := memo[n]
	if ok {
		return res
	}
	if n > 1 {
		res = fn(n-1) + fn(n-2)
	} else {
		res = n
	}
	memo[n] = res
	return res
}

func TestSum(t *testing.T) {
	t.Log(sum(100))
}

func TestFn(t *testing.T) {
	t.Log(fn(100))
}

func fnn(n int) int {
	if n < 1 {
		return n
	}
	p, q := 0, 1
	for i := 0; i < n; i++ {
		p, q = p+q, p
	}
	return p
}

func TestFnn(t *testing.T) {
	for i := 0; i < 50; i++ {
		n := fn(i)
		m := fnn(i)
		if n != m {
			t.Errorf("mismatch at %d: n = %d, m = %d ", i, n, m)
		}
		t.Logf("at %d: %d %d", i, n, m)
	}
}
