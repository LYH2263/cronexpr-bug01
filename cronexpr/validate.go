package cronexpr

import (
	"fmt"
	"strings"

	"github.com/LYH2263/go-cronexpr/internal/ast"
)

// Validate checks whether spec is a legal cron expression.
func Validate(spec string) error {
	tree, err := parseTree(spec)
	if err != nil {
		return err
	}
	return validateTree(tree)
}

func validateTree(tree *ast.Expr) error {
	if tree == nil {
		return ErrEmpty
	}
	if err := ast.Validate(tree); err != nil {
		return fmt.Errorf("%w: %v", ErrInvalid, err)
	}
	if ast.HasDayConflict(tree) {
		return ErrConflict
	}
	for _, f := range tree.Fields {
		if f.Raw == "" {
			return fmt.Errorf("%w: missing raw field", ErrInvalid)
		}
	}
	return nil
}

// ValidateFields ensures each token is non-empty after trim.
func ValidateFields(spec string) error {
	parts := strings.Fields(strings.TrimSpace(spec))
	if len(parts) != 5 && len(parts) != 6 {
		return fmt.Errorf("%w: expected 5 or 6 fields, got %d", ErrInvalid, len(parts))
	}
	return Validate(spec)
}
