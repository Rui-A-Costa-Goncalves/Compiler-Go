package ast

import(
	"goCompiler/token"
	"fmt"
)

type Node interface {
	TokenLiteral() string
	String() string
}

//Interfaces extend the Node interface
type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

//Make the program implement the node interface
func (p *Program) TokenLiteral() string {
	if len(p.Statements) == 0 {
		return ""
	}
	return p.Statements[0].TokenLiteral()
}

func (p *Program) String() string {
	res := ""
	for _, st := range(p.Statements) {
		res += st.String()
	}
	return res
}


// Let variableName = variableValue
type LetStatement struct {
	Token token.Token //LET TOKEN
	VariableName Identifier
	VariableValue Expression
}

func (l *LetStatement) TokenLiteral() string {
	return l.Token.Literal
}

func (l *LetStatement) String() string {
	return fmt.Sprintf("%s %s = %s;",l.TokenLiteral(),l.VariableName.String(),l.VariableValue.String())
}

//Empty function just to make the LetStatement a statementNode
func (l *LetStatement) statementNode() {}

type Identifier struct {
	Token token.Token //IDENTIFIER TOKEN
	Value string //variable name
}

func (i *Identifier) String() string {
	return i.Value
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

//Empty function to make the Identifier a expressionNode
func (i *Identifier) expressionNode() {}


type ReturnStatement struct {
	Token token.Token // RETURN TOKEN
	Value Expression
}

func (r *ReturnStatement) TokenLiteral() string{
	return r.Token.Literal
}

func (r *ReturnStatement) String() string {
	return fmt.Sprintf("%s %s;",r.TokenLiteral(),r.Value.String())
}

func (r *ReturnStatement) statementNode() {}


type ExpressionStatement struct {
	Token token.Token
	Expression Expression
}


func (e *ExpressionStatement) TokenLiteral() string{
	return e.Token.Literal
}

func (e *ExpressionStatement) String() string {
	return fmt.Sprintf("%s %s;",e.TokenLiteral(),e.Expression.String())
}

func (e *ExpressionStatement) statementNode() {}






