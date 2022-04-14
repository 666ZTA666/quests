package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//1. Остановка через контекст: <-ctx.Done() прилетел (смотри документацию пакета context).
	// Способ №1 завершен, аналогично можно сделать через таймаут или дедлайн.
	// Так давайте сделаем.
	// валидно
	myChan := make(chan bool)
	ctx, cancel := context.WithTimeout(context.TODO(), 3*time.Second)
	defer cancel()
	go func(ctx context.Context, myChan chan bool) {
		fmt.Println("Work of go-routine is beginning")
		myChan <- true
		for {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println("Go-routine still working")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx, myChan)

	if <-myChan {
		fmt.Println("I gonna kill you with context timeout")
	}
	time.Sleep(2 * time.Second)
	for i := 0; i < 10; i++ {
		fmt.Println("Main is still working")
		time.Sleep(500 * time.Millisecond)
	}
	// способ №1.2 завершен.

}
