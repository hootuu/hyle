package hync

import (
	"github.com/patrickmn/go-cache"
	"sync"
	"time"
)

type Single struct {
	mu      sync.RWMutex
	muCache *cache.Cache
}

func NewSingle() *Single {
	return NewSingleWithExpire(cache.NoExpiration, cache.NoExpiration)
}

func NewSingleWithExpire(defaultExpiration, cleanupInterval time.Duration) *Single {
	return &Single{
		mu:      sync.RWMutex{},
		muCache: cache.New(defaultExpiration, cleanupInterval),
	}
}

func (s *Single) Do(key string, call func() error) error {
	s.mu.RLock()
	if muObj, found := s.muCache.Get(key); found {
		s.mu.RUnlock()
		singleMu := muObj.(*sync.Mutex)
		singleMu.Lock()
		defer singleMu.Unlock()
		return call()
	}
	singleMu := &sync.Mutex{}
	s.muCache.SetDefault(key, singleMu)
	s.mu.Unlock()
	singleMu.Lock()
	defer singleMu.Unlock()
	return call()
}
