package lexer

import (
	"github.com/wawoon/go-monkey/token"
)

type Lexer struct {
	input        string
	position     int  // 入力における現在の位置
	readPosition int  // 次に読む文字
	ch           byte // 現在検査中の文字
}

// readTwoCharToken handles tokens that may consist of one or two characters
// depending on what follows the current position (e.g. '=' vs '==', '!' vs '!=').
// It does not advance past the current character when the token is a single-
// character token.  The caller is responsible for calling readChar() once at
// the end of NextToken().
func (l *Lexer) readTwoCharToken(expectedNext byte, twoCharType, singleCharType token.TokenType) token.Token {
    if l.peekChar() == expectedNext {
        ch := l.ch
        l.readChar()
        literal := string([]byte{ch, l.ch})
        return token.Token{Type: twoCharType, Literal: literal}
    }
    return newToken(singleCharType, l.ch)
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

	switch l.ch {
	case '=':
        tok = l.readTwoCharToken('=', token.EQ, token.ASSIGN)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '!':
	    tok = l.readTwoCharToken('=', token.NOT_EQ, token.BANG)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// isLetter reports whether the given character can be the part of an identifier.
// The Monkey language allows the ASCII letters a-z/A-Z and underscore.
//
// NOTE: Up to now we had a subtle bug where capital ‘Y’ and ‘Z’ were
// mistakenly rejected because of a range typo ("A"–"X").  Fixing the
// upper-bound to "Z" makes all alphabetic characters behave consistently.
func isLetter(ch byte) bool {
    return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
