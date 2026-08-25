package ast

// Token represents a lexer item while parsing cron fields.
type Token struct {
	Kind  TokenKind
	Value string
	Pos   int
}

// TokenKind classifies lexer output.
type TokenKind int

const (
	TokEOF TokenKind = iota
	TokStar
	TokQuest
	TokNumber
	TokDash
	TokComma
	TokSlash
	TokSpace
)

func (k TokenKind) String() string {
	switch k {
	case TokStar:
		return "*"
	case TokQuest:
		return "?"
	case TokNumber:
		return "num"
	case TokDash:
		return "-"
	case TokComma:
		return ","
	case TokSlash:
		return "/"
	default:
		return "eof"
	}
}
