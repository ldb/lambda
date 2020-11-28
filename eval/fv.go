package eval

import "github.com/ldb/lambda/ast"

// FV returns the Free Variables of an ast.Term.
// That is, a slice of ast.VariableTerm which are not bound by an abstraction.
func FV(term ast.Term) []*ast.VariableTerm {
	fv := make([]*ast.VariableTerm, 0)

	switch t := term.(type) {
	case *ast.VariableTerm:
		fv = vSetAppend(fv, t)
	case *ast.ApplicationTerm:
		fv = vSetAppend(fv, FV(t.Left)...)
		fv = vSetAppend(fv, FV(t.Right)...)
	case *ast.AbstractionTerm:
		fav := FV(t.Body)
		for i, v := range fav {
			if v.Value == t.Variable.Value {
				fav = vSetRemove(fav, i)
			}
		}
		fv = append(fv, fav...)
	}
	return fv
}

func vSetRemove(set []*ast.VariableTerm, i int) []*ast.VariableTerm {
	set[i] = set[len(set)-1]
	return set[:len(set)-1]
}

func vSetAppend(set []*ast.VariableTerm, items ...*ast.VariableTerm) []*ast.VariableTerm {
outer:
	for _, i := range items {
		for _, e := range set {
			if e.Value == i.Value {
				continue outer
			}
		}
		set = append(set, i)
	}
	return set
}
