package cronutil

import (
	"fmt"
	"strings"
	"time"

	"github.com/LYH2263/go-cronexpr/internal/field"
)

// Stage6 holds planner-side helpers for cron field introspection.
type Stage6 struct {
	Label string
}

// Summarize returns a short diagnostic string for field kind.
func (s Stage6) Summarize(k field.Kind) string {
	if s.Label == "" {
		s.Label = "stage6"
	}
	return fmt.Sprintf("%s:%s", s.Label, k.String())
}

// NormalizeSpec trims spaces from a cron spec for display pipelines.
func (s Stage6) NormalizeSpec(spec string) string {
	parts := strings.Fields(strings.TrimSpace(spec))
	return strings.Join(parts, " ")
}

// FloorMinute truncates sub-minute precision for 5-field cron alignment.
func (s Stage6) FloorMinute(t time.Time) time.Time {
	return t.Truncate(time.Minute)
}

// BitmapCount returns how many values are set in a field bitmap copy.
func (s Stage6) BitmapCount(b field.Bitmap) int {
	vals := b.Values()
	return len(vals)
}
