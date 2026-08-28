package model

import "time"

type Record struct {
	ID, Community, Title, Body, Status, Author string
	CreatedAt, UpdatedAt                       time.Time
	Priority                                   int
}
type User struct {
	ID, Name, Role string
	Active         bool
}
type Event struct {
	ID, RecordID, Kind, Detail string
	At                         time.Time
}
type Audit struct {
	ID, Actor, Action, RecordID string
	At                          time.Time
}
type Query struct {
	Status, Author, Text string
	Limit                int
}

func NewRecord(id, title, body, author string) *Record {
	now := time.Now().UTC()
	return &Record{ID: id, Community: "50", Title: title, Body: body, Author: author, Status: "immediate", CreatedAt: now, UpdatedAt: now, Priority: 1}
}
func (r *Record) Touch(status string) { r.Status = status; r.UpdatedAt = time.Now().UTC() }
func (r Record) Valid() bool {
	return r.ID != "" && r.Community == "50" && r.Title != "" && r.Body != "" && r.Author != ""
}
func NewUser(id, name, role string) User { return User{ID: id, Name: name, Role: role, Active: true} }
func NewEvent(id, recordID, kind, detail string) Event {
	return Event{ID: id, RecordID: recordID, Kind: kind, Detail: detail, At: time.Now().UTC()}
}
func NewAudit(id, actor, action, recordID string) Audit {
	return Audit{ID: id, Actor: actor, Action: action, RecordID: recordID, At: time.Now().UTC()}
}
