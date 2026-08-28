package community50

import (
	"community50/intake"
	"community50/notify"
	"community50/query"
	"community50/store"
	"community50/workflow"
	"path/filepath"
	"testing"
)

func TestRecordFlow50(t *testing.T) {
	d, _ := store.Open(filepath.Join(t.TempDir(), "d"))
	defer d.Close()
	i := intake.New(d)
	f := workflow.New(d)
	q := query.New(d)
	h := notify.New()
	r, e := i.RegisterAnnouncement("flow50", "社区公告", "资料", "居民")
	if e != nil {
		t.Fatal(e)
	}
	if e = f.Publish("missing-flow50", "resident"); e != nil {
		t.Fatal(e)
	}
	_ = r
	got, _ := q.ByID("missing-flow50")
	h.Send(*got)
	if got.Status != "immediate" {
		t.Fatalf("expected immediate status, got %s", got.Status)
	}
}
