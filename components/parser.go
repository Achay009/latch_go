package components

import "scoop/semantics"

/**
     *
     * For the parser, in understanding how language is interpreted in terms of PRECEDENCE and ASSOCIATIVITY
     * this are the rules we are following
     *
     * expression     → equality ;
     * equality       → comparison ( ( "!=" | "==" ) comparison )* ;
     * comparison     → term ( ( ">" | ">=" | "<" | "<=" ) term )* ;
     * term           → factor ( ( "-" | "+" ) factor )* ;
     * factor         → unary ( ( "/" | "*" ) unary )* ;
     * unary          → ( "!" | "-" ) unary
     *                | primary ;
     * primary        → NUMBER | STRING | "true" | "false" | "nil"
     *                | "(" expression ")" ;
     *
     *
     *
     *  Parser is written using the RECURSIVE DECENT PARSING which uses a top-down algo because its starts
     *  with the outermost grammar rule and works its way down into the nexted eexpressions before reaching a
     *  TERMINAL which is a literal (primary rule)
     *
     *
     *  A Recursive decent parser is a literal translation of the grammar's rule in the imperative code
     *
     *  each rule becomes a function the body of the rule translates to code roughly like
     *
     *
     *
		Grammar notation	Code representation
		Terminal	   =>     Code to match and consume a token
		Nonterminal    =>    Call to that rule’s function
		|	          =>      if or switch statement
		* or +	      =>      while or for loop
		?	         =>       if statement
     *
*/

type Parser struct {
	tokens  []semantics.Token
	current int
}

func InitParser(tokens []semantics.Token) *Parser {
	return &Parser{
		tokens:  tokens,
		current: 0,
	}
}

func (p *Parser) parse() {

}

func (p *Parser) expression()
