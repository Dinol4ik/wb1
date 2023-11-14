package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func workerStartRead(c chan int, workerNum int, ctx context.Context, group *sync.WaitGroup) {
	defer group.Done()
	for {
		select {
		case num := <-c:
			fmt.Printf("worker %d read %d\n", workerNum, num)
		case <-ctx.Done():
			fmt.Printf("worker %d finish read channel \n", workerNum)
			return
		}
	}
}

func main() {
	const workersNum = 3
	wg := &sync.WaitGroup{}
	channel := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(workersNum)
	ticker := time.NewTicker(2 * time.Second)

	for i := 0; i < workersNum; i++ {
		go workerStartRead(channel, i, ctx, wg)
	}

	go func() {
		for {
			select {
			case <-ticker.C:
				rNum := rand.Int()
				channel <- rNum
			case <-ctx.Done():
				fmt.Println("finish write")
				return
			}
		}
	}()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	fmt.Println("\ngot interrupt signal")
	cancel()
	wg.Wait()

}
