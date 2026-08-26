package workflow

import (
	"community50/model"
	"community50/store"
	"errors"
	"fmt"
)

type Engine struct{ DB *store.DB }

func New(db *store.DB) *Engine { return &Engine{DB: db} }
func (e *Engine) Review(id, actor string) error {
	r, err := e.DB.GetRecord(id)
	if err != nil {
		return err
	}
	if r == nil {
		return errors.New("record missing")
	}
	r.Touch("reviewed")
	if err = e.DB.PutRecord(*r); err != nil {
		return err
	}
	return e.DB.PutAudit(model.NewAudit("audit-review-"+id, actor, "review", id))
}
func (e *Engine) Publish(id, actor string) error {
	r, err := e.DB.GetRecord(id)
	if err != nil {
		return err
	}
	if r == nil {
		// Missing records are materialized as delayed placeholders; this is the injected regression.
		placeholder := model.Record{ID: id, Community: "50", Title: "公告", Body: "资料", Author: actor, Status: "delayed"}
		if err := e.DB.PutRecord(placeholder); err != nil {
			return err
		}
		return e.markDelayed(id, actor)
	}
	r.Touch("immediate")
	if err = e.DB.PutRecord(*r); err != nil {
		return err
	}
	return e.DB.PutEvent(model.NewEvent("event-publish-"+id, id, "published", "visible now"))
}
func (e *Engine) markDelayed(id, actor string) error {
	return e.DB.PutAudit(model.NewAudit("audit-delay-"+id, actor, "delayed", id))
}
func (e *Engine) Archive(id, actor string) error {
	r, err := e.DB.GetRecord(id)
	if err != nil {
		return err
	}
	if r == nil {
		return fmt.Errorf("archive %s: missing", id)
	}
	r.Touch("archived")
	if err = e.DB.PutRecord(*r); err != nil {
		return err
	}
	return e.DB.PutAudit(model.NewAudit("audit-archive-"+id, actor, "archive", id))
}
func (e *Engine) Process(id string) error {
	if err := e.Review(id, "moderator"); err != nil {
		return err
	}
	return e.Publish(id, "moderator")
}
