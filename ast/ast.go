package ast

import (
	"bytes"
	"github.com/ldb/lambda/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Term interface {
	Node
	Copy() Term
	termNode()
}

type LambdaTerm struct {
	Term Term
}

func (l *LambdaTerm) TokenLiteral() string {
	return l.Term.TokenLiteral()
}
func (l *LambdaTerm) String() string {
	return l.Term.String()
}

type VariableTerm struct {
	Token token.Token
	Value string
}

func (v *VariableTerm) termNode() {}
func (v *VariableTerm) TokenLiteral() string {
	return v.Token.Literal
}
func (v *VariableTerm) Copy() Term {
	vc := *v
	return Term(&vc)
}
func (v *VariableTerm) String() string {
	return v.Value
}

type ApplicationTerm struct {
	Token token.Token
	Left  Term
	Right Term
}

func (a *ApplicationTerm) termNode() {}
func (a *ApplicationTerm) TokenLiteral() string {
	return a.Token.Literal
}
func (a *ApplicationTerm) Copy() Term {
	ac := *a
	ac.Left = a.Left.Copy()
	ac.Right = a.Right.Copy()
	return Term(&ac)
}
func (a *ApplicationTerm) String() string {
	out := bytes.Buffer{}
	out.WriteString("(")
	out.WriteString(a.Left.String())
	out.WriteString(" ")
	out.WriteString(a.Right.String())
	out.WriteString(")")
	return out.String()
}

type AbstractionTerm struct {
	Token    token.Token
	Variable *VariableTerm
	Body     Term
}

func (a *AbstractionTerm) termNode() {}
func (a *AbstractionTerm) TokenLiteral() string {
	return a.Token.Literal
}
func (a *AbstractionTerm) Copy() Term {
	ac := *a
	ac.Variable = a.Variable.Copy().(*VariableTerm)
	ac.Body = a.Body.Copy()
	return Term(&ac)
}
func (a *AbstractionTerm) String() string {
	out := bytes.Buffer{}
	out.WriteString("(")
	out.WriteString("\\")
	out.WriteString(a.Variable.String())
	out.WriteString(".")
	out.WriteString(a.Body.String())
	out.WriteString(")")
	return out.String()
}
