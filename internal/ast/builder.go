package ast

import (
	"fmt"
	"strings"

	"github.com/LYH2263/go-cronexpr/internal/field"
)

// Builder constructs Expr AST from a cron string.
type Builder struct {
	spec   string
	closed bool
}

func NewBuilder(spec string) *Builder {
	return &Builder{spec: strings.TrimSpace(spec)}
}

func (b *Builder) Close() {
	b.closed = true
}

func (b *Builder) Build() (*Expr, error) {
	if b.closed {
		return nil, fmt.Errorf("builder closed")
	}
	parts := strings.Fields(b.spec)
	if len(parts) != 5 && len(parts) != 6 {
		return nil, fmt.Errorf("expected 5 or 6 fields, got %d", len(parts))
	}
	withSec := len(parts) == 6
	kinds := field.StandardKinds(withSec)
	fields := make([]Field, len(parts))
	for i, raw := range parts {
		f, err := ParseField(kinds[i], raw)
		if err != nil {
			return nil, fmt.Errorf("field %d: %w", i, err)
		}
		fields[i] = f
	}
	return &Expr{
		Fields:      fields,
		WithSeconds: withSec,
		RawSpec:     b.spec,
	}, nil
}
