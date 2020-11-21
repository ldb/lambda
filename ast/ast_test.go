package ast

import (
	"github.com/ldb/lambda/token"
	"testing"
)

func TestString(t *testing.T) {
	lt := LambdaTerm{
		Term: &ApplicationTerm{
			Token: token.Token{},
			Left: &AbstractionTerm{
				Token: token.Token{Kind: 0, Literal: ""},
				Variable: &VariableTerm{
					Token: token.Token{Kind: token.IDENT, Literal: "x"},
					Value: "x",
				},
				Body: &ApplicationTerm{
					Token: token.Token{Kind: 0, Literal: ""},
					Left: &AbstractionTerm{
						Token: token.Token{},
						Variable: &VariableTerm{
							Token: token.Token{Kind: token.IDENT, Literal: "y"},
							Value: "y",
						},
						Body: &VariableTerm{
							Token: token.Token{Kind: token.IDENT, Literal: "y"},
							Value: "y",
						},
					},
					Right: &VariableTerm{
						Token: token.Token{Kind: token.IDENT, Literal: "x"},
						Value: "x",
					},
				},
			},
			Right: &VariableTerm{
				Token: token.Token{Kind: token.IDENT, Literal: "z"},
				Value: "z",
			},
		},
	}
	if lt.String() != `((\x.((\y.y) x)) z)` {
		t.Errorf("lt.String() wrong. got = %s", lt.String())
	}
}
