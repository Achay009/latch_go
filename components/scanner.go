package components

import (
	"scoop/semantics"
	"strconv"
)

type Scanner struct {
	source             string
	tokens             []semantics.Token
	start              int
	current            int
	line               int
	reservedKeyWordMap map[string]semantics.TokenType
}

func InitScanner(source string) *Scanner {
	keywords := map[string]semantics.TokenType{
		"and":    semantics.AND,
		"class":  semantics.CLASS,
		"else":   semantics.ELSE,
		"false":  semantics.FALSE,
		"for":    semantics.FOR,
		"fun":    semantics.FUN,
		"if":     semantics.IF,
		"nil":    semantics.NIL,
		"or":     semantics.OR,
		"print":  semantics.PRINT,
		"return": semantics.RETURN,
		"super":  semantics.SUPER,
		"this":   semantics.THIS,
		"true":   semantics.TRUE,
		"var":    semantics.VAR,
		"while":  semantics.WHILE,
	}
	return &Scanner{
		source:             source,
		start:              0,
		current:            0,
		line:               1,
		reservedKeyWordMap: keywords,
	}
}

func (s *Scanner) ScanTokens() []semantics.Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.tokens = append(s.tokens, semantics.Token{
		TokenType: semantics.EOF,
		Lexeme:    "",
		Literal:   nil,
		Line:      s.line,
	})
	return s.tokens
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) advance() byte {
	currentChar := s.source[s.current]
	s.current++
	return currentChar
}

func (s *Scanner) addEmptyToken(tokenType semantics.TokenType) {
	s.addToken(tokenType, nil)
}

func (s *Scanner) addToken(tokenType semantics.TokenType, literal interface{}) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, semantics.Token{TokenType: tokenType, Lexeme: text, Literal: literal, Line: s.line})
}

func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return 0
	}
	return s.source[s.current]
}

func (s *Scanner) peekNext() byte {
	if s.current+1 >= len(s.source) {
		return 0
	}
	return s.source[s.current+1]
}

func (s *Scanner) scanToken() {
	character := rune(s.advance())
	switch character {
	case '(':
		s.addEmptyToken(semantics.LEFT_PAREN)
	case ')':
		s.addEmptyToken(semantics.RIGHT_PAREN)
	case '{':
		s.addEmptyToken(semantics.LEFT_BRACE)
	case '}':
		s.addEmptyToken(semantics.RIGHT_BRACE)
	case ',':
		s.addEmptyToken(semantics.COMMA)
	case '.':
		s.addEmptyToken(semantics.DOT)
	case '-':
		s.addEmptyToken(semantics.MINUS)
	case '+':
		s.addEmptyToken(semantics.PLUS)
	case ';':
		s.addEmptyToken(semantics.SEMICOLON)
	case '*':
		s.addEmptyToken(semantics.STAR)
	case '!':
		if s.match('=') {
			s.addEmptyToken(semantics.BANG_EQUAL)
		} else {
			s.addEmptyToken(semantics.BANG)
		}
	case '=':
		if s.match('=') {
			s.addEmptyToken(semantics.EQUAL_EQUAL)
		} else {
			s.addEmptyToken(semantics.EQUAL)
		}
	case '<':
		if s.match('=') {
			s.addEmptyToken(semantics.LESS_EQUAL)
		} else {
			s.addEmptyToken(semantics.LESS)
		}
	case '>':
		if s.match('=') {
			s.addEmptyToken(semantics.GREATER_EQUAL)
		} else {
			s.addEmptyToken(semantics.GREATER)
		}
	case '/':
		if s.match('/') {
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addEmptyToken(semantics.SLASH)
		}
	case ' ', '\r', '\t':
	case '\n':
		s.line++
	case '"':
		s.string()
	default:
		if s.isDigit(character) {
			s.number()
		} else if s.isAlpha(character) {
			s.identifier()
		} else {
			// suppose to throw error here from scoop
		}

	}

}

func (s *Scanner) string() {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}

	if s.isAtEnd() {
		// throw error here again ==> do this after refactor
		return
	}

	s.advance()
	value := s.source[s.start+1 : s.current-1]
	s.addToken(semantics.STRING, value)
}

func (s *Scanner) number() {
	for s.isDigit(rune(s.peek())) {
		s.advance()
	}

	if s.peek() == '.' && s.isDigit(rune(s.peekNext())) {

		s.advance()

		for s.isDigit(rune(s.peek())) {
			s.advance()
		}
	}
	number, err := strconv.ParseFloat(s.source[s.start:s.current], 64)
	if err != nil {
		// do something here
	}
	s.addToken(semantics.NUMBER, number)
}

func (s *Scanner) identifier() {
	for s.isAlphanumeric(rune(s.peek())) {
		s.advance()
	}

	text := s.source[s.start:s.current]
	tokenType, found := s.reservedKeyWordMap[text]

	if !found {
		tokenType = semantics.IDENTIFIER
	}

	s.addEmptyToken(tokenType)
}

func (s *Scanner) isDigit(char rune) bool {
	if char >= '0' && char <= '9' {
		return true
	}
	return false
}

func (s *Scanner) isAlphanumeric(peek rune) bool {
	return s.isAlpha(peek) || s.isDigit(peek)
}

func (s *Scanner) isAlpha(char rune) bool {
	if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char == '_' {
		return true
	}
	return false
}

func (s *Scanner) match(expected byte) bool {
	if s.isAtEnd() {
		return false
	}
	if s.source[s.current] != expected {
		return false
	}
	s.current++
	return true
}
