package eval

import (
	"fmt"
	"github.com/ldb/lambda/ast"
)

func BetaReduce(term ast.Term) TermSet {
	steps := TermSet{s: []ast.Term{term}}
	if isNF(term) {
		return steps
	}
	se := subExpressions(term)
	for _, e := range se.s {
		if !isRedex(e) {
			return steps
		}
		P := e.(*ast.ApplicationTerm)
		Q := P.Left.(*ast.AbstractionTerm)

		if vSetIn(FV(P.Right), Q.Variable) {
			// TODO: Perform Alpha Conversion First
			fmt.Printf("error: %s appears in FV(%s). Alpha Conversion necessary but not supported yet.\n", Q.Variable, P.Right)
			return steps
		}
		t := SubstituteFree(Q.Body, Q.Variable, P.Right)
		if t.String() == term.String() { // Fixed-Point reached. (e.g for terms like "(\x.(x x)) (\y.(y y))" ).
			fmt.Printf("fixed point reached for %s. Aborting reduction.\n", term.String())
			return steps
		}
		r := BetaReduce(t)
		steps.Add(r.Slice()...)
	}
	return steps
}

// isRedex returns true if the ast.Term T is a Redex, that is a Term of form ((\x.M) N).
func isRedex(T ast.Term) bool {
	P, ok := T.(*ast.ApplicationTerm)
	if !ok {
		return false
	}
	_, ok = P.Left.(*ast.AbstractionTerm)
	if !ok {
		return false
	}
	return true
}

// isNF returns true if the ast.Term T is in beta-normal form (beta-nf), that is it does not have a Redex as subexpression.
func isNF(T ast.Term) bool {
	se := subExpressions(T)
	for _, e := range se.s {
		if isRedex(e) {
			return false
		}
	}
	return true
}

// subExpressions returns subexpressions of ast.Term T.
func subExpressions(T ast.Term) TermSet {
	ts := TermSet{s: []ast.Term{}}

	switch t := T.(type) {
	case *ast.VariableTerm:
		ts.Add(t)
	case *ast.ApplicationTerm:
		ts.Add(t)
		ts.Add(subExpressions(t.Left).s...)
		ts.Add(subExpressions(t.Right).s...)
	case *ast.AbstractionTerm:
		ts.Add(t)
		ts.Add(subExpressions(t.Body).s...)
	}
	return ts
}
