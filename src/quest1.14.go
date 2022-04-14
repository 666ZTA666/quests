package main

import (
	"fmt"
	"reflect"
)

func main() {
	//Разработать программу, которая в рантайме способна определить тип
	//переменной: int, string, bool, channel из переменной типа interface{}
	mychan := make(chan interface{})
	// создали канал пустых интерфейсов

	go func(mychan chan interface{}) {
		// запускаем горутину, которая читает из канала пустые интерфейсы
		for a := range mychan {
			switch v := a.(type) {
			// v это тип переменной a
			case int:
				fmt.Print("через switch int '", v, "'\n")
			case bool:
				fmt.Print("через switch bool '", v, "'\n")
			case string:
				fmt.Print("через switch string '", v, "'\n")
			case chan struct{}:
				fmt.Print("через switch chan '", v, "'\n")
			default:
				fmt.Print("через switch unknown")
			}
			fmt.Print("через рефлексию: ", reflect.TypeOf(a), " '", reflect.ValueOf(a), "'\n")
			fmt.Printf("Через принтф: %T '%v'\n\n", a, a)
		}
		// типы данных выводятся, но так же могут быть сохранены где-то или аналогично switch case определять алгоритм работы.
	}(mychan)

	x := make(chan struct{})
	// создаем канал
	var (
		s string
		b bool
		i int
	)
	//создаем переменные типа строка, булева переменная и число.
	mychan <- x
	mychan <- s
	mychan <- b
	mychan <- i
	// отправляем все это в канал
	close(mychan)
	// закрываем канал.

}
