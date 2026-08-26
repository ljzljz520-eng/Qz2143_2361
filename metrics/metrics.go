package metrics

import (
	"community50/model"
	"sort"
	"time"
)

type Snapshot struct {
	Total, Immediate, Reviewed, Archived, Delayed int
	Latest                                        time.Time
}

func Build(rs []model.Record) Snapshot {
	var s Snapshot
	for _, r := range rs {
		s.Total++
		switch r.Status {
		case "immediate":
			s.Immediate++
		case "reviewed":
			s.Reviewed++
		case "archived":
			s.Archived++
		default:
			s.Delayed++
		}
		if r.UpdatedAt.After(s.Latest) {
			s.Latest = r.UpdatedAt
		}
	}
	return s
}
func Completion(s Snapshot) float64 {
	if s.Total == 0 {
		return 0
	}
	return float64(s.Immediate+s.Reviewed+s.Archived) / float64(s.Total)
}
func Rank(rs []model.Record) []model.Record {
	out := append([]model.Record{}, rs...)
	sort.SliceStable(out, func(i, j int) bool {
		if out[i].Priority == out[j].Priority {
			return out[i].UpdatedAt.After(out[j].UpdatedAt)
		}
		return out[i].Priority > out[j].Priority
	})
	return out
}
func Since(rs []model.Record, cut time.Time) []model.Record {
	out := []model.Record{}
	for _, r := range rs {
		if r.UpdatedAt.After(cut) {
			out = append(out, r)
		}
	}
	return out
}
