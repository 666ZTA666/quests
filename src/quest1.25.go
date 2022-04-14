package main

import (
	"fmt"

	"time"
)

func main() {
	var n int
	for {
		fmt.Println("введите количество секунд сна")
		_, err := fmt.Scanln(&n)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if n > 0 {
			break
		}
	}
	fmt.Println("ну теперь ждите")
	start := time.Now()
	MySleep(time.Duration(n) * time.Second)
	fmt.Println(time.Now().Sub(start))
}
func MySleep(d time.Duration) {
	<-time.After(d)
}
