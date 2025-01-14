package components

import (
	"fmt"
	"log"
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
	return fmt.Sprintf("Parse error at %+v: %s", e.token, e.message)
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

// func (p *Parser) Parse() (semantics.Expression, error) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			if parseError, ok := err.(*parseError); ok {
// 				panic(parseError) // already handled
// 			}
// 			panic(err)
// 		}
// 	}()

// 	expression := p.expression()
// 	return expression, nil
// }

func (p *Parser) Parse() ([]semantics.Statement, error) {
	statements := []semantics.Statement{}
	defer func() {
		if err := recover(); err != nil {
			if p, ok := err.(*parseError); ok {
				log.Println(p.Error())
				return
				// panic(p.Error())
				// return // already handled
			}
			panic(err.(error).Error())
		}
	}()
	// for statements
	// for !p.isAtEnd() {
	// 	statements = append(statements, p.statement())
	// }

	for !p.isAtEnd() {
		declaration, err := p.declaration()
		if err != nil {
			panic(err.Error())
		}
		statements = append(statements, declaration)
	}

	// fmt.Print(fmt.Sprintf("\nstatements in Parse : %+v", statements))

	return statements, nil
}

// variable declaration
func (p *Parser) declaration() (semantics.Statement, error) {

	defer func() error {
		if err := recover(); err != nil {
			if _, ok := err.(*parseError); ok {
				p.synchronise()
				// log.Println(p.Error())
				return nil
				// panic(p.Error())
				// return // already handled
			}
			panic(err.(error).Error())
		}
		return nil
	}()

	if p.match(semantics.VAR) {
		return p.varDeclaration(), nil
	}

	return p.statement(), nil
}

func (p *Parser) varDeclaration() semantics.Statement {
	name := p.consume(semantics.IDENTIFIER, "Expect variable name.")

	var initialiser semantics.Expression
	if p.match(semantics.EQUAL) {
		initialiser = p.expression()
	}

	p.consume(semantics.SEMICOLON, "Expect ';' after variable declaration")
	return semantics.InitVariableDeclaration(name, initialiser)
}

func (p *Parser) statement() semantics.Statement {
	if p.match(semantics.PRINT) {
		// log.Println("\nInside PRINT STATEMENT")
		return p.printStatement()
	}
	if p.match(semantics.IF) {
		return p.ifStatement()
	}
	if p.match(semantics.LEFT_BRACE) {
		// log.Println("\nInside BLOCK STATEMENT")
		blockStatements := p.block()
		return semantics.InitBlockStatement(blockStatements)
	}

	return p.expressionStatement()
}

func (p *Parser) ifStatement() semantics.Statement {
	var elseBranch semantics.Statement
	p.consume(semantics.LEFT_PAREN, "Expect '{' after 'if'.")
	condition := p.expression()
	p.consume(semantics.RIGHT_PAREN, "Expect '}' after 'if'.")

	thenBranch := p.statement()
	if p.match(semantics.ELSE) {
		elseBranch = p.statement()
	}

	return semantics.InitIFStatement(condition, thenBranch, elseBranch)
}

func (p *Parser) block() []semantics.Statement {
	statements := []semantics.Statement{}

	for !p.check(semantics.RIGHT_BRACE) && !p.isAtEnd() {
		declaration, err := p.declaration()
		if err != nil {
			panic("something happened when reading block statement")
		}
		statements = append(statements, declaration)
	}
	p.consume(semantics.RIGHT_BRACE, "Expect '}' after block.")
	return statements
}

func (p *Parser) printStatement() semantics.Statement {
	expr := p.expression()
	// log.Printf("inside PRINT RULE [%v]", fmt.Sprint(expr))
	// fmt.Printf(fmt.Sprintf("\ninside PRINT RULE [%+v]\n", expr))
	p.consume(semantics.SEMICOLON, "Expect ';' after value")
	return semantics.InitPrintStatement(expr)
}

func (p *Parser) expressionStatement() semantics.Statement {
	expr := p.expression()
	// fmt.Printf(fmt.Sprintf("\ninside EXPR RULE [%+v]\n", expr))
	p.consume(semantics.SEMICOLON, "Expect ';' after expression")
	return semantics.InitExpressionStatement(expr)
}

func (p *Parser) expression() semantics.Expression {
	return p.assignment()
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

	if p.match(semantics.NUMBER, semantics.STRING) {
		return semantics.InitLiteral(p.previous().Literal)
	}

	if p.match(semantics.IDENTIFIER) {
		return semantics.InitVariable(p.previous())
	}

	if p.match(semantics.LEFT_PAREN) {
		expr := p.expression()
		p.consume(semantics.RIGHT_PAREN, "Expect ')' after expression")
		return semantics.InitGrouping(expr)
	}
	// panic(p.peek(), )
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

func (p *Parser) synchronise() {
	p.advance()
	for !p.isAtEnd() {
		if p.previous().TokenType == semantics.SEMICOLON {
			return
		}

		switch p.peek().TokenType {
		case semantics.CLASS:
		case semantics.FUN:
		case semantics.VAR:
		case semantics.FOR:
		case semantics.IF:
		case semantics.WHILE:
		case semantics.PRINT:
		case semantics.RETURN:
			return

		}
		p.advance()
	}
}

func (p *Parser) assignment() semantics.Expression {
	expr := p.equality()

	if p.match(semantics.EQUAL) {
		equals := p.previous()
		value := p.assignment()

		if variable, ok := expr.(*semantics.Variable); ok {
			// name := semantics.Variable(expr.(*semantics.Variable))
			return &semantics.Assignment{Name: variable.Name, Value: value}
		}
		p.error(equals, "Invalid assignment type")
	}

	return expr
}

// NOTES
// Understanding variable Declaration as a statement and
// also the assignment and accessing of variable as an expression
// program        → declaration* EOF ;

// declaration    → varDecl
//                | statement ;

// statement      → exprStmt
//                | printStmt ;

// Expression for variable declaration
// primary        → "true" | "false" | "nil"
//                | NUMBER | STRING
//                | "(" expression ")"
//                | IDENTIFIER ;

// exprerssion tree turns into the below after variable assignment is added
// expression     → assignment ;
// assignment     → IDENTIFIER "=" assignment
//                | equality ;
