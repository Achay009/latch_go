package semantics

import "fmt"

type TokenType int

const (
	EOF TokenType = iota //END OF FILE

	//SINGLE CHARACTER TOKEN
	LEFT_PAREN
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	PLUS
	MINUS
	SEMICOLON
	SLASH
	STAR

	//ONE OR TWO CHARACTER TOKENS
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL
	IDENTIFIER
	STRING
	NUMBER

	//KEYWORDS
	AND
	CLASS
	FALSE
	FUN
	FOR
	IF
	ELSE
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE
)

type Token struct {
	TokenType TokenType
	Lexeme    string
	Literal   interface{}
	Line      int
}

func (t *Token) toString() {
	fmt.Printf("%v %v %v", t.TokenType, t.Lexeme, t.Literal)
	// return t.tokenType + " " + t.lexeme+ " "+ t.literal
}
