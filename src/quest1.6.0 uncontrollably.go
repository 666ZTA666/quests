package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	//0. Приложение померло.
	//не подконтрольно = не валидно
	//Есть канал, в котором информация о начале работы горутины
	myChan := make(chan bool)
	go func(myChan chan bool) {
		fmt.Println("Work of go-routine is beginning") //горутина начала работу
		myChan <- true                                 // в канал послали информацию
		for {
			fmt.Println("Still working") //демонстрируем работу горутины
			time.Sleep(500 * time.Millisecond)
		}
	}(myChan)
	if <-myChan {
		fmt.Println("I gonna kill you with exit") //Получили сигнали, сделали exit
		// горутина завершилась вместе с main. Это не есть хорошо на самом деле, но условие было написано не оч...
		os.Exit(0)
	}
}
