package model

import "strings"

func NormalizeTitle(v string) string { return strings.TrimSpace(strings.Join(strings.Fields(v), " ")) }
func (r Record) SearchText() string  { return strings.ToLower(r.Title + " " + r.Body + " " + r.Author) }
