package main

import (
	"context"
	"fmt"
	"math"
	"time"
)

func main() {
	//внизу есть код функции который возвращает в переменную значение введенное с консоли
	//с проверкой на некую "правильность" этого значения
	var n = Timeout()
	//создаем контекст с таймаутом. В следующем файле контекст будет с дедлайном
	ctx, cancel := context.WithTimeout(context.TODO(), time.Duration(n)*time.Second)
	defer cancel()
	fmt.Println("запускаем секундомер")
	start := time.Now() // записываем текущее время в переменную старт
	myChan := make(chan int)
	go myFunc(ctx, myChan) // запускаем функцию с контекстом, которая проверяет ошибку из контекста,
	// и если ошибка не пустая, то закрываем канал и горутину,
	// аналогочно можно было сделать через select и канал ctx.done
	for v := range myChan {
		fmt.Println("значение из канала:", v)
	} //Пока канал не закрыт мы выводим значения в stdout, после этого выводим время работы и завершаемся
	fmt.Println("реальное время работы", time.Now().Sub(start), "\nожидаемое время работы", time.Duration(n)*time.Second)
}
func myFunc(ctx context.Context, myChan chan int) {
	var j int
	for j < math.MaxInt {
		if ok := ctx.Err(); ok == nil {
			fmt.Print("отправляем в канал ", j, "\t\t")
			myChan <- j
			j++
		} else {
			close(myChan)
			return
		}
	}
}

func Timeout() (n int) {
	for {
		fmt.Println("введите целое количество секунд, через которое программа должна быть закрыта")
		_, err := fmt.Scanln(&n)
		if err != nil {
			fmt.Println(err)
		} else if n <= 0 {
			fmt.Println("значение должно быть больше 0")
		} else {
			fmt.Println(n, "секунд")
			break
		}
	}
	return
}
