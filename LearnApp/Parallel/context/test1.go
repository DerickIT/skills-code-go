package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World")
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Done")
				return
			default:
				fmt.Println("Working")
				return
			}
		}
	}(ctx)
	time.Sleep(10 * time.Second)
	fmt.Println("Cancelling the 监控")
	cancel()
	time.Sleep(5 * time.Second)
}
