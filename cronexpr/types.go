package cronexpr

import (
	"time"

	"github.com/LYH2263/go-cronexpr/internal/ast"
	"github.com/LYH2263/go-cronexpr/internal/tz"
)

// Expr is a compiled cron schedule.
type Expr struct {
	tree       *ast.Expr
	loc        *time.Location
	withSecond bool
	rawSpec    string
}

// Options configures parsing and evaluation.
type Options struct {
	Location *time.Location
	Locale   string
}

func (o Options) withDefaults() Options {
	if o.Location == nil {
		o.Location = time.Local
	}
	if o.Locale == "" {
		o.Locale = "zh"
	}
	return o
}

// Tree exposes the parsed AST for introspection tools.
func (e *Expr) Tree() *ast.Expr {
	if e == nil {
		return nil
	}
	return e.tree
}

// Location returns the evaluation timezone.
func (e *Expr) Location() *time.Location {
	if e == nil || e.loc == nil {
		return time.Local
	}
	return e.loc
}

// RawSpec returns the original cron string.
func (e *Expr) RawSpec() string {
	if e == nil {
		return ""
	}
	return e.rawSpec
}

// TZLayer returns the timezone helper bound to this expression.
func (e *Expr) TZLayer() *tz.Layer {
	return tz.NewLayer(e.loc)
}
