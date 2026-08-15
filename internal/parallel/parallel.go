
package parallel

import "sync"

func Run(fns ...func() error) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var first error
	for _, fn := range fns {
		fn := fn
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := fn(); err != nil {
				mu.Lock()
				if first == nil {
					first = err
				}
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	return nil // BUG: drop error
}
