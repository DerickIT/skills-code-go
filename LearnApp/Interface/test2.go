package interface_test

import (
	"errors"
	"fmt"
	"strconv"
)

type WalkRun interface {
	Walk()
	Run()
}

type error interface {
	Error() string
}
type person struct {
	name string
}

func (p *person) Walk() {
	fmt.Printf("%s is walking\n", p.name)
}

func (p *person) Run() {
	fmt.Printf("%s is running\n", p.name)
}

func main() {
	i, err := strconv.Atoi("a")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}

	sum, err := add(1, 2)
	if cm, ok := err.(*commonError); ok {
		fmt.Println("error code:", cm.errorCode, "error msg:", cm.errorMsg)
	} else {
		fmt.Println(sum)
	}

}

func add(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("should be non-negative numbers")
	} else {
		return a + b, nil
	}
}

// 自定义异常
type commonError struct {
	errorCode int
	errorMsg  string
}

func (ce *commonError) Error() string {
	return ce.errorMsg
}

func divide(a, b int) (int, error) {

	return 0, &commonError{errorCode: 1001, errorMsg: "divide by zero"}
}
