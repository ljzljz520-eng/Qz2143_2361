package audit

import (
	"community50/model"
	"strings"
)

func Describe(v model.Audit) string {
	return strings.Join([]string{v.Actor, v.Action, v.RecordID}, "/")
}
