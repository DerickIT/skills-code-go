package main

import "fmt"

/*
   1. 切片：切片是数组的一个引用，因此切片是引用类型。但自身是结构体，值拷贝传递。
   2. 切片的长度可以改变，因此，切片是一个可变的数组。
   3. 切片遍历方式和数组一样，可以用len()求长度。表示可用元素数量，读写操作不能超过该限制。
   4. cap可以求出slice最大扩张容量，不能超出数组限制。0 <= len(slice) <= len(array)，其中array是slice引用的数组。
   5. 切片的定义：var 变量名 []类型，比如 var str []string  var arr []int。
   6. 如果 slice == nil，那么 len、cap 结果都等于 0。
*/

func main() {
	//声名一个切片
	var s1 []int
	if s1 == nil {
		fmt.Println("s1 is nil")
	} else {
		fmt.Println("s1 is not nil")
	}
	// 	//声明一个切片并初始化
	s2 := []int{}

	//make()
	var s3 []int = make([]int, 0)
	fmt.Println(s1, s2, s3)

	var s4 []int = make([]int, 0, 0)
	fmt.Println(s4)
	//从数组切片
	s5 := []int{1, 2, 3}
	fmt.Println(s5)

	arr := [5]int{1, 2, 3, 4, 5}
	var s6 []int
	//前包后不包
	s6 = arr[1:4]
	fmt.Println(s6)

	data := [...]int{0, 1, 2, 3, 4, 10: 0}
	s := data[:2:3]

	s = append(s, 100, 200) // 一次 append 两个值，超出 s.cap 限制。

	fmt.Println(s, data)         // 重新分配底层数组，与原数组无关。
	fmt.Println(&s[0], &data[0]) // 比对底层数组起始指针。

	s11 := []int{0, 1, 2, 3, 8: 100} // 通过初始化表达式构造，可使用索引号。
	fmt.Println(s11, len(s11), cap(s11))

	s22 := make([]int, 6, 8) // 使用 make 创建，指定 len 和 cap 值。
	fmt.Println(s22, len(s22), cap(s22))

	s33 := make([]int, 6) // 省略 cap，相当于 cap = len。
	fmt.Println(s33, len(s33), cap(s33))

	slice1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	d1 := slice1[6:8]
	fmt.Println(d1, len(d1))
	d2 := slice1[:6:8]
	fmt.Println(d2, len(d2))
}
