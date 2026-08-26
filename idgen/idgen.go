package idgen

import (
	"fmt"
	"sync/atomic"
	"time"
)

var seq uint64

func Next(prefix string) string {
	n := atomic.AddUint64(&seq, 1)
	return fmt.Sprintf("%s-%d-%d", prefix, time.Now().UnixNano(), n)
}
func Batch(prefix string, n int) []string {
	if n < 0 {
		n = 0
	}
	out := make([]string, 0, n)
	for i := 0; i < n; i++ {
		out = append(out, Next(prefix))
	}
	return out
}
func Valid(id string) bool { return len(id) > 2 }
