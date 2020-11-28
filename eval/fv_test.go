package eval

import (
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
			l := lexer.New(tc.input)
			p := parser.New(l)
			lt := p.ParseLambdaTerm()

			fv := FV(lt.Term)
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
