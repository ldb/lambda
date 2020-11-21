package lexer

import (
	"github.com/ldb/lambda/token"
	"testing"
)

func TestLexer_Next(t *testing.T) {
	inputs := []string{`\x.x`,
		`\x.(x x)`,
		`\x.(x y)x`,
		`(\xy.x y)x y`,
		`(x y)(\x.x)z`,
	}

	testCases := [][]struct {
		expectedKind    token.Kind
		expectedLiteral string
	}{
		{
			{expectedKind: token.LAMBDA, expectedLiteral: "\\"},
			{expectedKind: token.IDENT, expectedLiteral: "x"},
			{expectedKind: token.DOT, expectedLiteral: "."},
			{expectedKind: token.IDENT, expectedLiteral: "x"},
			{expectedKind: token.EOF, expectedLiteral: ""},
		}, {
			{expectedKind: token.LAMBDA, expectedLiteral: "\\"},
			{expectedKind: token.IDENT, expectedLiteral: "x"},
			{expectedKind: token.DOT, expectedLiteral: "."},
			{expectedKind: token.LPAREN, expectedLiteral: "("},
			{expectedKind: token.IDENT, expectedLiteral: "x"},
			{expectedKind: token.SPACE, expectedLiteral: " "},
			{expectedKind: token.IDENT, expectedLiteral: "x"},
			{expectedKind: token.RPAREN, expectedLiteral: ")"},
			{expectedKind: token.EOF, expectedLiteral: ""},
		}, {
			{expectedKind: token.LAMBDA, expectedLiteral: "\\"},
			{expectedKind: token.IDENT, expectedLiteral: "x"},
			{expectedKind: token.DOT, expectedLiteral: "."},
			{expectedKind: token.LPAREN, expectedLiteral: "("},
			{expectedKind: token.IDENT, expectedLiteral: "x"},
			{expectedKind: token.SPACE, expectedLiteral: " "},
			{expectedKind: token.IDENT, expectedLiteral: "y"},
			{expectedKind: token.RPAREN, expectedLiteral: ")"},
			{expectedKind: token.IDENT, expectedLiteral: "x"},
			{expectedKind: token.EOF, expectedLiteral: ""},
		}, {
			{expectedKind: token.LPAREN, expectedLiteral: "("},
			{expectedKind: token.LAMBDA, expectedLiteral: "\\"},
			{expectedKind: token.IDENT, expectedLiteral: "x"},
			{expectedKind: token.IDENT, expectedLiteral: "y"},
			{expectedKind: token.DOT, expectedLiteral: "."},
			{expectedKind: token.IDENT, expectedLiteral: "x"},
			{expectedKind: token.SPACE, expectedLiteral: " "},
			{expectedKind: token.IDENT, expectedLiteral: "y"},
			{expectedKind: token.RPAREN, expectedLiteral: ")"},
			{expectedKind: token.IDENT, expectedLiteral: "x"},
			{expectedKind: token.SPACE, expectedLiteral: " "},
			{expectedKind: token.IDENT, expectedLiteral: "y"},
			{expectedKind: token.EOF, expectedLiteral: ""},
		}, {
			{expectedKind: token.LPAREN, expectedLiteral: "("},
			{expectedKind: token.IDENT, expectedLiteral: "x"},
			{expectedKind: token.SPACE, expectedLiteral: " "},
			{expectedKind: token.IDENT, expectedLiteral: "y"},
			{expectedKind: token.RPAREN, expectedLiteral: ")"},
			{expectedKind: token.LPAREN, expectedLiteral: "("},
			{expectedKind: token.LAMBDA, expectedLiteral: "\\"},
			{expectedKind: token.IDENT, expectedLiteral: "x"},
			{expectedKind: token.DOT, expectedLiteral: "."},
			{expectedKind: token.IDENT, expectedLiteral: "x"},
			{expectedKind: token.RPAREN, expectedLiteral: ")"},
			{expectedKind: token.IDENT, expectedLiteral: "z"},
			{expectedKind: token.EOF, expectedLiteral: ""},
		},
	}

	for i, input := range inputs {
		l := New(input)
		for i, tc := range testCases[i] {
			tok := l.Next()

			if tok.Kind != tc.expectedKind {
				t.Fatalf("testCases[%d] - kind wrong. expected=%q, got=%q", i, tc.expectedKind, tok.Kind)
			}
			if tok.Literal != tc.expectedLiteral {
				t.Fatalf("testCases[%d] - literal wrong. expected=%q, got=%q", i, tc.expectedLiteral, tok.Literal)
			}
		}
	}
}
