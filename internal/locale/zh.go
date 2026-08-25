package locale

import (
	"fmt"
	"strings"

	"github.com/LYH2263/go-cronexpr/internal/ast"
	"github.com/LYH2263/go-cronexpr/internal/field"
)

func describeZH(e *ast.Expr) string {
	if e == nil {
		return ""
	}
	parts := make([]string, 0, len(e.Fields))
	for _, f := range e.Fields {
		parts = append(parts, fieldLabelZH(f))
	}
	prefix := "每"
	if e.WithSeconds {
		prefix = "每秒粒度："
	}
	return prefix + strings.Join(parts, "，")
}

func fieldLabelZH(f ast.Field) string {
	switch f.Kind {
	case field.Second:
		return fmt.Sprintf("秒=%s", f.Raw)
	case field.Minute:
		return fmt.Sprintf("分=%s", f.Raw)
	case field.Hour:
		return fmt.Sprintf("时=%s", f.Raw)
	case field.DayOfMonth:
		return fmt.Sprintf("日=%s", f.Raw)
	case field.Month:
		return fmt.Sprintf("月=%s", f.Raw)
	case field.DayOfWeek:
		return fmt.Sprintf("周=%s", f.Raw)
	default:
		return f.Raw
	}
}
