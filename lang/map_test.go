package lang

import "testing"

type mm struct {
	m int
}

func TestMapValueRef(t *testing.T) {
	m := make(map[int]int, 8)

	m[1] = 1
	v := m[1]
	v = 2
	if v == m[1] {
		t.Log("map[key] return ref ")
	} else {
		t.Log("map[key] return value ")
	}
}

func TestMapStruct(t *testing.T) {
	m := make(map[int]mm, 8)
	m[1] = mm{1}
	m1 := m[1]
	m1.m = 2
	if m1.m == m[1].m {
		t.Log("map[key] return ref ")
	} else {
		t.Log("map[key] return value ")
	}
}
