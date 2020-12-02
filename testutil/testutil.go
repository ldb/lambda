package testutil

import (
	"github.com/ldb/lambda/ast"
	"github.com/ldb/lambda/lexer"
	"github.com/ldb/lambda/parser"
	"testing"
)

func Parse(t *testing.T, s string) ast.Term {
	l := lexer.New(s)
	p := parser.New(l)
	lt := p.ParseLambdaTerm()
	if err := p.Error(); err != nil {
		t.Fatalf("parser error: error parsing input %q: %v", s, err)
	}
	return lt.Term
}
