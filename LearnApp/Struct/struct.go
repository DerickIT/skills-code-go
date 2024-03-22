package main

import "fmt"

type NewInt int
type MyInt = int

type student struct {
	id   int
	name string
	age  int
}

func demo(ce []student) {
	//切片是引用传递，是可以改变值的
	ce[1].age = 999
	// ce = append(ce, student{3, "xiaowang", 56})
	// return ce
}

func main() {

	var a NewInt
	var b MyInt
	fmt.Printf("type of a:%T\n", a)
	fmt.Printf("type of b:%T\n", b)

	var ce []student //定义一个切片类型的结构体
	ce = []student{
		student{1, "xiaoming", 22},
		student{2, "xiaozhang", 33},
	}
	fmt.Println(ce)
	demo(ce)
	fmt.Println(ce)
}
