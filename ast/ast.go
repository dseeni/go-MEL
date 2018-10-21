package ast

import (
	"bytes"

	"github.com/nrtkbb/go-MEL/token"
)

// Node is top of AST interface
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement have some expression
type Statement interface {
	Node
	statementNode()
}

// Expression ...
type Expression interface {
	Node
	expressionNode()
}

// Program is represent the entire program
type Program struct {
	Statements []Statement
}

// TokenLiteral ...
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// String ...
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// ExpressionStatement ...
type ExpressionStatement struct {
	Token      token.Token // first token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

// TokenLiteral ...
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

// String ...
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// InfixExpression ...
type InfixExpression struct {
	Token    token.Token // infix token. ex) '+' or '-' ...
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode() {}

// TokenLiteral ...
func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}

// String ...
func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" ")
	out.WriteString(ie.Operator)
	out.WriteString(" ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

// PrefixExpression ...
type PrefixExpression struct {
	Token    token.Token // prefix token. ex) '-' or '!'
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {}

// TokenLiteral ...
func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

// String ...
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

// PostfixExpression ...
type PostfixExpression struct {
	Token    token.Token // postfix token. ex) '--' or '++'
	Operator string
	Left     Expression
}

func (pe *PostfixExpression) expressionNode() {}

// TokenLiteral ...
func (pe *PostfixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

// String ...
func (pe *PostfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Left.String())
	out.WriteString(pe.Operator)
	out.WriteString(")")

	return out.String()
}

// TernaryExpression ...
type TernaryExpression struct {
	Conditional Expression
	Token1      token.Token // '?'
	Operator1   string      // "?"
	TrueExp     Expression
	Token2      token.Token // ':'
	Operator2   string      // ":"
	FalseExp    Expression
}

func (te *TernaryExpression) expressionNode() {}

// TokenLiteral ...
func (te *TernaryExpression) TokenLiteral() string {
	return te.Token1.Literal
}

// String ...
func (te *TernaryExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(te.Conditional.String())
	out.WriteString(" " + te.Operator1 + " ")
	out.WriteString(te.TrueExp.String())
	out.WriteString(" " + te.Operator2 + " ")
	out.WriteString(te.FalseExp.String())
	out.WriteString(")")

	return out.String()
}

// StringStatement ...
type StringStatement struct {
	Token token.Token // token.StringDec
	Name  *Identifier
	Value Expression
}

func (ls *StringStatement) statementNode() {}

// TokenLiteral ...
func (ls *StringStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// String ...
func (ls *StringStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral())
	out.WriteString(" ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// Identifier is token.Ident
type Identifier struct {
	Token token.Token // token.Ident
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral ...
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// String ...
func (i *Identifier) String() string {
	return i.Value
}

// ReturnStatement ...
type ReturnStatement struct {
	Token       token.Token // token.Return
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral ...
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

// String ...
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral())
	out.WriteString(" ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// IntegerLiteral ...
type IntegerLiteral struct {
	Token token.Token // token.Int
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

// TokenLiteral ...
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

// String ...
func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

// StringLiteral ...
type StringLiteral struct {
	Token token.Token // token.String
	Value string
}

func (sl *StringLiteral) expressionNode() {}

// TokenLiteral ...
func (sl *StringLiteral) TokenLiteral() string {
	return sl.Token.Literal
}

// String ...
func (sl *StringLiteral) String() string {
	return sl.Token.Literal
}
