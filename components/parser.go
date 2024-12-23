package components

import (
	"fmt"
	"scoop/semantics"
)

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

type parseError struct {
	token   semantics.Token
	message string
}

func (e *parseError) Error() string {
	return fmt.Sprintf("Parse error at %v: %s", e.token, e.message)
}

func (p *Parser) error(token semantics.Token, message string) error {
	return &parseError{token: token, message: message}
}

func InitParser(tokens []semantics.Token) *Parser {
	return &Parser{
		tokens:  tokens,
		current: 0,
	}
}

func (p *Parser) Parse() (semantics.Expression, error) {
	defer func() {
		if err := recover(); err != nil {
			if parseError, ok := err.(*parseError); ok {
				panic(parseError) // already handled
			}
			panic(err)
		}
	}()

	expression := p.expression()
	return expression, nil
}

func (p *Parser) expression() semantics.Expression {
	return p.equality()
}

func (p *Parser) equality() semantics.Expression {
	expr := p.comparison()

	for p.match(semantics.BANG_EQUAL, semantics.EQUAL_EQUAL) {
		operator := p.previous()
		rightExpr := p.comparison()
		expr = semantics.InitBinary(expr, operator, rightExpr)
	}
	return expr
}

func (p *Parser) comparison() semantics.Expression {
	expr := p.term()

	for p.match(semantics.GREATER, semantics.GREATER_EQUAL, semantics.LESS, semantics.LESS_EQUAL) {
		operator := p.previous()
		rightExpr := p.term()
		expr = semantics.InitBinary(expr, operator, rightExpr)
	}
	return expr
}

func (p *Parser) term() semantics.Expression {
	expr := p.factor()

	for p.match(semantics.PLUS, semantics.MINUS) {
		operator := p.previous()
		rightExpr := p.factor()
		expr = semantics.InitBinary(expr, operator, rightExpr)
	}
	return expr
}

func (p *Parser) factor() semantics.Expression {
	expr := p.unary()

	for p.match(semantics.SLASH, semantics.STAR) {
		operator := p.previous()
		rightExpr := p.unary()
		expr = semantics.InitBinary(expr, operator, rightExpr)
	}
	return expr
}

func (p *Parser) unary() semantics.Expression {

	if p.match(semantics.BANG, semantics.MINUS) {
		operator := p.previous()
		rightExpr := p.unary()
		return semantics.InitUnary(operator, rightExpr)
	}
	return p.primary()
}

func (p *Parser) primary() semantics.Expression {
	if p.match(semantics.FALSE) {
		return semantics.InitLiteral(false)
	}
	if p.match(semantics.TRUE) {
		return semantics.InitLiteral(true)
	}
	if p.match(semantics.NIL) {
		return semantics.InitLiteral(nil)
	}

	if p.match(semantics.NUMBER) {
		return semantics.InitLiteral(p.previous().Literal)
	}
	if p.match(semantics.LEFT_PAREN) {
		expr := p.expression()
		p.consume(semantics.RIGHT_PAREN, "Expect ')' after expression")
		return semantics.InitGrouping(expr)
	}

	return nil
}

func (p *Parser) consume(tokenType semantics.TokenType, message string) semantics.Token {
	if p.check(tokenType) {
		return p.advance()
	}
	panic(p.error(p.peek(), message))
}

func (p *Parser) match(types ...semantics.TokenType) bool {
	for _, tokenType := range types {
		if p.check(tokenType) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) advance() semantics.Token {
	if !p.isAtEnd() {
		p.current++
	}

	return p.previous()

}

func (p *Parser) previous() semantics.Token {
	return p.tokens[p.current-1]
}

func (p *Parser) isAtEnd() bool {
	return p.peek().TokenType == semantics.EOF
}

func (p *Parser) peek() semantics.Token {
	return p.tokens[p.current]
}

func (p *Parser) check(tokenType semantics.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().TokenType == tokenType
}

/**
    * isAtEnd() checks if we’ve run out of tokens to parse.
    * peek() returns the current token we have yet to consume,
    * and previous() returns the most recently consumed token.
    * The latter makes it easier to use match() and then access the just-matched token.
* */
