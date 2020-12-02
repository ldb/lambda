package eval

import "github.com/ldb/lambda/ast"

type TermSet struct {
	s []ast.Term
}

func (ts *TermSet) Remove(i int) *TermSet {
	ts.s[i] = ts.s[len(ts.s)-1]
	ts.s = ts.s[:len(ts.s)-1]
	return ts
}

func (ts *TermSet) Add(items ...ast.Term) *TermSet {
	for _, i := range items {
		if ts.In(i) {
			continue
		}
		ts.s = append(ts.s, i)
	}
	return ts
}

func (ts *TermSet) In(item ast.Term) bool {
	for _, e := range ts.s {
		if e.String() == item.String() {
			return true
		}
	}
	return false
}

func (ts *TermSet) Slice() []ast.Term {
	return ts.s
}
