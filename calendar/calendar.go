package calendar

import (
	"community50/model"
	"time"
)

func DayStart(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}
func SameDay(a, b time.Time) bool {
	ay, am, ad := a.Date()
	by, bm, bd := b.Date()
	return ay == by && am == bm && ad == bd
}
func CreatedToday(rs []model.Record, now time.Time) []model.Record {
	out := []model.Record{}
	for _, r := range rs {
		if SameDay(r.CreatedAt, now) {
			out = append(out, r)
		}
	}
	return out
}
