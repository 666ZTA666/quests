package main

import (
	"fmt"
	"sync"
	"time"
)

//Sqad функция выводящая квадрат числа
func Sqad(u int) {
	fmt.Println(u * u)
}
func main() {
	const k = 5
	//Выводим время начала
	fmt.Println(time.Now())
	//Создадим срез
	var i = make([]int, k)
	//Для синхронизации дополнительных горутин и main используем sync.WaitGroup
	wg := new(sync.WaitGroup)
	//Мьютекс, чтобы обеспечить конкуретность работы
	mu := new(sync.Mutex)
	//В цикле заполняем срез значениями 2,4,6,8,10
	for j := 1; j <= k; j++ {
		i[j-1] = j * 2
		//fmt.Println(i[j-1])
	}
	fmt.Println(time.Now()) //Выводим время перед началом запуска параллельных вычислений
	wg.Add(k)               //Добавим 5 горутин
	for j := 0; j < k; j++ {
		// В цикле создаем горутины, передаем в горутину срез и индекс элемента, который надо возвести в квадрат
		// а так же указатель на wg чтобы отдавать сигнал о завершении работы горутины
		// и мьютекс, чтобы из паралельной работы сделать конкурентную
		go func(j int, i []int, wg *sync.WaitGroup, mu *sync.Mutex) {
			mu.Lock()   //Лочим
			Sqad(i[j])  //обращаемся
			mu.Unlock() //анлочим
			wg.Done()   //подаем сигнал о завершении
		}(j, i, wg, mu)
	}
	wg.Wait()               //ждем завершения всех горутин
	fmt.Println(time.Now()) //выводим время завершения работы
}
