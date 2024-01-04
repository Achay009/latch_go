package com.company;

class AbstractSyntaxTreePrinter implements Expression.Visitor<String> {

    String print(Expression expr){
        return expr.accept(this);
    }

    //Check that the Printer is returning expression and its tree connection
    public static void main(String[] args){
        Expression expression = new Expression.Binary(
                new Expression.Unary(new Token(Token.TokenType.BANG_EQUAL, "-", null, 1),
                        new Expression.Literal(123)
                ),
                new Token(Token.TokenType.STAR, "*", null, 1),
                new Expression.Grouping(new Expression.Literal(45.97))

        );

        System.out.println(new AbstractSyntaxTreePrinter().print(expression));
    }

    @Override
    public String visitBinaryExpression(Expression.Binary expression) {
        return parenthesize(expression.operator.lexeme, expression.left, expression.right);
    }

    @Override
    public String visitGroupingExpression(Expression.Grouping expression) {
        return parenthesize("group", expression.expression);
    }

    @Override
    public String visitLiteralExpression(Expression.Literal expression) {
        if (expression.value == null) return "nil";
        return expression.value.toString();
    }

    @Override
    public String visitUnaryExpression(Expression.Unary expression) {
        return parenthesize(expression.operator.lexeme, expression.right);
    }

    private String parenthesize(String lexeme, Expression... expressionsArgs) {
        StringBuilder builder = new StringBuilder();

        builder.append("(").append(lexeme);

        for(Expression expression : expressionsArgs){
            builder.append(" ");
            builder.append(expression.accept(this));
        }

        builder.append(")");

        return builder.toString();
    }
}
