package eval

import (
	"github.com/ldb/lambda/testutil"
	"testing"
)

func Test_isRedex(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"x", false},
		{"((u v) (x y))", false},
		{`((u v) \x.x)`, false},
		{`\y.(x u)`, false},
		{`(\x.y) u`, true},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			term := testutil.Parse(t, tc.input)
			r := isRedex(term)
			if r != tc.expected {
				t.Fatalf("unexpected result for input %s: got=%t, expected=%t", tc.input, r, tc.expected)
			}
		})
	}
}

func Test_isNF(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"x", true},
		{"((u v) (x y))", true},
		{`((u v) \x.x)`, true},
		{`\y.(x u)`, true},
		{`(\x.y) u`, false},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			term := testutil.Parse(t, tc.input)
			r := isNF(term)
			if r != tc.expected {
				t.Fatalf("unexpected result for input %s: got=%t, expected=%t", tc.input, r, tc.expected)
			}
		})
	}
}

func TestBetaReduce(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"x", ""},
		{"((u v) (x y))", ""},
		{`((u v) \x.x)`, "true"},
		{`\y.(x u)`, "true"},
		{`(\x.y) u`, "false"},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			term := testutil.Parse(t, tc.input)
			ss := BetaReduce(term)
			r := ss[len(ss)-1].String()
			if r != tc.expected {
				t.Fatalf("unexpected result for input %s: got=%s, expected=%s", tc.input, r, tc.expected)
			}
		})
	}
}
