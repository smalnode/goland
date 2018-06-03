package lang

import "testing"

type Interface interface {
	addOne()
}

type si struct {
	value int
}

func (s si) addOne() {
	s.value++
}

func callAddOne(i Interface) {
	i.addOne()
}

func TestPassInterface(t *testing.T) {
	s := si{}
	callAddOne(s)
	if s.value == 0 {
		t.Log("pass interface value ")
	} else {
		t.Log("pass interface ref ")
	}
}
