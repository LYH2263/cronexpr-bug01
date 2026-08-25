package cronexpr

import "time"

// ParseWithOptions compiles spec using optional timezone/locale defaults.
func ParseWithOptions(spec string, opts Options) (*Expr, error) {
	opts = opts.withDefaults()
	tree, err := parseTree(spec)
	if err != nil {
		return nil, err
	}
	if err := validateTree(tree); err != nil {
		return nil, err
	}
	return &Expr{
		tree:       tree,
		loc:        opts.Location,
		withSecond: tree.WithSeconds,
		rawSpec:    cloneString(spec),
	}, nil
}

func cloneString(s string) string {
	b := []byte(s)
	out := make([]byte, len(b))
	copy(out, b)
	return string(out)
}

// StandardFields returns field count (5 or 6).
func StandardFields(withSecond bool) int {
	if withSecond {
		return 6
	}
	return 5
}

// DefaultLocation is used when none is provided.
var DefaultLocation = time.Local
