package locale

import (
	"fmt"
	"strings"

	"github.com/LYH2263/go-cronexpr/internal/ast"
)

func describeEN(e *ast.Expr) string {
	if e == nil {
		return ""
	}
	parts := make([]string, 0, len(e.Fields))
	for _, f := range e.Fields {
		parts = append(parts, fmt.Sprintf("%s=%s", f.Kind, f.Raw))
	}
	return "Schedule: " + strings.Join(parts, ", ")
}
