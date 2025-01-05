package semantics

type Visitor interface {
	visitBinaryExpression(b *Binary) interface{}
	visitGroupingExpression(g *Grouping) interface{}
	visitLiteralExpression(l *Literal) interface{}
	visitUnaryExpression(u *Unary) interface{}
	visitVariableDeclarationExpression(v *Variable) interface{}
	visitAssignmentExpression(a *Assignment) interface{}
}

type Expression interface {
	Accept(visitor Visitor) interface{}
}

type Binary struct {
	left     Expression
	operator Token
	right    Expression
}

func InitBinary(left Expression, operator Token, right Expression) *Binary {
	return &Binary{
		left:     left,
		operator: operator,
		right:    right,
	}
}

func (b *Binary) Accept(visitor Visitor) interface{} {
	return visitor.visitBinaryExpression(b)
}

type Grouping struct {
	expression Expression
}

func (g *Grouping) Accept(visitor Visitor) interface{} {
	return visitor.visitGroupingExpression(g)
}

func InitGrouping(expression Expression) *Grouping {
	return &Grouping{
		expression: expression,
	}
}

type Literal struct {
	value interface{}
}

func (l *Literal) Accept(visitor Visitor) interface{} {
	return visitor.visitLiteralExpression(l)
}

func InitLiteral(value interface{}) *Literal {
	return &Literal{
		value: value,
	}
}

type Unary struct {
	operator Token
	right    Expression
}

func (u *Unary) Accept(visitor Visitor) interface{} {
	return visitor.visitUnaryExpression(u)
}

func InitUnary(operator Token, right Expression) *Unary {
	return &Unary{
		operator: operator,
		right:    right,
	}
}

// variable
type Variable struct {
	Name Token
}

func (v *Variable) Accept(visitor Visitor) interface{} {
	return visitor.visitVariableDeclarationExpression(v)
}

func InitVariable(token Token) *Variable {
	return &Variable{
		Name: token,
	}
}

type Assignment struct {
	Name  Token
	Value Expression
}

func (a *Assignment) Accept(visitor Visitor) interface{} {
	return visitor.visitAssignmentExpression(a)
	// return nil
}

func InitAssignment(name Token, value Expression) *Assignment {
	return &Assignment{
		Name:  name,
		Value: value,
	}
}

// A major difference between Expression and statements is that a statement does not determine
// the value of  and entity in programming languages but an expression does more so an expression
// in a value of some sort

// Statement produce what is called a SIDE EFFECT , which is the change in nature of a
// particular entity
