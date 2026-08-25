package tz

import "time"

// AdjustDST skips forward across DST gaps and deduplicates overlaps.
func (l *Layer) AdjustDST(t time.Time) time.Time {
	if l == nil || l.loc == nil {
		return t
	}
	_, off1 := t.In(l.loc).Zone()
	probe := t.Add(time.Hour)
	_, off2 := probe.In(l.loc).Zone()
	if off2-off1 > 0 {
		return t.Add(time.Duration(off2-off1) * time.Second)
	}
	return t
}

// IsDST reports whether t is in daylight saving for layer location.
func (l *Layer) IsDST(t time.Time) bool {
	name, _ := t.In(l.loc).Zone()
	return len(name) > 0 && name != "UTC" && name[len(name)-1] == 'D'
}
