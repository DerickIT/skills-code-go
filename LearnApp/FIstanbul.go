package main

import (
	"fmt"

	"rsc.io/quote"
)

func test(x [2]int) {
	fmt.Printf("x:%p\n", &x)
	x[1] = 1000
}

func main() {
	a := [3]int{1, 2}
	b := [...]int{1, 2, 3, 4}
	c := [5]int{3: 23, 1: 545}
	d := [...]struct {
		name string
		age  uint8
	}{
		{"user1", 10},
		{"user2", 5},
	}
	fmt.Println(a, b, c, d)
	fmt.Println(arr0, arr1, arr2, str)

	fmt.Println(quote.Go())
	fmt.Println("hhhhhh")
	println("hello world")

	arrff := [...]int{2, 3, 4, 6, 20}
	printArr(&arrff)
	fmt.Println(arrff)
}

var arr0 [5]int = [5]int{1, 2, 3}
var arr1 = [5]int{1, 2, 3, 4, 5}
var arr2 = [...]int{1, 2, 3, 4, 5, 6}
var str = [5]string{4: "hello", 1: "derick"}

// 多维数组遍历
func elementEach() {
	var f [2][3]int = [...][3]int{{1, 2, 3}, {23, 4, 2}}

	for k1, v1 := range f {

		for k2, v2 := range v1 {
			fmt.Printf("(%d,%d)=%d", k1, k2, v2)
		}
		fmt.Println()
	}
}

func printArr(arr *[5]int) {
	arr[0] = 10
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

//找出数组中和为给定值的两个元素的下标，例如数组[1,3,5,8,7]，
//找出两个元素之和等于8的下标分别是（0，4）和（1，2）

// 求元素和，是给定的值
func myTest(a [5]int, target int) {
	for i := 0; i < len(a); i++ {
		other := target - a[i]
		for j := i + 1; j < len(a); j++ {
			if a[j] == other {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}
}
