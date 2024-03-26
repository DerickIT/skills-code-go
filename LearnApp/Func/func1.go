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
	m := Manageer{User{1, "Tom"}, "Administrator"}
	fmt.Println(m.ToString())
	fmt.Println(m.User.ToString())
}
