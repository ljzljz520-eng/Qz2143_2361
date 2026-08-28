package workflow

import (
	"community50/model"
	"community50/store"
)

type Transition struct{ From, To string }

func Allowed(from, to string) bool {
	for _, x := range []Transition{{"immediate", "reviewed"}, {"reviewed", "immediate"}, {"reviewed", "archived"}, {"immediate", "archived"}} {
		if x.From == from && x.To == to {
			return true
		}
	}
	return false
}
func Snapshot(db *store.DB, id string) (string, error) {
	r, e := db.GetRecord(id)
	if e != nil || r == nil {
		return "", e
	}
	return r.Status, nil
}
func StateLabel(r *model.Record) string {
	if r == nil {
		return "missing"
	}
	return r.Status
}
