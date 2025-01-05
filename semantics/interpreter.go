package semantics

import (
	"fmt"
	"log"
	"strings"
)

// var env *Environment

type Interpreter struct {
	env *Environment
}

type RuntimeError struct {
	token   Token
	message string
}

func InitInterpreter() *Interpreter {
	return &Interpreter{env: InitEnvironment(nil)}
}

func (e *RuntimeError) Error() string {
	return fmt.Sprintf("Parse error at %v: %s", e.token, e.message)
}

func (p *Interpreter) error(token Token, message string) error {
	return &RuntimeError{token: token, message: message}
}

// for single expression
// func (p *Interpreter) Interprete(expr Expression) {
// 	value := p.evaluate(expr)
// 	fmt.Printf(p.stringify(value) + "\n")
// }

func (p *Interpreter) Interprete(expr []Statement) {
	log.Println("\ninside interpreter now...")
	for _, statement := range expr {
		p.execute(statement)
	}
}

func (p *Interpreter) execute(statement Statement) {
	statement.Accept(p)
}

func (p *Interpreter) stringify(objectA interface{}) string {
	if objectA == nil {
		return "nil"
	}

	if _, ok := objectA.(float64); ok {
		text := fmt.Sprintf("%v", objectA)
		if strings.HasSuffix(text, ".0") {
			return text[:len(text)-2]
		}
		return text
	}

	return fmt.Sprintf("%v", objectA)
}

func (p *Interpreter) visitExpressionStatement(exprStatement *ExpressionStatement) interface{} {
	p.evaluate(exprStatement.Expr)
	// fmt.Print(">>" + p.stringify(value) + "\n")
	return nil
}

func (p *Interpreter) visitPrintStatement(printStatement *Print) interface{} {
	value := p.evaluate(printStatement.Expr)
	fmt.Print(">> " + p.stringify(value) + "\n")
	return nil
}

func (p *Interpreter) visitBlockStatement(blockStatement *Block) interface{} {
	p.executeBlockStatement(blockStatement.Statements, InitEnvironment(p.env))
	return nil
}

func (p *Interpreter) executeBlockStatement(statements []Statement, env *Environment) {
	previous := p.env
	p.env = env
	for _, statement := range statements {
		p.execute(statement)
	}

	p.env = previous
}

func (p *Interpreter) visitVariableDeclarationStatement(varStatement *Var) interface{} {

	var value interface{}
	if varStatement.Initialiser != nil {
		value = p.evaluate(varStatement.Initialiser)
	}

	p.env.define(varStatement.Name.Lexeme, value)
	// fmt.Printf("This is the environment during declaration statement %v", p.env.values)

	return nil
}

func (p *Interpreter) visitAssignmentExpression(assignment *Assignment) interface{} {
	value := p.evaluate(assignment.Value)
	p.env.assign(assignment.Name, value)
	return value
}

func (p *Interpreter) visitVariableDeclarationExpression(varExpression *Variable) interface{} {
	// fmt.Printf("This is the environment during declaration expression %v", p.env.values)
	return p.env.get(varExpression.Name)
}

func (p *Interpreter) visitLiteralExpression(litExpr *Literal) interface{} {
	return litExpr.value
}

func (p *Interpreter) visitGroupingExpression(groupExpr *Grouping) interface{} {
	return p.evaluate(groupExpr.expression)
}

func (p *Interpreter) visitBinaryExpression(binExpr *Binary) interface{} {
	left := p.evaluate(binExpr.left)
	right := p.evaluate(binExpr.right)

	switch binExpr.operator.TokenType {
	case MINUS:
		p.checkNumberOperands(binExpr.operator, left, right)
		return float64(left.(float64)) - float64(right.(float64))
	case SLASH:
		p.checkNumberOperands(binExpr.operator, left, right)
		return float64(left.(float64)) / float64(right.(float64))
	case STAR:
		p.checkNumberOperands(binExpr.operator, left, right)
		return float64(left.(float64)) * float64(right.(float64))
	case PLUS:
		switch left := left.(type) {
		case float64:
			if right, ok := right.(float64); ok {
				return left + right
			}
		case string:
			if right, ok := right.(string); ok {
				return left + right
			}
		}
	case GREATER:
		p.checkNumberOperands(binExpr.operator, left, right)
		return float64(left.(float64)) > float64(right.(float64))
	case GREATER_EQUAL:
		p.checkNumberOperands(binExpr.operator, left, right)
		return float64(left.(float64)) >= float64(right.(float64))
	case LESS:
		p.checkNumberOperands(binExpr.operator, left, right)
		return float64(left.(float64)) < float64(right.(float64))
	case LESS_EQUAL:
		p.checkNumberOperands(binExpr.operator, left, right)
		return float64(left.(float64)) <= float64(right.(float64))
	case BANG_EQUAL:
		return !p.isEqual(left, right)
	case EQUAL_EQUAL:
		return p.isEqual(left, right)
	}
	return nil
}

func (p *Interpreter) visitUnaryExpression(unaryExpr *Unary) interface{} {
	right := p.evaluate(unaryExpr.right)

	switch unaryExpr.operator.TokenType {
	case MINUS:
		return -right.(float64)
	case BANG:
		return !p.isTruthy(right)
	}

	return nil
}

func (p *Interpreter) isEqual(objectA interface{}, objectB interface{}) bool {
	if objectA == nil && objectB == nil {
		return true
	}
	if objectA == nil {
		return false
	}

	return objectA == objectB
}

func (p *Interpreter) checkNumberOperand(operator Token, operand interface{}) {
	if _, ok := operand.(float64); !ok {
		panic(p.error(operator, "Operator must be a number"))
	}
}

func (p *Interpreter) checkNumberOperands(operator Token, operandA interface{}, operandB interface{}) {
	if _, ok1 := operandA.(float64); !ok1 {
		panic(p.error(operator, "Operator must be a numbers"))
	}

	if _, ok2 := operandB.(float64); !ok2 {
		panic(p.error(operator, "Operator must be a numbers"))
	}
}

func (p *Interpreter) isTruthy(object interface{}) bool {
	if object == nil {
		return false
	}
	if value, ok := object.(bool); ok {
		return bool(value)
	}
	return true
}

func (p *Interpreter) evaluate(expr Expression) interface{} {
	return expr.Accept(p)
}
