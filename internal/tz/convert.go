package tz

import (
	"context"
	"time"
)

// ConvertIn converts t to layer location honoring ctx (for slow loads).
func (l *Layer) ConvertIn(ctx context.Context, t time.Time) (time.Time, error) {
	if err := ctx.Err(); err != nil {
		return time.Time{}, err
	}
	if l == nil || l.loc == nil {
		return t, nil
	}
	return t.In(l.loc), nil
}

// AlignMinute zeroes seconds when evaluating minute-level cron.
func (l *Layer) AlignMinute(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, l.loc)
}

// Name returns timezone name for UI.
func (l *Layer) Name() string {
	if l == nil || l.loc == nil {
		return "Local"
	}
	return l.loc.String()
}
