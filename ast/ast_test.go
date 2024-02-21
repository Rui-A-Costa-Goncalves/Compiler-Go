package ast

import (
	"goCompiler/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		[]Statement{
			&LetStatement{
				Token: token.Token{
					Type : token.LET,
					Literal : "let",
				},
				VariableName: Identifier{
					Token : token.Token{
						Type : token.IDENTIFIER,
						Literal : "var",
					},
					Value : "var",
				},
				VariableValue: &Identifier{
					Token : token.Token{
						Type : token.IDENTIFIER,
						Literal : "newVar",
					},
					Value : "newVar",
				},
			},
		},
	}
	if program.String() != "let var = newVar;" {
		t.Fatalf("String function wrong, expected \"let var = newVar;\", got %s", program.String())
	}
}