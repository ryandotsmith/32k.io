package limit

import (
	"net/http"
	"sync"
)

const OneMB = 1e6 // 1MB

func MaxBytes(h http.Handler, max int64) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		r2 := new(http.Request)
		*r2 = *req
		r2.Body = http.MaxBytesReader(w, req.Body, max)
		h.ServeHTTP(w, r2)
	})
}

type Handler struct {
	next http.Handler
	l    *ConcurrencyLimiter
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !h.l.Acquire(r.RemoteAddr) {
		http.Error(w, "too many requests", http.StatusTooManyRequests)
		return
	}
	h.next.ServeHTTP(w, r)
}

func NewHandler(next http.Handler, limit int) *Handler {
	return &Handler{next: next, l: NewConcurrencyLimiter(limit)}
}

// A ConcurrencyLimiter limits how many concurrent requests
// are allowed to happen per key.
type ConcurrencyLimiter struct {
	limit int

	mu     sync.Mutex
	counts map[string]int
}

func NewConcurrencyLimiter(limit int) *ConcurrencyLimiter {
	return &ConcurrencyLimiter{
		limit:  limit,
		counts: make(map[string]int),
	}
}

// Acquire consumes an available slot for key,
// reporting whether one was available. If acquired,
// the caller is responsible for releasing the slot
// when the request is complete.
func (c *ConcurrencyLimiter) Acquire(key string) (ok bool) {
	c.mu.Lock()
	v := c.counts[key]
	if v < c.limit {
		c.counts[key] = v + 1
		ok = true
	}
	c.mu.Unlock()
	return ok
}

// Release returns an acquired slot back to the limiter.
func (c *ConcurrencyLimiter) Release(key string) {
	c.mu.Lock()
	c.counts[key] -= 1
	c.mu.Unlock()
}
