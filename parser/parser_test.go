package parser

import (
	"github.com/ldb/lambda/ast"
	"github.com/ldb/lambda/lexer"
	"testing"
)

func TestParseVariableTerm(t *testing.T) {
	testCases := []struct {
		input         string
		expectedValue string
	}{
		{"x", "x"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			l := lexer.New(tc.input)
			p := New(l)
			lt := p.ParseLambdaTerm()
			checkParserErrors(t, p)

			term, ok := lt.Term.(*ast.VariableTerm)
			if !ok {
				t.Fatalf("term is not ast.VariableTerm, got %T", lt.Term)
			}
			if term.Value != tc.expectedValue {
				t.Fatalf("unexpected ast.VariableTerm.Value, got %s", term.Value)
			}
		})
	}
}

func TestParsAbstractionTerm(t *testing.T) {
	testCases := []struct {
		input               string
		expectedVariable    string
		expectedBodyLiteral string
	}{
		{`(\x.x)`, "x", "x"},
		{`(\x.(x x))`, "x", "(x x)"},
		{`(\x.(\x.x))`, "x", `(\x.x)`},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			l := lexer.New(tc.input)
			p := New(l)
			lt := p.ParseLambdaTerm()
			checkParserErrors(t, p)

			term, ok := lt.Term.(*ast.AbstractionTerm)
			if !ok {
				t.Fatalf("term is not ast.VariableTerm, got %T", lt.Term)
			}
			if term.Variable.Value != tc.expectedVariable {
				t.Fatalf("unexpected ast.VariableTerm.Value, got %s", term.Variable.Value)
			}
			if term.Body.String() != tc.expectedBodyLiteral {
				t.Fatalf("unexpected ast.VariableTerm.Value, got %s", term.Body.String())
			}
		})
	}
}

func TestParseApplicationTerm(t *testing.T) {
	testCases := []struct {
		input              string
		expectedLHSLiteral string
		expectedRHSLiteral string
	}{
		{"(x x)", "x", "x"},
		{"x x", "x", "x"},
		{"(x y)", "x", "y"},
		{"(x (x x))", "x", "(x x)"},
		{"x (x x)", "x", "(x x)"},
		{"((x x) (x x))", "(x x)", "(x x)"},
		{"(x x) (x x)", "(x x)", "(x x)"},
		{`(x (\x.x))`, "x", `(\x.x)`},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			l := lexer.New(tc.input)
			p := New(l)
			lt := p.ParseLambdaTerm()
			checkParserErrors(t, p)

			term, ok := lt.Term.(*ast.ApplicationTerm)
			if !ok {
				t.Fatalf("term is not ast.VariableTerm, got %T", lt.Term)
			}
			if term.Left.String() != tc.expectedLHSLiteral {
				t.Fatalf("unexpected ast.VariableTerm.Left, got %s", term.Left.String())
			}
			if term.Right.String() != tc.expectedRHSLiteral {
				t.Fatalf("unexpected ast.VariableTerm.Right, got %s", term.Right.String())
			}
		})
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
