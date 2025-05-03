package hync

import "sync"

type On[T any] struct {
	mu sync.Mutex
	on []func(ctx *T)
}

func NewOn[T any](on ...func(ctx *T)) *On[T] {
	return &On[T]{
		on: on,
	}
}

func (w *On[T]) Add(on func(ctx *T)) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.on = append(w.on, on)
}

func (w *On[T]) On(ctx *T) {
	if len(w.on) == 0 {
		return
	}
	for _, on := range w.on {
		on(ctx)
	}
}
