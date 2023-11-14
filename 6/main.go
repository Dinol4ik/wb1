package main

import (
	"context"
	"fmt"
	"time"
)

func goroutineCancelWithChan(c chan struct{}) {
	for {
		select {
		default:
			fmt.Println("работа...")
		case <-c:
			fmt.Println("goroutine cancel with channel")
			return
		}
	}
}
func goroutineCancelWithCtx(ctx context.Context) {
	for {
		select {
		default:
			time.Sleep(100 * time.Millisecond)
			fmt.Println("работа...")
		case <-ctx.Done():
			fmt.Println("goroutine cancel with context")
			return
		}
	}
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan struct{})
	go goroutineCancelWithChan(c)
	go goroutineCancelWithCtx(ctx)
	c <- struct{}{} // завершение горутины по сигналу
	cancel()        // завершение горутины по контексту
	time.Sleep(2 * time.Second)
}
