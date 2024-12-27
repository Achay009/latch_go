package semantics

type StatementVisitor interface {
	visitExpressionStatement(expression *ExpressionStatement) interface{}

	visitPrintStatement(expression *Print) interface{}
}

type Statement interface {
	Accept(visitor StatementVisitor) interface{}
}

type Print struct {
	Expr Expression
}

func (p *Print) Accept(visitor StatementVisitor) interface{} {
	return visitor.visitPrintStatement(p)
}

func InitPrintStatement(expr Expression) *Print {
	return &Print{
		Expr: expr,
	}
}

type ExpressionStatement struct {
	Expr Expression
}

func (e *ExpressionStatement) Accept(visitor StatementVisitor) interface{} {
	return visitor.visitExpressionStatement(e)
}

func InitExpressionStatement(expr Expression) *ExpressionStatement {
	return &ExpressionStatement{
		Expr: expr,
	}
}

//Definittion
// program        → statement* EOF ;

// statement      → exprStmt
//                | printStmt ;

// exprStmt       → expression ";" ;
// printStmt      → "print" expression ";" ;
