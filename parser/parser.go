package parser

import (
	"fmt"
	"github.com/ldb/lambda/ast"
	"github.com/ldb/lambda/lexer"
	"github.com/ldb/lambda/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) ParseLambdaTerm() *ast.LambdaTerm {
	lt := &ast.LambdaTerm{}

	if p.curTokenIs(token.EOF) {
		return nil
	}

	lt.Term = p.parseTerm()
	return lt
}

func (p *Parser) parseTerm() ast.Term {
	switch p.curToken.Kind {
	case token.IDENT:
		return p.parseVariableTerm()
	case token.LPAREN:
		if p.peekTokenIs(token.LAMBDA) {
			return p.parseAbstractionTerm()
		}
		return p.parseApplicationTerm()
	default:
		p.currentTokenError()
		return nil
	}
}

func (p *Parser) parseVariableTerm() *ast.VariableTerm {
	return &ast.VariableTerm{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseAbstractionTerm() *ast.AbstractionTerm {
	// (
	t := &ast.AbstractionTerm{Token: p.curToken}
	p.nextToken() // \
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	t.Variable = p.parseVariableTerm()
	if !p.expectPeek(token.DOT) {
		return nil
	}
	p.nextToken() // .
	t.Body = p.parseTerm()
	p.nextToken() // )
	return t
}

func (p *Parser) parseApplicationTerm() *ast.ApplicationTerm {
	// (
	t := &ast.ApplicationTerm{Token: p.curToken}
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	t.Left = p.parseTerm() // LHS
	if !p.expectPeek(token.SPACE) {
		return nil
	}
	p.nextToken() // RHS
	t.Right = p.parseTerm()
	p.nextToken() // )
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
	e := fmt.Sprintf("unexpected next Token: %s, expected %s instead", p.peekToken.Kind, k)
	p.errors = append(p.errors, e)
}

func (p *Parser) currentTokenError() {
	e := fmt.Sprintf("unexpected Token: %s, expected variable or (", p.curToken.Kind.String())
	p.errors = append(p.errors, e)
}
