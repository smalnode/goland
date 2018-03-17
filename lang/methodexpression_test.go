package lang

import (
	"fmt"
	"testing"
)

type ms struct {
	value int
}

func (ms *ms) foo(v int) {
	fmt.Printf("method struct: [%d] ", ms.value*v)
}

func TestMethodExpression(t *testing.T) {
	(interface {
		foo(int)
	}).foo(&ms{}, 0)
}
