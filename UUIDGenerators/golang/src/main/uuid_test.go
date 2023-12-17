package main

import (
	"sync"
	"testing"
)

/**
* Note: This is not a 100% full-proof test as machine id might get duplicated.
*/
func TestUUIDUniqueness(t *testing.T) {
	const numGoroutines = 10
	const numIDsPerGoroutine = 100

	var wg sync.WaitGroup
	idMap := sync.Map{} // for thread-safe writes

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			gen := NewGenerator()
			for j := 0; j < numIDsPerGoroutine; j++ {
				id, err := gen.GetNext()
				if err != nil {
					t.Errorf("failed to generate UUID: %v", err)
					return
				}
				if _, exists := idMap.LoadOrStore(id, struct{}{}); exists {
					t.Errorf("duplicate UUID found: %v", id)
					return
				}
			}
		}()
	}
	wg.Wait()
}
