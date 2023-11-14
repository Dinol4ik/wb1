package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64

func AtomicCounter(wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt64(&counter, 1)
	fmt.Printf("goroutine atomic add +1, counter = %d\n", counter)
}

func MutexCounter(wg *sync.WaitGroup, mut *sync.Mutex) {
	defer wg.Done()
	mut.Lock()
	counter++
	mut.Unlock()
	fmt.Printf("goroutine mutex add +1, counter = %d\n", counter)
}

func main() {
	const workerPull = 5
	mutex := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	for i := 0; i < workerPull; i++ {
		wg.Add(1)
		go AtomicCounter(wg)
	}
	for i := 0; i < workerPull; i++ {
		wg.Add(1)
		go MutexCounter(wg, mutex)
	}
	wg.Wait()
}
