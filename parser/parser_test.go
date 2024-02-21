package parser

import (
	"goCompiler/ast"
	"goCompiler/lexer"
	"testing"
)

func TestBasicLetStatements (t *testing.T) {
	input := `
		let x = 4;
		let y = 20;
		let zz = 4214;
	`

	l := lexer.NewLexer(input)
	parser := NewParser(l)


	resProgram := parser.parseProgram()
	testParseErrors(t,parser)

	if resProgram == nil {
		t.Fatalf("ParseProgram() return nil")
	}
	if len(resProgram.Statements) != 3 {
		t.Fatalf("Incorrect number of statements, expected 3, got %v", len(resProgram.Statements))
	}
	expectedProgram := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"zz"},
	}

	for i, expected := range(expectedProgram) {
		if !testLetStatement(t, resProgram.Statements[i],expected.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T,s ast.Statement,identifier string) bool {
	if (s.TokenLiteral() != "let") {
		t.Fatalf("Not the correct token literal, expected let got %q",s.TokenLiteral())
		return false
	}

	letS, ok := s.(*ast.LetStatement)
	if !ok {
		t.Fatalf("Not the correct statement type, expected let got %T",s)
		return false
	}

	if letS.VariableName.Value != identifier {
		t.Fatalf("Not the correct identifier, expected %s got %s",identifier, letS.VariableName.Value)
		return false
	}

	if letS.VariableName.TokenLiteral() != identifier {
		t.Fatalf("Not the correct token literal, expected %s got %s",identifier, letS.VariableName.TokenLiteral())
		return false
	}
	return true
}

func testParseErrors(t *testing.T, parser *Parser) {
	errors := parser.errors
	if len(errors) == 0 {
		return
	}
	t.Errorf("%v errors in parser",len(errors))
	for _,error := range(errors) {
		t.Errorf(error)
	}
	t.FailNow()
}

func TestBasicReturnStatements (t *testing.T) {
	input := `
		return 4;
		return 20;
		return 4214;
	`

	l := lexer.NewLexer(input)
	parser := NewParser(l)


	resProgram := parser.parseProgram()
	testParseErrors(t,parser)

	if resProgram == nil {
		t.Fatalf("ParseProgram() return nil")
	}
	if len(resProgram.Statements) != 3 {
		t.Fatalf("Incorrect number of statements, expected 3, got %v", len(resProgram.Statements))
	}
	for _,s := range(resProgram.Statements) {
		retS, ok := s.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("Incorrect statement, expected return statement got %T",s)
		}
		if retS.TokenLiteral() != "return" {
			t.Errorf("Incorrect literal, expected return got %s",retS.TokenLiteral())
		}
	}
}