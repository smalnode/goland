package lang

import "testing"

type errorMsg struct {
	status  int
	message string
}

type unsupportedErrorMsg struct {
	errorMsg
	operation string
}

func TestPanic(t *testing.T) {
	defer func() {
		e := recover()
		t.Log("handle all unhandled panic, but doing nothing ", e)
	}()

	defer func() {
		e := recover()
		if v, ok := e.(errorMsg); ok {
			t.Logf("recover errorMsg: %v ", v)
			return
		}
		t.Log("repanic from errorMsg handler ")
		panic(e)
	}()

	defer func() {
		e := recover()
		if v, ok := e.(unsupportedErrorMsg); ok {
			t.Logf("recover unsupportedErrorMsg: %v ", v)
			return
		}
		t.Log("repanic from unsupportedErrorMsg handler ")
		panic(e)
	}()

	u := unsupportedErrorMsg{errorMsg{404, "NOT FOUND"}, "GET"}
	e := errorMsg{500, "INTERNAL ERROR"}

	t.Log(u)
	t.Log(e)
	panic(1)
}
