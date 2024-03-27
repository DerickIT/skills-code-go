package main

import (
	"fmt"
	"time"
)

func time1() {

	time1 := time.NewTimer(2 * time.Second)
	t1 := time.Now()
	fmt.Printf("t1:%v\n", t1)
	t2 := <-time1.C
	fmt.Printf("t2:%v\n", t2)
}

func time2() {

	time2 := time.NewTimer(time.Second)
	for {
		<-time2.C
		fmt.Println("时间到")
		//time2.Reset(time.Second)
	}

}

func time3() {
	time.Sleep(time.Second)
	timer3 := time.NewTimer(2 * time.Second)
	<-timer3.C
	fmt.Println("2秒到")

	// 2秒后，向timer.C写数据，有数据后，就可以读取
	<-time.After(2 * time.Second)
	fmt.Println("2秒到")

}

func time4() {
	timer4 := time.NewTimer(2 * time.Second)
	go func() {
		<-timer4.C
		fmt.Println("定时器执行了")
	}()
	b := timer4.Stop()
	if b {
		fmt.Println("timer4 停止了 已经关闭")
	}

}

func time5() {
	timer5 := time.NewTimer(3 * time.Second)
	timer5.Reset(1 * time.Second)
	fmt.Println(time.Now())
	fmt.Println(<-timer5.C)

}

func main() {
	time1()
	//time2()
	time3()
	time4()
	time5()

	ticker := time.NewTicker(1 * time.Second)
	i := 0
	go func() {
		for {
			i++
			fmt.Println(<-ticker.C)
			if i == 5 {
				ticker.Stop()
			}
		}
	}()
	for {

	}

	/*

			这段代码演示了如何使用Go语言的`time`包中的`Ticker`来实现定时器功能。下面是对这段代码的详细解释:

		1. `ticker := time.NewTicker(1 * time.Second)`
		   - 创建一个新的`Ticker`对象,它会每隔1秒钟发送一个时间值到其通道中。
		   - `time.Second`表示1秒钟的时间间隔。

		2. `i := 0`
		   - 定义一个整数变量`i`,并初始化为0。这个变量用于记录定时器触发的次数。

		3. `go func() { ... }()`
		   - 启动一个新的goroutine来执行匿名函数。
		   - 这个goroutine会独立于主goroutine运行。

		4. `for { ... }`
		   - 在匿名函数内部,使用一个无限循环来不断地从`ticker.C`通道中接收时间值。

		5. `i++`
		   - 每次从`ticker.C`通道中接收到时间值后,将`i`的值加1,表示定时器触发的次数增加。

		6. `fmt.Println(<-ticker.C)`
		   - 从`ticker.C`通道中接收时间值,并使用`fmt.Println`将其打印出来。
		   - `<-ticker.C`表示从`ticker.C`通道中接收一个时间值。

		7. `if i == 5 { ticker.Stop() }`
		   - 当`i`的值等于5时,表示定时器已经触发了5次。
		   - 调用`ticker.Stop()`方法来停止定时器,停止向`ticker.C`通道发送时间值。

		8. `for { }`
		   - 在主goroutine中,使用一个无限循环`for { }`来阻塞主goroutine的退出。
		   - 这样可以确保程序不会立即退出,而是等待其他goroutine完成。

		总体来说,这段代码的作用是:
		1. 创建一个定时器`ticker`,它每隔1秒钟会向其通道`ticker.C`发送一个时间值。
		2. 在一个新的goroutine中,通过无限循环从`ticker.C`通道中接收时间值,并打印出来。
		3. 当定时器触发了5次后,调用`ticker.Stop()`方法停止定时器。
		4. 在主goroutine中,使用无限循环`for { }`阻塞主goroutine的退出,以确保其他goroutine能够完成执行。

		这段代码展示了如何使用`Ticker`实现定时器功能,并在满足特定条件时停止定时器。同时,它也演示了如何使用goroutine来并发执行任务,以及如何阻塞主goroutine以等待其他goroutine完成。

	*/
}
