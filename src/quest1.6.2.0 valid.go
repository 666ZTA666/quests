package main

import (
	"fmt"
	"time"
)

func main() {
	//2. Остановка через сигнал из канала.
	// валидно
	// канал с сигналом о завершении работы
	closeChan := make(chan struct{})
	//канал с сигналом о начале работы горутины
	myChan := make(chan bool)
	go func(closeChan chan struct{}, myChan chan bool) {
		fmt.Println("Work of go-routine is beginning")
		myChan <- true // начинаем работу, отправляем сигнал
		for {
			select {
			case <-closeChan: // ждем сигнала на закрытие и выходим
				return
			default:
				fmt.Println("Go-routine still working") // Демонстрируем работу горутины
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(closeChan, myChan)

	if <-myChan {
		fmt.Println("I gonna kill you with chanel signal")
		// сигнал о начале пришел, отправляем сигнал о завершении
		closeChan <- struct{}{}
	}
	for i := 0; i < 5; i++ {
		fmt.Println("Main is still working") // демонстрация работы main
		time.Sleep(500 * time.Millisecond)
	}
	// метод 2.0 реализован
}
