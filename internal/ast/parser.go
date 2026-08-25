package ast

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/LYH2263/go-cronexpr/internal/field"
)

// ParseField parses one cron column into a Field node.
func ParseField(kind field.Kind, raw string) (Field, error) {
	bm := field.NewBitmap(kind)
	f := Field{Kind: kind, Raw: raw}
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return f, fmt.Errorf("empty field")
	}
	if raw == "*" {
		bm.SetAll()
		f.Star = true
		f.Bitmap = bm
		return f, nil
	}
	if raw == "?" {
		bm.SetAll()
		f.Quest = true
		f.Bitmap = bm
		return f, nil
	}
	parts := strings.Split(raw, ",")
	for _, part := range parts {
		if err := applyPart(&bm, part); err != nil {
			return f, err
		}
	}
	f.Bitmap = bm
	return f, nil
}

func applyPart(bm *field.Bitmap, part string) error {
	step := 1
	if strings.Contains(part, "/") {
		chunks := strings.Split(part, "/")
		if len(chunks) != 2 {
			return fmt.Errorf("bad step: %s", part)
		}
		part = chunks[0]
		s, err := strconv.Atoi(chunks[1])
		if err != nil || s <= 0 {
			return fmt.Errorf("bad step value: %s", chunks[1])
		}
		step = s
	}
	if part == "*" || part == "" {
		lo, hi := bm.Bounds()
		for v := lo; v <= hi; v += step {
			bm.Set(v)
		}
		return nil
	}
	if strings.Contains(part, "-") {
		ab := strings.Split(part, "-")
		if len(ab) != 2 {
			return fmt.Errorf("bad range: %s", part)
		}
		a, err1 := strconv.Atoi(ab[0])
		b, err2 := strconv.Atoi(ab[1])
		if err1 != nil || err2 != nil || a > b {
			return fmt.Errorf("bad range: %s", part)
		}
		for v := a; v <= b; v += step {
			bm.Set(v)
		}
		return nil
	}
	n, err := strconv.Atoi(part)
	if err != nil {
		return fmt.Errorf("bad number: %s", part)
	}
	bm.Set(n)
	return nil
}
