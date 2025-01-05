package semantics

import (
	"fmt"
	"strings"
)

type AbstractSyntaxTreePrinter struct {
}

func InitAbstractSyntaxTreePrinter() *AbstractSyntaxTreePrinter {
	return &AbstractSyntaxTreePrinter{}
}

func (a *AbstractSyntaxTreePrinter) Print(expression Expression) string {
	return expression.Accept(a).(string)
}

func (a *AbstractSyntaxTreePrinter) visitAssignmentExpression(assgn *Assignment) interface{} {
	return nil
}

func (a *AbstractSyntaxTreePrinter) visitVariableDeclarationExpression(variable *Variable) interface{} {
	return nil
}

func (a *AbstractSyntaxTreePrinter) visitBinaryExpression(binaryExpression *Binary) interface{} {
	return a.parenthesize(binaryExpression.operator.Lexeme, binaryExpression.left, binaryExpression.right)
}

func (a *AbstractSyntaxTreePrinter) visitGroupingExpression(groupExpression *Grouping) interface{} {
	return a.parenthesize("group", groupExpression.expression)
}

func (a *AbstractSyntaxTreePrinter) visitLiteralExpression(literalExpression *Literal) interface{} {
	if literalExpression.value == nil {
		return "nil"
	}

	return literalExpression.value
}

func (a *AbstractSyntaxTreePrinter) visitUnaryExpression(unaryExpression *Unary) interface{} {
	return a.parenthesize(unaryExpression.operator.Lexeme, unaryExpression.right)
}

func (a *AbstractSyntaxTreePrinter) parenthesize(lexeme string, expressionArgs ...Expression) string {
	builder := &strings.Builder{}
	builder.WriteString("(")
	builder.WriteString(lexeme)

	for _, expresssion := range expressionArgs {
		builder.WriteString(" ")
		builder.WriteString(fmt.Sprint(expresssion.Accept(a)))
	}

	builder.WriteString(")")

	return builder.String()
}
