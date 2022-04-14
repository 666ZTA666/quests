package main

import (
	"fmt"
	"time"
)

func main() {
	//2. Остановка через сигнал из канала.
	//Утечка горутины через sleep

	closeChan := make(chan bool)
	myChan := make(chan bool)
	go func(closeChan, myChan chan bool) {
		fmt.Println("Work of go-routine is beginning")
		myChan <- true
		for i := 1; i < 500; i++ {
			fmt.Println("Go-routine still working")
			time.Sleep(500 * time.Millisecond)
			if i%5 == 0 {
				<-closeChan
			}
		}
	}(closeChan, myChan)

	if <-myChan {
		fmt.Println("I gonna kill you with no chanel signal")
	}
	time.Sleep(2 * time.Second)
	for i := 0; i < 5; i++ {
		fmt.Println("Main is still working")
		time.Sleep(500 * time.Millisecond)
	}
	// Метод 2.1 реализован. По факту горутина не завершает свою работы до конца работы всей программы,
	// но засыпает в ожидании сигнала в канал, которого не будет.
}
