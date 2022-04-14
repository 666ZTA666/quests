package main

import (
	"fmt"
	"time"
)

func main() {
	//4. Рутина запаниковала
	//не подконтрольно = не валидно
	myChan := make(chan bool)
	go func(myChan chan bool) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Go-routine end, panic is recovered:", r)
			}
		}()
		fmt.Println("Work of go-routine is beginning")
		myChan <- true
		for i := 0; i < 6; i++ {
			fmt.Println("Go-routine still working")
			time.Sleep(500 * time.Millisecond)
			if i%2 == 0 {
				panic("just panic")
			}
		}
	}(myChan)
	if <-myChan {
		fmt.Println("I just wait for your panic")
	}
	for i := 0; i < 5; i++ {
		fmt.Println("Main is still working")
		time.Sleep(500 * time.Millisecond)
	}
	// Если не ловить панику рековером, то получится в принципе первый случай,
	// где программа падает вместе с горутинами, что уже было приведено в первом варианте.
}
