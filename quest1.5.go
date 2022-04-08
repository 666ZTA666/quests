package main

import (
	"context"
	"fmt"
	"math"
	"time"
)

func main() {
	var n = Timeout()
	//var ping = Ping()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Duration(n)*time.Second)
	defer cancel()
	fmt.Println("запускаем секундомер")
	start := time.Now()
	myChan := make(chan int)
	go myFunc(ctx, myChan /*, ping*/)
	for v := range myChan {
		fmt.Println("значение из канала:", v)
	}
	fmt.Println("реальное время работы", time.Now().Sub(start), "\nожидаемое время работы", time.Duration(n)*time.Second)
}
func myFunc(ctx context.Context, myChan chan int /*, ping int*/) {
	var j int
	for j < math.MaxInt {
		if ok := ctx.Err(); ok == nil {
			fmt.Print("отправляем в канал ", j, "\t\t")
			myChan <- j
			j++
			//time.Sleep(time.Duration(ping) * time.Millisecond)
		} else {
			close(myChan)
			break
		}
	}
}

/*func Ping() (ping int) {
	for {
		fmt.Println("введите целое количество миллисекунд, через которое будут отправляться данные в канал")
		_, err := fmt.Scanln(&ping)
		if err != nil {
			fmt.Println(err)
		} else if ping <= 0 {
			fmt.Println("значение должно быть больше 0, хотя бы 1)")
		} else {
			fmt.Println(ping, "миллисекунд")
			break
		}
	}
	return
}*/

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
