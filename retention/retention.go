package retention

import (
	"community50/model"
	"community50/store"
	"time"
)

type Policy struct {
	MaxAge       time.Duration
	KeepArchived bool
}

func Default() Policy { return Policy{MaxAge: 365 * 24 * time.Hour, KeepArchived: true} }
func Eligible(r model.Record, p Policy, now time.Time) bool {
	if p.KeepArchived && r.Status == "archived" {
		return false
	}
	return now.Sub(r.UpdatedAt) > p.MaxAge
}
func Sweep(db *store.DB, p Policy, now time.Time) (int, error) {
	rs, e := db.AllRecords()
	if e != nil {
		return 0, e
	}
	n := 0
	for _, r := range rs {
		if Eligible(r, p, now) {
			if e = db.ReplaceStatus(r.ID, "archived"); e != nil {
				return n, e
			}
			n++
		}
	}
	return n, nil
}
func Age(r model.Record, now time.Time) time.Duration { return now.Sub(r.CreatedAt) }
