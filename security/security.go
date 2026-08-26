package security

import (
	"community50/model"
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func Token(user, secret string) string {
	h := sha256.Sum256([]byte(user + ":" + secret))
	return hex.EncodeToString(h[:])
}
func AllowedRole(role, action string) bool {
	rules := map[string][]string{"admin": {"read", "write", "archive"}, "moderator": {"read", "write", "archive"}, "resident": {"read", "write"}, "guest": {"read"}}
	for _, x := range rules[role] {
		if x == action {
			return true
		}
	}
	return false
}
func Sanitize(v string) string {
	return strings.Map(func(r rune) rune {
		if r < ' ' {
			return -1
		}
		return r
	}, v)
}
func CanEdit(u model.User, r model.Record) bool {
	return u.Active && (u.Role == "admin" || u.Role == "moderator" || u.ID == r.Author)
}
