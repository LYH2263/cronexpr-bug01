package ast

import "unicode"

// Lexer tokenizes one cron field fragment.
type Lexer struct {
	input string
	pos   int
}

func NewLexer(input string) *Lexer {
	return &Lexer{input: input}
}

func (l *Lexer) Next() Token {
	if l.pos >= len(l.input) {
		return Token{Kind: TokEOF}
	}
	ch := rune(l.input[l.pos])
	l.pos++
	switch ch {
	case '*':
		return Token{Kind: TokStar, Value: "*", Pos: l.pos - 1}
	case '?':
		return Token{Kind: TokQuest, Value: "?", Pos: l.pos - 1}
	case '-':
		return Token{Kind: TokDash, Value: "-", Pos: l.pos - 1}
	case ',':
		return Token{Kind: TokComma, Value: ",", Pos: l.pos - 1}
	case '/':
		return Token{Kind: TokSlash, Value: "/", Pos: l.pos - 1}
	default:
		if unicode.IsDigit(ch) {
			start := l.pos - 1
			for l.pos < len(l.input) && unicode.IsDigit(rune(l.input[l.pos])) {
				l.pos++
			}
			return Token{Kind: TokNumber, Value: l.input[start:l.pos], Pos: start}
		}
	}
	return Token{Kind: TokEOF}
}

func (l *Lexer) Rest() string {
	if l.pos >= len(l.input) {
		return ""
	}
	return l.input[l.pos:]
}
