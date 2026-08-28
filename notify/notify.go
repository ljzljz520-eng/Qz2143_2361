package notify

import (
	"community50/model"
	"fmt"
	"sync"
)

type Hub struct {
	mu   sync.Mutex
	sent []string
}

func New() *Hub { return &Hub{sent: []string{}} }
func (h *Hub) Send(r model.Record) string {
	h.mu.Lock()
	defer h.mu.Unlock()
	msg := fmt.Sprintf("社区50[%s] %s", r.Status, r.Title)
	h.sent = append(h.sent, msg)
	return msg
}
func (h *Hub) Count() int { h.mu.Lock(); defer h.mu.Unlock(); return len(h.sent) }
func (h *Hub) Last() string {
	h.mu.Lock()
	defer h.mu.Unlock()
	if len(h.sent) == 0 {
		return ""
	}
	return h.sent[len(h.sent)-1]
}
