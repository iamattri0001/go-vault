package dbutils

import (
	"sync"
)

type Operation func() error

// RunParallel executes multiple DB operations concurrently.
// Returns the first error encountered (if any).
func RunParallel(ops ...Operation) error {
	var wg sync.WaitGroup
	errCh := make(chan error, len(ops))

	for _, op := range ops {
		wg.Add(1)
		go func(op Operation) {
			defer wg.Done()
			if err := op(); err != nil {
				errCh <- err
			}
		}(op)
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			return err
		}
	}
	return nil
}
