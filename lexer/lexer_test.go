package lexer

import (
	"goCompiler/token"
	"testing"
)

func TestNextTokenSingleChar(t *testing.T) {
	inputToTest := "+={}(),;-/*!<>"

	expectedResult := []struct {
		expectedType token.TokenType
		expectedLiteral string
	}{
		{token.PLUS,"+"},
		{token.ASSIGN,"="},
		{token.LEFT_BRACE,"{"},
		{token.RIGHT_BRACE,"}"},
		{token.LEFT_PARENTHESIS,"("},
		{token.RIGHT_PARENTHESIS,")"},
		{token.COMMA,","},
		{token.SEMICOLON,";"},
		{token.MINUS,"-"},
		{token.DIVIDE,"/"},
		{token.MULTIPLY,"*"},
		{token.NOT,"!"},
		{token.SMALLER,"<"},
		{token.BIGGER,">"},
	}

	lexer := NewLexer(inputToTest)

	for i,expectedTok := range(expectedResult) {
		detectedToken := lexer.NextToken()

		if detectedToken.Type != expectedTok.expectedType {
			t.Fatalf("test %q got the wrong tokenType, expected %q got %q",i+1,expectedTok.expectedType,detectedToken.Type)
		}

		if detectedToken.Literal != expectedTok.expectedLiteral {
			t.Fatalf("test %q literal got the wrong literal, expected %q got %q",i+1,expectedTok.expectedLiteral,detectedToken.Literal)
		}
	}
}

func TestNextTokenInitializer(t *testing.T) {
	inputToTest := `let var = 4;
		let el = 11;
		let add = fn(x,y) {
			x + y;
		};
		let result = add(var,el);

		if !(5*3 < 12) {
			return true;
			} else {
			return false;
			}
		10 == 10;
		10 != 9;
	
	`

	expectedResult := []struct {
		expectedType token.TokenType
		expectedLiteral string
	}{
		{token.LET,"let"},
		{token.IDENTIFIER,"var"},
		{token.ASSIGN,"="},
		{token.INT,"4"},
		{token.SEMICOLON,";"},
		{token.LET,"let"},
		{token.IDENTIFIER,"el"},
		{token.ASSIGN,"="},
		{token.INT,"11"},
		{token.SEMICOLON,";"},
		{token.LET,"let"},
		{token.IDENTIFIER,"add"},
		{token.ASSIGN,"="},
		{token.FUNCTION,"fn"},
		{token.LEFT_PARENTHESIS, "("},
		{token.IDENTIFIER, "x"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "y"},
		{token.RIGHT_PARENTHESIS, ")"},
		{token.LEFT_BRACE, "{"},
		{token.IDENTIFIER, "x"},
		{token.PLUS, "+"},
		{token.IDENTIFIER, "y"},
		{token.SEMICOLON,";"},
		{token.RIGHT_BRACE,"}"},
		{token.SEMICOLON,";"},
		{token.LET,"let"},
		{token.IDENTIFIER,"result"},
		{token.ASSIGN,"="},
		{token.IDENTIFIER,"add"},
		{token.LEFT_PARENTHESIS, "("},
		{token.IDENTIFIER, "var"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "el"},
		{token.RIGHT_PARENTHESIS, ")"},
		{token.SEMICOLON,";"},
		{token.IF,"if"},
		{token.NOT,"!"},
		{token.LEFT_PARENTHESIS, "("},
		{token.INT,"5"},
		{token.MULTIPLY,"*"},
		{token.INT,"3"},
		{token.SMALLER,"<"},
		{token.INT,"12"},
		{token.RIGHT_PARENTHESIS, ")"},
		{token.LEFT_BRACE,"{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON,";"},
		{token.RIGHT_BRACE,"}"},
		{token.ELSE,"else"},
		{token.LEFT_BRACE,"{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON,";"},
		{token.RIGHT_BRACE,"}"},
		{token.INT,"10"},
		{token.EQUAL,"=="},
		{token.INT,"10"},
		{token.SEMICOLON,";"},
		{token.INT,"10"},
		{token.NOT_EQUAL,"!="},
		{token.INT,"9"},
		{token.SEMICOLON,";"},
		
	}

	lexer := NewLexer(inputToTest)

	for i,expectedTok := range(expectedResult) {
		detectedToken := lexer.NextToken()

		if detectedToken.Type != expectedTok.expectedType {
			t.Fatalf("test %q got the wrong tokenType, expected %q got %q",i+1,expectedTok.expectedType,detectedToken.Type)
		}

		if detectedToken.Literal != expectedTok.expectedLiteral {
			t.Fatalf("test %q literal got the wrong literal, expected %q got %q",i+1,expectedTok.expectedLiteral,detectedToken.Literal)
		}
	}
}

func TestNextTokenLineDetection(t *testing.T) {
	inputToTest := `let var = 4;
		let el = 11;
		let add = fn(x,y) {
			x + y;
		};`
		expectedResult := []struct {
			expectedType token.TokenType
			expectedLiteral string
			expectedLine int
		}{
			{token.LET,"let",1},
			{token.IDENTIFIER,"var",1},
			{token.ASSIGN,"=",1},
			{token.INT,"4",1},
			{token.SEMICOLON,";",1},
			{token.LET,"let",2},
			{token.IDENTIFIER,"el",2},
			{token.ASSIGN,"=",2},
			{token.INT,"11",2},
			{token.SEMICOLON,";",2},
			{token.LET,"let",3},
			{token.IDENTIFIER,"add",3},
			{token.ASSIGN,"=",3},
			{token.FUNCTION,"fn",3},
			{token.LEFT_PARENTHESIS, "(",3},
			{token.IDENTIFIER, "x",3},
			{token.COMMA, ",",3},
			{token.IDENTIFIER, "y",3},
			{token.RIGHT_PARENTHESIS, ")",3},
			{token.LEFT_BRACE, "{",3},
			{token.IDENTIFIER, "x",4},
			{token.PLUS, "+",4},
			{token.IDENTIFIER, "y",4},
			{token.SEMICOLON,";",4},
			{token.RIGHT_BRACE,"}",5},
			{token.SEMICOLON,";",5},
		}
	lexer := NewLexer(inputToTest)

	for i,expectedTok := range(expectedResult) {
		detectedToken := lexer.NextToken()

		if detectedToken.Type != expectedTok.expectedType {
			t.Fatalf("test %q got the wrong tokenType, expected %q got %q",i+1,expectedTok.expectedType,detectedToken.Type)
		}

		if detectedToken.Literal != expectedTok.expectedLiteral {
			t.Fatalf("test %q literal got the wrong literal, expected %q got %q",i+1,expectedTok.expectedLiteral,detectedToken.Literal)
		}
		if detectedToken.Line != expectedTok.expectedLine {
			t.Fatalf("test %q line got the wrong line, expected %q got %q",i+1,expectedTok.expectedLine,detectedToken.Line)
		}
	}
}