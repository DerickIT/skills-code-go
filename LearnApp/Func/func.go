package main

import (
	"fmt"
)

type S struct {
	T
}

type T struct {
	int
}

func (t T) testT() {
	fmt.Println("如类型 S 包含匿名字段 T，则 S 和 *S 方法集包含 T 方法。")
}

func main() {
	s1 := S{T{1}}
	s2 := &s1
	fmt.Printf("s1 is : %v\n", s1)
	s1.testT()
	fmt.Printf("s2 is : %v\n", s2)
	s2.testT()
}
