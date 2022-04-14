package main

import (
	"fmt"
	"time"
)

func main() {
	//2. Остановка через сигнал из канала.
	// валидно

	closeChan := make(chan bool)
	myChan := make(chan bool)
	go func(closeChan, myChan chan bool) {
		fmt.Println("Work of go-routine is beginning")
		myChan <- true
		for {
			select {
			case <-closeChan:
				return
			default:
				fmt.Println("Go-routine still working")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(closeChan, myChan)

	if <-myChan {
		fmt.Println("I gonna kill you with chanel signal")
		closeChan <- true
	}
	for i := 0; i < 5; i++ {
		fmt.Println("Main is still working")
		time.Sleep(500 * time.Millisecond)
	}
	// метод 2.0 реализован
}
