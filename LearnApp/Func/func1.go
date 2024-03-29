package main

import (
	"fmt"
)

type User struct {
	id   int
	name string
}

type Manageer struct {
	User
	title string
}

func (u *User) ToString() string {

	return fmt.Sprintf("User: %p, %v", u, u)

}

func (m *Manageer) ToString() string {

	return fmt.Sprintf("Manageer: %p, %v", m, m)
}

func main80() {

	pos, neg := addr(), addr()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}

	m := Manageer{User{1, "Tom"}, "Administrator"}
	fmt.Println(m.ToString())
	fmt.Println(m.User.ToString())
}

func addr() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		result := a
		a, b = b, a+b
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
