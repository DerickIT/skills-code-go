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

func main() {

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
