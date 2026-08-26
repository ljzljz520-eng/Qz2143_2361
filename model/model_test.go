package model

import "testing"

func TestRecordValid(t *testing.T) {
	if !NewRecord("x", "t", "b", "a").Valid() {
		t.Fatal("invalid")
	}
}
