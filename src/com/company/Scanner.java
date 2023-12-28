package com.company;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import com.company.Token.*;
import static com.company.Token.TokenType.*;

public class Scanner {

    private final String source;
    private final List<Token> tokens = new ArrayList<>();
    private int start = 0;
    private int current = 0;
    private int line = 1;

    private static final Map<String, TokenType> reservedKeywordsMap;

    static {
        reservedKeywordsMap = new HashMap<>();
        reservedKeywordsMap.put("and", AND);
        reservedKeywordsMap.put("class",  CLASS);
        reservedKeywordsMap.put("else",   ELSE);
        reservedKeywordsMap.put("false",  FALSE);
        reservedKeywordsMap.put("for",    FOR);
        reservedKeywordsMap.put("fun",    FUN);
        reservedKeywordsMap.put("if",     IF);
        reservedKeywordsMap.put("nil",    NIL);
        reservedKeywordsMap.put("or",     OR);
        reservedKeywordsMap.put("print",  PRINT);
        reservedKeywordsMap.put("return", RETURN);
        reservedKeywordsMap.put("super",  SUPER);
        reservedKeywordsMap.put("this",   THIS);
        reservedKeywordsMap.put("true",   TRUE);
        reservedKeywordsMap.put("var",    VAR);
        reservedKeywordsMap.put("while",  WHILE);
    }


    public Scanner(String source) {
        this.source = source;
    }

    public List<Token> scanTokens() {

        while (!isAtEnd()){
            //Have we reached the end of a line we are scanning ??
            start = current;
            scanToken();
        }

        tokens.add(new Token(EOF, "", null, line));
        return tokens;
    }

    private void scanToken() {
        char c = advance();
        switch (c){
            case '(': addToken(LEFT_PAREN); break;
            case ')': addToken(RIGHT_PAREN); break;
            case '{': addToken(LEFT_BRACE); break;
            case '}': addToken(RIGHT_BRACE); break;
            case ',': addToken(COMMA); break;
            case '.': addToken(DOT); break;
            case '-': addToken(MINUS); break;
            case '+': addToken(PLUS); break;
            case ';': addToken(SEMICOLON); break;
            case '*': addToken(STAR); break;
            case '!':
                addToken(match('=') ? BANG_EQUAL : BANG);
                break;
            case '=':
                addToken(match('=') ? EQUAL_EQUAL : EQUAL);
                break;
            case '<' :
                addToken(match('=') ? LESS_EQUAL : LESS);
                break;
            case '>' :
                addToken(match('=') ? GREATER_EQUAL : GREATER);
                break;
            case '/':
                if (match('/')){
                    while (peek() != '\n' && !isAtEnd()) advance();
                }else{
                    addToken(SLASH);
                }
                break;
            case ' ':
            case '\r':
            case '\t':
                break;
            case '\n':
                line++;
                break;
            case '"':
                string();
                break;
            default:
                if(isDigit(c)){
                    number();
                }else if (isAlpha(c)) {
                    // Identify Keywords and reserved words like "var", "nil", "Boolean"
                    identifier();
                }else {
                    Latch.error(line, "Unexpected error");
                }
                break;
        }
    }

    private void identifier() {
        while (isAlphaNumeric(peek())) advance();

        String text = source.substring(start, current);
        TokenType tokenType = reservedKeywordsMap.get(text);

        if (tokenType == null) tokenType = IDENTIFIER;

        addToken(tokenType);
    }

    private boolean isAlphaNumeric(char peek) {
        return isAlpha(peek) || isDigit(peek);
    }

    private boolean isAlpha(char c) {
        return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_';
    }

    // NUMBER LITERALS
    private void number() {
        while (isDigit(peek())) advance();

        //look for fractional number(part) !!
        if(peek() == '.' && isDigit(peekNext())){
            //consume the '.'
            advance();

            // to get the remaining decimal number
            while(isDigit(peek())) advance();
        }

        addToken(NUMBER, Double.parseDouble(source.substring(start, current)));
    }

    // Usually check the next token
    private char peekNext() {
        if (current + 1 >= source.length()) return '\0';
        return source.charAt(current + 1);
    }

    private boolean isDigit(char c) {

        return c >= '0'  && c <=9;
    }


    //STRING LITERALS
    private void string() {
        while (peek() != '"' && !isAtEnd()){
            if (peek() == '\n') line++;
            advance();
        }

        if (isAtEnd()){
            Latch.error(line, "Unterminated string");
            return;
        }



        // The closing "
        advance();

        String value = source.substring(start + 1, current - 1);
        addToken(STRING, value);
    }

    private char peek() {
        if (isAtEnd()) return '\0';
        return source.charAt(current);
    }

    private boolean match(char expected) {
        if (isAtEnd()) return false;
        if (source.charAt(current) != expected) return false;

        current++;
        return true;
    }

    private void addToken(Token.TokenType type) {
        addToken(type, null);
    }

    private void addToken(Token.TokenType type, Object literal){
        String text = source.substring(start, current);
        tokens.add(new Token(type, text, literal, line));
    }

    private char advance() {
        return source.charAt(current++);
    }

    private boolean isAtEnd() {
        return current >= source.length();
    }
}
