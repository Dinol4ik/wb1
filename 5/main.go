package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func workerRead(c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("workerRead finished")
}
func workerWrite(c chan int, wg *sync.WaitGroup, second int) {
	defer wg.Done()
	timer := time.NewTimer(time.Duration(second) * time.Second)
	ticker := time.NewTicker(1 * time.Millisecond)
	defer timer.Stop()
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			rNum := rand.Int()
			c <- rNum
		case <-timer.C:
			close(c)
			fmt.Println("timer is down")
			return
		}
	}
}
func main() {
	sec := 2
	channel := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go workerWrite(channel, wg, sec)
	go workerRead(channel, wg)
	wg.Wait()
}
