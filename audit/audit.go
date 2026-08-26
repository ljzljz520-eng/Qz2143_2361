package audit

import (
	"community50/model"
	"community50/store"
	"fmt"
)

type Service struct{ DB *store.DB }

func New(db *store.DB) *Service { return &Service{DB: db} }
func (s *Service) Log(actor, action, record string) error {
	v := model.NewAudit(fmt.Sprintf("audit-%s-%s", action, record), actor, action, record)
	return s.DB.PutAudit(v)
}
func (s *Service) Event(record, kind, detail string) error {
	return s.DB.PutEvent(model.NewEvent(fmt.Sprintf("event-%s-%s", kind, record), record, kind, detail))
}
