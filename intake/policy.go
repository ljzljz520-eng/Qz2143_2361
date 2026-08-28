package intake

import (
	"community50/model"
	"strings"
)

func Clean(r *model.Record) {
	r.Title = model.NormalizeTitle(r.Title)
	r.Body = strings.TrimSpace(r.Body)
	r.Community = "50"
}
func Priority(title string) int {
	if strings.Contains(title, "紧急") {
		return 3
	}
	if strings.Contains(title, "重要") {
		return 2
	}
	return 1
}
