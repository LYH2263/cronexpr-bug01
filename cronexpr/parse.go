package cronexpr

import (
	"strings"

	"github.com/LYH2263/go-cronexpr/internal/ast"
)

// Parse compiles a cron expression into an Expr.
func Parse(spec string) (*Expr, error) {
	return ParseWithOptions(spec, Options{Location: DefaultLocation})
}

// MustParse panics if spec is invalid.
func MustParse(spec string) *Expr {
	e, err := Parse(spec)
	if err != nil {
		panic(err)
	}
	return e
}

func parseTree(spec string) (*ast.Expr, error) {
	spec = strings.TrimSpace(spec)
	if spec == "" {
		return nil, ErrEmpty
	}
	b := ast.NewBuilder(spec)
	defer b.Close()
	return b.Build()
}

// ParseReader parses cron text from a multi-line reader string.
func ParseReader(body string) (*Expr, error) {
	lines := strings.Split(body, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		return Parse(line)
	}
	return nil, ErrEmpty
}
