package main

import (
	"fmt"
	"time"
)

func main() {
	//7. указатель, как глобальная переменная, только круче.
	// Изначально хотел сделать такую же штуку по глобальной переменной, но мне сказали,
	// что ГП это не круто.

	myChan := make(chan bool)
	f := false                           // создаем переменную типа bool в которой по дефолту хранится 0 == false
	go func(myChan chan bool, f *bool) { // переменную передаем по указателю, чтобы узнать об изменении.
		fmt.Println("Work of go-routine is beginning")
		myChan <- true // начали работу, отправили сигнал
		for {
			if *f { // в бесконечном цикле проверяем не поменялось ли значение переменной
				return // если переменная стала true == 1, то уходим
			} else { // иначе имитируем работу горутины
				fmt.Println("Go-routine still working")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(myChan, &f)

	if <-myChan {
		fmt.Println("I gonna kill you with pointer")
		// получили сигнал о начале работы горутины, ждем чуток, и меняем значение переменной.
		time.Sleep(1 * time.Second)
		f = true
	}

	for i := 0; i < 5; i++ {
		fmt.Println("Main is still working") // имитируем работу main
		time.Sleep(500 * time.Millisecond)
	}
	// в принципе не сильно дорогой метод для закрытия кучи горутин, но закрыть канал пустых структур всё таки дешевле,
	// но метод валидный и, если захотеть можно и ему найти применения.
}
