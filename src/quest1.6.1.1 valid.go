package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//1. Остановка через контекст: <-ctx.Done() прилетел
	// Способ №1 завершен, аналогично можно сделать через таймаут или дедлайн.
	// Так давайте сделаем.
	// валидно
	myChan := make(chan bool)
	// Создаем контекст с дэдлайном в 3 секунды
	ctx, cancel := context.WithDeadline(context.TODO(), time.Now().Add(time.Duration(3)*time.Second))
	defer cancel()

	go func(ctx context.Context, myChan chan bool) {
		fmt.Println("Work of go-routine is beginning") // начинаем работу
		myChan <- true                                 // передаем сигнал о начале работы
		for {
			select {
			case <-ctx.Done(): // По контексту выходим
				return
			default:
				fmt.Println("Go-routine still working") //Демонстрируем работу горутины
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx, myChan)

	if <-myChan {
		fmt.Println("I gonna kill you with context deadline") // когда приходит сигнал...
		// Да мы в общем ничего не делаем, там время само все сделает
	}
	time.Sleep(2 * time.Second) // ждем 2 из 3х секунд
	for i := 0; i < 10; i++ {
		fmt.Println("Main is still working") // демонстрация работы main
		time.Sleep(500 * time.Millisecond)
	}
	// способ №1.1 завершен, аналогично можно сделать через таймаут в версии 1.2.

}
