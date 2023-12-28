package com.company;

public class Token {

    final int line;
    final Object literal;
    final String lexeme;
    final TokenType type;

    public Token(TokenType type, String lexeme, Object literal, int line) {
        this.line = line;
        this.literal = literal;
        this.lexeme = lexeme;
        this.type = type;
    }

    @Override
    public String toString() {
        return type + " " + lexeme + " " + literal;
    }

    enum TokenType {
        // Single-character tokens.
        LEFT_PAREN, RIGHT_PAREN, LEFT_BRACE, RIGHT_BRACE,
        COMMA, DOT, MINUS, PLUS, SEMICOLON, SLASH, STAR,

        // One or two character tokens.
        BANG, BANG_EQUAL,
        EQUAL, EQUAL_EQUAL,
        GREATER, GREATER_EQUAL,
        LESS, LESS_EQUAL,

        // Literals.
        IDENTIFIER, STRING, NUMBER,

        // Keywords.
        AND, CLASS, ELSE, FALSE, FUN, FOR, IF, NIL, OR,
        PRINT, RETURN, SUPER, THIS, TRUE, VAR, WHILE,

        EOF
    }




}
