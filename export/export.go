package export

import (
	"community50/model"
	"encoding/json"
	"fmt"
	"strings"
)

func JSON(rs []model.Record) ([]byte, error) { return json.MarshalIndent(rs, "", "  ") }
func CSV(rs []model.Record) string {
	var b strings.Builder
	b.WriteString("id,title,status,author\n")
	for _, r := range rs {
		fmt.Fprintf(&b, "%s,%s,%s,%s\n", r.ID, strings.ReplaceAll(r.Title, ",", " "), r.Status, r.Author)
	}
	return b.String()
}
func Group(rs []model.Record) map[string][]model.Record {
	m := map[string][]model.Record{}
	for _, r := range rs {
		m[r.Status] = append(m[r.Status], r)
	}
	return m
}
func IDs(rs []model.Record) []string {
	out := make([]string, 0, len(rs))
	for _, r := range rs {
		out = append(out, r.ID)
	}
	return out
}
