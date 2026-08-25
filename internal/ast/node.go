package ast

import "github.com/LYH2263/go-cronexpr/internal/field"

// Expr is the root AST node for a cron schedule.
type Expr struct {
	Fields      []Field
	WithSeconds bool
	RawSpec     string
}

// Field is one cron column with bitmap constraints.
type Field struct {
	Kind   field.Kind
	Bitmap field.Bitmap
	Raw    string
	Star   bool
	Quest  bool
}

// Clone returns a deep copy of the expression tree.
func (e *Expr) Clone() *Expr {
	if e == nil {
		return nil
	}
	cp := &Expr{
		WithSeconds: e.WithSeconds,
		RawSpec:     e.RawSpec,
		Fields:      make([]Field, len(e.Fields)),
	}
	for i, f := range e.Fields {
		cp.Fields[i] = Field{
			Kind:   f.Kind,
			Bitmap: f.Bitmap.Clone(),
			Raw:    f.Raw,
			Star:   f.Star,
			Quest:  f.Quest,
		}
	}
	return cp
}
