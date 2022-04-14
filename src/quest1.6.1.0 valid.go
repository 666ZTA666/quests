package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//1. Остановка через контекст: <-ctx.Done() прилетел.
	//валидно

	//Есть канал, в котором информация о начале работы горутины
	myChan := make(chan bool)
	// создаем контекст с функцией отмены
	ctx, cancel := context.WithCancel(context.TODO())
	go func(ctx context.Context, myChan chan bool) {
		fmt.Println("Work of go-routine is beginning") //горутина начала работу
		myChan <- true                                 // в канал послали информацию
		for {                                          // имитация бурной деятельности в ожидании завершения
			select {
			case <-ctx.Done(): //при получении сигнала - выходим из функции, запущенной в отдельной горутине
				return
			default:
				fmt.Println("Go-routine still working") //демонстрируем работу горутины
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx, myChan)

	if <-myChan {
		fmt.Println("I gonna kill you with context cancel") //Получили сигнали, сделали cancel
		cancel()
	}
	for i := 0; i < 10; i++ {
		fmt.Println("Main is still working") // демонстрируем работу main
		time.Sleep(500 * time.Millisecond)
	}
	// способ №1 завершен, аналогично можно сделать через таймаут или дедлайн

}
