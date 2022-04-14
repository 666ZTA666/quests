package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//5. Рутина засыпает в ожидании sync.WaitGroup.
	//Утечка горутины через sleep
	wg := new(sync.WaitGroup)
	// создаем группу ожидания из двух потоков
	wg.Add(2)
	myChan := make(chan bool)
	go func(myChan chan bool, wg *sync.WaitGroup) {
		fmt.Println("Work of go-routine is beginning")
		myChan <- true // начали работу, отправили сиганал
		for i := 1; i < 6; i++ {
			fmt.Println("Go-routine still working") // имитация работы горутины
			time.Sleep(500 * time.Millisecond)
			if i%4 == 0 {
				wg.Wait() // ждем завершения работы всей группы
			}
		}
	}(myChan, wg)

	if <-myChan {
		fmt.Println("I gonna kill you with WaitGroup")
		// получили сигнал, но не закончили ни одной функции из wg
	}
	time.Sleep(2 * time.Second) // ждем пока горутина не заснет
	for i := 0; i < 5; i++ {
		fmt.Println("Main is still working") // имитация работы main
		time.Sleep(500 * time.Millisecond)
	}
	// метод 5 реализован. Как и в методе с каналом без сигнала горутина засыпает,
	// ожидая завершения всех членов группы. И завешается уже вместе с основной горутиной программы.
}
