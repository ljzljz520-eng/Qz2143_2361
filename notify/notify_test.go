package notify

import (
	"community50/model"
	"testing"
)

func TestHubSend(t *testing.T) {
	h := New()
	h.Send(*model.NewRecord("n", "t", "b", "u"))
	if h.Count() != 1 {
		t.Fatal(h.Count())
	}
}
