package eval

import (
	"github.com/ldb/lambda/ast"
	"github.com/ldb/lambda/lexer"
	"github.com/ldb/lambda/parser"
	"testing"
)

func TestFV(t *testing.T) {
	testCases := []struct {
		input    string
		expected []string
	}{
		{"x", []string{"x"}},
		{`\x.x`, []string{}},
		{`\x.y`, []string{"y"}},
		{"x y", []string{"x", "y"}},
		{`\x.(x y) y`, []string{"y"}},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			term := parse(tc.input)
			fv := FV(term)
			if len(fv) != len(tc.expected) {
				t.Fatalf("unequal lengh: got=%d, expected=%d", len(fv), len(tc.expected))
			}
			for i, v := range fv {
				if tc.expected[i] != v.Value {
					t.Fatalf("unexpected free variable at %d: got=%q, expected=%q", i, tc.expected[i], v.Value)
				}
			}
		})
	}
}

func TestSubstitute(t *testing.T) {
	testCases := []struct {
		inputM   string
		inputx   string
		inputN   string
		expected string
	}{
		{"x", "x", "z", "z"},
		{"y", "x", "z", "y"},
		{"((u v) (x y))", "x", "z", "((u v) (z y))"},
		{`((u v) \x.x)`, "x", "z", `((u v) (\x.x))`},
		{`\y.(x u)`, "x", "z", `(\y.(z u))`},
		{`\x.(x u)`, "x", "z", `(\x.(x u))`},
	}

	for _, tc := range testCases {
		t.Run(tc.inputM, func(t *testing.T) {
			M := parse(tc.inputM)
			xt := parse(tc.inputx)
			N := parse(tc.inputN)

			x, ok := xt.(*ast.VariableTerm)
			if !ok {
				t.Fatal("inputx is not a valid *ast.VariableTerm")
			}

			e := Substitute(M, x, N)
			if e.String() != tc.expected {
				t.Fatalf("unexpected substitution result: got=%s, expected=%s", e.String(), tc.expected)
			}
		})
	}
}

func parse(s string) ast.Term {
	l := lexer.New(s)
	p := parser.New(l)
	lt := p.ParseLambdaTerm()
	return lt.Term
}
