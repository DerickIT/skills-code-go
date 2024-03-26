package main

type Sayer interface {
	say()
}

type Mover interface {
	move()
}

type dog struct{}

type cat struct{}

func (d *dog) say() {
	println("汪汪汪")
}

func (d *dog) move() {
	println("狗会动")
}

func (c cat) say() {
	println("喵喵喵")
}
func (c cat) move() {
	println("猫会动")
}

func main() {
	var s Sayer
	var m Mover
	a := cat{}
	b := &dog{}
	s = a
	m = a
	s.say()
	m.move()
	s = b
	m = b
	s.say()
	m.move()
}
