package hync

import "sync"

type On struct {
	mu sync.Mutex
	on []func()
}

func NewOn(on ...func()) *On {
	return &On{
		on: on,
	}
}

func (w *On) Add(on func()) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.on = append(w.on, on)
}

func (w *On) On() {
	if len(w.on) == 0 {
		return
	}
	for _, on := range w.on {
		on()
	}
}
