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
		myChan <- true // начали свою недолгосрочную работу
		for i := 0; i < 6; i++ {
			fmt.Println("Go-routine still working") // имитация работы горутины
			time.Sleep(500 * time.Millisecond)      // буквально через 3 секунды горутина закончит работу
		}
	}(myChan)

	if <-myChan {
		fmt.Println("I gonna kill you with long time waiting")
		// режим хатико)
	}
	time.Sleep(3 * time.Second) // ждем пока горутина закончит
	for i := 0; i < 5; i++ {
		fmt.Println("Main is still working") // имитация работы main
		time.Sleep(500 * time.Millisecond)
	}
	// метод 3 реализован. Сигнала из вне нет, горутина просто сама завершилась.
}
