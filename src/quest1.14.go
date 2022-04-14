package main

import (
	"fmt"
	"reflect"
)

func main() {
	//Разработать программу, которая в рантайме способна определить тип
	//переменной: int, string, bool, channel из переменной типа interface{}
	mychan := make(chan interface{})

	go func(mychan chan interface{}) {
		for a := range mychan {
			switch v := a.(type) {
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
	}(mychan)

	x := make(chan struct{})
	var (
		s string
		b bool
		i int
	)
	mychan <- x
	mychan <- s
	mychan <- b
	mychan <- i

	close(mychan)

}
