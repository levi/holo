package parser

import (
	"github.com/levi/holo/expr"
	"github.com/levi/holo/token"
)

// Parser parses a flat sequence of tokens into an AST, reporting errors when encountered
type Parser struct {
	tokens  []*token.Token
	current int
}

// NewParser allocates a new parser with a sequence of tokens to parse
func NewParser(tokens []*token.Token) *Parser {
	p := new(Parser)
	p.tokens = tokens
	return p
}

// Parse parses the token sequence
func (p *Parser) Parse() ([]expr.Stmt, error) {
	var statements []expr.Stmt
	for !p.isAtEnd() {
		statement, err := p.statement()
		if err != nil {
			return statements, err
		}
		statements = append(statements, statement)
	}
	return statements, nil
}

func (p *Parser) statement() (expr.Expr, error) {
	if p.match(token.PrintToken) {
		return p.printStatement()
	}
	return p.expressionStatement()
}

func (p *Parser) expressionStatement() (expr.Stmt, error) {
	value, err := p.expression()
	if err != nil {
		return nil, err
	}
	p.consume(token.SemicolonToken, "Expected ';' after value.")
	return expr.NewExpression(value), nil
}

func (p *Parser) printStatement() (expr.Stmt, error) {
	value, err := p.expression()
	if err != nil {
		return nil, err
	}
	p.consume(token.SemicolonToken, "Expected ';' after value.")
	return expr.NewPrint(value), nil
}

func (p *Parser) expression() (expr.Expr, error) {
	return p.equality()
}

func (p *Parser) equality() (expr.Expr, error) {
	e, err := p.comparison()
	if err != nil {
		return nil, err
	}

	for p.match(token.BangEqualToken, token.EqualEqualToken) {
		operator := p.previous()
		right, err := p.comparison()
		if err != nil {
			return nil, err
		}
		e = expr.NewBinary(e, *operator, right)
	}

	return e, nil
}

func (p *Parser) comparison() (expr.Expr, error) {
	e, err := p.addition()
	if err != nil {
		return nil, err
	}

	for p.match(token.GreaterToken, token.GreaterEqualToken, token.LessToken, token.LessEqualToken) {
		operator := p.previous()
		right, err := p.addition()
		if err != nil {
			return nil, err
		}
		e = expr.NewBinary(e, *operator, right)
	}
	return e, nil
}

func (p *Parser) addition() (expr.Expr, error) {
	e, err := p.multiplication()
	if err != nil {
		return nil, err
	}

	for p.match(token.MinusToken, token.PlusToken) {
		operator := p.previous()
		right, err := p.multiplication()
		if err != nil {
			return nil, err
		}
		e = expr.NewBinary(e, *operator, right)
	}
	return e, nil
}

func (p *Parser) multiplication() (expr.Expr, error) {
	e, err := p.unary()
	if err != nil {
		return nil, err
	}
	for p.match(token.SlashToken, token.StarToken) {
		operator := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		e = expr.NewBinary(e, *operator, right)
	}
	return e, nil
}

func (p *Parser) unary() (expr.Expr, error) {
	if p.match(token.BangToken, token.MinusToken) {
		operator := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		return expr.NewUnary(*operator, right), nil
	}

	return p.primary()
}

func (p *Parser) primary() (expr.Expr, error) {
	if p.match(token.FalseToken) {
		return expr.NewLiteral(false), nil
	} else if p.match(token.TrueToken) {
		return expr.NewLiteral(true), nil
	} else if p.match(token.NilToken) {
		return expr.NewLiteral(nil), nil
	} else if p.match(token.NumberToken, token.StringToken) {
		return expr.NewLiteral(p.previous().Literal), nil
	} else if p.match(token.LeftParenToken) {
		e, err := p.expression()
		if err != nil {
			return nil, err
		}

		_, err = p.consume(token.RightParenToken, "Expected ')' after expression")
		if err != nil {
			return nil, err
		}

		return expr.NewGrouping(e), nil
	}

	return nil, NewParseError(*(p.peek()), "Expected expression.")
}

func (p *Parser) match(types ...string) bool {
	for _, t := range types {
		if p.check(t) {
			p.advance()
			return true
		}
	}

	return false
}

func (p *Parser) consume(token, description string) (*token.Token, error) {
	if p.check(token) {
		return p.advance(), nil
	}
	return nil, NewParseError(*(p.peek()), description)
}

func (p *Parser) check(t string) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().TokenType == t
}

func (p *Parser) advance() *token.Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.peek().TokenType == token.EOFToken
}

func (p *Parser) peek() *token.Token {
	return p.tokens[p.current]
}

func (p *Parser) previous() *token.Token {
	return p.tokens[p.current-1]
}

func (p *Parser) sync() {
	p.advance()
	for !p.isAtEnd() {
		if p.previous().TokenType == token.SemicolonToken {
			return
		}

		switch p.peek().TokenType {
		case token.ClassToken:
		case token.FnToken:
		case token.VarToken:
		case token.ForToken:
		case token.IfToken:
		case token.WhileToken:
		case token.PrintToken:
		case token.ReturnToken:
			return
		}

		p.advance()
	}
}
