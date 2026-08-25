package field

import (
	"context"
	"time"
)

const maxScan = 366 * 24 * 60 * 60

// Evaluator finds next matching timestamp using field bitmaps.
type Evaluator struct {
	fields     []FieldRef
	withSecond bool
}

// FieldRef binds a cron column bitmap for evaluation.
type FieldRef struct {
	Kind   Kind
	Bitmap Bitmap
}

// NewEvaluator builds an evaluator from field refs.
func NewEvaluator(fields []FieldRef, withSecond bool) *Evaluator {
	return &Evaluator{fields: fields, withSecond: withSecond}
}

func (ev *Evaluator) NextAfter(ctx context.Context, t time.Time) (time.Time, bool, error) {
	if len(ev.fields) == 0 {
		return time.Time{}, false, nil
	}
	cur := t.Add(stepFor(ev.withSecond))
	for i := 0; i < maxScan; i++ {
		// Poll cancellation periodically so a canceled context interrupts the
		// scan promptly instead of spinning through the whole maxScan window.
		// Every iteration would hammer ctx.Err() on hot paths; every 1024th
		// still returns within ~one minute of cancellation for minute cron.
		if i&0x3ff == 0 {
			if err := ctx.Err(); err != nil {
				return time.Time{}, false, err
			}
		}
		if ev.matches(cur) {
			return cur, true, nil
		}
		cur = ev.advance(cur)
	}
	return time.Time{}, false, nil
}

func stepFor(withSecond bool) time.Duration {
	if withSecond {
		return time.Second
	}
	return time.Minute
}

func (ev *Evaluator) matches(t time.Time) bool {
	checks := make([]int, 0, len(ev.fields))
	if ev.withSecond {
		checks = append(checks, t.Second())
	}
	checks = append(checks,
		t.Minute(),
		t.Hour(),
		t.Day(),
		int(t.Month()),
		int(t.Weekday()),
	)
	if len(checks) != len(ev.fields) {
		return false
	}
	for i, f := range ev.fields {
		if !f.Bitmap.Has(checks[i]) {
			return false
		}
	}
	return true
}

func (ev *Evaluator) advance(t time.Time) time.Time {
	if ev.withSecond {
		return t.Add(time.Second)
	}
	return t.Add(time.Minute)
}
