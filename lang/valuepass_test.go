package lang

import "testing"

func passStrut(ss s) {
	ss.v++
}

func passStructPtr(ss *s) {
	ss.v++
}

func (ss s) addByStruct() {
	ss.v++
}

func (ss *s) addByStructPtr() {
	ss.v++
}

func TestStruct(t *testing.T) {
	var ss s
	ss = s{"A", 100}
	passStrut(ss)
	if ss.v == 100 {
		t.Log("pass struct is value ")
	} else {
		t.Log("pass struct is reference ")
	}

	ss = s{"A", 100}
	ss.addByStruct()
	if ss.v == 100 {
		t.Log("method pass struct is value ")
	} else {
		t.Log("method pass struct is reference ")
	}

	ss = s{"A", 100}
	passStructPtr(&ss)
	if ss.v == 100 {
		t.Log("pass struct ptr is value ")
	} else {
		t.Log("pass struct ptr is reference ")
	}

	ss = s{"A", 100}
	(&ss).addByStructPtr()
	if ss.v == 100 {
		t.Log("method pass struct ptr is value ")
	} else {
		t.Log("method pass struct ptr is reference ")
	}
}

func addArray(a [2]int) {
	for k := range a {
		a[k]++
	}
}

func addSlice(a []int) {
	for k := range a {
		a[k]++
	}
}

func TestArray(t *testing.T) {
	a := [2]int{0, 1}
	addArray(a)
	for k, v := range a {
		if v == k {
			t.Log("pass array is value ")
		} else {
			t.Log("pass array is reference ")
		}
	}

	a = [2]int{0, 1}
	addSlice(a[:])
	for k, v := range a {
		if v == k {
			t.Log("pass slice is value ")
		} else {
			t.Log("pass slice is reference ")
		}
	}
}

func addQP(q q) {
	q.p.pp++
}

func addQPPtr(q *q) {
	q.p.pp++
}

func TestPassNestedStruct(t *testing.T) {
	q := q{}
	addQP(q)
	if q.p.pp == 0 {
		t.Log("nested struct pass by value ")
	} else {
		t.Log("nested struct pass by reference ")
	}

	q.p.pp = 0
	addQPPtr(&q)
	if q.p.pp == 0 {
		t.Log("nested struct pass by value ")
	} else {
		t.Log("nested struct pass by reference ")
	}
}
