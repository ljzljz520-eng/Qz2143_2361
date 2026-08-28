package intake

import (
	"community50/store"
	"path/filepath"
	"testing"
)

func TestRegisterAnnouncement(t *testing.T) {
	d, _ := store.Open(filepath.Join(t.TempDir(), "d"))
	defer d.Close()
	s := New(d)
	r, e := s.RegisterAnnouncement("a", " title ", "body", "u")
	if e != nil || r.Community != "50" {
		t.Fatal(e, r)
	}
}
