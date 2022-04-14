package main

import (
	"fmt"
	"math"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var c int
	w, err := os.Create("log_to_4_quest.txt")
	if err != nil {
		fmt.Println("ошибка открытия лога")
		return
	}
	for {
		fmt.Println("введите количество воркеров)")
		_, err := fmt.Fprintln(w, "введите количество воркеров)")
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = fmt.Scanln(&c)
		if err != nil {
			return
		}
		_, err = fmt.Fprintln(w, c)
		if err != nil {
			fmt.Println(err)
			return
		}
		if c <= 0 {
			fmt.Println("количество воркеров должно быть больше 0")
			_, err := fmt.Fprintln(w, "количество воркеров должно быть больше 0")
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			break
		}
	}
	mychan := make(chan int)
	chanelClosing := make(chan int)
	for i := 1; i < c+1; i++ {
		fmt.Print("запуск воркера №", i, "\n")
		_, err := fmt.Fprint(w, "запуск воркера №", i, "\n")
		if err != nil {
			fmt.Println(err)
			return
		}
		go func(i int) {
			var v int
			var ok bool
			for {
				if v, ok = <-mychan; ok {
					fmt.Print("воркер №", i, ":", v, "\n")
					_, err := fmt.Fprint(w, "воркер №", i, ":", v, "\n")
					if err != nil {
						fmt.Println(err)
						return
					}
				} else {
					fmt.Print("канал для воркера №", i, " закрыт, завершаем работу\n")
					_, err := fmt.Fprint(w, "канал для воркера №", i, " закрыт, завершаем работу\n")
					if err != nil {
						fmt.Println(err)
						return
					}
					chanelClosing <- i
					break
				}
			}
		}(i)
	}

	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		var i int
		for i < math.MaxInt {
			select {
			case <-signalChan:
				fmt.Println(time.Now())
				close(mychan)
				fmt.Println("закрываем канал передачи данных")
				_, err = fmt.Fprintln(w, "закрываем канал передачи данных")
				if err != nil {
					fmt.Println(err)
					return
				}
				cleanupDone <- true
				return
			default:
				mychan <- i
				i++
				//time.Sleep(time.Second)
			}
		}
	}()
	<-cleanupDone
	for i := 0; i < c; i++ {
		val := <-chanelClosing
		fmt.Print("воркер №", val, " прекратил работу\n")
		_, err = fmt.Fprint(w, "воркер №", val, " прекратил работу\n")
		if err != nil {
			fmt.Println(err)
		}
	}
	err = w.Close()
	if err != nil {
		fmt.Println("ошибка закрытия лог файла", err)
	}
	fmt.Println(time.Now())
}
