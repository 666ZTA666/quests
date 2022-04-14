package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//6. Рутина ждёт разблокировки мьютекса.
	//Утечка горутины через sleep
	mu := new(sync.Mutex)
	myChan := make(chan bool)
	go func(myChan chan bool, mu *sync.Mutex) {
		fmt.Println("Work of go-routine is beginning")
		myChan <- true // начали работу, отправили сигнал
		for i := 0; i < 100; i++ {
			mu.Lock()                               // "условно" лочим мьютекс на записи какой-нибудь переменной
			mu.Unlock()                             // анлочим мьютекс после записи
			fmt.Println("Go-routine still working") //имитация работы горутины
			time.Sleep(500 * time.Millisecond)
		}
	}(myChan, mu)

	if <-myChan {
		fmt.Println("I gonna kill you with mutex")
		// получили сигнал, дадим горутине немного времени и залочим мьютекс у себя в main
		time.Sleep(3 * time.Second)
		mu.Lock()
	}
	for i := 0; i < 5; i++ {
		fmt.Println("Main is still working") // имитация работы main
		time.Sleep(500 * time.Millisecond)
	}
	// Метод 6 реализован. Опять-таки, как и с waitgroup, как и с пустым каналом,
	// из которого ведется чтение в горутине, все эти методы лишь усыпляют горутину
	// до завершения программы, где она завершается вместе с main'ом.
}
