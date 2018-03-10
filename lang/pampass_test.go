package lang

import (
	"testing"
)

type structWithPtr struct {
	intPtr *int
	strPtr *string
}

func addStructWithPtr(s structWithPtr) {
	*s.intPtr++
	*s.strPtr += "X"
}

func TestAddStructWithPtr(t *testing.T) {
	s := structWithPtr{new(int), new(string)}
	t.Logf("before: *s.intPtr = %d, *s.strPtr = %s ", *s.intPtr, *s.strPtr)
	addStructWithPtr(s)
	t.Logf("after: *s.intPtr = %d, *s.strPtr = %s ", *s.intPtr, *s.strPtr)
}
