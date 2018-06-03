package closure

import (
	"fmt"
	"sort"
	"testing"
)

type age int
type height float64

// Person is a ordernary struct
type Person struct {
	Name   string
	Age    age
	Height height
}

// By return whether lhs < rhs
type By func(lhs, rhs *Person) bool

// Sort receive a sort func and a list
func (by By) Sort(persons []Person) {
	ps := &personSorter{
		persons: persons,
		by:      by,
	}
	sort.Sort(ps)
}

// a wrapper struct
type personSorter struct {
	persons []Person
	by      By
}

func (ps *personSorter) Len() int {
	return len(ps.persons)
}

func (ps *personSorter) Less(i, j int) bool {
	return ps.by(&ps.persons[i], &ps.persons[j])
}

func (ps *personSorter) Swap(i, j int) {
	ps.persons[i], ps.persons[j] = ps.persons[j], ps.persons[i]
}

func TestClosure(t *testing.T) {
	var ps = []Person{
		{"A", 1, 1.0},
		{"B", 2, 2.0},
		{"C", 3, 3.0},
		{"D", 4, 4.0},
		{"E", 5, 5.0},
	}

	decreasingHeight := func(lhs, rhs *Person) bool {
		return rhs.Height < lhs.Height
	}
	for _, p := range ps {
		fmt.Printf("%v\n", p)
	}

	By(decreasingHeight).Sort(ps)
	var pre, cur Person
	for k, v := range ps {
		pre = cur
		cur = v
		if &pre != nil && pre.Height < cur.Height {
			t.Errorf("element %d:%v < %d:%v ", k-1, pre, k, v)
		}
	}
}
