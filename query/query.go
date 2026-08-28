package query

import (
	"community50/model"
	"community50/store"
	"sort"
	"strings"
)

type Service struct{ DB *store.DB }

func New(db *store.DB) *Service { return &Service{DB: db} }
func (s *Service) Find(q model.Query) ([]model.Record, error) {
	all, e := s.DB.AllRecords()
	if e != nil {
		return nil, e
	}
	out := make([]model.Record, 0)
	for _, r := range all {
		if q.Status != "" && r.Status != q.Status {
			continue
		}
		if q.Author != "" && r.Author != q.Author {
			continue
		}
		if q.Text != "" && !strings.Contains(strings.ToLower(r.Title+" "+r.Body), strings.ToLower(q.Text)) {
			continue
		}
		out = append(out, r)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].UpdatedAt.After(out[j].UpdatedAt) })
	if q.Limit > 0 && len(out) > q.Limit {
		out = out[:q.Limit]
	}
	return out, nil
}
func (s *Service) ByID(id string) (*model.Record, error) { return s.DB.GetRecord(id) }
func Summarize(rs []model.Record) map[string]int {
	m := map[string]int{}
	for _, r := range rs {
		m[r.Status]++
	}
	return m
}
