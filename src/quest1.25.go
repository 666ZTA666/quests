package main

import (
	"fmt"

	"time"
)

func main() {
	var n int
	for {
		fmt.Println("введите количество секунд сна")
		_, err := fmt.Scanln(&n) //считали записали
		if err != nil {
			fmt.Println(err)
			continue
		}
		if n > 0 { // проверка на 0
			break
		}
	}
	fmt.Println("ну теперь ждите")
	start := time.Now()                     // отсечка времени
	MySleep(time.Duration(n) * time.Second) // мой сон
	fmt.Println(time.Now().Sub(start))      // время после сна
}
func MySleep(d time.Duration) {
	<-time.After(d) // сигнал из канала по прошествии времени.
}
