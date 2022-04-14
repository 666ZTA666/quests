package main

import "fmt"

type Human struct {
	name string
	age  uint
}

func (h *Human) Hello() *Human {
	fmt.Println("hello, my name is", h.name)
	return h
}
func (h *Human) GetAge() *Human {
	fmt.Println("my age is", h.age)
	return h
}
func (h *Human) SetName(name string) *Human {
	h.name = name
	return h
}
func (h *Human) SetAge(age uint) *Human {
	h.age = age
	return h
}

type Action struct {
	Human
}

func main() {
	var john = Human{name: "John", age: 18}
	john.Hello().GetAge()
	SomeAction := new(Action)
	SomeAction.SetName("action").SetAge(5).Hello().GetAge()
}
