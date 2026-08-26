package notify

import (
	"community50/model"
	"strings"
)

func Audience(r model.Record) string {
	if strings.Contains(r.Title, "居民") {
		return "residents"
	}
	return "community50"
}
func Ready(r model.Record) bool { return r.Status == "immediate" || r.Status == "reviewed" }
