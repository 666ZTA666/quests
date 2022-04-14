package main

import (
	"fmt"
	"sync"
)

func main() {
	mapa := make(map[int]string) // создаем map: ключ - число, значение - строка
	// создаем рв мьютекс
	mu := new(sync.RWMutex)
	myxChan := make(chan int)
	// в цикле создаем по горутине, в каждую передаем ее номер, рв мьютекс и канал
	for i := 1; i <= 5; i++ {
		go func(i int, mu *sync.RWMutex, myxChan chan int) {
			mu.Lock()                   // лочим на запись
			mapa[i] = fmt.Sprint(i * 5) // пишем
			mu.Unlock()                 // анлочим на запись
			fmt.Println("записали", i)  // выводим информацию, что что-то записали
			myxChan <- i                // передаем в канал ключ по которому есть информация
		}(i, mu, myxChan)
	}
	for i := 0; i < 5; i++ {
		// получаем из канала ключ по которому лежит значение
		x := <-myxChan
		mu.RLock()                           // лочим на чтение
		fmt.Println("прочитали", x, mapa[x]) // читаем
		mu.RUnlock()                         // анлочим на чтение
	}
	//посмотрел sync.map, он здесь нахрен не нужен, у меня дома столько процессорных Ядер нет(
}
