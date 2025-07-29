package hync

import "sync"

func WgFunc(wg *sync.WaitGroup, f func()) {
	if wg != nil {
		wg.Add(1)
	}
	go func() {
		defer func() {
			if wg != nil {
				wg.Done()
			}
		}()
		f()
	}()
}
