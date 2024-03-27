package main

import (
	"fmt"
	"runtime"
	"time"
)

// func main() {
// 	go func(s string) {
// 		for i := 0; i < 2; i++ {
// 			fmt.Println(s)
// 		}
// 	}("world")

// 	for i := 0; i < 2; i++ {
// 		runtime.Gosched()
// 		fmt.Println("hello")
// 	}
// }

func m1ain() {
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			runtime.Goexit()
			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
	for {

	}

}

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func ma1in() {
	runtime.GOMAXPROCS(2)
	go a()
	go b()
	time.Sleep(time.Second)
}
