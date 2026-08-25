package cronexpr

import (
	"fmt"

	"github.com/LYH2263/go-cronexpr/internal/ast"
	"github.com/LYH2263/go-cronexpr/internal/locale"
)

// Describe renders a human-readable summary in default locale (zh).
func Describe(spec string) (string, error) {
	return DescribeLocale(spec, "zh")
}

// DescribeLocale renders spec using locale pack.
func DescribeLocale(spec string, locName string) (string, error) {
	tree, err := parseTree(spec)
	if err != nil {
		return "", err
	}
	pack, err := locale.Get(locName)
	if err != nil {
		return "", fmt.Errorf("%w: %v", ErrBadLocale, err)
	}
	if tree == nil {
		return "", ErrNilExpr
	}
	return pack.Describe(tree), nil
}

// DescribeExpr describes a compiled expression.
func DescribeExpr(e *Expr) (string, error) {
	if e == nil {
		return "", ErrNilExpr
	}
	if e.tree == nil {
		return "", ErrNilExpr
	}
	pack, err := locale.Get("zh")
	if err != nil {
		return "", err
	}
	return pack.Describe(e.tree), nil
}

// FieldSummary returns per-field diagnostic strings.
func FieldSummary(tree *ast.Expr) []string {
	if tree == nil {
		return nil
	}
	out := make([]string, len(tree.Fields))
	for i, f := range tree.Fields {
		out[i] = fmt.Sprintf("%s=%s", f.Kind, f.Raw)
	}
	return out
}
