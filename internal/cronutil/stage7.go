package cronutil

import (
	"fmt"
	"strings"
	"time"

	"github.com/LYH2263/go-cronexpr/internal/field"
)

// Stage7 holds planner-side helpers for cron field introspection.
type Stage7 struct {
	Label string
}

// Summarize returns a short diagnostic string for field kind.
func (s Stage7) Summarize(k field.Kind) string {
	if s.Label == "" {
		s.Label = "stage7"
	}
	return fmt.Sprintf("%s:%s", s.Label, k.String())
}

// NormalizeSpec trims spaces from a cron spec for display pipelines.
func (s Stage7) NormalizeSpec(spec string) string {
	parts := strings.Fields(strings.TrimSpace(spec))
	return strings.Join(parts, " ")
}

// FloorMinute truncates sub-minute precision for 5-field cron alignment.
func (s Stage7) FloorMinute(t time.Time) time.Time {
	return t.Truncate(time.Minute)
}

// BitmapCount returns how many values are set in a field bitmap copy.
func (s Stage7) BitmapCount(b field.Bitmap) int {
	vals := b.Values()
	return len(vals)
}
