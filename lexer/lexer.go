package lexer

import "github.com/ldb/lambda/token"

type Lexer struct {
	input   string
	pos     int
	readPos int
	ch      byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPos]
	}
	l.pos = l.readPos
	l.readPos += 1
}

func (l *Lexer) Next() token.Token {
	var tok token.Token

	switch l.ch {
	case '\\':
		tok = newToken(token.LAMBDA, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '.':
		tok = newToken(token.DOT, l.ch)
	case ' ':
		tok = newToken(token.SPACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Kind = token.EOF
	default:
		if isLetter(l.ch) {
			tok = newToken(token.IDENT, l.ch)
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func newToken(kind token.Kind, ch byte) token.Token {
	return token.Token{Kind: kind, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}
