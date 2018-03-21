package lang

import "testing"

// Fooer a interface can foo
type Fooer interface {
	Foo() string
}

type cat struct{}

func (c cat) Foo() string {
	return "meow"
}

type tiger struct {
	cat
}

func (t tiger) Foo() string {
	return "hmmmm"
}

type panther struct {
	cat
}

func TestFooer(t *testing.T) {
	var foo Fooer
	foo = tiger{}
	t.Log(foo.Foo())
	if tt, ok := foo.(tiger); ok {
		t.Log(tt.cat.Foo())
	}

	foo = panther{}
	t.Log(foo.Foo())
	p := panther{}
	t.Log(p.Foo())
	t.Log(p.cat.Foo())
}
