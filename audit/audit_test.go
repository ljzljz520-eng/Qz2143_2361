package audit

import (
	"community50/store"
	"path/filepath"
	"testing"
)

func TestLog(t *testing.T) {
	d, _ := store.Open(filepath.Join(t.TempDir(), "d"))
	defer d.Close()
	if e := New(d).Log("u", "x", "r"); e != nil {
		t.Fatal(e)
	}
}
