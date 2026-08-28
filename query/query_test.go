package query

import (
	"community50/model"
	"community50/store"
	"path/filepath"
	"testing"
)

func TestFindByStatus(t *testing.T) {
	d, _ := store.Open(filepath.Join(t.TempDir(), "d"))
	defer d.Close()
	d.PutRecord(*model.NewRecord("q", "t", "b", "u"))
	r, e := New(d).Find(model.Query{Status: "immediate"})
	if e != nil || len(r) != 1 {
		t.Fatal(e, len(r))
	}
}
