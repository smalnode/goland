package lang_test

import (
	"testing"
)

func TestRuneIndex(t *testing.T) {
	// go test -v
	s := "Go语言"
	// 1 + 1 + 3 +  8
	// rune cane split string by language char
	if len(s) != 8 {
		t.Errorf("len(s) expected %d, got %d ", 8, len(s))
	}
	r := []rune(s)
	if len(r) != 4 {
		t.Errorf("len(rune(s)) expected %d, got %d ", 4, len(r))
	}
	t.Logf("default format, r[0]=%#v, r[1]=%#v, r[2]=%#v, r[3]=%#v ", r[0], r[1], r[2], r[3])
	t.Logf("string format, r[0]=%s, r[1]=%s, r[2]=%s, r[3]=%s ", string(r[0]), string(r[1]), string(r[2]), string(r[3]))
	t.Logf("len([]byte(string(r[0]))) = %d ", len([]byte(string(r[0]))))
	t.Logf("len([]byte(string(r[1]))) = %d ", len([]byte(string(r[1]))))
	t.Logf("len([]byte(string(r[2]))) = %d ", len([]byte(string(r[2]))))
	t.Logf("len([]byte(string(r[3]))) = %d ", len([]byte(string(r[3]))))
}
