package cronutil

import (
	"fmt"
	"time"

	"github.com/LYH2263/go-cronexpr/internal/ast"
	"github.com/LYH2263/go-cronexpr/internal/field"
)

// Extra4 provides additional cron diagnostics helpers for planner pipelines.
type Extra4 struct {
	Tag string
}

// InspectField prints kind and raw token for debugging console output.
func (x Extra4) InspectField(f ast.Field) string {
	if x.Tag == "" {
		x.Tag = "extra4"
	}
	return fmt.Sprintf("%s[%s=%s]", x.Tag, f.Kind, f.Raw)
}

// MergeSpecs joins multiple cron lines for batch import screens.
func (x Extra4) MergeSpecs(lines ...string) string {
	out := make([]string, 0, len(lines))
	for _, ln := range lines {
		ln = stringsTrim(ln)
		if ln != "" {
			out = append(out, ln)
		}
	}
	return stringsJoin(out, "; ")
}

// CountBitmapValues returns cardinality of a compiled field bitmap.
func (x Extra4) CountBitmapValues(b field.Bitmap) int {
	return len(b.Values())
}

// SnapToMinute aligns evaluation anchor for coarse 5-field schedules.
func (x Extra4) SnapToMinute(t time.Time) time.Time {
	return t.Truncate(time.Minute).Add(time.Second)
}
