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
	wg.Add(2)
	myChan := make(chan bool)
	go func(myChan chan bool, wg *sync.WaitGroup) {
		fmt.Println("Work of go-routine is beginning")
		myChan <- true
		for i := 1; i < 6; i++ {
			fmt.Println("Go-routine still working")
			time.Sleep(500 * time.Millisecond)
			if i%4 == 0 {
				wg.Wait()
			}
		}
	}(myChan, wg)

	if <-myChan {
		fmt.Println("I gonna kill you with WaitGroup")
	}
	time.Sleep(2 * time.Second)
	for i := 0; i < 5; i++ {
		fmt.Println("Main is still working")
		time.Sleep(500 * time.Millisecond)
	}
	// метод 5 реализован. Как и в методе с каналом без сигнала горутина засыпает,
	// ожидая завершения всех членов группы. И завешается уже вместе с основной горутиной программы.
}
