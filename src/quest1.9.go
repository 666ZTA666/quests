package main

import (
	"fmt"
)

func main() {
	mychan := make(chan int)        // первый канал чисел
	myxchan := make(chan int)       //второй канал чисел
	exitchan := make(chan struct{}) // канал закрытия
	go func(mychan chan int) {
		for i := 0; i <= 100; i++ {
			mychan <- i // первый воркер отправляет в канал 101 число, от 0 до 100 включительно.
		}
		fmt.Println("First worker is over. First chanel closed.")
		close(mychan) // как закончит, он закроет за собой канал.
	}(mychan)
	go func(mychan, myxchan chan int) {
		for val := range mychan { // второй воркер работает пока канал из которого он читает не закрыт.
			myxchan <- val * val
		}
		fmt.Println("Second worker is over. Second chanel closed.")
		close(myxchan) // а потом закрывает за собой канал, в который писал данные.
	}(mychan, myxchan)

	go func(myxchan chan int, exitchan chan struct{}) {
		for val := range myxchan { // третий воркер читает данные из канала, пока тот не закрыт и выводит их в stdout
			fmt.Println(val)
		}
		fmt.Println("All work is over.") // а потом отправляет в main сигнал о завершении работы.
		exitchan <- struct{}{}
	}(myxchan, exitchan)
	<-exitchan // ждем сигнала завершения работы.
	// видел классный аналог, где функцию последнего воркера выполняем main-горутина, это уменьшает нагрузку на память
	//выделяемую на новые горутины.
}
