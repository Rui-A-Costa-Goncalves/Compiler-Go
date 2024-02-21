package parser

import (
	"goCompiler/lexer"
	"goCompiler/ast"
	"goCompiler/token"
	"fmt"
)

type Parser struct {
	l *lexer.Lexer
	curToken token.Token
	nextToken token.Token
	errors []string
	prefixOperatorParserFns map[token.TokenType]prefixOperatorParserFn
	infixOperatorParserFns map[token.TokenType]infixOperatorParserFn
}

type prefixOperatorParserFn func() ast.Expression
type infixOperatorParserFn func(ast.Expression) ast.Expression // the Expression that it receives as argument is the expression before the operator

func NewParser(l *lexer.Lexer) *Parser {
	p := Parser{l :l,errors: []string{}}
	//populate the first two tokens
	p.getToken()
	p.getToken()
	return &p
}

func (p *Parser) getToken() {
	token := p.l.NextToken()
	p.curToken = p.nextToken
	p.nextToken = token
}

//function that will read the tokens and create the ast
func (p *Parser) parseProgram() *ast.Program {
	var prog ast.Program
	prog.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		switch (p.curToken.Type) {
		case token.LET:
			letStat := p.parseLetStatement()
			if letStat != nil {
				prog.Statements = append(prog.Statements, letStat)
			}
		case token.RETURN:
			retStat := p.parseReturnStatement()
			if retStat != nil {
				prog.Statements = append(prog.Statements, retStat)
			}
		}
		p.getToken()
	}
	return &prog
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	if (p.nextToken.Type != token.IDENTIFIER) {
		msg := fmt.Sprintf("expected next token to be %s, got %s at line %v column %v",token.IDENTIFIER,p.nextToken.Type,p.nextToken.Line,p.nextToken.Column)
		p.errors = append(p.errors,msg)
		return nil
	}
	statement_token := p.curToken
	p.getToken()
	identifier := ast.Identifier{Token: p.curToken,Value: p.curToken.Literal}

	if (p.nextToken.Type != token.ASSIGN) {
		msg := fmt.Sprintf("expected next token to be %s, got %s at line %v column %v",token.ASSIGN,p.nextToken.Type,p.nextToken.Line,p.nextToken.Column)
		p.errors = append(p.errors,msg)
		return nil
	}
	p.getToken()
	
	for (p.curToken.Type != token.SEMICOLON) {
		p.getToken()
	}
	return &ast.LetStatement{Token: statement_token, VariableName: identifier}
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	retStat := ast.ReturnStatement{Token: p.curToken}
	p.getToken()

	for p.curToken.Type != token.SEMICOLON {
		p.getToken()
	}
	return &retStat
}

func (p *Parser) addPrefixFunction(toktype token.TokenType, prefixFunc prefixOperatorParserFn) {
	_,ok := p.prefixOperatorParserFns[toktype]
	if !ok {
		p.prefixOperatorParserFns[toktype] = prefixFunc
	}
}

func (p *Parser) addInfixFunction(toktype token.TokenType, infixFunc infixOperatorParserFn) {
	_,ok := p.infixOperatorParserFns[toktype]
	if !ok {
		p.infixOperatorParserFns[toktype] = infixFunc
	}
}