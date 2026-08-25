package tz

import "time"

// Layer wraps timezone conversions for cron evaluation.
type Layer struct {
	loc *time.Location
}

func NewLayer(loc *time.Location) *Layer {
	if loc == nil {
		loc = time.Local
	}
	return &Layer{loc: loc}
}

func (l *Layer) Location() *time.Location {
	return l.loc
}

func (l *Layer) ToLocal(t time.Time) time.Time {
	return t.In(l.loc)
}

func (l *Layer) ToUTC(t time.Time) time.Time {
	return t.UTC()
}
