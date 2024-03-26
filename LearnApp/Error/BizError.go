package main

import (
	"fmt"
	"os"
	"time"
)

func test01() {
	a := [5]int{1, 2, 3, 4, 5}
	a[1] = 10
	fmt.Println(a)
	index := 11
	a[index] = 11
	fmt.Println(a)
}

func getCircleArea(radius float32) (area float32) {
	if radius < 0 {
		panic("半径不能为负数")
	}
	return 3.14 * radius * radius
}

func test02() {
	getCircleArea(-5)
}

func test03() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("发生错误", err)
		}
	}()
	getCircleArea(-5)
	fmt.Println("没有执行")
}

func test04() {
	test03()
	fmt.Println("test04执行完毕")
}

type PathError struct {
	path       string
	op         string
	createTime string
	message    string
}

func (p *PathError) Error() string {
	return fmt.Sprintf("path=%s,op=%s,createTime=%s,message=%s", p.path, p.op, p.createTime, p.message)
}

func Open(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return &PathError{
			path:       filename,
			op:         "read",
			message:    err.Error(),
			createTime: fmt.Sprintf("%v", time.Now()),
		}
	}
	defer file.Close()
	return nil
}

func main() {

	test04()

	err := Open("test.txt")
	switch v := err.(type) {
	case *PathError:
		fmt.Println("get path error,", v)
	default:
	}
}
