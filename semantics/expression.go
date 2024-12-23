package semantics

type Visitor interface {
	visitBinaryExpression(b *Binary) interface{}
	visitGroupingExpression(g *Grouping) interface{}
	visitLiteralExpression(l *Literal) interface{}
	visitUnaryExpression(u *Unary) interface{}
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
