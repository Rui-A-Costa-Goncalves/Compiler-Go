package token

type TokenType string

type Token struct{
	Type TokenType
	Literal string
	Line int
	Column int
	File string
}

const (
	ILLEGAL = "ILLEGAL" //stores all the tokens for wich we dont have the type
	EOF = "EOF"

	//Initializer
	LET = "LET"

	//Identifier
	IDENTIFIER = "IDENTIFIER"

	//types
	INT = "INT"
	FUNCTION = "FUNCTION"

	//operators
	ASSIGN = "="
	PLUS = "+"
	MINUS = "-"
	DIVIDE = "/"
	MULTIPLY = "*"
	EQUAL = "=="
	NOT = "!"
	NOT_EQUAL = "!="
	SMALLER = "<"
	BIGGER = ">"

	//special characters
	SEMICOLON = ";"
	COMMA = ","
	LEFT_PARENTHESIS = "("
	RIGHT_PARENTHESIS = ")"
	LEFT_BRACE = "{"
	RIGHT_BRACE = "}"

	//boolean values
	TRUE = "true"
	FALSE = "false"
	
	//special keywords
	IF = "if"
	ELSE = "else"
	RETURN = "return"
)

var keywords = map[string]TokenType {
	"let" : LET,
	"fn": FUNCTION,
	"if" : IF,
	"else" : ELSE,
	"return" : RETURN,
	"true" : TRUE,
	"false" : FALSE,
}

func GetLetterTokenType (word string) TokenType {
	tokType, ok := keywords[word]
	if ok {
		return tokType
	}
	return IDENTIFIER
}