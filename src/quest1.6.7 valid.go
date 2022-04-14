package main

import (
	"fmt"
	"time"
)

func main() {
	//7. указатель, как глобальная переменная, только круче.
	// Изначально хотел сделать такую же штуку по глобальной переменной, но мне сказали,
	// что ГП это не круто.

	myChan := make(chan bool)
	f := false
	go func(myChan chan bool, f *bool) {
		fmt.Println("Work of go-routine is beginning")
		myChan <- true
		for {
			if *f {
				return
			} else {
				fmt.Println("Go-routine still working")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(myChan, &f)

	if <-myChan {
		fmt.Println("I gonna kill you with pointer")
		time.Sleep(1 * time.Second)
		f = true
	}

	for i := 0; i < 5; i++ {
		fmt.Println("Main is still working")
		time.Sleep(500 * time.Millisecond)
	}
	// метод 3 реализован
}
