package semantics

type Visitor interface {
	visitBinaryExpression(b *Binary) Binary
	visitGroupingExpression(g *Grouping) Grouping
	visitLiteralExpression(l *Literal) Literal
	visitUnaryExpression(u *Unary) Unary
}

type Expression interface {
	accept(visitor Visitor) interface{}
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

func (b *Binary) accept(visitor Visitor) Binary {
	return visitor.visitBinaryExpression(b)
}

type Grouping struct {
	expression Expression
}

func (g *Grouping) accept(visitor Visitor) Grouping {
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

func (l *Literal) accept(visitor Visitor) Literal {
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

func (u *Unary) accept(visitor Visitor) Unary {
	return visitor.visitUnaryExpression(u)
}

func InitUnary(operator Token, right Expression) *Unary {
	return &Unary{
		operator: operator,
		right:    right,
	}
}
