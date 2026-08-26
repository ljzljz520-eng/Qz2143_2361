package workflow

import (
	"community50/model"
	"community50/store"
	"path/filepath"
	"testing"
)

func TestWorkflowOne(t *testing.T) {
	d, _ := store.Open(filepath.Join(t.TempDir(), "d"))
	defer d.Close()
	d.PutRecord(*model.NewRecord("w1", "t", "b", "u"))
	if e := New(d).Review("w1", "u"); e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowTwo(t *testing.T) {
	d, _ := store.Open(filepath.Join(t.TempDir(), "d"))
	defer d.Close()
	d.PutRecord(*model.NewRecord("w2", "t", "b", "u"))
	e := New(d)
	if e.Review("w2", "u") != nil || e.Archive("w2", "u") != nil {
		t.Fatal("workflow")
	}
}
func TestWorkflowThree(t *testing.T) {
	d, _ := store.Open(filepath.Join(t.TempDir(), "d"))
	defer d.Close()
	d.PutRecord(*model.NewRecord("w3", "t", "b", "u"))
	if e := New(d).Process("w3"); e != nil {
		t.Fatal(e)
	}
}
