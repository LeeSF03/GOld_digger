package parser

import (
	"GOld_digger/ast"
	"GOld_digger/lexer"
	"GOld_digger/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token // current Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
