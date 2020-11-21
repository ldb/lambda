package token

type Kind int

type Token struct {
	Kind    Kind
	Literal string
}

const (
	ILLEGAL Kind = iota // Illegal token
	EOF                 // EOF
	LAMBDA              // "\"
	LPAREN              // "("
	RPAREN              // ")"
	DOT                 // "."
	IDENT               // any ASCII Char sequence
	SPACE               // " "
)

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",
	LAMBDA:  "LAMBDA",
	LPAREN:  "LPAREN",
	RPAREN:  "RPAREN",
	DOT:     "DOT",
	IDENT:   "IDENT",
	SPACE:   "SPACE",
}

func (k Kind) String() string {
	if k < 0 || Kind(len(tokens)) < k {
		return ""
	}
	return tokens[k]
}
