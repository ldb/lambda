package eval

import "github.com/ldb/lambda/ast"

// FV returns the Free Variables (Notation FV(M)) of an ast.Term M.
// That is, a slice of ast.VariableTerm which are not bound by an abstraction.
func FV(M ast.Term) []*ast.VariableTerm {
	fv := make([]*ast.VariableTerm, 0)

	switch t := M.(type) {
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
	for _, i := range items {
		if vSetIn(set, i) {
			continue
		}
		set = append(set, i)
	}
	return set
}

func vSetIn(set []*ast.VariableTerm, item *ast.VariableTerm) bool {
	for _, e := range set {
		if e.Value == item.Value {
			return true
		}
	}
	return false
}

// substituteBound substitutes N for free occurrences of x in M, notation M[x:=N]:
// x[x:=N]       === N;
// y[x:=N]       === y;
// (P Q)[x:=N]   === (P[x:=N])(Q[x:=N]);
// (\y.P)[x:=N]  === (\y.(P[x:=N]); provided x !== y and y not in FV(N)
// (\x.(P[x:=N]) === (\x.P)
func SubstituteFree(M ast.Term, x *ast.VariableTerm, N ast.Term) ast.Term {
	if !vSetIn(FV(M), x) {
		return M
	}

	switch t := M.(type) {
	case *ast.VariableTerm:
		if t.Value == x.Value {
			M = N
		}
	case *ast.ApplicationTerm:
		t.Left = SubstituteFree(t.Left, x, N)
		t.Right = SubstituteFree(t.Right, x, N)
	case *ast.AbstractionTerm:
		if vSetIn(FV(N), t.Variable) {
			break
		}
		t.Body = SubstituteFree(t.Body, x, N)
	}
	return M
}
