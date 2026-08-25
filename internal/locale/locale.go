package locale

import (
	"fmt"

	"github.com/LYH2263/go-cronexpr/internal/ast"
)

// Pack renders AST to human text.
type Pack struct {
	Name string
}

var registry = map[string]*Pack{
	"zh": {Name: "zh"},
	"en": {Name: "en"},
}

func Get(name string) (*Pack, error) {
	p, ok := registry[name]
	if !ok {
		return nil, fmt.Errorf("unknown locale %q", name)
	}
	return p, nil
}

func (p *Pack) Describe(e *ast.Expr) string {
	if e == nil {
		return ""
	}
	switch p.Name {
	case "en":
		return describeEN(e)
	default:
		return describeZH(e)
	}
}
