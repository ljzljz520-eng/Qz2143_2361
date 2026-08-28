package api

import (
	"community50/intake"
	"community50/query"
	"community50/store"
	"community50/workflow"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

func TestHealth(t *testing.T) {
	d, _ := store.Open(filepath.Join(t.TempDir(), "d"))
	defer d.Close()
	s := New(intake.New(d), workflow.New(d), query.New(d))
	w := httptest.NewRecorder()
	s.health(w, httptest.NewRequest("GET", "/health", nil))
	if w.Code != 200 {
		t.Fatal(w.Code)
	}
}
