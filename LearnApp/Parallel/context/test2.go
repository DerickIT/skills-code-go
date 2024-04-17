package main

import (
	"fmt"
	"time"
)

func main1() {

	fmt.Println("Hello World")
	stop := make(chan bool)
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("Done")
				return
			default:
				fmt.Println("Working")
				time.Sleep(2 * time.Second)
			}
		}

	}()

	time.Sleep(10 * time.Second)
	fmt.Println("Cancelling the 监控")
	stop <- true

	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}
