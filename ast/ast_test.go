package ast

import (
	"monkey/token"
	"testing"
)

func TestString(t *testing.T) {
	program := ProgramNode{
		Statements: []StatementNode{
			&LetStatementNode{
				Token: token.Token{Type: token.LET, Literal: "let"},

				Name: IdentifierNode{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},

				Value: &IdentifierNode{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
