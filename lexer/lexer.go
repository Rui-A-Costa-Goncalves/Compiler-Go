package lexer

import (
	"goCompiler/token"
)


type Lexer struct {
	input string
	position int
	nextPosition int
	parsingLine int
	char byte
}

func NewLexer(input string) *Lexer {
	return &Lexer{input: input, parsingLine: 1}
}

func (l* Lexer) ReadChar() {
	if (l.nextPosition >= len(l.input)) {
		l.char = 0
	} else {
		l.position = l.nextPosition
		l.char = l.input[l.position]
		l.nextPosition ++
	}
}

func (l* Lexer) ReadNextNotWhiteSpaceChar() {
	l.ReadChar()
	for l.char == ' ' || l.char == '\n' || l.char == '\t' || l.char == '\r' {
		if l.char == '\n' {
			l.parsingLine++
		}
		l.ReadChar()
	}
}

func (l* Lexer) BacktrackLexerPointer(char byte) {
	l.position --
	l.nextPosition --
	l.char = char
}

func isLetter (b byte) bool {
	if (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || b == '_' {
		return true
	}
	return false
}

func generateLetterToken (l *Lexer) token.Token {
	myChar := l.char
	word := ""
	starting_column := l.position +1
	var lastChar byte
	for isLetter(myChar) {
		word += string(myChar)
		l.ReadChar()
		lastChar = myChar
		myChar = l.char
	}
	l.BacktrackLexerPointer(lastChar)
	return token.Token {Type : token.GetLetterTokenType(word), Literal : word, Line : l.parsingLine, Column : starting_column}
}

func isNumber (b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	}
	return false
}
func generateNumberToken (l *Lexer) token.Token {
	myChar := l.char
	word := ""
	starting_column := l.position +1
	var lastChar byte
	for isNumber(myChar) {
		word += string(myChar)
		l.ReadChar()
		lastChar = myChar
		myChar = l.char
	}
	l.BacktrackLexerPointer(lastChar)
	return  token.Token {Type : token.INT, Literal : word, Line : l.parsingLine, Column : starting_column}
}

func (l * Lexer) peekNextChar() byte {
	if (l.nextPosition >= len(l.input)) {
		return 0
	} else {
		return l.input[l.nextPosition]
	}
}



func (l *Lexer) NextToken() token.Token {

	l.ReadNextNotWhiteSpaceChar()
	switch (l.char){
	case '=':
		nextChar := l.peekNextChar()
		if (nextChar == '=') {
			prevChar := l.char
			starting_column := l.position +1
			//fix the pointers
			l.ReadChar()
			newChar := l.char
			newLiteral := string(prevChar) + string(newChar)
			return token.Token {Type : token.EQUAL, Literal : newLiteral, Line : l.parsingLine, Column : starting_column}
		} else {
			return token.Token {Type : token.ASSIGN, Literal : string(l.char), Line : l.parsingLine, Column : l.position+1}
		}
		//do the equal verification
	case '!':
		nextChar := l.peekNextChar()
		if (nextChar == '=') {
			prevChar := l.char
			starting_column := l.position +1
			//fix the pointers
			l.ReadChar()
			newChar := l.char
			newLiteral := string(prevChar) + string(newChar)
			return token.Token {Type : token.NOT_EQUAL, Literal : newLiteral, Line : l.parsingLine, Column : starting_column}
		} else {
			return token.Token {Type : token.NOT, Literal : string(l.char), Line : l.parsingLine, Column : l.position+1}
		}
	case '+':
		return token.Token {Type : token.PLUS, Literal : string(l.char), Line : l.parsingLine, Column : l.position+1}
	case '-':
		return token.Token {Type : token.MINUS, Literal : string(l.char), Line : l.parsingLine, Column : l.position+1}
	case '/':
		return token.Token {Type : token.DIVIDE, Literal : string(l.char), Line : l.parsingLine, Column : l.position+1}
	case '*':
		return token.Token {Type : token.MULTIPLY, Literal : string(l.char), Line : l.parsingLine, Column : l.position+1}
	case ';':
		return token.Token {Type : token.SEMICOLON, Literal : string(l.char), Line : l.parsingLine, Column : l.position+1}
	case ',':
		return token.Token {Type : token.COMMA, Literal : string(l.char), Line : l.parsingLine, Column : l.position+1}
	case '(':
		return token.Token {Type : token.LEFT_PARENTHESIS, Literal : string(l.char), Line : l.parsingLine, Column : l.position+1}
	case ')':
		return token.Token {Type : token.RIGHT_PARENTHESIS, Literal : string(l.char), Line : l.parsingLine, Column : l.position+1}
	case '{':
		return token.Token {Type : token.LEFT_BRACE, Literal : string(l.char), Line : l.parsingLine, Column : l.position+1}
	case '}':
		return token.Token {Type : token.RIGHT_BRACE, Literal : string(l.char), Line : l.parsingLine, Column : l.position+1}
	case '<':
		return token.Token {Type : token.SMALLER, Literal : string(l.char), Line : l.parsingLine, Column : l.position+1}
	case '>':
		return token.Token {Type : token.BIGGER, Literal : string(l.char), Line : l.parsingLine, Column : l.position+1}
	case 0:
		return token.Token {Type : token.EOF, Literal :"", Line : l.parsingLine, Column : l.position+1} 
	default:
		if (isLetter(l.char)) {
			return generateLetterToken(l)
		}
		if (isNumber(l.char)) {
			return generateNumberToken(l)
		}
		return token.Token {Type : token.ILLEGAL, Literal: "", Line : l.parsingLine, Column : l.position+1} 
	}
}