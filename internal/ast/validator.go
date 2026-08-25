package ast

import (
	"fmt"

	"github.com/LYH2263/go-cronexpr/internal/field"
)

// Validate checks structural constraints on the AST.
func Validate(e *Expr) error {
	if e == nil {
		return fmt.Errorf("nil expr")
	}
	if len(e.Fields) != 5 && len(e.Fields) != 6 {
		return fmt.Errorf("bad field count %d", len(e.Fields))
	}
	for _, f := range e.Fields {
		lo, hi := f.Bitmap.Bounds()
		vals := f.Bitmap.Values()
		if len(vals) == 0 && !f.Star && !f.Quest {
			return fmt.Errorf("%s has no allowed values", f.Kind)
		}
		for _, v := range vals {
			if v < lo || v > hi {
				return fmt.Errorf("%s out of range: %d", f.Kind, v)
			}
		}
	}
	return nil
}

// HasDayConflict reports dom/dow both constrained (neither * nor ?).
func HasDayConflict(e *Expr) bool {
	if e == nil || len(e.Fields) < 5 {
		return false
	}
	base := 0
	if e.WithSeconds {
		base = 1
	}
	dom := e.Fields[base+2]
	dow := e.Fields[base+4]
	domOpen := dom.Star || dom.Quest
	dowOpen := dow.Star || dow.Quest
	return !domOpen && !dowOpen
}

// FieldByKind finds a field node by kind.
func FieldByKind(e *Expr, k field.Kind) (Field, bool) {
	for _, f := range e.Fields {
		if f.Kind == k {
			return f, true
		}
	}
	return Field{}, false
}
