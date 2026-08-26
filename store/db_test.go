package store

import (
	"community50/model"
	"path/filepath"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := filepath.Join(t.TempDir(), "x.db")
	d, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	if e = d.PutRecord(*model.NewRecord("r1", "公告", "内容", "u")); e != nil {
		t.Fatal(e)
	}
	d.Close()
	d, e = Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer d.Close()
	r, e := d.GetRecord("r1")
	if e != nil || r == nil {
		t.Fatalf("%v %#v", e, r)
	}
}
