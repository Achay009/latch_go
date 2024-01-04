package com.company;

import java.util.List;
import com.company.Token;

abstract class Expression {

  interface Visitor<R> {
    R visitBinaryExpression(Binary expression);
    R visitGroupingExpression(Grouping expression);
    R visitLiteralExpression(Literal expression);
    R visitUnaryExpression(Unary expression);
  }
/**
 * Representing code
 *
 * After SCANNER has sectioned our code into tokens, we have to know how to make sense of what has been
 * given, there in lies the SYNTACTIC GRAMMAR as the SCANNER has has done the LEXICAL GRAMMAR
 * for us to create SYNTACTIC GRAMMAR the below is derived below and dwe create classes and subclasses from
 * the derived formular below
 *
 *
 *
 * expression     → literal
 *                | unary
 *                | binary
 *                | grouping ;
 *
 * literal        → NUMBER | STRING | "true" | "false" | "nil" ;
 * grouping       → "(" expression ")" ;
 * unary          → ( "-" | "!" ) expression ;
 * binary         → expression operator expression ;
 * operator       → "==" | "!=" | "<" | "<=" | ">" | ">="
 *                | "+"  | "-"  | "*" | "/" ;
 *
 */
  static class Binary extends Expression {
    Binary(Expression left, Token operator, Expression right) {
      this.left = left;
      this.operator = operator;
      this.right = right;
    }

    final Expression left;
    final Token operator;
    final Expression right;

    @Override
    <R> R accept(Visitor<R> visitor) {
      return visitor.visitBinaryExpression(this);
    }
  }

  static class Grouping extends Expression {
    Grouping(Expression expression) {
      this.expression = expression;
    }

    final Expression expression;

    @Override
    <R> R accept(Visitor<R> visitor) {
      return visitor.visitGroupingExpression(this);
    }
  }

  static class Literal extends Expression {
    Literal(Object value) {
      this.value = value;
    }

    final Object value;

    @Override
    <R> R accept(Visitor<R> visitor) {
      return visitor.visitLiteralExpression(this);
    }
  }

  static class Unary extends Expression {
    Unary(Token operator, Expression right) {
      this.operator = operator;
      this.right = right;
    }

    final Token operator;
    final Expression right;

    @Override
    <R> R accept(Visitor<R> visitor) {
      return visitor.visitUnaryExpression(this);
    }
  }


  abstract <R> R accept(Visitor<R> visitor);
}
