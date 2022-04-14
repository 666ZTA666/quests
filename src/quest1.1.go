package main

import "fmt"

//Human структура с произвольным набором полей
type Human struct {
	name string
	age  uint
}

//Методы структуры Human
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

//Action структура "наследующая" от Human методы
type Action struct {
	Human
}
type ActionHuman struct {
	h Human
}

func main() {
	//Создаем и инициализируем структуру Human
	var john = Human{name: "John", age: 18}
	//Используем методы структуры
	john.Hello().GetAge()
	//Создаем структуру Action
	SomeAction := new(Action)
	//Инициализируем поля встроенными в human методами и выводим значения этих полей.
	SomeAction.Hello().GetAge().SetName("action").SetAge(5).Hello().GetAge()
	//Через встраивание
	SomeNewAction := new(ActionHuman)
	SomeNewAction.h.SetName("new action").SetAge(10).Hello().GetAge()
}
