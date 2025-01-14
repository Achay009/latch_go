package semantics

type StatementVisitor interface {
	visitExpressionStatement(statement *ExpressionStatement) interface{}

	visitPrintStatement(statement *Print) interface{}

	visitVariableDeclarationStatement(statement *Var) interface{}

	visitBlockStatement(block *Block) interface{}

	visitIFStatement(conditional *If) interface{}
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

// var declaration statement
type Var struct {
	Name        Token
	Initialiser Expression
}

func (v *Var) Accept(visitor StatementVisitor) interface{} {
	return visitor.visitVariableDeclarationStatement(v)
}

func InitVariableDeclaration(token Token, expr Expression) *Var {
	return &Var{
		Initialiser: expr,
		Name:        token,
	}
}

//Definittion
// program        → statement* EOF ;

// statement      → exprStmt
//                | printStmt ;

// exprStmt       → expression ";" ;
// printStmt      → "print" expression ";" ;

// Defnining Block statement
// statement      → exprStmt
//                | printStmt
//                | block ;

// block          → "{" declaration* "}" ;

type Block struct {
	Statements []Statement
}

func (v *Block) Accept(visitor StatementVisitor) interface{} {
	return visitor.visitBlockStatement(v)
}

func InitBlockStatement(states []Statement) *Block {
	return &Block{
		Statements: states,
	}
}

type If struct {
	Condition  Expression
	ThenBranch Statement
	ElseBranch Statement
}

func (i *If) Accept(visitor StatementVisitor) interface{} {
	return visitor.visitIFStatement(i)
}

func InitIFStatement(condition Expression, thenBranch Statement, elseBranch Statement) *If {
	return &If{
		Condition:  condition,
		ThenBranch: thenBranch,
		ElseBranch: elseBranch,
	}
}
