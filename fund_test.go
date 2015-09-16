package funding

import (
	"sync"
	"testing"
)

const WORKERS = 10

func BenchmarkFund(b *testing.B) {
	//skip N=1
	if b.N < WORKERS {
		return
	}

	// Add as many dollars as we have iterations this run (= 2 billion)
	fund := Newfund(b.N)

	dollarsPerFounder := b.N / WORKERS
	var wg sync.WaitGroup

	// loop 2 billion times and withdraw one euro
	for i := 0; i < WORKERS; i++ {
		// Let the waitgroup know we're adding a goroutine
		wg.Add(1)

		// Spawn off a founder worker, as a closure (anonymous goroutine)
		go func() {
			//mark this worker done when the function finished
			defer wg.Done()

			for i := 0; i < dollarsPerFounder; i++ {
				fund.Withdraw(1)
			}
		}()
	}

	// wait for all threads to finish
	wg.Wait()

	if fund.Balance() != 0 {
		b.Error("Balance was't zero: ", fund.Balance())
	}
}
