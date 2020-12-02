package parser

import (
	"fmt"
	"github.com/ldb/lambda/ast"
	"github.com/ldb/lambda/lexer"
	"github.com/ldb/lambda/token"
)

type ParseError struct {
	Position int
	msg      error
}

func (e *ParseError) Error() string {
	return e.msg.Error()
}

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	error     error
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l: l,
	}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) Error() error {
	return p.error
}

func (p *Parser) ParseLambdaTerm() *ast.LambdaTerm {
	firstToken := p.curToken

	var t ast.Term
	switch p.curToken.Kind {
	case token.IDENT:
		t = p.parseVariableTerm()
	case token.LAMBDA:
		t = p.parseAbstractionTerm()
	case token.LPAREN:
		t = p.parseTerm()
	}

	if p.peekTokenIs(token.EOF) {
		return &ast.LambdaTerm{Term: t}
	}
	p.expectPeek(token.SPACE)
	p.nextToken()
	lt := &ast.LambdaTerm{
		Term: &ast.ApplicationTerm{
			Token: firstToken,
			Left:  t,
			Right: p.parseTerm(),
		},
	}

	p.expectPeek(token.EOF)
	return lt
}

func (p *Parser) parseTerm() ast.Term {
	switch p.curToken.Kind {
	case token.IDENT:
		return p.parseVariableTerm()
	case token.LAMBDA:
		return p.parseAbstractionTerm()
	case token.LPAREN:
		var t ast.Term
		p.nextToken()
		if p.curTokenIs(token.LAMBDA) {
			t = p.parseAbstractionTerm()
		} else {
			t = p.parseApplicationTerm()
		}
		if !p.expectPeek(token.RPAREN) {
			return nil
		}
		return t
	default:
		p.currentTokenError(token.LPAREN, token.IDENT)
		return nil
	}
}

func (p *Parser) parseVariableTerm() *ast.VariableTerm {
	return &ast.VariableTerm{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseAbstractionTerm() *ast.AbstractionTerm {
	t := &ast.AbstractionTerm{Token: p.curToken}
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	t.Variable = p.parseVariableTerm()
	if !p.expectPeek(token.DOT) {
		return nil
	}
	p.nextToken() // .
	t.Body = p.parseTerm()
	return t
}

func (p *Parser) parseApplicationTerm() *ast.ApplicationTerm {
	t := &ast.ApplicationTerm{Token: p.curToken}
	t.Left = p.parseTerm() // LHS
	if !p.expectPeek(token.SPACE) {
		return nil
	}
	p.nextToken() // RHS
	t.Right = p.parseTerm()
	return t
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.Next()
}

func (p *Parser) curTokenIs(k token.Kind) bool {
	return p.curToken.Kind == k
}

func (p *Parser) peekTokenIs(k token.Kind) bool {
	return p.peekToken.Kind == k
}

func (p *Parser) expectPeek(k token.Kind) bool {
	if !p.peekTokenIs(k) {
		p.peekTokenError(k)
		return false
	}
	p.nextToken()
	return true
}

func (p *Parser) peekTokenError(k token.Kind) {
	p.error = &ParseError{
		msg:      fmt.Errorf("unexpected next Token at position %d: expected %s, got %s instead", p.peekToken.Position, k, p.peekToken.Kind),
		Position: p.peekToken.Position,
	}
}

func (p *Parser) currentTokenError(k ...token.Kind) {
	p.error = &ParseError{
		msg:      fmt.Errorf("unexpected Token at position %d: expected %s, got expected %s instead", p.curToken.Position, k, p.curToken.Kind),
		Position: p.curToken.Position,
	}
}
