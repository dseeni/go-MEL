package lexer

import (
	"github.com/nrtkbb/go-MEL/token"
)

type Lexer struct {
	input        []rune
	position     int
	readPosition int
	rune         rune
}

func New(input string) *Lexer {
	l := &Lexer{input: []rune(input)}
	l.readRune()
	return l
}

func (l *Lexer) readRune() {
	if l.readPosition >= len(l.input) {
		l.rune = 0
	} else {
		l.rune = rune(l.input[l.readPosition])
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekRune() rune {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func newToken(tokenType token.TokenType, r rune) token.Token {
	return token.Token{Type: tokenType, Literal: string(r)}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.rune {
	case '=':
		tok = l.peekRuneCheck('=', token.EQ, token.ASSIGN)
	case '!':
		tok = l.peekRuneCheck('=', token.NOT_EQ, token.BANG)
	case '<':
		tok = l.peekRuneCheck('<', token.LTENSOR, token.LT)
	case '>':
		tok = l.peekRuneCheck('>', token.RTENSOR, token.GT)
	case ';':
		tok = newToken(token.SEMICOLON, l.rune)
	case '(':
		tok = newToken(token.LPAREN, l.rune)
	case ')':
		tok = newToken(token.RPAREN, l.rune)
	case '{':
		tok = newToken(token.LBRACE, l.rune)
	case '}':
		tok = newToken(token.RBRACE, l.rune)
	case '[':
		tok = newToken(token.LBRACKET, l.rune)
	case ']':
		tok = newToken(token.RBRACKET, l.rune)
	case ',':
		tok = newToken(token.COMMA, l.rune)
	case '+':
		tok = newToken(token.PLUS, l.rune)
	case '-':
		tok = newToken(token.MINUS, l.rune)
	case '/':
		tok = newToken(token.SLASH, l.rune)
	case '*':
		tok = newToken(token.ASTERISK, l.rune)
	case '$':
		tok.Literal = l.readIdentifier()
		tok.Type = token.IDENT
		return tok
	case '"':
		tok.Literal = l.readString()
		tok.Type = token.STRING_DATA
		return tok
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.rune) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.rune) || '.' == l.rune && isDigit(l.peekRune()) {
			if '0' == l.rune && 'x' == l.peekRune() {
				tok.Type = token.INT_16DATA
				tok.Literal = l.readHexadecimalNumber()
			} else {
				tok.Type, tok.Literal = l.readNumber()
			}
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.rune)
		}
	}

	l.readRune()
	return tok
}

func (l *Lexer) peekRuneCheck(peek rune, trueType, falseType token.TokenType) (token.Token) {
	if peek == l.peekRune() {
		ch := l.rune
		l.readRune()
		literal := string(ch) + string(l.rune)
		return token.Token{Type: trueType, Literal: literal}
	} else {
		return newToken(falseType, l.rune)
	}
}

func (l *Lexer) readString() string {
	position := l.position
	l.readRune() // '"'
	for '"' != l.rune {
		if '\\' == l.rune {
			l.readRune() // '\\'
		}
		l.readRune()
	}
	l.readRune() // '"'
	return string(l.input[position:l.position])
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	l.readRune() // '$'
	for isIdentifier(l.rune) {
		l.readRune()
	}
	return string(l.input[position:l.position])
}

func isLetter(r rune) bool {
	return 'a' <= r && r <= 'z' || 'A' <= r && r <= 'Z' || '_' == r
}

func isIdentifier(r rune) bool {
	return isLetter(r) || '0' <= r && r <= '9'
}

func (l *Lexer) readHexadecimalNumber() string {
	position := l.position
	l.readRune() // '0'
	l.readRune() // 'x'
	for isHexadecimalDigit(l.rune) {
		l.readRune()
	}
	return string(l.input[position:l.position])
}

func isHexadecimalDigit(r rune) bool {
	return '0' <= r && r <= '9' || 'a' <= r && r <= 'f' || 'A' <= r && r <= 'F'
}

func (l *Lexer) readNumber() (token.TokenType, string) {
	var typ token.TokenType
	typ = token.INT_DATA
	position := l.position
	for isDigit(l.rune) {
		l.readRune()
	}
	if '.' == l.rune {
		typ = token.FLOAT_DATA
		l.readRune()
		for isDigit(l.rune) {
			l.readRune()
		}
	}
	if 'e' == l.rune || 'E' == l.rune {
		if '-' == l.peekRune() || '+' == l.peekRune() {
			l.readRune() // 'e' or 'E'
			l.readRune() // '-' or '+'
			for isDigit(l.rune) {
				l.readRune()
			}
		}
	}
	return typ, string(l.input[position:l.position])
}

func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

func (l *Lexer) skipWhitespace() {
	for ' ' == l.rune || '\t' == l.rune || '\n' == l.rune || '\r' == l.rune {
		l.readRune()
	}
}
