package timeline

import (
	"community50/model"
	"sort"
	"time"
)

type Item struct {
	At    time.Time
	Label string
}

func RecordItems(r model.Record) []Item {
	return []Item{{r.CreatedAt, "created"}, {r.UpdatedAt, r.Status}}
}
func Sort(xs []Item) []Item {
	out := append([]Item{}, xs...)
	sort.Slice(out, func(i, j int) bool { return out[i].At.Before(out[j].At) })
	return out
}
func Merge(all ...[]Item) []Item {
	out := []Item{}
	for _, xs := range all {
		out = append(out, xs...)
	}
	return Sort(out)
}
func Labels(xs []Item) []string {
	out := []string{}
	for _, x := range xs {
		out = append(out, x.Label)
	}
	return out
}
