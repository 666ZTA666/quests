package main

import (
	"fmt"
	"time"
)

func main() {
	//2. Остановка через сигнал из канала.
	//Утечка горутины через sleep
	// небуферизированный канал, для блокировки горутины
	closeChan := make(chan struct{})
	myChan := make(chan bool)
	go func(closeChan chan struct{}, myChan chan bool) {
		fmt.Println("Work of go-routine is beginning")
		myChan <- true // начали работу, отправили сигнал
		for i := 1; i < 500; i++ {
			fmt.Println("Go-routine still working")
			time.Sleep(500 * time.Millisecond)
			if i%5 == 0 { // спустя некоторое время решили прочитать из канала, в который никто не пишет и заснули
				<-closeChan
			}
		}
	}(closeChan, myChan)

	if <-myChan {
		fmt.Println("I gonna kill you with no chanel signal")
		// получаем сигнал о начале работы горутины и НЕ посылаем ей сигнал.
	}
	time.Sleep(2 * time.Second)
	for i := 0; i < 5; i++ {
		fmt.Println("Main is still working") // Демонстрация работы main
		time.Sleep(500 * time.Millisecond)
	}
	// Метод 2.1 реализован. По факту горутина не завершает свою работы до конца работы всей программы,
	// но засыпает в ожидании сигнала в канал, которого не будет.
}
