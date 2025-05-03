package hync

import "sync"

type Line struct {
	mu sync.Mutex
}

func NewLine() *Line {
	return &Line{}
}

func (line *Line) Do(call func() error) error {
	line.mu.Lock()
	defer line.mu.Unlock()
	return call()
}
