package main

import (
	"fmt"
	"sync"
	"time"
)

//Sqads функция, которая возвращает нам квадрат числа и выводит его в stdin.
func Sqads(u int) (v int) {
	v = u * u
	fmt.Println(v)
	return v
}

func main() {
	const k = 5
	//Выводим время начала работы
	fmt.Println(time.Now())
	//Инициализируем срез на 5 элементов
	var i = make([]int, k)
	//Создаем WaitGroup
	wg := new(sync.WaitGroup)
	//
	mu := new(sync.Mutex)
	//Заполняем срез значениями 2,4,6,8,10
	for j := 1; j <= k; j++ {
		i[j-1] = j * 2
	}
	fmt.Println(time.Now())  // выводим время перед началом работы доп. горутин
	wg.Add(k)                // добавляем в WG 5 горутин
	var v int                // переменная для суммы квадратов
	for j := 0; j < k; j++ { //цикл на запуск горутин
		go func(j int, i []int, v *int, wg *sync.WaitGroup, mu *sync.Mutex) {
			mu.Lock()         //лочим
			*v += Sqads(i[j]) //Обрабатываем
			mu.Unlock()       //анлочим
			wg.Done()         //готово
		}(j, i, &v, wg, mu)
	}
	wg.Wait()               //ждем все горутины
	fmt.Println(v)          //выводим сумму квадратов
	fmt.Println(time.Now()) //и время завершения
}
