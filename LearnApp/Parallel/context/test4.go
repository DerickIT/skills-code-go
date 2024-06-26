package main

import (
	"context"
	"fmt"
	"time"
)

var key string = "name"

func main4() {

	ctx, cancel := context.WithCancel(context.Background())
	valueCtx := context.WithValue(ctx, key, "【监控1】")
	go watch1(valueCtx)
	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")

	cancel()
	time.Sleep(5 * time.Second)
}

func watch1(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Value(key), "监控退出，停止了...")
			return
		default:
			fmt.Println(ctx.Value(key), "goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}
