package api

import (
	"community50/intake"
	"community50/model"
	"community50/query"
	"community50/workflow"
	"encoding/json"
	"net/http"
)

type Server struct {
	Intake *intake.Service
	Flow   *workflow.Engine
	Query  *query.Service
}

func New(i *intake.Service, f *workflow.Engine, q *query.Service) *Server {
	return &Server{Intake: i, Flow: f, Query: q}
}
func (s *Server) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", s.health)
	mux.HandleFunc("/records", s.records)
	mux.HandleFunc("/records/publish", s.publish)
	return mux
}
func (s *Server) health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}
func (s *Server) records(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var in struct{ ID, Title, Body, Author string }
		if json.NewDecoder(r.Body).Decode(&in) != nil {
			http.Error(w, "bad json", 400)
			return
		}
		x, e := s.Intake.RegisterAnnouncement(in.ID, in.Title, in.Body, in.Author)
		if e != nil {
			http.Error(w, e.Error(), 400)
			return
		}
		_ = json.NewEncoder(w).Encode(x)
		return
	}
	q := model.Query{Status: r.URL.Query().Get("status"), Text: r.URL.Query().Get("q")}
	xs, e := s.Query.Find(q)
	if e != nil {
		http.Error(w, e.Error(), 500)
		return
	}
	_ = json.NewEncoder(w).Encode(xs)
}
func (s *Server) publish(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "missing id", 400)
		return
	}
	if e := s.Flow.Publish(id, "api"); e != nil {
		http.Error(w, e.Error(), 404)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
