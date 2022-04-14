package main

import (
	"fmt"
	"sync"
)

func main() {
	mapa := make(map[int]string)
	mu := new(sync.RWMutex)
	myxChan := make(chan int)
	for i := 1; i <= 5; i++ {
		go func(i int, mu *sync.RWMutex, myxChan chan int) {
			mu.Lock()
			mapa[i] = fmt.Sprint(i * 5)
			mu.Unlock()
			fmt.Println("записали", i)
			myxChan <- i
		}(i, mu, myxChan)
	}
	for i := 0; i < 5; i++ {
		x := <-myxChan
		mu.RLock()
		fmt.Println("прочитали", x, mapa[x])
		mu.RUnlock()
	}
	//посмотрел sync.map, он здесь нахрен не нужен, у меня дома столько процессорных Ядер нет(
}
