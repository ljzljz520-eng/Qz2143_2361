package intake

import (
	"community50/model"
	"community50/store"
	"errors"
	"fmt"
)

type Service struct{ DB *store.DB }

func New(db *store.DB) *Service { return &Service{DB: db} }
func (s *Service) Register(r *model.Record) error {
	if r == nil {
		return errors.New("nil record")
	}
	if !r.Valid() {
		return errors.New("invalid record")
	}
	if err := s.DB.PutRecord(*r); err != nil {
		return err
	}
	return nil
}
func (s *Service) RegisterAnnouncement(id, title, body, author string) (*model.Record, error) {
	r := model.NewRecord(id, title, body, author)
	if err := s.Register(r); err != nil {
		return nil, err
	}
	return r, nil
}
func (s *Service) EnsureUser(id, name, role string) (*model.User, error) {
	u := model.NewUser(id, name, role)
	if err := s.DB.PutUser(u); err != nil {
		return nil, fmt.Errorf("user: %w", err)
	}
	return &u, nil
}
func (s *Service) ValidateCommunity(r *model.Record) bool { return r != nil && r.Community == "50" }
