package main

import (
	"fmt"
	"time"
)

func main() {
	//3. Рутина завершила работу штатно.
	// не валидно, сигнала из "вне" нет.

	myChan := make(chan bool)
	go func(myChan chan bool) {
		fmt.Println("Work of go-routine is beginning")
		myChan <- true
		for i := 0; i < 6; i++ {
			fmt.Println("Go-routine still working")
			time.Sleep(500 * time.Millisecond)
		}
	}(myChan)

	if <-myChan {
		fmt.Println("I gonna kill you with long time waiting")
	}
	time.Sleep(3 * time.Second)
	for i := 0; i < 5; i++ {
		fmt.Println("Main is still working")
		time.Sleep(500 * time.Millisecond)
	}
	// метод 3 реализован
}
